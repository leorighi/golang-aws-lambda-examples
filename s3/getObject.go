package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func handler() {
	svc, _ := S3Connection()
	result, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("myBucket"),
		Key:    aws.String("file.txt"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Download complete.")
}

func main() {
	lambda.Start(handler)
}

func S3Connection() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return &s3.Client{}, err
	}
	svc := s3.NewFromConfig(cfg)
	return svc, nil
}
