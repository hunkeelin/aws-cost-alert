package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

func handler(ctx context.Context) (string, error) {
	cost, err := getCost()
	if err != nil {
		return "", err
	}
	if os.Getenv("PHONENUMBER") == "" {
		return "", fmt.Errorf("Please specify env variable PHONENUMBER including the country code")
	}
	err = sendSms(&sendSmsInput{
		message:     fmt.Sprintf("The aws cost so far for today is $%s", cost),
		phoneNumber: os.Getenv("PHONENUMBER"),
	})
	return "v1", nil
}

func main() {
	lambda.Start(handler)
}
