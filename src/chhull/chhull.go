package main

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

// CHEvent is the input json for the lambda
type CHEvent struct {
	Key string `json:"key"`
}

// CHResponse is the output for the lambda
type CHResponse struct {
	Key string `json:"key"`
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
	var payload CHEvent
	err := json.Unmarshal([]byte(event.Body), &payload)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid JSON.",
			StatusCode: 400,
		}, nil
	}

	// TODO: working here
	// read the thing into memory
	req, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Key:    aws.String(payload.Key),
		Bucket: aws.String(uploadBucketName),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid model.",
			StatusCode: 400,
		}, nil
	}

	reader := bufio.NewReader(req.Body)

	firstLine, err := reader.ReadString('\n')
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error reading file from S3.",
			StatusCode: 400,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       firstLine,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
