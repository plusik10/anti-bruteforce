package usecase

import (
	"context"

	"github.com/plusik10/anti-bruteforce/internal/entity"
)

var _ NetManager = (*NetManagerUsecase)(nil)

type NetManagerUsecase struct {
	repo NetManagerRepo
}

func (n NetManagerUsecase) DeleteIPFromStorage(ctx context.Context, ip string) error {
	err := n.repo.RemoveIP(ctx, ip)
	if err != nil {
		return err
	}
	return nil
}

func (n NetManagerUsecase) Auth(ctx context.Context, net entity.Net) (bool, error) {
	err := n.repo.CheckIPToWhiteList(ctx, net.IP)
	if err != nil {
		return false, err
	}
	return true, nil
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
