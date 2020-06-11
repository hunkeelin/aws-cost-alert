package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"time"
)

func getCost() (string, error) {
	sess, err := session.NewSession(nil)
	if err != nil {
		return "", err
	}
	svc := costexplorer.New(sess)
	result, err := svc.GetCostAndUsage(&costexplorer.GetCostAndUsageInput{
		Metrics:     []*string{aws.String("UnblendedCost")},
		Granularity: aws.String("DAILY"),
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(fmt.Sprintf(time.Now().AddDate(0, 0, -1).UTC().Format("2006-01-02"))),
			End:   aws.String(fmt.Sprintf(time.Now().UTC().Format("2006-01-02"))),
		},
	})
	if err != nil {
		return "", err
	}
	if len(result.ResultsByTime) != 1 {
		return "", fmt.Errorf("Result by time should be 1")
	}
	return *result.ResultsByTime[0].Total["UnblendedCost"].Amount, nil
}
