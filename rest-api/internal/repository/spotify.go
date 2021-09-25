package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	spt "github.com/zmb3/spotify"
	"net/http"
	"net/url"
	"rest-api/internal/cache"
	"syscall"
)

type Spotify interface {
	GetAudioAnalysisByTrackID(ctx context.Context, id string) (*spt.AudioFeatures, error)
}

type spotify struct {
	resty *resty.Client
	Cache cache.TokenCache
}

// RestyAuthenticatedClientConfig represents configuration to create a NewRestyAuthenticatedClient()
type RestyAuthenticatedClientConfig struct {
	// HostURL to the Rest API that we should connect to
	HostURL string
	// TokenCache
	TokenCache cache.TokenCache
}

// NewRestyAuthenticatedClient returns a *resty.Client
func NewRestyAuthenticatedClient(config *RestyAuthenticatedClientConfig) (*resty.Client, error) {
	tokenCache := config.TokenCache
	if _, err := url.Parse(config.HostURL); err != nil || config.HostURL == "" {
		return nil, errors.New("bad HostURL")
	}
	r := resty.New()
	r.SetHostURL(config.HostURL)

	r.OnBeforeRequest(func(cli *resty.Client, req *resty.Request) error {
		if !tokenCache.Valid() {
			err := tokenCache.UpdateToken()
			if err != nil {
				return err
			}
		}

		cli.SetAuthToken(tokenCache.GetAccessToken())
		req.SetAuthToken(tokenCache.GetAccessToken())

		return nil
	})

	r.OnAfterResponse(func(_ *resty.Client, res *resty.Response) error {
		return nil
	})

	r.SetRetryCount(5)

	// Condition function will be provided with *resty.Response as a
	// parameter. It is expected to return (bool, error) pair.
	// Resty will retry in case condition returns true or non nil error.
	r.AddRetryCondition(func(r *resty.Response, err error) bool {
		if r.StatusCode() == http.StatusUnauthorized {
			log.Warn().Msg("resty http client retry conditions met, status-code == 401, trying to get a new token")
			err = tokenCache.UpdateToken()
			return true
		}
		return false
	})

	return r, nil
}

// GET https://api.spotify.com/v1/audio-features/{id}
func (s spotify) GetAudioAnalysisByTrackID(ctx context.Context, id string) (*spt.AudioFeatures, error) {
	resp, err := s.resty.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetPathParams(map[string]string{"id": id}).
		Get("/v1/audio-features/{id}")

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("unexpected status code: " + resp.Status())
	}

	var audioFeature *spt.AudioFeatures
	err = json.Unmarshal(resp.Body(), &audioFeature)
	if err != nil && err != syscall.EPIPE {
		return nil, err
	}

	return audioFeature, nil
}

func NewSpotifyAPIRepository(restyClient *resty.Client) Spotify {
	return &spotify{
		resty: restyClient,
	}
}
