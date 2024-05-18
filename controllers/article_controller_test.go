package controllers_test

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestArticleListHandler(t *testing.T) {
	var tests = []struct {
		name       string
		query      string
		resultCode int
	}{
		{name: "query num", query: "1", resultCode: http.StatusOK},
		{name: "query noNum", query: "a", resultCode: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("/article/list?page=%s", tt.query)
			req := httptest.NewRequest("GET", url, nil)
			w := httptest.NewRecorder()
			r := mux.NewRouter()
			r.HandleFunc("/article/list", aCon.GetArticleListHandler)
			r.ServeHTTP(w, req)
			if w.Code != tt.resultCode {
				t.Errorf("got %d, want %d", w.Code, tt.resultCode)
			}
		})
	}
}
