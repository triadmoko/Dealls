//go:build wireinject
// +build wireinject

package injector

import (
	"app/domain"
	"app/repository/repo_user"
	"app/service/service_auth"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)
var newSetUserRepository = wire.NewSet(
	repo_user.NewRepository,
	wire.Bind(
		new(domain.RepositoryUser),
		new(*repo_user.Repository),
	),
)

func InitializedAuth(db *gorm.DB, logger *logrus.Logger) *service_auth.ServiceAuth {
	wire.Build(newSetUserRepository, service_auth.NewService)
	return nil
}
