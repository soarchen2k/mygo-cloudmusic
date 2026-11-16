package app

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/egoist/mygo"
)

// logName is the log's file in the app's logs directory, and logLimit how
// big it grows before the next launch moves it to logName.1, the one
// before that dropped.
const (
	logName  = "mygo-cloudmusic.log"
	logLimit = 1 << 20
)

// logDir is where the log is kept, once useLogFile opened it.
var logDir string

// useLogFile writes what the app logs, as why a song did not play, to a
// file in its logs directory as well as to stderr.
func useLogFile() {
	dir, err := mygo.App.Path(mygo.PathLogs)
	if err != nil {
		log.Println("log:", err)
		return
	}
	f, err := openLog(dir)
	if err != nil {
		log.Println("log:", err)
		return
	}
	logDir = dir
	log.SetOutput(io.MultiWriter(os.Stderr, f))
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Printf("start %s %s", mygo.App.Name(), mygo.App.Version())
}

// openLog opens the log in dir to append to, first moving it aside once it
// has grown past logLimit.
func openLog(dir string) (*os.File, error) {
	path := filepath.Join(dir, logName)
	if info, err := os.Stat(path); err == nil && info.Size() > logLimit {
		os.Rename(path, path+".1")
	}
	return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
}
