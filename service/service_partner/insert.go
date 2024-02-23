package service_partner

import (
	"app/constant"
	partnerv1 "app/gen/proto/partner/v1"
	"app/model"
	"app/pkg"
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
)

func (s *ServicePartner) SwipePartner(ctx context.Context, req *connect.Request[partnerv1.RequestSwipePartner]) (*connect.Response[partnerv1.ResponseSwipePartner], error) {
	userMeta, ok := ctx.Value("user").(pkg.MetaToken)
	if !ok {
		return nil, connect.NewError(connect.CodeUnauthenticated, constant.ErrAuthorization)
	}

	result, err := s.repoPartner.Create(ctx, model.Interest{
		UserID:         userMeta.ID,
		ID:             uuid.NewString(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		DeletedAt:      nil,
		IsInterest:     req.Msg.IsInterest,
		InterestUserID: req.Msg.PartnerId,
	})
	if err != nil {
		s.logger.Error("s.repoPartner.Create", err)
		return nil, connect.NewError(connect.CodeInternal, constant.ErrInternalServer)
	}
	pkg.Prt(result)
	return nil, nil
}
