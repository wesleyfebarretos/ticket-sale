package gateway_service

import (
	"context"
	"errors"

	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/repository/implementation/gateway_repository"
)

func GetActive(ctx context.Context) (*gateway_repository.GetActiveResponse, error) {
	gateway := gateway_repository.New().GetActive(ctx)

	if gateway == nil {
		return nil, errors.New("Non active gateway")
	}

	return gateway, nil
}
