package usecase

import (
	"context"

	"github.com/plusik10/anti-bruteforce/internal/entity"
)

var _ NetManager = (*NetManagerUsecase)(nil)

type NetManagerUsecase struct {
	repo NetManagerRepo
}

func (n NetManagerUsecase) DeleteIpFromStorage(ctx context.Context, ip string) error {
	if err := n.repo.RemoveIP(ctx, ip); err != nil {
		return err
	}
	return nil
}

func (n NetManagerUsecase) Auth(ctx context.Context, net entity.Net) (bool, error) {
	_, _ = ctx, net // TODO: REMOVE
	panic("implement me")
}

func (n *NetManagerUsecase) AddIPToBlackList(ctx context.Context, ip string) error {
	err := n.repo.InsertIP(ctx, ip, true)
	if err != nil {
		return err
	}
	return nil
}

func (n NetManagerUsecase) AddIPToWhiteList(ctx context.Context, ip string) error {
	err := n.repo.InsertIP(ctx, ip, false)
	if err != nil {
		return err
	}
	return nil
}

func (n NetManagerUsecase) ClearBucket() error {
	panic("implement me")
}

func NewNetManagerUsecase(repo NetManagerRepo) *NetManagerUsecase {
	return &NetManagerUsecase{repo: repo}
}
