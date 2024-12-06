package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/heroiclabs/nakama-common/runtime"
)

// Payload structure for parsing incoming JSON
type IncomingPayload struct {
	RequestID string          `json:"RequestID"`
	Body      json.RawMessage `json:"Body"`
}

// OutgoingPayload structure for returning JSON responses
type OutgoingPayload struct {
	RequestID string                 `json:"requestID"`
	Payload   map[string]interface{} `json:"payload"`
}

// func rpcPing(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
func executeLambda(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	logger.Debug("Ping received with payload: %s", payload)

	// Parse the incoming payload
	var incoming IncomingPayload
	err := json.Unmarshal([]byte(payload), &incoming)
	if err != nil {
		logger.Error("Error parsing payload: %v", err)
		return "", err
	}

	// Extract requestID and decode Body
	requestID := incoming.RequestID
	var rawBody string
	err = json.Unmarshal(incoming.Body, &rawBody) // Decode the JSON string
	if err != nil {
		logger.Error("Error decoding Body field: %v", err)
		return "", err
	}

	// Step 2: Unmarshal the decoded string into a Go map
	var requestBody map[string]interface{}
	err = json.Unmarshal([]byte(rawBody), &requestBody) // Decode into map
	if err != nil {
		logger.Error("Error parsing body: %v", err)
		return "", err
	}

	// Prepare the HTTP request to the Lambda function
	const lambdaFunctionUrl = "https://gon7npt4xmt4lomya5gqvsat4i0besqw.lambda-url.ap-south-1.on.aws/"
	bodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		logger.Error("Error serializing request body: %v", err)
		return "", err
	}

	req, err := http.NewRequest("POST", lambdaFunctionUrl, bytes.NewBuffer(bodyBytes))
	if err != nil {
		logger.Error("Error creating HTTP request: %v", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	// Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error calling Lambda function: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body: %v", err)
		return "", err
	}

	logger.Info("Response from Lambda: %s", string(respBody))

	// Construct the outgoing response
	responsePayload := OutgoingPayload{
		RequestID: requestID,
		Payload: map[string]interface{}{
			"OpCode":  "LambdaResponse",
			"Message": string(respBody),
		},
	}

	// Serialize and return the response
	responseBytes, err := json.Marshal(responsePayload)
	if err != nil {
		logger.Error("Error serializing response: %v", err)
		return "", err
	}

	return string(responseBytes), nil
}
