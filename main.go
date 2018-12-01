package main

import (
	"fmt"
	"net/http"

	"github.com/go-training/facebook-account-kit/config"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/resty.v1"
)

// AuthSuccess success response for get access token
type AuthSuccess struct {
	ID                      string `json:"id"`
	AccessToken             string `json:"access_token"`
	TokenRefreshIntervalSec int    `json:"token_refresh_interval_sec"`
}

// AuthError error response for get access token
type AuthError struct {
	Error struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      int    `json:"code"`
		FbtraceID string `json:"fbtrace_id"`
	} `json:"error"`
}

// Me query user from access token
type Me struct {
	Email struct {
		Address string `json:"address"`
	} `json:"email"`
	Phone struct {
		Number         string `json:"number"`
		CountryPrefix  string `json:"country_prefix"`
		NationalNumber string `json:"national_number"`
	} `json:"phone"`
	ID          string `json:"id"`
	Application struct {
		ID string `json:"id"`
	} `json:"application"`
}

var (
	tokenExchangeURL = "https://graph.accountkit.com/%s/access_token?" +
		"grant_type=authorization_code&code=%s&&access_token=AA|%s|%s"
	getMeURL = "https://graph.accountkit.com/%s/me?access_token=%s"
)

func main() {
	conf := config.MustLoad()

	router := gin.Default()
	router.Static("/images", "./images")
	router.LoadHTMLGlob("templates/*")

	// user login page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title":   "facebook accountkit example",
			"appID":   conf.Facebook.AppID,
			"version": conf.Facebook.Version,
		})
	})

	router.POST("/login", func(c *gin.Context) {
		code := c.PostForm("code")
		url := fmt.Sprintf(tokenExchangeURL, conf.Facebook.Version, code, conf.Facebook.AppID, conf.Facebook.Secret)
		authSuccess := &AuthSuccess{}
		authError := &AuthError{}
		resp, _ := resty.R().
			SetResult(authSuccess).
			SetError(authError).
			Get(url)
		fmt.Printf("\nResponse Body: %v", resp)

		if resp.StatusCode() == http.StatusOK && authSuccess.AccessToken != "" {
			user := &Me{}
			// Get Account Kit information
			url := fmt.Sprintf(getMeURL, conf.Facebook.Version, authSuccess.AccessToken)
			resp, _ := resty.R().
				SetResult(user).
				SetError(authError).
				Get(url)
			fmt.Printf("\nResponse Body: %v", resp)

			c.HTML(http.StatusOK, "success.html", gin.H{
				"title":          "facebook accountkit example",
				"email":          user.Email.Address,
				"phone":          user.Phone.Number,
				"countryPrefix":  user.Phone.CountryPrefix,
				"nationalNumber": user.Phone.NationalNumber,
				"id":             user.ID,
				"applicationID":  user.Application.ID,
				"url":            url,
			})
			return
		}
	})

	router.Run(":" + conf.HTTP.Port)
}
