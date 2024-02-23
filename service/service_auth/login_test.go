package service_auth

import (
	"app/domain/mocks"
	authv1 "app/gen/proto/auth/v1"
	"app/model"
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestServiceAuth_Login(t *testing.T) {
	var mockRepoUser = new(mocks.RepositoryUser)
	var srv = NewService(&logrus.Logger{}, mockRepoUser)
	ctx := context.Background()
	username := "testuser"
	password := "1"
	id := uuid.NewString()

	mockRepoUser.On("DetailByUsername", ctx, username).Return(model.User{
		ID:       id,
		Username: username,
		Password: "$2a$04$.j8bw6aqScH/CeVXetnDQ.5s/6NZZubXxKrNAbthotJ2e7jl.8BXO",
	}, nil)

	req := &connect.Request[authv1.LoginRequest]{
		Msg: &authv1.LoginRequest{
			Username: username,
			Password: password,
		},
	}
	response, err := srv.Login(ctx, req)
	assert.NotEqual(t, "", response.Msg.Token)
	assert.NoError(t, err)
	mockRepoUser.AssertExpectations(t)
}
