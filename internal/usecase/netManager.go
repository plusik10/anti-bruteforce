package usecase

import (
	"context"
	"github.com/plusik10/anti-bruteforce/internal/entity"
)

var _ NetManager = (*NetManagerUsecase)(nil)

type NetManagerUsecase struct {
	repo NetManagerRepo
}

func (n NetManagerUsecase) Auth(ctx context.Context, net entity.Net) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NetManagerUsecase) AddIpToBlackList(ctx context.Context, ip string) error {
	err := n.repo.InsertIP(ctx, ip, true)
	if err != nil {
		return err
	}
	return nil
}

func (n *NetManagerUsecase) DeleteFromBlackList(ctx context.Context, ip string) error {
	err := n.repo.RemoveIP(ctx, ip)
	if err != nil {
		return err
	}
	return nil
}

func (n NetManagerUsecase) AddIpToWhiteList(ctx context.Context, ip string) error {
	err := n.repo.InsertIP(ctx, ip, false)
	if err != nil {
		return err
	}
	return nil
}

func (n NetManagerUsecase) DeleteFromWhiteList(ctx context.Context, ip string) error {
	err := n.repo.RemoveIP(ctx, ip)
	if err != nil {
		return err
	}
	return nil
}

func (n NetManagerUsecase) ClearBucket() error {
	//TODO implement me
	panic("implement me")
}

func NewNetManagerUsecase(repo NetManagerRepo) *NetManagerUsecase {
	return &NetManagerUsecase{repo: repo}
}
