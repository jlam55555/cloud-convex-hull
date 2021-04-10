package main

// use v2 of the aws go sdk
import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

// MyEvent is the input json for the lambda
type MyEvent struct {
	Type string `json:"type"`
	Key  string `json:"key"`
}

// these are specified at linker-time
// ref: https://stackoverflow.com/a/28460195/2397327
var awsRegion, uploadBucketName string

var s3Client *s3.Client
var psClient *s3.PresignClient

// init is run once on lambda startup
func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}

	s3Client = s3.NewFromConfig(cfg)
	psClient = s3.NewPresignClient(s3Client)
}

// HandleRequest presigns an S3 URL for either a PUT or GET request
func HandleRequest(ctx context.Context, payload MyEvent) (string, error) {
	if payload.Type == "PUT" {
		input := &s3.PutObjectInput{
			Bucket: aws.String(uploadBucketName),
			Key:    aws.String(payload.Key),
		}

		psRequest, err :=
			psClient.PresignPutObject(context.TODO(), input)
		if err != nil {
			return "", err
		}

		return psRequest.URL, nil
	} else {
		input := &s3.GetObjectInput{
			Bucket: aws.String(uploadBucketName),
			Key:    aws.String(payload.Key),
		}

		psRequest, err :=
			psClient.PresignGetObject(context.TODO(), input)
		if err != nil {
			return "", err
		}

		return psRequest.URL, nil
	}
}

// main starts
func main() {
	lambda.Start(HandleRequest)
}
