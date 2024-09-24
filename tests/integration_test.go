package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAdmin(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/admins", "application/json", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
