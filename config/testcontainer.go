package config

import (
	"context"
	"fmt"
	"os"
	"testing"

	gatewayadapters "github.com/duducv/go-clean-arch/internal/infra/gateway"
	persistenceadapters "github.com/duducv/go-clean-arch/internal/infra/persistence"
	"github.com/duducv/go-clean-arch/internal/infra/tracing"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestContainerSetupConfig struct {
	Ctx      context.Context
	Adapters *RepositoryAdapters
	CleanUp  func()
}

func NewTestContainerSetupConfig(t *testing.T) *TestContainerSetupConfig {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForAll(
			wait.ForLog("database system is ready to accept connections"),
			wait.ForListeningPort("5432/tcp"),
		),
	}
	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}
	initSql, err := os.ReadFile("../../build/database/schema/init.sql")
	if err != nil {
		t.Fatal("Erro ao ler o arquivo init.sql:", err)
	}
	tmpFile, err := os.CreateTemp("", "init-*.sql")
	if err != nil {
		t.Fatal("Erro ao criar arquivo temporário:", err)
	}
	defer os.Remove(tmpFile.Name())
	if _, err := tmpFile.Write(initSql); err != nil {
		t.Fatal("Erro ao escrever no arquivo temporário:", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatal("Erro ao fechar o arquivo temporário:", err)
	}
	if err := postgresContainer.CopyFileToContainer(ctx, tmpFile.Name(), "/docker-entrypoint-initdb.d/init.sql", 0644); err != nil {
		t.Fatal("Erro ao copiar o arquivo para o contêiner:", err)
	}
	_, _, err = postgresContainer.Exec(ctx, []string{"psql", "-U", "user", "-d", "testdb", "-f", "/docker-entrypoint-initdb.d/init.sql"})
	if err != nil {
		t.Fatal("Erro ao executar o script SQL:", err)
	}
	port, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatal(err)
	}
	dbURL := fmt.Sprintf("postgres://user:password@localhost:%s/testdb?sslmode=disable", port.Port())
	pgConn := OpenPGXConnection(dbURL)
	tracingAdapter := tracing.NewNotracingAdapter()
	adapters := NewRepositoryAdapters(
		persistenceadapters.NewTicketRepositorySQLAdapter(pgConn, tracingAdapter),
		persistenceadapters.NewEventRepositorySQLAdapter(pgConn, tracingAdapter),
		persistenceadapters.NewTransactionRepositorySQLAdapter(pgConn, tracingAdapter),
		gatewayadapters.NewFakePaymentGateway(),
		persistenceadapters.NewUnitOfWorkSQLAdapter(*pgConn),
		tracingAdapter,
	)
	return &TestContainerSetupConfig{
		Ctx:      ctx,
		Adapters: adapters,
		CleanUp: func() {
			postgresContainer.Terminate(ctx)
		},
	}
}
