package main

import (
	"context"
	"github.com/Bchain/serverless-go-crud/internal/api_service"
	"github.com/Bchain/serverless-go-crud/internal/dyndb_service"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var dbService *dyndb_service.DynDBService

func init() {
	var err error
	dbService, err = dyndb_service.CreateDynamoDbService()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Started dynDbClient %v\n", dbService)
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return api_service.SuccessResponse("you have called the db list service"), nil
}

func main() {
	lambda.Start(Handler)
}
