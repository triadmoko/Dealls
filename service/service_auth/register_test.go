package service_auth

import (
	"app/domain/mocks"
	authv1 "app/gen/proto/auth/v1"
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestServiceAuth_Register(t *testing.T) {
	var mockRepoUser = new(mocks.RepositoryUser)
	var srv = NewService(&logrus.Logger{}, mockRepoUser)
	ctx := context.Background()

	req := &connect.Request[authv1.RegisterRequest]{
		Msg: &authv1.RegisterRequest{
			Username: "testuser",
			Password: "testpassword",
		},
	}
	res, err := srv.Register(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
