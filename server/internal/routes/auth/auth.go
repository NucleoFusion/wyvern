package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"wyvern-server/internal/models"
	"wyvern-server/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type GithubResponse struct {
	ID int `json:"id"` // Only require ID at this point, more data when setting up profile
}

func AddAuthRoutes(r *gin.Engine) {
	grp := r.Group("/auth")

	grp.GET("/github", func(c *gin.Context) {
		godotenv.Load(".env")

		clientID := os.Getenv("CLIENT_ID")

		authURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&scope=user", clientID)
		c.Redirect(http.StatusTemporaryRedirect, authURL)
	})

	r.GET("/auth/callback", func(c *gin.Context) {
		// Getting Context
		app, err := utils.GetContext(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Getting Vars
		code := c.Query("code")

		godotenv.Load(".env")
		clientID := os.Getenv("CLIENT_ID")
		clientSecret := os.Getenv("CLIENT_SECRET")

		// Getting Access Token from Github
		tokenResp, _ := http.PostForm("https://github.com/login/oauth/access_token", url.Values{
			"client_id":     {clientID},
			"client_secret": {clientSecret},
			"code":          {code},
		})

		body, _ := io.ReadAll(tokenResp.Body)
		values, _ := url.ParseQuery(string(body))
		accessToken := values.Get("access_token")

		// Getting User Info
		req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
		req.Header.Set("Authorization", "BEARER "+accessToken)
		req.Header.Set("Accept", "application/vnd.github+json")

		client := &http.Client{}
		resp, _ := client.Do(req)
		userData, _ := io.ReadAll(resp.Body)

		// Parsing ID from Github Response
		var githubResp GithubResponse
		err = json.Unmarshal(userData, &githubResp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Creating User if not exists
		userExists := false
		var userID int
		err = app.Pg.QueryRow("INSERT INTO users(github_id) VALUES ($1) RETURNING id", githubResp.ID).Scan(&userID)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				userExists = true

				// Fetching User if Exists
				err = app.Pg.QueryRow(
					"SELECT id FROM users WHERE github_id = $1",
					githubResp.ID,
				).Scan(&userID)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "user exists, but fetch failed"})
					return
				}
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		// If User Exists
		uuidToken := "" // Will be used to check if already existed
		if userExists {
			val, err := app.Rdb.Get(context.Background(), fmt.Sprintf("gh:%d", githubResp.ID)).Result()
			if err == redis.Nil {
				// Do nothing, so a new uuid will be created
			} else if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			} else {
				uuidToken = val
			}
		}

		// If Entry in redis existed for user (Refreshing token)
		if uuidToken != "" {
			err := app.Rdb.Set(context.Background(), "sess:"+uuidToken, accessToken, 24*5*time.Hour).Err()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.Redirect(http.StatusPermanentRedirect, "http://localhost:5173/home")
			return
		}

		// Getting a UUID, since none exists
		var uuidStr string

		for {
			uuidStr = uuid.NewString()
			exists, err := app.Rdb.Exists(context.Background(), "sess:"+uuidStr).Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			if exists == 0 {
				break
			}
		}

		// Setting AccessToken & UUID entry in Redis
		err = app.Rdb.Set(context.Background(), "sess:"+uuidStr, accessToken, 24*5*time.Hour).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = app.Rdb.Set(context.Background(), fmt.Sprintf("gh:%d", githubResp.ID), uuidStr, 24*5*time.Hour).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Setting Cookie
		cookieData := models.UserCookie{
			UUID:     uuidStr,
			UserID:   userID,
			GithubID: githubResp.ID,
		}
		data, _ := json.Marshal(cookieData)

		c.SetCookie("wyvern_session", string(data), 7*24*60*60, "/", "localhost", false, true) // TODO: change for hosting
		fmt.Println(string(data))                                                              // TODO: ONLY FOR DEV, NOT IN PROD
		c.Redirect(http.StatusPermanentRedirect, "http://localhost:5173/home")
	})
}
