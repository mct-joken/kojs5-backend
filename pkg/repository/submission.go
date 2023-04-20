package repository

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type SubmissionRepository interface {
	// FindSubmissionByID 提出を1つ取得
	FindSubmissionByID(id id.SnowFlakeID) *domain.Submission

	// FindSubmissionResultByID 提出結果を1つ取得
	FindSubmissionResultByID(id id.SnowFlakeID) *domain.SubmissionResult
}