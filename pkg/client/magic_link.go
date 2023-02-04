package client

import (
	"github.com/denysvitali/devops-security-cli/pkg/client/models"
	"time"
)

type TokenPermission string

type MagicLink struct {
	Id          int               `json:"id"`
	FirstChars  string            `json:"first_chars"`
	Status      models.Status     `json:"status"`
	Permissions []TokenPermission `json:"permissions"`
	CreatedAt   time.Time         `json:"created_at"`
	ExpiresAt   time.Time         `json:"expires_at"`
}

func (c *Client) GetMagicLinks() (*[]MagicLink, error) {
	return Get[[]MagicLink](c, "/api/v1/magic-links", nil)
}
