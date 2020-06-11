package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type sendSmsInput struct {
	message     string
	phoneNumber string
}

func sendSms(s *sendSmsInput) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		return err
	}
	svc := sns.New(sess)
	_, err = svc.Publish(&sns.PublishInput{
		Message:     aws.String(s.message),
		PhoneNumber: aws.String(s.phoneNumber),
	})
	if err != nil {
		return err
	}
	return nil
}
