import * as cdk from "@aws-cdk/core";
import * as ec2 from "@aws-cdk/aws-ec2";
import * as rds from "@aws-cdk/aws-rds";
import * as secret from "@aws-cdk/aws-secretsmanager";

interface RdsStackProps extends cdk.StackProps {
  vpc: ec2.Vpc;
  rdsSg: ec2.SecurityGroup;
  rdsProxySg: ec2.SecurityGroup;
}

export class AwsCdkRdsStack extends cdk.Stack {
  private readonly vpc: ec2.Vpc;
  private readonly rdsSg: ec2.SecurityGroup;
  private readonly rdsProxySg: ec2.SecurityGroup;

  public readonly secret: secret.Secret;
  public readonly rdsInstance: rds.DatabaseInstance;
  public readonly rdsProxy: rds.DatabaseProxy;

  constructor(scope: cdk.Construct, id: string, props: RdsStackProps) {
    super(scope, id, props);

    this.vpc = props.vpc;
    this.rdsSg = this.createRdsSecurityGroup();
    //this.rdsSg = props.rdsSg;
    this.rdsProxySg = props.rdsProxySg;

    this.secret = this.generateSecrets();
    this.rdsInstance = this.createRdsInstance();
    this.rdsProxy = this.createRdsProxy();
  }

  private createRdsSecurityGroup(): ec2.SecurityGroup {
    const sg = new ec2.SecurityGroup(this, "Proxy to DB Connection", {
      vpc: this.vpc,
    });
    return sg;
  }

  private generateSecrets(): secret.Secret {
    return new secret.Secret(this, "DBCredentialsSecret", {
      secretName: "stream-rds-credentials",
      generateSecretString: {
        secretStringTemplate: JSON.stringify({
          username: "admin",
        }),
        excludePunctuation: true,
        includeSpace: false,
        generateStringKey: "password",
      },
    });
  }
  private createRdsInstance(): rds.DatabaseInstance {
    const rdsInstance = new rds.DatabaseInstance(this, "DBInstance", {
      engine: rds.DatabaseInstanceEngine.mysql({
        version: rds.MysqlEngineVersion.VER_5_7_33,
      }),
      credentials: rds.Credentials.fromSecret(this.secret),
      instanceType: ec2.InstanceType.of(
        ec2.InstanceClass.T3,
        ec2.InstanceSize.SMALL
      ),
      port: 3306,
      vpc: this.vpc,
      vpcSubnets: {
        subnetType: ec2.SubnetType.ISOLATED,
      },
      multiAz: false,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      deletionProtection: false,
      parameterGroup: new rds.ParameterGroup(this, "ParameterGroup", {
        engine: rds.DatabaseInstanceEngine.mysql({
          version: rds.MysqlEngineVersion.VER_5_7_31,
        }),
        parameters: {
          character_set_client: "utf8mb4",
          character_set_server: "utf8mb4",
        },
      }),
    });

    return rdsInstance;
  }

  private createRdsProxy(): rds.DatabaseProxy {
    return this.rdsInstance.addProxy("stream-rds-proxy", {
      secrets: [this.secret],
      debugLogging: true,
      vpc: this.vpc,
      securityGroups: [this.rdsProxySg],
      requireTLS: false,
    });
  }
}
