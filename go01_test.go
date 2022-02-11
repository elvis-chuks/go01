package go01_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/elvis-chuks/go01"
)

func TestNotifyClient(t *testing.T) {
	messages := []string{"hi", "ho"}
	response := go01.NotifyClient("http://localhost:5000/api/v1/", messages, time.Second)
	fmt.Println(response)
}
