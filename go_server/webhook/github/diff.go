package github

import (
	"context"
	"fmt"
	"log/slog"
)

type ChangeResponse struct {
	Filename  string `json:"filename"`
	Status    string `json:"status"`
	Patch     string `json:"patch"` // Contains the actual diff, empty if diff is too large or file is binary
	Additions int    `json:"additions"`
	Deletions int    `json:"deletions"`
}

func (client *Client) GetChanges(ctx context.Context, owner, repoName string, prNumber int, installID string) ([]ChangeResponse, error) {
	var page int = 1
	var allFiles []ChangeResponse

	for {
		var pageFiles []ChangeResponse

		urlPath := fmt.Sprintf("/repos/%s/%s/pulls/%d/files?per_page=100&page=%d", owner, repoName, prNumber, page)

		_, respErr := client.DoRequest(ctx, urlPath, "GET", installID, nil, &pageFiles)

		if respErr != nil {
			slog.Error("Unable to fetch changes", "ERROR", respErr)
			return nil, respErr
		}

		allFiles = append(allFiles, pageFiles...)

		if len(pageFiles) < 100 {
			break
		}
		page++
	}

	return allFiles, nil
}
