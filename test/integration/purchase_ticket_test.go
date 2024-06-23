package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/duducv/go-clean-arch/config"
	"github.com/duducv/go-clean-arch/internal/infra/rest/routes"
	"github.com/stretchr/testify/assert"
)

func TestPurchaseTicketSuccess(t *testing.T) {
	containerConfig := config.NewTestContainerSetupConfig(t)
	defer containerConfig.CleanUp()
	input := map[string]any{
		"eventId":         "clxnut5z2000008lg9aombhm6",
		"email":           "john.doe@gmail.com",
		"creditCardToken": "987654321",
	}
	inputAsBytes, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/purchase_ticket", bytes.NewBuffer(inputAsBytes))
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	router := routes.ConfigRouter()
	routes.ApplyRoutes(router, containerConfig.Adapters)
	router.ServeHTTP(w, req)
	responseBody := map[string]any{}
	if err := json.NewDecoder(w.Body).Decode(&responseBody); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, w.Code, http.StatusCreated)
	assert.Equal(t, responseBody["status"], "approved")
	assert.NotEmpty(t, responseBody["ticketId"])
}
