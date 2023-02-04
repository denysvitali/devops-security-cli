package client_test

import (
	"github.com/denysvitali/devops-security-cli/pkg/client"
	"github.com/denysvitali/devops-security-cli/pkg/client/models"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const MY_TOKEN = "dummy-token"

func getClient(t *testing.T) *client.Client {
	c, err := client.New(MY_TOKEN)
	if err != nil {
		t.Fatalf("unable to get client: %v", err)
	}
	return c
}

func TestClient_GetMagicLinks(t *testing.T) {
	gock.New(client.ApiEndpoint).
		Get("/api/v1/magic-links").
		Reply(200).
		File("../../resources/test/magic-links.json")
	c := getClient(t)
	magicLinks, err := c.GetMagicLinks()
	if err != nil {
		t.Fatalf("unable to get magic links: %v", err)
	}
	assert.NotNil(t, magicLinks)
	assert.Len(t, *magicLinks, 1)
	assert.Equal(t, 7, (*magicLinks)[0].Id)
	assert.Equal(t, "abc12", (*magicLinks)[0].FirstChars)
	assert.Equal(t, models.ValidStatus, (*magicLinks)[0].Status)
	assert.Equal(t, []client.TokenPermission{
		"consumeLicenses",
		"listMagicLinks",
		"createMagicLinks",
		"revokeMagicLinks",
		"viewUsageInformation",
	}, (*magicLinks)[0].Permissions)
	assert.Equal(t,
		time.Date(2023, 1, 22, 8,32,5, 175000000, time.UTC),
		(*magicLinks)[0].CreatedAt,
	)
	assert.Equal(t,
		time.Date(2023, 2, 28, 23,59,59, 0, time.UTC),
		(*magicLinks)[0].ExpiresAt,
	)
}
