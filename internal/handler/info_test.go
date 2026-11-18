package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {
	v := "/info/v1"
	handler := http.HandlerFunc(GetHealth)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", v, nil)

	handler.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

}
