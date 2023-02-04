package main

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"github.com/denysvitali/devops-security-cli/pkg/client"
	"strings"
)

type MagicLinksCmd struct {
}

func doMagicLinks(c *client.Client) {
	style := GetTableStyle()

	magicLinks, err := c.GetMagicLinks()
	if err != nil {
		logger.Errorf("unable to get magic links: %v", err)
		return
	}

	if magicLinks == nil {
		logger.Errorf("magicLinks is null")
		return
	}

	t := simpletable.New()
	t.SetStyle(style)
	t.Header = &simpletable.Header{Cells: []*simpletable.Cell{
		{
			Align: simpletable.AlignRight,
			Text:  "ID",
		},
		{
			Align: simpletable.AlignLeft,
			Text:  "First Chars",
		},
		{
			Align: simpletable.AlignLeft,
			Text:  "Created At",
		},
		{
			Align: simpletable.AlignLeft,
			Text:  "Expires At",
		},
		{
			Align: simpletable.AlignLeft,
			Text:  "Permissions",
		},
		{
			Align: simpletable.AlignLeft,
			Text: "Status",
		},
	}}
	for _, v := range *magicLinks {
		t.Body.Cells = append(t.Body.Cells, []*simpletable.Cell{
			{
				Align: simpletable.AlignRight,
				Text:  fmt.Sprintf("%d", v.Id),
			},
			{
				Align: simpletable.AlignLeft,
				Text:  v.FirstChars,
			},
			{
				Align: simpletable.AlignLeft,
				Text:  v.CreatedAt.Format("2006-01-02 15:04:05"),
			},
			{
				Align: simpletable.AlignLeft,
				Text:  v.ExpiresAt.Format("2006-01-02 15:04:05"),
			},
			{
				Align: simpletable.AlignLeft,
				Text:  formatPermissions(v.Permissions),
			},
			{
				Align: simpletable.AlignLeft,
				Text: string(v.Status),
			},
		})
	}
	t.Println()
}

func formatPermissions(permissions []client.TokenPermission) string {
	p := []string{}
	for _, v := range permissions {
		p = append(p, string(v))
	}
	return strings.Join(p, ",")
}
