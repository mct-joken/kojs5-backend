package user

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type Data struct {
	id       id.SnowFlakeID
	name     string
	email    string
	password string
	role     int
}

func NewData(id id.SnowFlakeID, name string, email string, password string, role int) *Data {
	return &Data{id: id, name: name, email: email, password: password, role: role}
}

func (d Data) GetID() id.SnowFlakeID {
	return d.id
}

func (d Data) GetName() string {
	return d.name
}

func (d Data) GetEmail() string {
	return d.email
}

func (d Data) GetPassword() string {
	return d.password
}

func (d Data) IsAdmin() bool {
	return d.role == 0
}

// DataToDomain DTOをドメインモデルに変換
func DataToDomain(in Data) domain.User {
	u, _ := domain.NewUser(in.GetID(), in.GetName(), in.GetEmail())
	if in.IsAdmin() {
		u.SetAdmin()
	}
	return *u
}

// DomainToData ドメインモデルをDTOに変換
func DomainToData(in domain.User) Data {
	role := 0
	if !in.IsAdmin() {
		role = 1
	}
	return *NewData(in.GetID(), in.GetName(), in.GetEmail(), in.GetPassword(), role)
}