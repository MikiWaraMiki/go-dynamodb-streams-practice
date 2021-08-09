import * as cdk from "@aws-cdk/core";
import * as dynamodb from "@aws-cdk/aws-dynamodb";
import * as iam from "@aws-cdk/aws-iam";
import * as lambda from "@aws-cdk/aws-lambda";
import { Code, Runtime } from "@aws-cdk/aws-lambda";
import {
  AssetHashType,
  Duration,
} from "@aws-cdk/aws-cloudwatch/node_modules/@aws-cdk/core";
import { resolve } from "path";

export class AwsCdkStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // DynamoDB
    const eventStoreTable = new dynamodb.Table(this, "EventStoreTable", {
      tableName: "event-store",
      partitionKey: {
        name: "eventProviderId",
        type: dynamodb.AttributeType.STRING,
      },
      billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      stream: dynamodb.StreamViewType.NEW_AND_OLD_IMAGES,
    });

    const providerStoreTable = new dynamodb.Table(this, "ProviderStoreTable", {
      tableName: "provider-store",
      partitionKey: {
        name: "eventProviderId",
        type: dynamodb.AttributeType.STRING,
      },
      billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
    });

    // Iam
    const lambdaRole = new iam.Role(this, "LambdaIamRole", {
      roleName: "stream-practice-lambda-role",
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
      managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName("CloudWatchLogsFullAccess"),
        iam.ManagedPolicy.fromAwsManagedPolicyName("AmazonDynamoDBFullAccess"),
      ],
    });

    // Lambda
    const commanderFunction = new lambda.Function(this, "CommanderFunction", {
      functionName: "commander",
      handler: "main",
      runtime: lambda.Runtime.GO_1_X,
      role: lambdaRole,
      environment: {
        ENVIRONMENT: "prod",
      },
      timeout: Duration.seconds(30),
      code: Code.fromAsset(resolve(__dirname, "../../src/commander/"), {
        assetHashType: AssetHashType.OUTPUT,
        bundling: {
          image: Runtime.GO_1_X.bundlingImage,
          command: [
            "bash",
            "-c",
            "GOOS=linux GOARCH=amd64 go build -o /asset-output/main",
          ],
          user: "root",
        },
      }),
    });
    const readModelUpdaterFunction = new lambda.Function(
      this,
      "ReadModelUpdaterFunction",
      {
        functionName: "readmodel-updater",
        handler: "main",
        runtime: lambda.Runtime.GO_1_X,
        role: lambdaRole,
        environment: {
          // NOTE: YOU NEED CHANGE
          DB_USER: "root",
          DB_PASSWORD: "password",
          HOSTNAME: "localhost",
          DB_NAME: "examples",
          PORT: "3306",
        },
        timeout: Duration.seconds(60),
        code: Code.fromAsset(
          resolve(__dirname, "../../src/readmodel_updater/"),
          {
            assetHashType: AssetHashType.OUTPUT,
            bundling: {
              image: Runtime.GO_1_X.bundlingImage,
              command: [
                "bash",
                "-c",
                "GOOS=linux GOARCH=amd64 go build -o /asset-output/main",
              ],
              user: "root",
            },
          }
        ),
      }
    );
    readModelUpdaterFunction.addEventSourceMapping(
      "UpdateReadModelSourceMapping",
      {
        eventSourceArn: eventStoreTable.tableStreamArn,
        batchSize: 10,
        startingPosition: lambda.StartingPosition.LATEST,
      }
    );
  }
}
