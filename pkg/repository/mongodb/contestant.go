package mongodb

import (
	"context"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository/mongodb/entity"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
	"go.mongodb.org/mongo-driver/bson"
)

type ContestantRepository struct {
	client Client
}

func (c ContestantRepository) JoinContest(d domain.Contestant) error {
	role := 0
	if d.IsAdmin() {
		role = 1
	}
	e := entity.Contestant{
		ID:        d.GetID(),
		ContestID: d.GetContestID(),
		UserID:    d.GetUserID(),
		Role:      role,
		Point:     d.GetPoint(),
	}
	_, err := c.client.Cli.Database("kojs").Collection("contestant").InsertOne(context.Background(), e)
	if err != nil {
		return err
	}

	return nil
}

func (c ContestantRepository) FindContestantByID(id id.SnowFlakeID) *domain.Contestant {
	filter := &bson.M{"_id": id}

	result := c.client.Cli.Database("kojs").Collection("contestant").FindOne(context.Background(), filter)

	var contestant entity.Contestant
	if err := result.Decode(&contestant); err != nil {
		return nil
	}
	res := contestant.ToDomain()
	return &res
}

func (c ContestantRepository) FindContestantByUserID(id id.SnowFlakeID) []domain.Contestant {
	filter := &bson.M{"userID": id}

	cursor, err := c.client.Cli.Database("kojs").Collection("contestant").Find(context.Background(), filter)
	if err != nil {
		return []domain.Contestant{}
	}

	var contestant []entity.Contestant
	if err := cursor.All(context.Background(), &contestant); err != nil {
		return nil
	}
	res := make([]domain.Contestant, len(contestant))
	for i, v := range contestant {
		res[i] = v.ToDomain()
	}
	return res
}

func (c ContestantRepository) FindContestantByContestID(id id.SnowFlakeID) []domain.Contestant {
	filter := &bson.M{"contestID": id}
	cursor, err := c.client.Cli.Database("kojs").Collection("contestant").Find(context.Background(), filter)
	if err != nil {
		return []domain.Contestant{}
	}

	var contestant []entity.Contestant
	if err := cursor.All(context.Background(), &contestant); err != nil {
		return nil
	}
	res := make([]domain.Contestant, len(contestant))
	for i, v := range contestant {
		res[i] = v.ToDomain()
	}
	return res
}

func NewContestantRepository(cli Client) *ContestRepository {
	return &ContestRepository{client: cli}
}