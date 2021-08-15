import * as cdk from "@aws-cdk/core";
import * as lambda from "@aws-cdk/aws-lambda";
import * as ec2 from "@aws-cdk/aws-ec2";
import * as rds from "@aws-cdk/aws-rds";
import * as iam from "@aws-cdk/aws-iam";
import * as path from "path";

interface LambdaStackProps extends cdk.StackProps {
  commanderSrcPath: string;
  updaterSrcPath: string;
  vpc: ec2.Vpc;
  lambdaSg: ec2.SecurityGroup;
  rdsProxy: rds.DatabaseProxy;
  dynamoStreamArn: string;
}

export class AwsCdkLambdaStack extends cdk.Stack {
  public readonly commanderLambda: lambda.Function;
  public readonly readModelUpdaterLambda: lambda.Function;
  public readonly lambdaIamRole: iam.Role;

  private readonly vpc: ec2.Vpc;
  private readonly lambdaSg: ec2.SecurityGroup;
  private readonly rdsProxy: rds.DatabaseProxy;

  constructor(scope: cdk.Construct, id: string, props: LambdaStackProps) {
    super(scope, id, props);

    this.vpc = props.vpc;
    this.lambdaSg = props.lambdaSg;
    this.rdsProxy = props.rdsProxy;

    this.lambdaIamRole = this.createLambdaIamRole();
    this.commanderLambda = this.createCommanderLambda(props.commanderSrcPath);
    this.readModelUpdaterLambda = this.createReadModelUpdatedLambda(
      props.updaterSrcPath,
      props.dynamoStreamArn
    );
  }

  private createLambdaIamRole(): iam.Role {
    return new iam.Role(this, "LambdaIamRole", {
      roleName: "stream-practive-lambda-role",
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
      managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName("CloudwatchFullAccess"),
        iam.ManagedPolicy.fromAwsManagedPolicyName("AmazonDynamoDBFullAccess"),
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaVPCAccessExecutionRole"
        ),
      ],
    });
  }

  private createCommanderLambda(srcPath: string): lambda.Function {
    return new lambda.Function(this, "LambdaCommander", {
      functionName: "commander",
      handler: "main",
      runtime: lambda.Runtime.GO_1_X,
      role: this.lambdaIamRole,
      environment: {
        ENVIRONMENT: "prod",
      },
      code: lambda.Code.fromAsset(path.resolve(__dirname, srcPath), {
        assetHashType: cdk.AssetHashType.OUTPUT,
        bundling: {
          image: lambda.Runtime.GO_1_X.bundlingImage,
          command: [
            "bash",
            "-c",
            "GOOS=linux GOARCH=amd64 go build -o /asset-output/main",
          ],
          user: "root",
        },
      }),
    });
  }

  private createReadModelUpdatedLambda(
    srcPath: string,
    streamArn: string
  ): lambda.Function {
    const lambdaFunc = new lambda.Function(this, "LambdaReadModelUpdater", {
      functionName: "updater",
      handler: "main",
      runtime: lambda.Runtime.GO_1_X,
      role: this.lambdaIamRole,
      code: lambda.Code.fromAsset(path.resolve(__dirname, srcPath), {
        assetHashType: cdk.AssetHashType.OUTPUT,
        bundling: {
          image: lambda.Runtime.GO_1_X.bundlingImage,
          command: [
            "bash",
            "-c",
            "GOOS=linux GOARCH=amd64 go build -o /asset-output/main",
          ],
          user: "root",
        },
      }),
      vpc: this.vpc,
      securityGroups: [this.lambdaSg],
      environment: {
        // TODO: アプリ側でSecrets Managerから取得する。検証目的のため環境変数にベタがき
        DB_HOSTNAME: this.rdsProxy.endpoint,
        DB_NAME: "examples",
        DB_PASSWORD: "sample",
        DB_USER: "admin",
      },
    });
    lambdaFunc.addEventSourceMapping("UpdateReadModelSourceMapping", {
      eventSourceArn: streamArn,
      batchSize: 10,
      startingPosition: lambda.StartingPosition.LATEST,
    });

    return lambdaFunc;
  }
}
