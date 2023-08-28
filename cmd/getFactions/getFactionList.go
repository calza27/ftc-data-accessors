package main

import (
	"context"
	"fireteam-core/internal/constants"
	"fireteam-core/internal/helpers"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(getFactionListHandler)
}

func getFactionListHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	s3Service, err := helpers.InitS3Service(ctx, constants.S3_BUCKET)
	if err != nil {
		log.Printf("Error establishing S3 service: %s", err.Error())
		return helpers.BuildResponse(500, "An internal error occurred")
	}
	data, err := s3Service.Get(constants.FACTIONS_FILE)
	if err != nil {
		errMsg := fmt.Sprintf("Error getting factions file contents from S3: %s", err.Error())
		log.Print(errMsg)
		return helpers.BuildResponse(500, errMsg)
	}
	return helpers.BuildResponse(200, string(data))
}
