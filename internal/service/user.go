package service

import (
	"calling-bill/ent"
	entCall "calling-bill/ent/call"
	entUser "calling-bill/ent/user"
	"calling-bill/internal/helpers"
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

type UserService struct {
	DB     *ent.Client
	Logger *zap.Logger
}

type CreateCallData struct {
	Duration int `json:"call_duration"`
}

func (s *UserService) AddCall(ctx context.Context, userId int, createCallData CreateCallData) error {

	newCall := s.DB.Call.
		Create().
		SetDuration(createCallData.Duration).
		SetBlockCount(helpers.CalculateBlockCount(createCallData.Duration)).
		SetUserID(userId)

	_, err := newCall.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

type BillingData struct {
	CallCount  int `sql:"call_count" json:"call_count"`
	BlockCount int `sql:"block_count" json:"block_count"`
}

func (s *UserService) GetBilling(ctx context.Context, userId int) (*BillingData, error) {
	var billingData []BillingData
	err := s.DB.Call.
		Query().
		Where(entCall.HasUserWith(entUser.ID(userId))).
		Modify(func(s *sql.Selector) {
			s.Select(
				sql.As(sql.Count("*"), "call_count"),
				sql.As(sql.Sum("block_count"), "block_count"),
			)
		}).
		Scan(ctx, &billingData)

	if err != nil {
		return nil, err
	}
	if len(billingData) != 1 {
		return nil, errors.New("error when get billing data")
	}
	return &billingData[0], nil
}

func (s *UserService) GetUserIDFromUsername(ctx context.Context, username string) (*int, error) {
	userId, err := s.DB.User.Query().Where(entUser.Username(username)).FirstID(ctx)
	if err != nil {
		s.Logger.Error(fmt.Sprintf("error when get userId of username: %s\n", username))
		return nil, err
	}

	return &userId, nil
}
