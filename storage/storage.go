package storage

import (
	"context"
	"user/api/models"
)

type IStorage interface {
	CloseDB()
	UserStorage() UserStorage
}

type UserStorage interface {
	Create(ctx context.Context, user models.AddUser) (int64, error)
	CreateUsers(ctx context.Context, users models.AddUsers) (error)
	// Update(ctx context.Context, user models.User) (int64, error)
	// UpdateUsers(ctx context.Context, users models.UpdateUsers) error
	// Delete(ctx context.Context, id int64) error
	// DeleteUsers(ctx context.Context, ids models.DeleteUsers) error
	// GetById(ctx context.Context, id int64) (models.User, error)
	// GetList(ctx context.Context, req models.GetListRequest) (models.GetListResponse, error)
}