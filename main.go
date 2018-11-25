package main

import (
	"fmt"
	"net/http"

	"github.com/go-training/facebook-account-kit/config"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/resty.v1"
)

type AuthSuccess struct {
	ID                      string `json:"id"`
	AccessToken             string `json:"access_token"`
	TokenRefreshIntervalSec int    `json:"token_refresh_interval_sec"`
}

type AuthError struct {
	Error struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      int    `json:"code"`
		FbtraceID string `json:"fbtrace_id"`
	} `json:"error"`
}

type Me struct {
	Email struct {
		Address string `json:"address"`
	} `json:"email"`
	ID          string `json:"id"`
	Application struct {
		ID string `json:"id"`
	} `json:"application"`
}

var tokenExchangeURL = "https://graph.accountkit.com/%s/access_token?" +
	"grant_type=authorization_code&code=%s&&access_token=AA|%s|%s"

func main() {
	conf := config.MustLoad()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// user login page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "facebook accountkit example",
		})
	})

	// redire to [GET] /login page if login as email.
	router.GET("/login", func(c *gin.Context) {
		code := c.Query("code")
		url := fmt.Sprintf(tokenExchangeURL, conf.Facebook.Version, code, conf.Facebook.AppID, conf.Facebook.Secret)
		authSuccess := &AuthSuccess{}
		authError := &AuthError{}
		resp, err := resty.R().
			SetResult(authSuccess).
			SetError(authError).
			Get(url)

		// explore response object
		fmt.Printf("\nError: %v", err)
		fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
		fmt.Printf("\nResponse Status: %v", resp.Status())
		fmt.Printf("\nResponse Time: %v", resp.Time())
		fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
		fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())

		me := &Me{}
		if resp.StatusCode() == http.StatusOK && authSuccess.AccessToken != "" {
			// Get Account Kit information
			meURL := "https://graph.accountkit.com/" + conf.Facebook.Version + "/me?" +
				"access_token=" + authSuccess.AccessToken
			resp, err := resty.R().
				SetHeader("Accept", "application/json").
				SetResult(me).
				SetError(authError).
				Get(meURL)
			fmt.Printf("\nError: %v", err)
			fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
			fmt.Printf("\nResponse Status: %v", resp.Status())
			fmt.Printf("\nResponse Time: %v", resp.Time())
			fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
			fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())
		}

		c.HTML(http.StatusOK, "success.html", gin.H{
			"title":         "facebook accountkit example",
			"email":         me.Email.Address,
			"id":            me.ID,
			"applicationID": me.Application.ID,
		})
	})
	router.Run(":8080")
}
