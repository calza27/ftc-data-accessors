package helpers

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
)

func ValidateJsonBody(bodyStr string) (map[string]interface{}, error) {
	var bodyJSON map[string]interface{}
	err := json.Unmarshal([]byte(bodyStr), &bodyJSON)
	if err != nil {
		return nil, err
	}
	return bodyJSON, nil
}

func BuildResponse(statusCode int, message string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       message,
		StatusCode: statusCode,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"content-type":                "application/json",
		},
	}, nil
}

func GetPathParameter(request events.APIGatewayProxyRequest, parameterName string) (string, error) {
	if request.PathParameters == nil {
		return "", errors.New("no path parameters")
	}
	pathParameter, ok := request.PathParameters[parameterName]
	if !ok {
		return "", errors.New("No path parameter with name " + parameterName)
	}
	return pathParameter, nil
}
