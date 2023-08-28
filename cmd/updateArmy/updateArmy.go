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
	lambda.Start(updateArmyHander)
}

func updateArmyHander(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Handler started")
	s3Service, err := helpers.InitS3Service(ctx, constants.S3_BUCKET)
	if err != nil {
		log.Printf("Error establishing S3 service: %s", err.Error())
		return helpers.BuildResponse(500, "An internal error occurred")
	}
	army, err := helpers.GetPathParameter(request, constants.ARMY_PATH)
	if err != nil {
		return helpers.BuildResponse(400, "Missing army path parameter")
	}
	armyFileName := army + ".json"
	bodyStr := request.Body
	_, err = helpers.ValidateJsonBody(bodyStr)
	if err != nil {
		return helpers.BuildResponse(400, "Invalid JSON supplied!")
	}
	if err := s3Service.Put(armyFileName, []byte(bodyStr)); err != nil {
		errMsg := fmt.Sprintf("Error updating army file contents to S3: %s", err.Error())
		log.Print(errMsg)
		return helpers.BuildResponse(500, errMsg)
	}
	data, err := s3Service.Get(armyFileName)
	if err != nil {
		errMsg := fmt.Sprintf("Error getting updated army file contents from S3: %s", err.Error())
		log.Print(errMsg)
		return helpers.BuildResponse(500, errMsg)
	}
	return helpers.BuildResponse(200, string(data))
}
