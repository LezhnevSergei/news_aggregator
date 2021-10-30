package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/LezhnevSergei/news_aggregator/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_HandleNewsCreate(t *testing.T) {
	s := NewServer(teststore.New())

	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"title":      "TestTitle",
				"created_at": "2017-11-20T11:20:10+01:00",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid payload",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/news", b)

			s.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}

	//rec := httptest.NewRecorder()
	//req, _ := http.NewRequest(http.MethodPost, "/news", nil)
	//s := NewServer(teststore.New())
	//s.ServeHTTP(rec, req)
	//
	//assert.Equal(t, rec.Code, http.StatusOK)
}
