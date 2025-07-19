package repository

import (
	"context"
	"effective_mobile/internal/model"
	"effective_mobile/pkg/logger"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CreateSubscription(c context.Context, sub *model.Subscription) error {

	query := `INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id`

	err := r.db.QueryRow(c, query, sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, sub.EndDate).Scan(&sub.ID)
	if err != nil {
		logger.Error("failed to create subscription", zap.Error(err))
		return err
	}

	logger.Info("subscription created successfully")
	return nil
}

func (r *Repo) GetSubscriptionByID(c context.Context, id int64) (*model.Subscription, error) {

	query := `SELECT id, service_name, price, user_id, start_date, end_date
			FROM subscriptions WHERE id = $1`

	sub := &model.Subscription{}
	err := r.db.QueryRow(c, query, id).Scan(
		&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserID, &sub.StartDate, &sub.EndDate)
	if err != nil {
		logger.Error("failed to get subscription", zap.Error(err))
		return nil, err
	}

	logger.Info("subscription found successfully")
	return sub, nil
}

func (r *Repo) UpdateSubscription(c context.Context, sub *model.Subscription) error {

	query := `UPDATE subscriptions
			SET service_name=$1, price=$2, end_date=$3 WHERE id=$4`

	_, err := r.db.Exec(c, query, sub.ServiceName, sub.Price, sub.EndDate, sub.ID)
	if err != nil {
		logger.Error("failed to update subscription", zap.Error(err))
		return err
	}

	logger.Info("subscription updated successfully")
	return nil
}

func (r *Repo) DeleteSubscription(c context.Context, id int64) error {

	query := `DELETE FROM subscriptions WHERE id = $1`

	_, err := r.db.Exec(c, query, id)
	if err != nil {
		logger.Error("failed to delete subscription", zap.Error(err))
		return err
	}

	return nil
}

func (r *Repo) ListSubscription(c context.Context, filter *model.Filter) ([]model.Subscription, error) {

	query := `SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE 1=1`
	args := []interface{}{}
	argsIdx := 1

	if filter.UserID != nil {
		query += fmt.Sprintf(" AND user_id = $%d", argsIdx)
		args = append(args, *filter.UserID)
		argsIdx++
	}

	if filter.ServiceName != nil {
		query += fmt.Sprintf(" AND service_name = $%d", argsIdx)
		args = append(args, *filter.ServiceName)
		argsIdx++
	}

	if filter.Price != nil {
		query += fmt.Sprintf(" AND price = $%d", argsIdx)
		args = append(args, *filter.Price)
		argsIdx++
	}

	if filter.StartDate != nil {
		query += fmt.Sprintf(" AND start_date >= $%d", argsIdx)
		args = append(args, *filter.StartDate)
		argsIdx++
	}

	if filter.EndDate != nil {
		query += fmt.Sprintf(" AND (end_date IS NULL OR end_date <= $%d)", argsIdx)
		args = append(args, *filter.EndDate)
		argsIdx++
	}

	if filter.Limit != nil {
		query += fmt.Sprintf(" LIMIT $%d", argsIdx)
		args = append(args, *filter.Limit)
		argsIdx++
	}

	if filter.Offset != nil {
		query += fmt.Sprintf(" OFFSET $%d", argsIdx)
		args = append(args, *filter.Offset)
		argsIdx++
	}

	rows, err := r.db.Query(c, query, args...)
	if err != nil {
		logger.Error("failed sending request", zap.Error(err))
	}
	defer rows.Close()

	var subs []model.Subscription

	for rows.Next() {
		var sub model.Subscription
		err := rows.Scan(&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserID, &sub.StartDate, &sub.EndDate)
		if err != nil {
			logger.Error("failed receiving", zap.Error(err))
			return nil, err
		}
		subs = append(subs, sub)
	}

	return subs, nil
}

func (r *Repo) TotalCost(c context.Context, filter *model.Filter) (int, error) {

	query := `SELECT COALESCE(SUM(price), 0) FROM subscriptions WHERE 1=1`
	args := []interface{}{}
	argsIdx := 1

	if filter.UserID != nil {
		query += fmt.Sprintf(" AND user_id = $%d", argsIdx)
		args = append(args, *filter.UserID)
		argsIdx++
	}

	if filter.ServiceName != nil {
		query += fmt.Sprintf(" AND service_name = $%d", argsIdx)
		args = append(args, *filter.ServiceName)
		argsIdx++
	}

	if filter.StartDate != nil {
		query += fmt.Sprintf(" AND start_date >= $%d", argsIdx)
		args = append(args, *filter.StartDate)
		argsIdx++
	}

	if filter.EndDate != nil {
		query += fmt.Sprintf(" AND (end_date IS NULL OR end_date <= $%d)", argsIdx)
		args = append(args, *filter.EndDate)
		argsIdx++
	}

	var totalCost int
	err := r.db.QueryRow(c, query, args...).Scan(&totalCost)
	if err != nil {
		logger.Error("failed sending request", zap.Error(err))
		return 0, err
	}

	return totalCost, nil
}
