package main

import (
	"os"
	"testing"
)

func TestSendSms(t *testing.T) {
	if os.Getenv("PHONENUMBER") == "" {
		return
	}
	err := sendSms(&sendSmsInput{
		message:     "test",
		phoneNumber: os.Getenv("PHONENUMBER"),
	})
	if err != nil {
		t.Error(err.Error())
	}
}
