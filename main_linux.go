package main

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/ebitengine/purego"
)

// trayLib is the library of the tray icon, which the app ships in lib next
// to its executable (scripts/tray-libs.sh) for systems without it.
const trayLib = "libayatana-appindicator3.so.1"

// relaunched holds LD_LIBRARY_PATH as it was before the app started again
// with its own libraries, to give back to what it runs.
const relaunched = "MYGO_CLOUDMUSIC_LD_LIBRARY_PATH"

// useBundledLibs starts the app again with its own libraries, where the
// system has no tray library to load: the dynamic loader only reads
// LD_LIBRARY_PATH as a process starts. Started again, it puts the variable
// back as it was, for what the app opens, as the browser; the loader keeps
// what it read.
func useBundledLibs() {
	if old, ok := os.LookupEnv(relaunched); ok {
		os.Unsetenv(relaunched)
		if old == "" {
			os.Unsetenv("LD_LIBRARY_PATH")
		} else {
			os.Setenv("LD_LIBRARY_PATH", old)
		}
		return
	}
	exe, err := os.Executable()
	if err != nil {
		return
	}
	dir := filepath.Join(filepath.Dir(exe), "lib")
	if _, err := os.Stat(filepath.Join(dir, trayLib)); err != nil {
		return
	}
	// The names MyGo loads it by.
	for _, name := range []string{trayLib, "libappindicator3.so.1"} {
		if h, err := purego.Dlopen(name, purego.RTLD_LAZY|purego.RTLD_LOCAL); err == nil {
			purego.Dlclose(h)
			return
		}
	}
	syscall.Exec(exe, os.Args, withLibs(os.Environ(), dir))
}

// withLibs is env with dir first in LD_LIBRARY_PATH, and what it was kept
// in relaunched.
func withLibs(env []string, dir string) []string {
	old := ""
	out := make([]string, 0, len(env)+2)
	for _, kv := range env {
		if v, ok := strings.CutPrefix(kv, "LD_LIBRARY_PATH="); ok {
			old = v
			continue
		}
		out = append(out, kv)
	}
	path := dir
	if old != "" {
		path += ":" + old
	}
	return append(out, "LD_LIBRARY_PATH="+path, relaunched+"="+old)
}
