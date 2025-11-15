package models

import "time"

type PullRequestModel struct {
	PullRequestId     string
	PullRequestName   string
	Status            string
	AuthorId          string
	NeedMoreReviewers bool
	CreatedAt         time.Time
	MergedAt          *time.Time
}
