#!/usr/bin/env node
import "source-map-support/register";
import * as cdk from "@aws-cdk/core";
import { AwsCdkStack } from "../lib/aws-cdk-stack";
import { AwsCdkVpcStack } from "../lib/aws-cdk-vpc";
import { AwsCdkBastionStack } from "../lib/aws-cdk-bastion";
import { AwsCdkSgStack } from "../lib/aws-cdk-sg";
import { AwsCdkRdsStack } from "../lib/aws-cdk-rds";
import { AwsCdkLambdaStack } from "../lib/aws-cdk-lambda";
import { AwsCdkDynamoDBStack } from "../lib/aws-cdk-dynamodb";

const app = new cdk.App();

const vpcStack = new AwsCdkVpcStack(app, "DynamoVpcStack");

const sgStack = new AwsCdkSgStack(app, "DynamoSecurityGroupStack", {
  vpc: vpcStack.vpc,
});

const bastionStack = new AwsCdkBastionStack(app, "DynamoBastionStack", {
  vpc: vpcStack.vpc,
  bastionSg: sgStack.bastionSg,
});

const rdsStack = new AwsCdkRdsStack(app, "DynamoRdsStack", {
  vpc: vpcStack.vpc,
  rdsSg: sgStack.rdsSg,
  rdsProxySg: sgStack.rdsProxySg,
});

const dynamoDBStack = new AwsCdkDynamoDBStack(app, "DynamoDBStack");

const lambdaStack = new AwsCdkLambdaStack(app, "DynamoLambdaStack", {
  commanderSrcPath: "../../src/commander/",
  updaterSrcPath: "../../src/readmodel_updater/",
  vpc: vpcStack.vpc,
  lambdaSg: sgStack.lambdaSg,
  rdsProxy: rdsStack.rdsProxy,
  dynamoStreamArn: dynamoDBStack.eventStoreTable.tableStreamArn ?? "",
});
