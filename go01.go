package go01

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	FailedMessagesQueue []FailedMessage // failed messages queue
	SuccessfulMessages  int
)

func NotifyClient(clientUrl string, messages []string, duration time.Duration) Response {
	// split messages into chunks and send via goroutine
	messageChunks := ChunkSlice(messages, 20)

	var wg sync.WaitGroup

	for i, messageChunk := range messageChunks {
		fmt.Println("Starting notification worker : ", i+1)
		wg.Add(1)
		go SendHttpNotification(clientUrl, messageChunk, &wg, duration)
	}

	wg.Wait()

	return Response{FailedMessages: FailedMessagesQueue,
		SuccessfulMessagesCount: SuccessfulMessages,
		FailedMessagesCount:     len(FailedMessagesQueue)}
}

func SendHttpNotification(client string, messages []string, wg *sync.WaitGroup, duration time.Duration) {
	// loop message and send request
	defer wg.Done()

	var reqWg sync.WaitGroup

	for _, message := range messages {
		// create request body

		reqBody := RequestBody{
			Id:      "",
			Message: message,
		}

		reqBodyJson, err := json.Marshal(reqBody)

		if err != nil {
			log.Println(err)
		}

		// make request
		reqWg.Add(1)
		go func(client, message string) {
			defer reqWg.Done()
			resp, err := http.Post(client, "application/json", bytes.NewBuffer(reqBodyJson))
			if err != nil {
				FailedMessagesQueue = append(FailedMessagesQueue, FailedMessage{
					Message: message,
					Id:      reqBody.Id,
					Error:   err,
				})
			} else {
				if resp.StatusCode >= 400 {
					FailedMessagesQueue = append(FailedMessagesQueue, FailedMessage{
						Message: message,
						Id:      reqBody.Id,
						Error:   errors.New(resp.Status),
					})
				} else {
					SuccessfulMessages += 1
				}
			}
		}(client, message)

		time.Sleep(duration) // duration
	}
	time.Sleep(time.Second) // some time to clear goroutines
	reqWg.Wait()
}

func ChunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for {
		if len(slice) == 0 {
			break
		}

		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}
