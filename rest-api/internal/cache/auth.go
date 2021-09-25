package cache

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	oauth2cc "golang.org/x/oauth2/clientcredentials"
	"net/url"
	"sync"
)

// TokenCache interface for TokenCache
type TokenCache interface {
	GetAccessToken() string
	Valid() bool
	UpdateToken() error
}

// NewTokenCache constructor for the token cache
func NewTokenCache(clientID string, clientSecret string, tokenURL string) (TokenCache, error) {
	if clientID == "" {
		return nil, errors.New("clientID is empty")
	}

	if clientSecret == "" {
		return nil, errors.New("clientSecret is empty")
	}

	if _, err := url.Parse(tokenURL); err != nil || tokenURL == "" {
		return nil, errors.New("tokenURL is empty")
	}

	ctx := context.Background()

	oauth2config := oauth2cc.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
	}

	tokenCache := &RestyTokenCache{
		Ctx:          ctx,
		Oauth2config: oauth2config,
	}

	return tokenCache, nil
}

// restyTokenCache implements TokenCache
type RestyTokenCache struct {
	token        *oauth2.Token
	Oauth2config oauth2cc.Config
	mux          sync.Mutex
	Ctx          context.Context
}

func (c *RestyTokenCache) setToken(token *oauth2.Token) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access this property
	c.token = token
	c.mux.Unlock()
}

func (c *RestyTokenCache) UpdateToken() error {
	token, err := c.Oauth2config.Token(c.Ctx)

	if err != nil {
		log.Error().
			Msg("resty token cache experienced an error while retrieving a new token")
		return err
	}
	c.setToken(token)
	return nil
}

func (c *RestyTokenCache) Valid() bool {
	return c.token.Valid()
}

func (c *RestyTokenCache) GetAccessToken() string {
	if !c.Valid() {
		return ""
	}
	fmt.Println(c.token.AccessToken)
	return c.token.AccessToken
}
