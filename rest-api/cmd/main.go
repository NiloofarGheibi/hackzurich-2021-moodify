package main

import (
	"net/http"
	"rest-api/internal/cache"
	"rest-api/internal/config"
	"rest-api/internal/controllers"
	"rest-api/internal/repository"
	"rest-api/internal/router"
	"rest-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Cannot start service due to lack of configuration")
	}

	conn := repository.InitDatabaseConnection(conf.Database.URL, conf.Database.UserName, conf.Database.Password)
	repo := repository.NewDatabaseRepository(conn)

	// Create TokenCache
	tokenCache, err := cache.NewTokenCache(conf.Spotify.ClientID, conf.Spotify.ClientSecret, conf.Spotify.TokenURL)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("unable to create token cache")
	}

	// Check if auth token is accessible
	err = tokenCache.UpdateToken()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("unable to create spotify service to service auth token")
	}

	spotifyConf := &repository.RestyAuthenticatedClientConfig{HostURL: conf.Spotify.Host, TokenCache: tokenCache}
	authSpotifyClient, err := repository.NewRestyAuthenticatedClient(spotifyConf)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("unable to create authenticated spotify client")
	}
	spotify := repository.NewSpotifyAPIRepository(authSpotifyClient)

	userService := services.NewUserService(repo)
	musicService := services.NewMusicService(repo, spotify)

	userController := controllers.NewUserController(userService)
	radioController := controllers.NewRadioController(musicService)

	r := router.Create()
	r.SetHealthEndPoint(health)
	r.InitUserRoutes(userController)
	r.InitRadioRoutes(radioController)

	log.Fatal().
		Err(r.Router.Run()).
		Msg("Cannot start service")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "alive",
	})
}
