package service

import (
	"user/pkg/logger"
	"user/storage"
)

type IServiceManager interface {
	User() userService
}

type Service struct {
	userService  userService
	logger          logger.ILogger
}

func New(storage storage.IStorage, logger logger.ILogger) Service {
	services := Service{}
	services.userService = NewUserService(storage, logger)
	services.logger = logger

	return services
}

func (s Service) User() userService {
	return s.userService
}