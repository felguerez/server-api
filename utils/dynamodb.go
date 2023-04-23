package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"os"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

type Item struct {
	ExpiresAt    int64  `json:"expiresAt,omitempty"`
	AccessToken  string `json:"accessToken" json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken" json:"refreshToken,omitempty"`
	TokenType    string `json:"tokenType" json:"tokenType,omitempty"`
	Id           string `json:"id,omitempty"`
}

var TableName = aws.String(os.Getenv("TABLE_NAME"))

// Initialize a session that the SDK will use to load
// credentials from the shared credentials file ~/.aws/credentials
// and region from the shared configuration file ~/.aws/config.
var sess = session.Must(session.NewSessionWithOptions(session.Options{
	SharedConfigState: session.SharedConfigEnable,
}))

var svc = dynamodb.New(sess)

// PutItem godoc
// @Summary Adds a Spotify accessToken to the db
// @Description Writes an accessToken and expiresAt to the db for user id "felguerez" (embedded in item.Id)
// @Tags spotify
func PutItem(item Item) map[string]*dynamodb.AttributeValue {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: TableName,
	}
	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
	fmt.Println("Successfully added item to db:" + string(item.AccessToken))
	return av
}

func GetItem(key string) (*Item, error) {
	input := &dynamodb.GetItemInput{
		TableName: TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(key),
			},
		},
	}
	result, err := svc.GetItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var item Item
	if err := dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		return nil, err
	}
	return &item, nil
}
