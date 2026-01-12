package external

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"health-app/internal/domain/patient"
)

type HISAPIClient struct {
	client *http.Client
}

type HISAPI interface {
	GetPatientByID(ctx context.Context, baseURL, id string) (*patient.Patient, error)
	SearchPatients(ctx context.Context, baseURL string, filters patient.PatientsSearchRequest) ([]*patient.Patient, error)
}

func NewHISAPIClient() HISAPI {
	return &HISAPIClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *HISAPIClient) GetPatientByID(ctx context.Context, baseURL, id string) (*patient.Patient, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/patient/search/%s", baseURL, id),
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external api error: %d", resp.StatusCode)
	}

	type Response[T any] struct {
		Data T `json:"data"`
	}
	var res Response[patient.Patient]
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (c *HISAPIClient) SearchPatients(ctx context.Context, baseURL string, filters patient.PatientsSearchRequest) ([]*patient.Patient, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/patient/search", baseURL),
		nil,
	)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"national_id":   filters.NationalID,
		"passport_id":   filters.PassportID,
		"first_name":    filters.FirstName,
		"middle_name":   filters.MiddleName,
		"last_name":     filters.LastName,
		"date_of_birth": filters.DateOfBirth,
		"phone_number":  filters.PhoneNumber,
		"email":         filters.Email,
	}

	q := req.URL.Query()
	for key, value := range params {
		if v := strings.TrimSpace(value); v != "" {
			q.Add(key, v)
		}
	}

	req.URL.RawQuery = q.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("external api error: %d", resp.StatusCode)
	}

	type Response[T any] struct {
		Data T `json:"data"`
	}
	var res Response[[]*patient.Patient]
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res.Data, nil
}
