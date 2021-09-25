package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Training interface {
	GetMood(ctx context.Context, features string) (int, error)
}

type training struct {
	resty *resty.Client
}

// NewRestyClient returns a *resty.Client
func NewRestyClient(url string) (*resty.Client, error) {
	r := resty.New()
	r.SetHostURL(url)
	return r, nil
}

func NewTrainingAPIRepository(restyClient *resty.Client) Training {
	return &training{
		resty: restyClient,
	}
}

func (t training) GetMood(ctx context.Context, features string) (int, error) {
	uJSON, err := json.Marshal(&map[string]string{"features": features})
	if err != nil {
		return -1, err
	}

	resp, err := t.resty.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(uJSON).
		Post("/mood")

	if err != nil {
		return -1, err
	}

	if resp.IsError() {
		return -1, fmt.Errorf("unexpected status code: " + resp.Status())
	}

	var m *mood
	err = json.Unmarshal(resp.Body(), &m)
	if err != nil {
		return -1, err
	}

	return m.mood, nil
}

type mood struct {
	mood int `json:"mood"`
}
