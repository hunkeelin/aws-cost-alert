package main

import (
	"testing"
)

func TestGetCost(t *testing.T) {
	_, err := getCost()
	if err != nil {
		t.Error(err.Error())
	}
}
