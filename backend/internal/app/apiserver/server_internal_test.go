package apiserver

import (
	"bytes"
	"encoding/json"
	"github.com/LezhnevSergei/news_aggregator/internal/app/models"
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
}

func TestServer_GetList(t *testing.T) {
	s := NewServer(teststore.New())
	news := map[string]string{
		"title":      "TestTitle",
		"created_at": "2017-11-20T11:20:10+01:00",
	}
	t.Run("test get list", func(t *testing.T) {
		rec := httptest.NewRecorder()
		b := &bytes.Buffer{}
		json.NewEncoder(b).Encode(news)
		req, err := http.NewRequest(http.MethodGet, "/news", b)

		s.ServeHTTP(rec, req)

		n := models.News{}
		err = json.Unmarshal(b.Bytes(), &n)

		assert.NoError(t, err)
		assert.Equal(t, n.Title, "TestTitle")
	})
}

func TestServer_GetNewsByTitle(t *testing.T) {
	s := NewServer(teststore.New())
	news := map[string]string{
		"title":      "TestTitle",
		"created_at": "2017-11-20T11:20:10+01:00",
	}
	t.Run("test get news by title", func(t *testing.T) {
		rec := httptest.NewRecorder()
		b := &bytes.Buffer{}
		json.NewEncoder(b).Encode(news)
		req, err := http.NewRequest(http.MethodGet, "/news?search=Test", b)

		s.ServeHTTP(rec, req)

		n := models.News{}
		err = json.Unmarshal(b.Bytes(), &n)

		assert.NoError(t, err)
		assert.Equal(t, n.Title, "TestTitle")
	})
}
