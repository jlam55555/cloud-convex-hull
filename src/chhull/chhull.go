package main

import (
	"bufio"
	"bytes"
	"chutil"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golang/geo/r3"
	"github.com/markus-wa/quickhull-go"
	"log"
	"objio"
	"strings"
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

// QuickHullGoTest tries out the prewritten Go library for quickhull
func QuickHullGoTest(points [][3]float64) ([][3]float64, [][3]int) {
	// convert points to r3.Vector
	pointsR3 := make([]r3.Vector, len(points))
	for i, v := range points {
		pointsR3[i] = r3.Vector{v[0], v[1], v[2]}
	}

	qh := quickhull.QuickHull{}

	ch := qh.ConvexHull(pointsR3, true, false, 0)

	// convert to vertices, faces as needed for objio
	vertices := make([][3]float64, 0)
	faces := make([][3]int, 0)
	for i, triangle := range ch.Triangles() {
		vertices = append(vertices,
			[3]float64{triangle[0].X, triangle[0].Y, triangle[0].Z},
			[3]float64{triangle[1].X, triangle[1].Y, triangle[1].Z},
			[3]float64{triangle[2].X, triangle[2].Y, triangle[2].Z},
		)
		faces = append(faces, [3]int{3*i + 1, 3*i + 2, 3*i + 3})
	}

	return vertices, faces
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

	// make sure key field is present
	if payload.Key == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Missing field \"Key\".",
			StatusCode: 400,
		}, nil
	}

	// read the thing into memory
	req, err := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Key:    aws.String(payload.Key + ".obj"),
		Bucket: aws.String(uploadBucketName),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid model.",
			StatusCode: 400,
		}, nil
	}

	points, err := objio.Parse(bufio.NewReader(req.Body))
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Reading I/O error.",
			StatusCode: 400,
		}, nil
	}
	if err = req.Body.Close(); err != nil {
		panic(err)
	}

	buf := bytes.NewBufferString("")

	// perform quickhull
	vertices, faces := QuickHullGoTest(points)
	if err = objio.Dump(buf, vertices, faces); err != nil {
		panic(err)
	}

	// write file back to new model in s3
	newKey := chutil.GenKey()
	_, err = s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Key:    aws.String(newKey),
		Bucket: aws.String(uploadBucketName),
		Body:   strings.NewReader(buf.String()),
	})
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Writing I/O error.",
			StatusCode: 400,
		}, nil
	}

	jsn, err := json.Marshal(CHResponse{
		Key: newKey,
	})
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsn),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
