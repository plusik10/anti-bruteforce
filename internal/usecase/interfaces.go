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
		Upsert(ctx context.Context, ip string, isBlack bool) error
		RemoveIp(ctx context.Context, ip string)
	}
)
