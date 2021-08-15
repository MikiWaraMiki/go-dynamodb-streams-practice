import * as cdk from "@aws-cdk/core";
import * as ec2 from "@aws-cdk/aws-ec2";

export class AwsCdkVpcStack extends cdk.Stack {
  public readonly vpc: ec2.Vpc;

  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    this.vpc = this.createVpc();

    this.setupGatewayEndpoint();
    this.setupInterfaceEndpoint();
  }

  private createVpc(): ec2.Vpc {
    return new ec2.Vpc(this, "ExampleVpc", {
      cidr: "172.17.0.0/16",
      natGateways: 0,
      maxAzs: 2,
      subnetConfiguration: [
        { name: "public", cidrMask: 24, subnetType: ec2.SubnetType.PUBLIC },
        { name: "private", cidrMask: 24, subnetType: ec2.SubnetType.ISOLATED },
      ],
    });
  }

  private setupGatewayEndpoint() {
    this.vpc.addGatewayEndpoint("DynamoDBEndPointForReadModelUpdater", {
      service: ec2.GatewayVpcEndpointAwsService.DYNAMODB,
      subnets: [{ subnetType: ec2.SubnetType.ISOLATED }],
    });
  }

  private setupInterfaceEndpoint() {
    this.vpc.addInterfaceEndpoint("SecretsManagerVpcEndpoint", {
      service: ec2.InterfaceVpcEndpointAwsService.SECRETS_MANAGER,
    });
  }
}
