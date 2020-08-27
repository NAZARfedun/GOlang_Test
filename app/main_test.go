package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Nazar_Test/app/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestPostPlayersListFail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockLogger := mocks.NewMockLogger(mockCtrl)
	mockLogger.EXPECT().Errorln("invalid character '{' looking for beginning of object key string")

	handlers := Handlers{mockLogger}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	bodyReader := strings.NewReader(`{{}}`)
	req, err := http.NewRequest("POST", "/PostPlayersList", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	r := *mux.NewRouter()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.PostPlayersList)
	r.Handle("/PostPlayersList", handler).Methods("POST")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	if rr.Body.String() != `invalid character '{' looking for beginning of object key string` {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "invalid character '{' looking for beginning of object key string")
	}
}

func TestPostPlayersListSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	mockLogger := mocks.NewMockLogger(mockCtrl)

	handlers := Handlers{mockLogger}

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	bodyReader := strings.NewReader(`{
		"playerList" : [
			{
				"Nickname" : "Nazar",
				"Score"    : 4
			},
			{
				"Nickname" : "Misha",
				"Score"    : 1
			}
		]
	}`)
	req, err := http.NewRequest("POST", "/PostPlayersList", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	r := *mux.NewRouter()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.PostPlayersList)
	r.Handle("/PostPlayersList", handler).Methods("POST")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	if rr.Body.String() != `{[{Misha 1} {Nazar 4}]}` {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), `{[{Misha 1} {Nazar 4}]}`)
	}
}

func TestPostValueFail(t *testing.T) {
	mockCrtl := gomock.NewController(t)
	mockLogger := mocks.NewMockLogger(mockCrtl)
	mockLogger.EXPECT().Errorln("invalid character '{' looking for beginning of object key string")

	handlers := Handlers{mockLogger}

	bodyReader := strings.NewReader(`{{}}`)
	req, err := http.NewRequest("POST", "/PostValue", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	r := *mux.NewRouter()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.PostValue)
	r.Handle("/PostValue", handler).Methods("POST")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	if rr.Body.String() != `invalid character '{' looking for beginning of object key string` {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), `invalid character '{' looking for beginning of object key string`)
	}
}

func TestPostValueSuccess(t *testing.T) {
	mockCrtl := gomock.NewController(t)
	mockLogger := mocks.NewMockLogger(mockCrtl)

	handlers := Handlers{mockLogger}

	bodyReader := strings.NewReader(`{
		"Nickname" : "Nazar",
		"Score"    : 343
	}`)
	req, err := http.NewRequest("POST", "/PostValue", bodyReader)
	if err != nil {
		t.Fatal(err)
	}

	r := *mux.NewRouter()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.PostValue)
	r.Handle("/PostValue", handler).Methods("POST")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	if rr.Body.String() != `{Nazar 343}` {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), `{Nazar 343}`)
	}
}
