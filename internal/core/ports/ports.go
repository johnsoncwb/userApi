package ports

import (
	"context"
	"github.com/johnsoncwb/userApi/internal/core/domain"
)

type WebExternal interface {
	GetAll(context.Context) ([]domain.User, error)
	GetById(context.Context) (domain.User, error)
}

type UserService interface {
	GetAll(context.Context) ([]domain.User, error)
	GetById(context.Context) (domain.User, error)
}
