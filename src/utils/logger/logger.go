package logger

import (
	"log"
	"os"
	"path"
	"utils/tools"
)

const (
	LevelTrace = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

var (
	level = LevelTrace
)

func GetLevel(l int) int {
	return l
}

func SetLevel(l int) {
	level = l
}

func Trace(msg string) {
	if level > LevelTrace {
		return
	}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LstdFlags)
	logger.Println("[T]" + msg)
}

func Debug(msg string) {
	if level > LevelDebug {
		return
	}
	logfilename := tools.GetDateYYYYMMDD() + ".log"
	logfilepath := path.Join(tools.GetExecutePath(), "logs", logfilename)
	logfile, err := os.OpenFile(logfilepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer logfile.Close()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	logger := log.New(logfile, "", log.Ldate|log.Ltime|log.LstdFlags)
	logger.Println("[D]" + msg)
}

func Info(msg string) {
	if level > LevelInfo {
		return
	}
	logfilename := tools.GetDateYYYYMMDD() + ".log"
	logfilepath := path.Join(tools.GetExecutePath(), "logs", logfilename)
	logfile, err := os.OpenFile(logfilepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer logfile.Close()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	logger := log.New(logfile, "", log.Ldate|log.Ltime|log.LstdFlags)
	logger.Println("[I]" + msg)
}

func Warning(msg string) {
	if level > LevelWarning {
		return
	}
	logfilename := tools.GetDateYYYYMMDD() + ".log"
	logfilepath := path.Join(tools.GetExecutePath(), "logs", logfilename)
	logfile, err := os.OpenFile(logfilepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer logfile.Close()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	logger := log.New(logfile, "", log.Ldate|log.Ltime|log.LstdFlags)
	logger.Println("[W]" + msg)
}

func Error(msg string) {
	if level > LevelError {
		return
	}
	logfilename := tools.GetDateYYYYMMDD() + ".log"
	logfilepath := path.Join(tools.GetExecutePath(), "logs", logfilename)
	logfile, err := os.OpenFile(logfilepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer logfile.Close()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	logger := log.New(logfile, "", log.Ldate|log.Ltime|log.LstdFlags)
	logger.Println("[E]" + msg)
}

func Fatal(msg string) {
	if level > LevelFatal {
		return
	}
	logfilename := tools.GetDateYYYYMMDD() + ".log"
	logfilepath := path.Join(tools.GetExecutePath(), "logs", logfilename)
	logfile, err := os.OpenFile(logfilepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer logfile.Close()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	logger := log.New(logfile, "", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("[F]" + msg)
}
