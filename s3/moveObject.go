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
	_, err := svc.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String("my-bucket"),
		CopySource: aws.String("mybucket/file.txt"),
		Key:        aws.String("copyFile.txt"),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Object copied.")
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
