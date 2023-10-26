package repo

import (
	"context"
	"fmt"

	"github.com/plusik10/anti-bruteforce/internal/usecase"
	"github.com/plusik10/anti-bruteforce/pkg/postgres"
)

var _ usecase.NetManagerRepo = (*NetManagerPostgresRepo)(nil)

type NetManagerPostgresRepo struct {
	*postgres.Postgres
}

func NewPostgres(pg *postgres.Postgres) *NetManagerPostgresRepo {
	return &NetManagerPostgresRepo{pg}
}

func (n *NetManagerPostgresRepo) InsertIP(ctx context.Context, ip string, isBlock bool) error {
	insert := n.Builder.Insert("ip_list")
	if isBlock {
		insert = insert.Columns("ip,block_ip").Values(ip, 1)
	} else {
		insert = insert.Columns("ip").Values(ip)
	}

	sql, args, err := insert.ToSql()
	if err != nil {
		return fmt.Errorf("NetManagerRepo - Upsert - n.Builder: %w", err)
	}
	_, err = n.Pool.Exec(ctx, sql, args)
	if err != nil {
		return fmt.Errorf("NetManagerRepo - Upsert - n.PoolExec: %w", err)
	}
	return nil
}

func (n *NetManagerPostgresRepo) RemoveIP(ctx context.Context, ip string) error {
	sql, args, err := n.Builder.Delete("").From("ip_list").Where("ip=", ip).ToSql()
	if err != nil {
		return fmt.Errorf("NetManagerRepo - RemoveIp - n.Builder: %w", err)
	}
	_, err = n.Pool.Exec(ctx, sql, args)
	if err != nil {
		return fmt.Errorf("NetManagerRepo - RemoveIp - n.PoolExec: %w", err)
	}
	return nil
}
