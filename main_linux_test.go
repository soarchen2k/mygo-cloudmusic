package main

import (
	"slices"
	"testing"
)

// Starting again with the app's libraries puts them before those the
// system was told of, and keeps what LD_LIBRARY_PATH was, unset or not, to
// put back once started.
func TestWithLibs(t *testing.T) {
	got := withLibs([]string{"HOME=/h", "LD_LIBRARY_PATH=/opt/x", "LANG=C"}, "/app/lib")
	want := []string{"HOME=/h", "LANG=C", "LD_LIBRARY_PATH=/app/lib:/opt/x", relaunched + "=/opt/x"}
	if !slices.Equal(got, want) {
		t.Errorf("with a path: %q", got)
	}
	got = withLibs([]string{"HOME=/h"}, "/app/lib")
	want = []string{"HOME=/h", "LD_LIBRARY_PATH=/app/lib", relaunched + "="}
	if !slices.Equal(got, want) {
		t.Errorf("without one: %q", got)
	}
}
