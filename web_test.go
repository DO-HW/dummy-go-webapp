package main

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T){

    req, _ := http.NewRequest("GET", "", nil)
    w := httptest.NewRecorder()

    hello(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Home page didn't return %v", http.StatusOK)
    }

    assert.Equal(t, w.Body.String(), "Greeting World!", "body should be Hello World!");

}
