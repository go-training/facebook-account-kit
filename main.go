package main

import (
	"fmt"
	"log"
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

func main() {
	conf := config.MustLoad()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "facebook accountkit example",
		})
	})
	router.GET("/login", func(c *gin.Context) {
		code := c.Query("code")
		token_exchange_url := "https://graph.accountkit.com/" + conf.Facebook.Version + "/access_token?" +
			"grant_type=authorization_code&code=" + code +
			"&access_token=AA|" + conf.Facebook.AppID + "|" + conf.Facebook.Secret
		log.Println(token_exchange_url)
		authSuccess := &AuthSuccess{}
		authError := &AuthError{}
		resp, err := resty.R().
			SetHeader("Accept", "application/json").
			SetResult(authSuccess).
			SetError(authError).
			Get(token_exchange_url)

		// explore response object
		fmt.Printf("\nError: %v", err)
		fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
		fmt.Printf("\nResponse Status: %v", resp.Status())
		fmt.Printf("\nResponse Time: %v", resp.Time())
		fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
		fmt.Printf("\nResponse Body: %v", resp) // or resp.String() or string(resp.Body())
		fmt.Printf("\nauthSuccess: %#v\n", authSuccess)
		fmt.Printf("\nauthError: %#v\n", authError)

		if resp.StatusCode() == http.StatusOK && authSuccess.AccessToken != "" {
			// Get Account Kit information
			meURL := "https://graph.accountkit.com/" + conf.Facebook.Version + "/me?" +
				"access_token=" + authSuccess.AccessToken
			resp, err := resty.R().
				SetHeader("Accept", "application/json").
				// SetResult(authSuccess).
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
			"title": "facebook accountkit example",
		})
	})
	router.Run(":8080")
}
