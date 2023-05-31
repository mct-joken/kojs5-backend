package mongodb

import (
	"context"
	"fmt"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository/mongodb/entity"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
	"go.mongodb.org/mongo-driver/bson"
)

type ProblemRepository struct {
	cli Client
}

func (p ProblemRepository) CreateProblem(in domain.Problem) error {
	sets := in.GetCaseSets()
	setsEntity := make([]entity.CaseSet, len(sets))
	for i, v := range sets {
		caseEntity := make([]entity.Case, len(v.GetCases()))
		for j, k := range v.GetCases() {
			caseEntity[j] = entity.Case{
				ID:        k.GetID(),
				CaseSetID: k.GetCasesetID(),
				In:        k.GetIn(),
				Out:       k.GetOut(),
			}
		}

		setsEntity[i] = entity.CaseSet{
			ID:    v.GetID(),
			Name:  v.GetName(),
			Point: v.GetPoint(),
			Cases: caseEntity,
		}
	}
	e := entity.Problem{
		ID:          in.GetProblemID(),
		ContestID:   in.GetContestID(),
		Index:       in.GetIndex(),
		Title:       in.GetTitle(),
		Text:        in.GetText(),
		Point:       in.GetPoint(),
		MemoryLimit: in.GetMemoryLimit(),
		TimeLimit:   in.GetTimeLimit(),
		CaseSets:    setsEntity,
	}

	_, err := p.cli.Cli.Database("kojs").Collection("problem").InsertOne(context.Background(), e)
	if err != nil {
		return err
	}

	return nil
}

func (p ProblemRepository) FindProblemByID(id id.SnowFlakeID) *domain.Problem {
	result := p.cli.Cli.Database("kojs").Collection("problem").FindOne(context.Background(), &bson.M{"_id": id})

	var problem entity.Problem
	if err := result.Decode(&problem); err != nil {
		fmt.Println(err)
		return nil
	}
	res := problem.ToDomain()
	return &res
}

func (p ProblemRepository) FindProblemByTitle(name string) *domain.Problem {
	result := p.cli.Cli.Database("kojs").Collection("problem").FindOne(context.Background(), &bson.M{"name": name})

	var problem entity.Problem
	if err := result.Decode(&problem); err != nil {
		return nil
	}
	res := problem.ToDomain()
	return &res
}

func (p ProblemRepository) FindCaseSetByID(id id.SnowFlakeID) *domain.Caseset {
	filter := &bson.M{"casesets.id": id}
	cursor := p.cli.Cli.Database("kojs").Collection("problem").FindOne(context.Background(), filter)

	var problem entity.Problem
	if err := cursor.Decode(&problem); err != nil {
		return nil
	}
	res := problem.ToDomain()
	for _, v := range res.GetCaseSets() {
		if v.GetID() == id {
			return &v
		}
	}
	return nil
}

func (p ProblemRepository) FindCaseByID(id id.SnowFlakeID) *domain.Case {
	cursor := p.cli.Cli.Database("kojs").Collection("problem").FindOne(context.Background(), &bson.M{"casesets.cases.id": id})

	var problem entity.Problem
	if err := cursor.Decode(&problem); err != nil {
		fmt.Println(err)
		return nil
	}
	res := problem.ToDomain()
	for _, v := range res.GetCaseSets() {
		for _, k := range v.GetCases() {
			if k.GetID() == id {
				return &k
			}
		}
	}

	return nil
}

func NewProblemRepository(cli Client) *ProblemRepository {
	return &ProblemRepository{cli: cli}
}