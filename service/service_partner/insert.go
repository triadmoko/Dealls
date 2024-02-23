package service_partner

import (
	partnerv1 "app/gen/proto/partner/v1"
	"context"

	"connectrpc.com/connect"
)

func (s *ServicePartner) SwipePartner(ctx context.Context, req *connect.Request[partnerv1.RequestSwipePartner]) (*connect.Response[partnerv1.ResponseSwipePartner], error) {
	return nil, nil
}
