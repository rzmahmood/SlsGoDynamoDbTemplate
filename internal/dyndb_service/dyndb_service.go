package dyndb_service

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"
)

type DynDBService struct {
	dbSession *session.Session
	dbClient *dynamodb.DynamoDB
	tableName string
}

func CreateDynamoDbService() (*DynDBService, error){
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, err
	}
	dynDbClient := dynamodb.New(sess)
	if dynDbClient == nil {
		return nil, errors.New("failed to initialize dynamodb session")
	}

	dbTableName:= os.Getenv("DYNAMODB_TABLE")
	log.Printf("Using table name %v\n", dbTableName)

	return &DynDBService{
		dbSession: sess,
		dbClient:  dynDbClient,
		tableName: dbTableName,

	}, nil
}

func (d *DynDBService) StoreAddress(addr string) bool {
	ethAddr := AddressItem{EthAddress: addr}
	dbItem, err := dynamodbattribute.MarshalMap(ethAddr)
	if err != nil {
		log.Println("Failed to MarshalMap the Ethereum Address in StoreAddress")
		return false
	}
	input := dynamodb.PutItemInput{
		Item: dbItem,
		TableName: aws.String(d.tableName),
	}
	_, err = d.dbClient.PutItem(&input)
	if err != nil {
		fmt.Printf("Got error calling PutItem: %v\n", err.Error())
		return false
	}
	return true
}

func (d *DynDBService) GetAddresses() {

}



