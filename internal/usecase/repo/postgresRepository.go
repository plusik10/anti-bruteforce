package repo

import (
	"context"
	"fmt"
	"log"

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

func (n *NetManagerPostgresRepo) CheckIPToWhiteList(ctx context.Context, ip string) error {
	sql, _, err := n.Builder.
		Select("count(1)").
		From("ip_list").
		Where("ip=?", ip).
		Where("block_ip=0").
		ToSql()
	if err != nil {
		return fmt.Errorf("NetManagerRepo CheckIPToWhiteList - n.Builder: %w", err)
	}

	var count int

	row := n.Pool.QueryRow(ctx, sql)
	err = row.Scan(&count)
	if err != nil {
		return fmt.Errorf("NetManagerRepo CheckIPToWhiteList - row.Scan: %w, sql: %s", err, sql)
	}

	if count == 0 {
		return fmt.Errorf("IP not found in whitelist")
	}
	return nil
}

func (n *NetManagerPostgresRepo) CheckIPToBlackList(ctx context.Context, ip string) error {
	panic("Not implement")
}

func (n *NetManagerPostgresRepo) InsertIP(ctx context.Context, ip string, isBlock bool) error {
	var isBlockInt int
	if isBlock {
		isBlockInt = 1
	}

	sql, args, err := n.Builder.
		Insert("ip_list").
		Columns("ip, block_ip").
		Values(ip, isBlockInt).ToSql()
	if err != nil {
		return fmt.Errorf("NetManagerRepo - Upsert - n.Builder: %w", err)
	}

	_, err = n.Pool.Exec(ctx, sql, args[0], args[1])
	if err != nil {
		return fmt.Errorf("NetManagerRepo - Upsert - n.PoolExec: %w query: %s ", err, sql)
	}
	return nil
}

func (n *NetManagerPostgresRepo) RemoveIP(ctx context.Context, ip string) error {
	log.Println(ip)
	sql, args, err := n.Builder.
		Delete("").
		From("ip_list").
		Where("ip=?", ip).ToSql()
	if err != nil {
		return fmt.Errorf("NetManagerRepo - RemoveIp - n.Builder: %w, sql = %s", err, sql)
	}

	_, err = n.Pool.Exec(ctx, sql, args[0])
	if err != nil {
		return fmt.Errorf("NetManagerRepo - RemoveIp - n.PoolExec: %w", err)
	}
	return nil
}
