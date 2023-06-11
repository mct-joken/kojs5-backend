package service

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestCaseSetService_IsExists(t *testing.T) {
	r := inmemory.NewProblemRepository(dummyData.ProblemArray)
	s := NewCaseSetService(r)

	// trueのとき
	assert.Equal(t, true, s.IsExists(*dummyData.ExistsCasesetData))
	// falseのとき
	assert.Equal(t, false, s.IsExists(*dummyData.NotExistsCasesetData))
}
