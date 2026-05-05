// cmd/api/main_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
    got := add(2, 3)
    if got != 5 {
        t.Errorf("want 5, got %d", got)
    }
}
