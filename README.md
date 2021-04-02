# cfn-ssh-key-import
Custom CloudFormation resource to import an SSH public key to the AWS EC2 service.

## Requirements
- [SAM CLI](https://aws.amazon.com/serverless/sam/) (built with 1.21.1)
- [Go](https://golang.org/) (built with 1.15.8)

## Usage
The custom resource is very simple to deploy and use.

1. Build and deploy the SAM application:
   ```bash
   sam build
   sam deploy --guided
   ```
   For more information about the resources deployed as part of this process, see the [SAM template](./template.yaml).
2. Call the custom resource in your CloudFormation template with the public key you wish to import - see the [example template](./example/template.yaml).