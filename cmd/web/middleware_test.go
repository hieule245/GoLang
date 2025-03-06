package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurve(t *testing.T) {
	var myH MyHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do something
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var myH MyHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do something
	default:
		t.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}
