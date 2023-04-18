package domain

import (
	"time"
	"unicode/utf8"

	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type Contest struct {
	id          id.SnowFlakeID
	title       string
	description string
	startAt     time.Time
	endAt       time.Time
}

// ContestTitleLengthError コンテストタイトルの文字数エラー
type ContestTitleLengthError struct {
	message string
}

func (e ContestTitleLengthError) Error() string {
	return e.message
}

// ContestDescriptionLengthError コンテストタイトルの文字数エラー
type ContestDescriptionLengthError struct {
	message string
}

func (e ContestDescriptionLengthError) Error() string {
	return e.message
}

// ContestDateInvalidError コンテスト開始/終了時刻が不正のときのエラー
type ContestDateInvalidError struct {
}

func (e ContestDateInvalidError) Error() string {
	return ""
}

/*
NewContest
不変値: ID
*/
func NewContest(cID id.SnowFlakeID) *Contest {
	return &Contest{
		id: cID,
	}
}

func (c *Contest) GetID() id.SnowFlakeID {
	return c.id
}

func (c *Contest) GetTitle() string {
	return c.title
}

func (c *Contest) GetDescription() string {
	return c.description
}

func (c *Contest) GetStartAt() time.Time {
	return c.startAt
}

func (c *Contest) GetEndAt() time.Time {
	return c.endAt
}

func (c *Contest) SetTitle(title string) error {
	/*
		Title文字数 制約
		5文字以上 128文字以下
	*/
	if utf8.RuneCountInString(title) < 5 || utf8.RuneCountInString(title) > 128 {
		return ContestTitleLengthError{}
	}

	c.title = title
	return nil
}

func (c *Contest) SetDescription(description string) error {
	/*
		Description文字数 制約
		10文字以上 50000文字以下
	*/
	if utf8.RuneCountInString(description) < 10 || utf8.RuneCountInString(description) > 50000 {
		return ContestDescriptionLengthError{}
	}
	c.description = description
	return nil
}

func (c *Contest) SetStartAt(at time.Time) error {
	/*
		StartAt/EndAt 制約
		EndAtはStartAtより1分以上後にしなければいけない
	*/
	if at.After(c.endAt) || at.Sub(c.endAt) < 60*time.Second {
		return ContestDateInvalidError{}
	}

	c.startAt = at
	return nil
}

func (c *Contest) SetEndAt(at time.Time) error {
	if at.Before(c.startAt) || at.Sub(c.endAt) < 60*time.Second {
		return ContestDateInvalidError{}
	}

	c.endAt = at
	return nil
}