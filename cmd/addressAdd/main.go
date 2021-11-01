package main

import (
	"context"
	"fmt"
	"github.com/Bchain/serverless-go-crud/internal/api_service"
	"github.com/Bchain/serverless-go-crud/internal/dyndb_service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
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
	ethAddress, ok := request.QueryStringParameters["ethAddress"]
	if !ok {
		return api_service.FailureResponse(http.StatusBadRequest, "failed due to missing ethAddress"), nil
	}
	log.Printf("Received request with ethAddress %v\n", ethAddress)

	ok = dbService.StoreAddress(ethAddress)
	if !ok {
		return api_service.FailureResponse(http.StatusBadRequest, "failed to store address"), nil
	}

	return api_service.SuccessResponse(fmt.Sprintf("successfully stored ethAddress %v", ethAddress)), nil
}

func main() {
	lambda.Start(Handler)
}
