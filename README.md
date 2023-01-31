# Golang AWS Lambda Examples
This repository contains several Golang-based Lambda examples that interact with different AWS services. Each example is self-contained and demonstrates how to perform specific tasks with AWS services.

## Requirements
Before you can run the examples, you need to have the following software installed on your computer:
* Go programming language 1.19

## Getting started
1. Clone the repository to your computer: git clone https://github.com/leorighi/golang-aws-lambda-examples.git
2. Navigate to the example you want to run: cd golang-aws-lambda-examples/leorighi
3. Build the example: go build
4. Deploy and run.
## Examples
### DynamoDB
* **createDynamoClient**: The createDynamoClient is a helper example/function that establishes a connection to the DynamoDB and returns the dynamodb.Client.
* **moveObject**: This example implements an AWS Lambda function in the Go programming language that adds an item to a DynamoDB table. The function specifies the default configuration and credentials, and creates a new client with the dynamodb.NewFromConfig method. The handler function specifies the item to be added and calls the PutItem method on the DynamoDB client. If the item is successfully added, the result is printed to the console. The DynamoConnection function returns a pointer to a dynamodb.Client that can be used to interact with the DynamoDB service
### S3
* **createS3Client**: The createS3Client is a helper example/function that establishes a connection to the S3 service and returns the s3.Client.
* **getObject**: This example implements an AWS Lambda function to download an object from an S3 bucket. The function connects to the S3 service using the AWS SDK v2, specifying the default configuration and credentials, and retrieves the specified object using the GetObject method of the s3.Client.
* **deleteObject**: This example implements an AWS Lambda function to delete an object from an S3 bucket. The function connects to the S3 service using the AWS SDK v2, specifying the default configuration and credentials, and deletes the specified object using the DeleteObject method of the s3.Client.
* **listAllObjects**: This example implements an AWS Lambda function that lists objects in an S3 bucket.The function starts by connecting to the S3 service using the S3Connection method and specifying the default configuration and credentials. The ListObjectsV2 method is then called on the S3 client to list the objects in the specified bucket. The objects keys are then printed.
* **moveObject**: This example implements an AWS Lambda function that copies an object from one S3 bucket to another. The function starts by connecting to the S3 service using the S3Connection method and specifying the default configuration and credentials. The CopyObject method is then called on the S3 client to copy the specified object to a new key in the same or a different bucket. The result of the copy operation is printed indicating that the object was copied is displayed.
