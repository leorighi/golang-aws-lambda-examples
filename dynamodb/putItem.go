package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func handler() {
	svc, _ := DynamoConnection()
	input := &dynamodb.PutItemInput{
		Item: map[string]dynamodb.AttributeValue{
			"UserId": {
				S: aws.String("123456"),
			},
			"UserName": {
				S: aws.String("SydBarrett"),
			},
		},
		TableName: aws.String("Users"),
	}

	result, err := svc.PutItem(context.Background(), input)
	if err != nil {
		fmt.Println("Error putting item: ", err)
		return
	}

	fmt.Println("Item added: ", result)
}

func main() {
	lambda.Start(handler)
}

func DynamoConnection() (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}
	svc := dynamodb.NewFromConfig(cfg)
	return svc, nil
}
