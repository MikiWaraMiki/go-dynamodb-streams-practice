import * as cdk from "@aws-cdk/core";
import * as ec2 from "@aws-cdk/aws-ec2";

interface SgStackProps extends cdk.StackProps {
  vpc: ec2.Vpc;
}

export class AwsCdkSgStack extends cdk.Stack {
  public readonly rdsSg: ec2.SecurityGroup;
  public readonly lambdaSg: ec2.SecurityGroup;
  public readonly bastionSg: ec2.SecurityGroup;
  public readonly rdsProxySg: ec2.SecurityGroup;
  private readonly vpc: ec2.Vpc;

  constructor(scope: cdk.Construct, id: string, props: SgStackProps) {
    super(scope, id, props);

    this.vpc = props.vpc;
    this.bastionSg = this.createBastionSecurityGroup();
    this.lambdaSg = this.createLambdaSecurityGroup();
    this.rdsProxySg = this.createRdsProxySecurityGroup();
    this.rdsSg = this.createRdsSecurityGroup();
  }

  private createBastionSecurityGroup(): ec2.SecurityGroup {
    return new ec2.SecurityGroup(this, "Bastion", {
      vpc: this.vpc,
      description: "For Bastion",
    });
  }

  private createLambdaSecurityGroup(): ec2.SecurityGroup {
    return new ec2.SecurityGroup(this, "LambdaToRdsProxy", {
      vpc: this.vpc,
      description: "For Lambda To Proxy",
    });
  }

  private createRdsProxySecurityGroup(): ec2.SecurityGroup {
    const sg = new ec2.SecurityGroup(this, "Rds Proxy Security Group", {
      vpc: this.vpc,
      description: "Rds Proxy Sg",
    });
    sg.addIngressRule(
      this.lambdaSg,
      ec2.Port.tcp(3306),
      "allow db connection from lambda"
    );

    return sg;
  }

  private createRdsSecurityGroup(): ec2.SecurityGroup {
    const sg = new ec2.SecurityGroup(this, "Proxy to DB Connection", {
      vpc: this.vpc,
    });
    sg.addIngressRule(
      this.rdsProxySg,
      ec2.Port.tcp(3306),
      "allow db connection from rds proxy"
    );
    sg.addIngressRule(
      this.bastionSg,
      ec2.Port.tcp(3306),
      "allow mysql from bastion"
    );
    return sg;
  }
}
