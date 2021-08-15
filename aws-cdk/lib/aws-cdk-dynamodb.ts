import * as cdk from "@aws-cdk/core";
import * as dynamodb from "@aws-cdk/aws-dynamodb";

export class AwsCdkDynamoDBStack extends cdk.Stack {
  public readonly eventStoreTable: dynamodb.Table;
  public readonly eventProviderTable: dynamodb.Table;

  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    this.eventStoreTable = this.createEventStoreTable();
    this.eventProviderTable = this.createEventProviderTable();
  }

  private createEventStoreTable(): dynamodb.Table {
    return new dynamodb.Table(this, "EventStoreTable", {
      tableName: "event-store",
      partitionKey: {
        name: "eventProviderId",
        type: dynamodb.AttributeType.STRING,
      },
      billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
      stream: dynamodb.StreamViewType.NEW_AND_OLD_IMAGES,
    });
  }

  private createEventProviderTable(): dynamodb.Table {
    return new dynamodb.Table(this, "ProviderStoreTable", {
      tableName: "provider-store",
      partitionKey: {
        name: "eventProviderId",
        type: dynamodb.AttributeType.STRING,
      },
      billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
    });
  }
}
