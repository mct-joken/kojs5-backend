package mongodb

import (
	"context"
	"fmt"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository/mongodb/entity"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
	"go.mongodb.org/mongo-driver/bson"
)

type SubmissionRepository struct {
	client Client
}

func (s SubmissionRepository) FindSubmissionByID(id id.SnowFlakeID) (*domain.Submission, error) {
	filter := &bson.M{"_id": id}

	result := s.client.Cli.Database("kojs").Collection("submission").FindOne(context.Background(), filter)

	var submission entity.Submission
	if err := result.Decode(&submission); err != nil {
		return nil, fmt.Errorf("failed to decode submission data: %w", err)
	}

	res := submission.ToDomain()
	return &res, nil
}

func NewSubmissionRepository(client Client) *SubmissionRepository {
	return &SubmissionRepository{client: client}
}
