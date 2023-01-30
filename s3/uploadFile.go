package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func handler() {
	svc, _ := S3Connection()
	file, err := os.Open("/path/to/file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("myBucket"),
		Key:    aws.String("file.txt"),
		Body:   file,
	})
	if err != nil {
		panic(err)
	}
}

func main() {
	lambda.Start(handler)
}

func S3Connection() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}
	svc := s3.NewFromConfig(cfg)
	return svc, nil
}
