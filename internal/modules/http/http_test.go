package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := &config{Port: "8080"}
	mod := New(cfg)

	assert.NotNil(t, mod.router, "Router should not be nil")
	assert.Equal(t, cfg, mod.config, "Config should be correctly assigned")
}

func TestRegisterAndServeHTTP(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		handler      http.HandlerFunc
		requestPath  string
		expectedCode int
		method       string
	}{
		{
			name:         "Valid path GET request",
			path:         "/test",
			handler:      func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) },
			requestPath:  "/test",
			expectedCode: http.StatusOK,
			method:       "GET",
		},
		{
			name:         "Invalid path GET request",
			path:         "/test",
			handler:      func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) },
			requestPath:  "/invalid",
			expectedCode: http.StatusNotFound,
			method:       "GET",
		},
		{
			name:         "Valid path POST request",
			path:         "/post",
			handler:      func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusCreated) },
			requestPath:  "/post",
			expectedCode: http.StatusCreated,
			method:       "POST",
		},
		{
			name:         "Path with PUT request",
			path:         "/put",
			handler:      func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusNoContent) },
			requestPath:  "/put",
			expectedCode: http.StatusNoContent,
			method:       "PUT",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &config{Port: "8080"}
			mod := New(cfg)
			mod.Register(tt.path, tt.handler)

			req, err := http.NewRequest(tt.method, tt.requestPath, nil)
			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()
			mod.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expectedCode, recorder.Code, "Handler returned wrong status code for "+tt.name)
		})
	}
}
