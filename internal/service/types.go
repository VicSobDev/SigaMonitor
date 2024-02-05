package service

import (
	"context"

	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
	ctx    context.Context
}

func NewService(ctx context.Context) (*Service, error) {

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Service{
		logger: logger,
		ctx:    ctx,
	}, nil
}
