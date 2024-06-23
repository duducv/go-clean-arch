package dto_test

import (
	"testing"

	"github.com/duducv/go-clean-arch/internal/core/application/constants"
	"github.com/duducv/go-clean-arch/internal/core/application/dto"
	"github.com/stretchr/testify/assert"
)

func TestNewErrorOutput(t *testing.T) {
	dtoErrors := []string{"Not Found."}
	errorOutPut := dto.NewErrorOutput("sql: row not found.", constants.DomainLayer, 500, dtoErrors...)

	assert.Equal(t, errorOutPut.Raw, "sql: row not found.")
	assert.Equal(t, errorOutPut.Layer, constants.DomainLayer)
	assert.Equal(t, errorOutPut.StatusCode, 500)
	assert.Len(t, errorOutPut.Message, 1)
	assert.Contains(t, errorOutPut.Message, "Not Found.")
}
