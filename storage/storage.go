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
	CreateMany(ctx context.Context, users models.AddUsers) error
	Update(ctx context.Context, user models.UpdateUser) (int64, error)
	UpdateMany(ctx context.Context, users models.UpdateUsers) error
	Delete(ctx context.Context, user models.DeleteUser) error
	DeleteMany(ctx context.Context, ids models.DeleteUsers) error
	GetById(ctx context.Context, id int64) (models.GetUser, error)
	GetList(ctx context.Context, req models.GetListRequest) (models.GetListResponse, error)
}