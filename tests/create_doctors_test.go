package main

import (
	"bytes"
	doctor "crud-api/models/doctor"
	routes "crud-api/routes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDoctors(t *testing.T) {
	router := routes.SetupRouter()
	doctor := doctor.NewDoctor("João Silva", "Heart Surgery", 42, true)
	doctorJson, _ := json.Marshal(doctor)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/doctors", bytes.NewBuffer(doctorJson))
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}
