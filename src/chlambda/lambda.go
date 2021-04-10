package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"time"
)

type MyEvent struct {
	Name string `json:"name"`
}

// these are specified at linker-time
// ref: https://stackoverflow.com/a/28460195/2397327
var awsRegion, uploadBucketName string
var svc *s3.S3

// setup commands
func init() {
	// create aws session
	sess, err := session.NewSession()
	if err != nil {
		log.Fatalln(err)
	}

	// create s3 service client
	svc = s3.New(sess)
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(uploadBucketName),
		Key:    aws.String("TestKey"),
	})

	urlStr, err := req.Presign(5 * time.Minute)
	if err != nil {
		//log.Fatalln(err)
		return "", err
	}

	return urlStr, nil
	//return fmt.Sprintf("Hello, %s!", name.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
