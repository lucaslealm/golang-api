package main

import (
	routes "crud-api/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDoctors(t *testing.T) {
	router := routes.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/doctors", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
