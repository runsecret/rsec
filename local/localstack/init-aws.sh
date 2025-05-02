#!/bin/bash

# Setup AWS Secret Manager secrets locally for integration testing
# This script should be placed in ready.d directory to run at startup

set -e

echo "Setting up mock AWS Secret Manager secrets for local testing..."

# Setup local endpoint
LOCAL_ENDPOINT="http://localhost:4566"

# Configure AWS CLI for local use
export AWS_ACCESS_KEY_ID="test"
export AWS_SECRET_ACCESS_KEY="test"
export AWS_DEFAULT_REGION="us-east-1"

# Check if localstack is running
if ! curl -s $LOCAL_ENDPOINT > /dev/null; then
    echo "LocalStack endpoint not available. Please ensure LocalStack is running."
    exit 1
fi

# Create test secrets
create_secret() {
    local name=$1
    local value=$2

    echo "Creating/updating secret: $name"

    # Check if secret exists (silently)
    if aws --endpoint-url=$LOCAL_ENDPOINT secretsmanager describe-secret --secret-id "$name" --region us-east-1 --output json &> /dev/null; then
        # Secret exists, update it silently
        aws --endpoint-url=$LOCAL_ENDPOINT secretsmanager update-secret \
            --secret-id "$name" \
            --secret-string "$value" \
            --region us-east-1 \
            --output json > /dev/null
    else
        # Secret doesn't exist, create it silently
        aws --endpoint-url=$LOCAL_ENDPOINT secretsmanager create-secret \
            --name "$name" \
            --secret-string "$value" \
            --region us-east-1 \
            --output json > /dev/null
    fi
}

# Define test secrets
create_secret "basicPassword" 'password1234'
create_secret "test/database/credentials" '{"username":"testuser","password":"testpassword","host":"localhost","port":5432,"dbname":"testdb"}'
create_secret "test/api/keys" '{"api_key":"sk-test-12345","api_secret":"abcdef123456"}'
create_secret "test/oauth/client" '{"client_id":"test-client-id","client_secret":"test-client-secret"}'
create_secret "test/encryption/keys" '{"primary":"AES256Key-32Characters1234567890","secondary":"AES256Key-32Characters0987654321"}'
create_secret "test/service/endpoints" '{"service1":"http://localhost:8081","service2":"http://localhost:8082"}'

echo "AWS Secret Manager test secrets have been successfully created!"
