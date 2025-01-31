package controller_test

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/server/controller"
	"github.com/mct-joken/kojs5-backend/pkg/utils/seed"

	"github.com/mct-joken/kojs5-backend/pkg/application/user"
	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
	"github.com/mct-joken/kojs5-backend/pkg/utils/mail/dummy"
	"github.com/stretchr/testify/assert"
)

func TestUserController_Create(t *testing.T) {
	r := inmemory.NewUserRepository(seed.NewSeeds().Users)
	u := service.NewUserService(r)
	s := user.NewCreateUserService(r, *u, dummy.NewMailer(), "")
	f := user.NewFindUserService(r)
	c := controller.NewUserController(r, *s, *f)

	res, _ := c.Create(model.CreateUserRequestJSON{
		Name:     "miyoshi",
		Email:    "me@example.jp",
		Password: "hello",
	})

	assert.Equal(t, model.CreateUserResponseJSON{
		ID:    res.ID,
		Name:  "miyoshi",
		Email: "me@example.jp",
	}, res)
}
