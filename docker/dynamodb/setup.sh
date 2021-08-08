#!/bin/bash

PROFILE=$1
# Create EventStore
aws dynamodb create-table --endpoint-url http://localhost:8000 \
--table-name event-store \
--attribute-definitions AttributeName=eventProviderId,AttributeType=S \
--key-schema AttributeName=eventProviderId,KeyType=HASH \
--provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
--stream-specification StreamEnabled=true,StreamViewType=NEW_IMAGE \
--profile $PROFILE

# Create ProviderStore
aws dynamodb create-table --endpoint-url http://localhost:8000 \
--table-name provider-store \
--attribute-definitions AttributeName=eventProviderId,AttributeType=S \
--key-schema AttributeName=eventProviderId,KeyType=HASH \
--provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
--profile $PROFILE
