package models

type UserCookie struct {
	UUID     string `json:"uuid"`
	UserID   int    `json:"user_id"`
	GithubID int    `json:"github_id"`
}

type UserModel struct {
	ID        int    `json:"id" db:"id"`
	GithubID  int    `json:"github_id" db:"github_id"`
	AvatarURL string `json:"avatar_url" db:"avatar_url"`

	// Custom stuff NOT GITHUBS'!!
	DisplayName string `json:"display_name" db:"display_name"`
	Bio         string `json:"bio" db:"bio"`
}
