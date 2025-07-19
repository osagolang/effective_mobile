package service

import (
	"context"
	"effective_mobile/internal/model"
	"effective_mobile/internal/repository"
	"effective_mobile/pkg/logger"
	"go.uber.org/zap"
	"time"
)

type Service struct {
	repo *repository.Repo
}

func NewService(repo *repository.Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateSubscription(c context.Context, request *model.CreateSubscriptionRequest) (*model.Subscription, error) {

	startDate, err := time.Parse("01-2006", request.StartDate)
	if err != nil {
		logger.Warn("incorrect start date", zap.Error(err))
		return nil, err
	}

	sub := &model.Subscription{
		ServiceName: request.ServiceName,
		Price:       request.Price,
		UserID:      request.UserID,
		StartDate:   startDate,
	}

	if err := s.repo.CreateSubscription(c, sub); err != nil {
		logger.Error("error sending model", zap.Error(err))
		return nil, err
	}
	return sub, nil
}

func (s *Service) UpdateSubscription(c context.Context, id int64, request *model.UpdateSubscriptionRequest) (*model.Subscription, error) {

	sub, err := s.repo.GetSubscriptionByID(c, id)
	if err != nil {
		return nil, err
	}
	if sub == nil {
		return nil, nil
	}

	if request.ServiceName != nil {
		sub.ServiceName = *request.ServiceName
	}

	if request.Price != nil {
		sub.Price = *request.Price
	}

	if request.EndDate != nil {
		endDate, err := time.Parse("01-2006", *request.EndDate)
		if err != nil {
			logger.Warn("incorrect end date", zap.Error(err))
		}
		sub.EndDate = &endDate
	}

	if err := s.repo.UpdateSubscription(c, sub); err != nil {
		logger.Error("error sending model", zap.Error(err))
		return nil, err
	}

	return sub, nil
}
