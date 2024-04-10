package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, want, got T) {
	t.Helper()

	if want != got {
		t.Errorf("want %v got %v", want, got)
	}
}

func StringContains(t *testing.T, want, got string) {
	t.Helper()

	if !strings.Contains(got, want) {
		t.Errorf("got: %q; expected to contain: %q", got, want)
	}
}
