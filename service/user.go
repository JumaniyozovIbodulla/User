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

func (u userService) CreateMany(ctx context.Context, users models.AddUsers) error {

	err := u.storage.UserStorage().CreateMany(ctx, users)
	if err != nil {
		u.logger.Error("failed to create new users: ", logger.Error(err))
		return err
	}
	return nil
}

func (u userService) Update(ctx context.Context, user models.UpdateUser) (int64, error) {
	id, err := u.storage.UserStorage().Update(ctx, user)

	if err != nil {
		u.logger.Error("failed to update an user: ", logger.Error(err))
		return id, err
	}
	return id, nil
}

func (u userService) UpdateMany(ctx context.Context, users models.UpdateUsers) error {
	err := u.storage.UserStorage().UpdateMany(ctx, users)

	if err != nil {
		u.logger.Error("failed to update users: ", logger.Error(err))
		return err
	}
	return nil
}

func (u userService) Delete(ctx context.Context, user models.DeleteUser) error {
	err := u.storage.UserStorage().Delete(ctx, user)

	if err != nil {
		u.logger.Error("failed to delete a user: ", logger.Error(err))
		return err
	}
	return nil
}


func (u userService) DeleteMany(ctx context.Context, users models.DeleteUsers) error {

	err := u.storage.UserStorage().DeleteMany(ctx, users)
	if err != nil {
		u.logger.Error("failed to delete users: ", logger.Error(err))
		return err
	}
	return nil
}

func (u userService) GetById(ctx context.Context, id int64) (models.GetUser, error) {

	user, err := u.storage.UserStorage().GetById(ctx, id)
	if err != nil {
		u.logger.Error("failed to get an user: ", logger.Error(err))
		return user, err
	}
	return user, nil
}

func (u userService) GetList(ctx context.Context, req models.GetListRequest) (models.GetListResponse, error) {

	users, err := u.storage.UserStorage().GetList(ctx, req)
	if err != nil {
		u.logger.Error("failed to get users: ", logger.Error(err))
		return users, err
	}
	return users, nil
}