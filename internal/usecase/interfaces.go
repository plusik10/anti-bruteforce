package usecase

import (
	"context"

	"github.com/plusik10/anti-bruteforce/internal/entity"
)

type (
	NetManager interface {
		Auth(ctx context.Context, net entity.Net) (bool, error)
		AddIPToBlackList(ctx context.Context, ip string) error
		DeleteIPFromStorage(ctx context.Context, ip string) error
		AddIPToWhiteList(ctx context.Context, ip string) error
		ClearBucket() error
	}

	NetManagerRepo interface {
		CheckIPToWhiteList(ctx context.Context, ip string) error
		CheckIPToBlackList(ctx context.Context, ip string) error
		InsertIP(ctx context.Context, ip string, isBlock bool) error
		RemoveIP(ctx context.Context, ip string) error
	}
)
