package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func init() {
	Routes()
}
func TestRoutes(t *testing.T) {
	t.Log("start........")
	{
		request, e := http.NewRequest("GET", "/sendjson", nil)
		if e != nil {
			t.Fatal(e)
		}
		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, request)

		if rw.Code != 200 {
			t.Fatal("not 200", rw.Code)
		}
		u := struct {
			Name  string
			Email string
		}{}
		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal(err)
		}
		if u.Name != "Bill" {
			t.Error(u.Name)
		}
		if u.Email != "bill@bili.com" {
			t.Error(u.Email)
		}
	}
}
