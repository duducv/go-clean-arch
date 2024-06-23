package persistenceadapters

import (
	"context"

	"github.com/duducv/go-clean-arch/internal/core/application/constants"
	applicationrepository "github.com/duducv/go-clean-arch/internal/core/application/repository"
	"github.com/jackc/pgx/v5"
)

type UnitOfWorkSQLAdapter struct {
	pgConn pgx.Tx
}

func (uow *UnitOfWorkSQLAdapter) Begin(ctx context.Context) (context.Context, error) {
	var err error
	uow.pgConn, err = uow.pgConn.Conn().BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, constants.TXKey, &uow.pgConn), nil
}

func (uow *UnitOfWorkSQLAdapter) Commit(ctx context.Context) error {
	return uow.pgConn.Commit(ctx)
}

func (uow *UnitOfWorkSQLAdapter) Rollback(ctx context.Context) error {
	return uow.pgConn.Rollback(ctx)
}

func NewUnitOfWorkSQLAdapter(pgConn pgx.Tx) applicationrepository.UnitOfWorkRepository {
	return &UnitOfWorkSQLAdapter{
		pgConn: pgConn,
	}
}
