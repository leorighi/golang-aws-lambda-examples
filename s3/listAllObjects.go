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
	result, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String("myBucket"),
	})
	if err != nil {
		panic(err)
	}

	for _, object := range result.Contents {
		fmt.Println(*object.Key)
	}
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
