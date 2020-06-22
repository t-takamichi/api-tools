package main

import (
	"testing"
)

func TestFilterUrlParameter(t *testing.T) {
	result := FilterUrlParameter("http://127.0.0.1?a=1&b=1")
	if v, ok := result["a"]; ok {
		if v != "1" {
			t.Errorf("a=1 = %s", v)
		}
	}
	if v, ok := result["b"]; ok {

		if v != "1" {
			t.Errorf("a=1 = %s", v)
		}
	}

}
