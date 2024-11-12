package goverseerr

import (
	"fmt"
)

type IssueCount struct {
	Total    int `json:"total"`
	Video    int `json:"video"`
	Audio    int `json:"audio"`
	Subtitle int `json:"subtitle"`
	Other    int `json:"other"`
	Open     int `json:"open"`
	Closed   int `json:"closed"`
}

func (o *Overseerr) GetIssueCounts() (*IssueCount, error) {
	var issueCount *IssueCount
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetResult(&issueCount).Get("/issue/count")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("received non-200 status code (%d)", resp.StatusCode())
	}
	return issueCount, nil
}
