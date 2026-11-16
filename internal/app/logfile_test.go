package app

import (
	"os"
	"path/filepath"
	"testing"
)

// A log grown past its limit is moved aside at launch, replacing the one
// moved there before, and a new one begun; a smaller one is added to.
func TestOpenLogMovesBigLogAside(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, logName)
	os.WriteFile(path+".1", []byte("oldest"), 0o644)
	os.WriteFile(path, make([]byte, logLimit+1), 0o644)
	f, err := openLog(dir)
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString("new\n")
	f.Close()
	if b, _ := os.ReadFile(path); string(b) != "new\n" {
		t.Errorf("log begun with %d bytes", len(b))
	}
	if info, _ := os.Stat(path + ".1"); info == nil || info.Size() != logLimit+1 {
		t.Error("the big log was not moved aside")
	}

	f, _ = openLog(dir)
	f.WriteString("more\n")
	f.Close()
	if b, _ := os.ReadFile(path); string(b) != "new\nmore\n" {
		t.Errorf("small log not added to: %q", b)
	}
}
