package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type APIClient struct {
	BASE_URI string
	client   http.Client
}

func NewApiClient() *APIClient {
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return &APIClient{
		client:   client,
		BASE_URI: os.Getenv("API_BASE_URI"),
	}
}

func (a *APIClient) ScheduleTask(data LogisticsRequestDTO) (*OptimizeScheduleResponse, error) {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := a.client.Post(a.BASE_URI+"/schedule", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response OptimizeScheduleResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil

}

func (a *APIClient) ProductionPlan(data ProductionPlanRequest) (*ProductionResponse, error) {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := a.client.Post(a.BASE_URI+"/generate_production_plan",
		"application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response ProductionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil

}
