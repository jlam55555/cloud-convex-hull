package main

// use v2 of the aws go sdk
import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"strings"
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

func getPresignedPutUrl(event events.APIGatewayProxyRequest, payload MyEvent) (
	events.APIGatewayProxyResponse, error) {

	input := &s3.PutObjectInput{
		Bucket: aws.String(uploadBucketName),
		Key:    aws.String(payload.Key),
	}

	psRequest, err :=
		psClient.PresignPutObject(context.TODO(), input)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       psRequest.URL,
		StatusCode: 200,
	}, nil
}

func getPresignedGetUrl(event events.APIGatewayProxyRequest, payload MyEvent) (
	events.APIGatewayProxyResponse, error) {

	input := &s3.GetObjectInput{
		Bucket: aws.String(uploadBucketName),
		Key:    aws.String(payload.Key),
	}

	psRequest, err :=
		psClient.PresignGetObject(context.TODO(), input)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       psRequest.URL,
		StatusCode: 200,
	}, nil
}

// HandleRequest presigns an S3 URL for either a PUT or GET request
func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (
	events.APIGatewayProxyResponse, error) {

	// make sure content-type is set to application/json
	// (any text type is fine, but can cause problems if no type specified)
	if event.Headers["content-type"] != "application/json" {
		return events.APIGatewayProxyResponse{
			Body:       "Request content-type header must be application/json.",
			StatusCode: 400,
		}, nil
	}

	// make sure input json is valid
	var payload MyEvent
	err := json.Unmarshal([]byte(event.Body), &payload)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid JSON.",
			StatusCode: 400,
		}, nil
	}

	// make sure key is specified
	if payload.Key == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Input field 'key' must not be empty.",
			StatusCode: 400,
		}, nil
	}

	// make sure type is specified
	payload.Type = strings.ToUpper(payload.Type)
	if payload.Type != "PUT" && payload.Type != "GET" {
		return events.APIGatewayProxyResponse{
			Body:       "Input field 'type' must be 'PUT' or 'GET'.",
			StatusCode: 400,
		}, nil
	}

	if payload.Type == "PUT" {
		return getPresignedPutUrl(event, payload)
	}
	return getPresignedGetUrl(event, payload)
}

func main() {
	lambda.Start(HandleRequest)
}
