package usecase

import (
	"context"

	"github.com/plusik10/anti-bruteforce/internal/entity"
)

type (
	NetManager interface {
		Auth(ctx context.Context, net entity.Net) (bool, error)
		AddIpToBlackList(ctx context.Context, ip string) error
		DeleteFromBlackList(ctx context.Context, ip string) error
		AddIpToWhiteList(ctx context.Context, ip string) error
		DeleteFromWhiteList(ctx context.Context, ip string) error
		ClearBucket() error
	}

	NetManagerRepo interface {
		InsertIP(ctx context.Context, ip string, isBlock bool) error
		RemoveIP(ctx context.Context, ip string) error
	}
)
