package hub

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"wyvern-server/internal/models"
	"wyvern-server/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateHub(c *gin.Context) {
	// Getting Contexts
	app, err := utils.GetContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, token, err := utils.GetSessionContext(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Getting Params
	owner := c.Query("owner")
	repo := c.Query("repo")
	if owner == "" || repo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid/missing parameters").Error()})
		return
	}

	ownerID, err := GetRepoInfo(token, repo, owner)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Checking if user is owner of repo
	if ownerID == -1 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "you are not owner of repo"})
		return
	}

	var exists bool
	err = app.Pg.QueryRow(`SELECT EXISTS (
		SELECT 1 FROM hub WHERE owner_id = $1 AND repo = $2
		)`, ownerID, repo).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "hub for repo already exists"})
		return
	}

	var hubRow models.Hub
	err = app.Pg.QueryRow(`INSERT INTO hub(repo, owner_id) VALUES ($1, $2) RETURNING id, repo, owner_id`, repo, ownerID).Scan(&hubRow.ID, &hubRow.Repo, &hubRow.OwnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, hubRow)
}

func GetRepoInfo(token, repo, owner string) (int64, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("User-Agent", "wyvern")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return -1, errors.New("repo doesnt exist")
	}

	var repoData struct {
		Owner struct {
			Login string `json:"login"`
			ID    int64  `json:"id"`
		} `json:"owner"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&repoData); err != nil {
		return -1, err
	}

	if !strings.EqualFold(repoData.Owner.Login, owner) {
		return -1, nil
	}

	return repoData.Owner.ID, nil
}
