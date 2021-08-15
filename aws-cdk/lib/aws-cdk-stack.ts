import * as cdk from "@aws-cdk/core";
import { AwsCdkVpcStack } from "./aws-cdk-vpc";

export class AwsCdkStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const vpcStack = new AwsCdkVpcStack(scope, id, props);
  }
}
