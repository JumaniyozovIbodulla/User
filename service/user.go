package service

import (
	"context"
	"user/api/models"
	"user/pkg/logger"
	"user/storage"
)

type userService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewUserService(storage storage.IStorage, logger logger.ILogger) userService {
	return userService{
		storage: storage,
		logger:  logger,
	}
}

func (u userService) Create(ctx context.Context, user models.AddUser) (int64, error) {

	id, err := u.storage.UserStorage().Create(ctx, user)
	if err != nil {
		u.logger.Error("failed to create a new user: ", logger.Error(err))
		return 0, err
	}
	return id, nil
}

func (u userService) CreateUsers(ctx context.Context, users models.AddUsers) error {

	err := u.storage.UserStorage().CreateUsers(ctx, users)
	if err != nil {
		u.logger.Error("failed to create new users: ", logger.Error(err))
		return err
	}
	return nil
}