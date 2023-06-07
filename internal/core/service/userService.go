package service

import (
	"context"

	"github.com/johnsoncwb/userApi/internal/adapters/external"
	"github.com/johnsoncwb/userApi/internal/core/domain"
	logutils "github.com/johnsoncwb/userApi/internal/utils/logUtils"
)

type UserService struct {
	repo external.Requests
}

func (u *UserService) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := u.repo.GetAll(ctx)
	if err != nil {
		// config.GlobalConfig.NewRelicLogger.Error().Fields(map[string]interface{}{
		// 	"message": "erro ao requisitar todos os usuários",
		// 	"local":   "service",
		// }).Err(err)
		return []domain.User{}, err
	}

	// config.GlobalConfig.NewRelicLogger.Info().Fields(map[string]interface{}{
	// 	"total": len(users),
	// }).Msg("total de usuários encontrados")

	return users, nil

}

func (u *UserService) GetById(ctx context.Context) (domain.User, error) {

	logger := logutils.LoggerFromContext(ctx)

	user, err := u.repo.GetById(ctx)
	if err != nil {
		// config.GlobalConfig.NewRelicLogger.Error().Fields(map[string]interface{}{
		// 	"message": "erro ao requisitar usuário",
		// 	"local":   "service",
		// }).Err(err)
		return domain.User{}, err
	}

	logger.WithFields(map[string]interface{}{
		"id":       user.Id,
		"name":     user.Name,
		"username": user.Username,
		"address":  user.Address,
		"phone":    user.Phone,
		"website":  user.Website,
		"company":  user.Company,
	}).Info("dados do usuário encontrado")

	return user, nil
}

func NewUserService(repo external.Requests) *UserService {
	return &UserService{repo: repo}
}
