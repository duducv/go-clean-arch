package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func OpenPGXConnection(dbUrl string) *pgx.Tx {
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatalf("unable to connect to postgres: %v", err)
	}
	tx := NewPgxConn(conn)
	return &tx
}

type PgxConn struct {
	conn *pgx.Conn
}

func NewPgxConn(conn *pgx.Conn) pgx.Tx {
	return PgxConn{
		conn: conn,
	}
}

func (p PgxConn) Begin(ctx context.Context) (pgx.Tx, error) {
	return p, nil
}

func (p PgxConn) Commit(ctx context.Context) error {
	return nil
}

func (p PgxConn) Rollback(ctx context.Context) error {
	return nil
}

func (p PgxConn) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return p.conn.CopyFrom(ctx, tableName, columnNames, rowSrc)
}

func (p PgxConn) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return p.conn.SendBatch(ctx, b)
}

func (p PgxConn) LargeObjects() pgx.LargeObjects {
	return pgx.LargeObjects{}
}

func (p PgxConn) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	return p.conn.Prepare(ctx, name, sql)
}

func (p PgxConn) Exec(ctx context.Context, sql string, arguments ...any) (commandTag pgconn.CommandTag, err error) {
	return p.conn.Exec(ctx, sql, arguments...)
}

func (p PgxConn) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return p.conn.Query(ctx, sql, args...)
}

func (p PgxConn) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return p.conn.QueryRow(ctx, sql, args...)
}

func (p PgxConn) Conn() *pgx.Conn {
	return p.conn
}
