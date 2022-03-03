package main

import (
	"context"
	"testing"
)

func TestInvoke(t *testing.T) {
	ctx := context.Background()
	if err := Invoke(ctx); err != nil {
		t.Error(err)
	}
}
