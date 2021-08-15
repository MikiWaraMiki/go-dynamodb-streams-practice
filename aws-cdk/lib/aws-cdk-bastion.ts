import * as cdk from "@aws-cdk/core";
import * as ec2 from "@aws-cdk/aws-ec2";

interface BastionStackProps extends cdk.StackProps {
  vpc: ec2.Vpc;
  bastionSg: ec2.SecurityGroup;
}

export class AwsCdkBastionStack extends cdk.Stack {
  public readonly ec2: ec2.BastionHostLinux;
  private readonly vpc: ec2.Vpc;
  private readonly bastionSg: ec2.SecurityGroup;

  constructor(scope: cdk.Construct, id: string, props: BastionStackProps) {
    super(scope, id, props);

    this.vpc = props.vpc;
    this.bastionSg = props.bastionSg;

    this.ec2 = this.createBastionHost();
  }

  private createBastionHost(): ec2.BastionHostLinux {
    const host = new ec2.BastionHostLinux(this, "BastionHost", {
      vpc: this.vpc,
      instanceType: ec2.InstanceType.of(
        ec2.InstanceClass.T3,
        ec2.InstanceSize.MICRO
      ),
      securityGroup: this.bastionSg,
      subnetSelection: {
        subnetType: ec2.SubnetType.PUBLIC,
      },
    });
    host.instance.addUserData("yum -y update", "yum install -y mysql jq htop");
    return host;
  }
}
