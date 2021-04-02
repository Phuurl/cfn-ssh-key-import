package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func importKeyPair(ec2Svc *ec2.EC2, name string, material string) (err error) {
	importInput := &ec2.ImportKeyPairInput{
		KeyName: aws.String(name),
		PublicKeyMaterial: []byte(material),
	}
	_, err = ec2Svc.ImportKeyPair(importInput)
	handleError(err)

	return
}

func deleteKeyPair(ec2Svc *ec2.EC2, name string) (err error) {
	deleteInput := &ec2.DeleteKeyPairInput{
		KeyName: aws.String(name),
	}
	_, err = ec2Svc.DeleteKeyPair(deleteInput)
	handleError(err)

	return
}

func handler(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	keyName, _ := event.ResourceProperties["KeyName"].(string)
	keyMaterial, _ := event.ResourceProperties["KeyMaterial"].(string)

	sess, _ := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	ec2Svc := ec2.New(sess)

	if string(event.RequestType) == "Create" {
		err = importKeyPair(ec2Svc, keyName, keyMaterial)
	} else if string(event.RequestType) == "Update" {
		_ = deleteKeyPair(ec2Svc, keyName)
		err = importKeyPair(ec2Svc, keyName, keyMaterial)
	} else if string(event.RequestType) == "Delete" {
		err = deleteKeyPair(ec2Svc, keyName)
	} else {
		err = fmt.Errorf("unknown RequestType %s", string(event.RequestType))
		handleError(err)
	}

	if err == nil {
		physicalResourceID = keyName
	}
	return
}

func main() {
	lambda.Start(cfn.LambdaWrap(handler))
}
