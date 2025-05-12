# PDF Generator Lambda Function

A Go-based AWS Lambda function that generates PDF shipping labels based on provided input data.

## Project Overview

This project is an AWS Lambda function written in Go that generates PDF shipping labels. It uses AWS Serverless Application Model (SAM) for local development and deployment.

## Prerequisites

- [Go](https://golang.org/doc/install) (version 1.21.4 or later)
- [AWS CLI](https://aws.amazon.com/cli/) (configured with appropriate credentials)
- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
- [Docker](https://docs.docker.com/get-docker/) (for local testing)

## Project Structure

```
.
├── pdf-generator/          # Main application code
│   ├── cmd/                # Lambda handler entry point
│   │   └── main.go         # Main Lambda function code
│   ├── pkg/                # Application packages
│   │   ├── handlers/       # HTTP request handlers
│   │   ├── helpers/        # PDF generation utilities
│   │   └── types/          # Type definitions
│   ├── go.mod              # Go module definition
│   └── go.sum              # Go module checksums
├── events/                 # Test event payloads
│   └── event.json          # Sample API Gateway event
├── layers/                 # Lambda layer resources
│   └── images/             # Images for PDF generation
├── template.yaml          # SAM template
├── samconfig.toml         # SAM configuration
└── Makefile               # Build automation
```

## Local Development

### Building the Project

To build the project locally:

```bash
make build
# or
sam build
```

### Testing Locally

To test the lambda function locally with the sample event:

```bash
sam local invoke PDFGeneratorFunction -e events/event.json
```

## Deployment

### Manual Deployment

To deploy the Lambda function to AWS:

1. Navigate to the `pdf-generator/cmd` directory:

```bash
cd pdf-generator/cmd
```

2. Build the Lambda function binary:

For x86_64 (amd64) architecture:
```bash
GOARCH=amd64 GOOS=linux go build -o bootstrap main.go
```

For ARM64 architecture:
```bash
GOARCH=arm64 GOOS=linux go build -o bootstrap main.go
```

3. Create a deployment package by zipping the binary:

```bash
zip function.zip bootstrap
```

4. Deploy the function using AWS CLI:

```bash
aws lambda update-function-code \
    --function-name <your-function-name> \
    --zip-file fileb://function.zip
```

### SAM Deployment

To deploy using SAM:

1. Build the application:

```bash
sam build
```

2. Deploy the application:

```bash
sam deploy
```

For interactive deployment with prompts:

```bash
sam deploy --guided
```

## AWS Lambda Configuration

The Lambda function is defined in `template.yaml` with the following configuration:

- Runtime: go1.x
- Architecture: x86_64
- Memory: 256MB
- Timeout: 29 seconds
- Custom Layer: Contains resources needed for PDF generation

## Improving CI/CD

Here are recommendations for improving the CI/CD pipeline:

### 1. GitHub Actions Workflow

Create a `.github/workflows/deploy.yml` file:

```yaml
name: Deploy Lambda

on:
  push:
    branches: [ main ]
  workflow_dispatch:

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
          
      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v2
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: ${{ secrets.AWS_REGION }}
          
      - name: Build binary
        run: |
          cd pdf-generator/cmd
          GOARCH=amd64 GOOS=linux go build -o bootstrap main.go
          zip function.zip bootstrap
          
      - name: Deploy Lambda
        run: |
          aws lambda update-function-code \
            --function-name ${{ secrets.LAMBDA_FUNCTION_NAME }} \
            --zip-file fileb://pdf-generator/cmd/function.zip
```

### 2. AWS CodePipeline Setup

1. Create a CodeCommit repository or connect GitHub repository to AWS CodePipeline
2. Set up CodeBuild with a `buildspec.yml`:

```yaml
version: 0.2

phases:
  install:
    runtime-versions:
      golang: 1.21
  build:
    commands:
      - cd pdf-generator/cmd
      - GOARCH=amd64 GOOS=linux go build -o bootstrap main.go
      - zip function.zip bootstrap
  post_build:
    commands:
      - aws lambda update-function-code --function-name $LAMBDA_FUNCTION_NAME --zip-file fileb://function.zip

artifacts:
  files:
    - pdf-generator/cmd/function.zip
```

### 3. Infrastructure as Code

Consider using AWS CDK, Terraform, or updating your SAM template to manage all infrastructure components:

```bash
# Install and set up AWS CDK
npm install -g aws-cdk
cdk init app --language typescript

# Deploy with CDK
cdk deploy
```

### 4. Automated Testing

Add unit and integration tests for your Lambda function:

```bash
# Run tests before deployment
go test ./...

# Add test step to CI/CD pipeline
```

## Best Practices for Lambda Updates

1. **Use Version Control**: Always commit changes before deployment.

2. **Implement Versioning**: Use Lambda versioning and aliases:
   ```bash
   aws lambda publish-version --function-name <function-name>
   aws lambda create-alias --function-name <function-name> --name prod --function-version <version>
   ```

3. **Staged Deployments**: Deploy to staging environment first:
   ```bash
   aws lambda update-alias --function-name <function-name> --name staging --function-version <version>
   ```

4. **Monitoring**: Set up CloudWatch alarms and X-Ray tracing.

5. **Rollback Plan**: Keep previous versions available for rollback.

## Troubleshooting

- **Deployment Issues**: Check CloudFormation stack events in the AWS Console.
- **Runtime Errors**: Check CloudWatch logs for Lambda execution logs.
- **Local Testing**: Use `sam logs` to retrieve logs from deployed Lambda. 