# Serverless Golang Boilerplate

This project is a serverless application built with AWS CDK and Golang. It includes Lambda functions, API Gateway, and Cognito for user authentication.

## Features

- **AWS Lambda**: Serverless functions written in Go.
- **API Gateway**: HTTP APIs to expose Lambda functions.
- **Cognito**: User authentication and management.
- **AWS CDK**: Infrastructure as code for deploying resources.

## Prerequisites

1. **AWS CLI**: Install and configure the AWS CLI with a valid profile.
2. **Node.js**: Required for AWS CDK.
3. **Go**: Install Go (version 1.23 or later).
4. **Make**: Ensure `make` is available on your system.

## Project Structure

- `lambda/`: Contains the Lambda function code.
- `cdk/`: CDK application for defining AWS resources.
- `constants/env.go`: Environment variables used in the project.
- `.github/workflows/deploy.yml`: GitHub Actions workflow for CI/CD.

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/pedroSoaresll/serverless-golang-boilerplate.git
   cd serverless-golang-boilerplate
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   npm install -g aws-cdk
   ```

3. Set up environment variables:
   Define the following variables in your environment or `.env` file:
   - `ENV`: Deployment environment (e.g., `dev`, `prod`).
   - `CDK_DEFAULT_ACCOUNT`: AWS account ID.
   - `CDK_DEFAULT_REGION`: AWS region.

## Useful Commands

### Build and Deploy

- **Build Lambda Functions**:

  ```bash
  make build-lambda
  ```

- **Bootstrap CDK**:

  ```bash
  make cdk-bootstrap
  ```

- **Deploy Resources**:

  ```bash
  make cdk-deploy
  ```

- **Destroy Resources**:

  ```bash
  make cdk-destroy
  ```

### CI/CD

- **Validate Stack**:

  ```bash
  make ci-synth
  ```

- **Deploy in CI**:

  ```bash
  make ci-deploy
  ```

## Deployment Workflow

1. **Development**: Push changes to a feature branch. A pull request triggers deployment to the `dev` environment.
2. **Production**: Merging to the `main` branch triggers deployment to the `prod` environment.

## API Endpoints

- **GET /hello**: Returns a greeting message.
- **GET /hello/{name}**: Returns a personalized greeting.

## Cleanup

To remove all resources:

```bash
make cdk-destroy
```

## Contributing

Feel free to submit issues or pull requests to improve this project.

## License

This project is licensed under the MIT License.
