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

func Trace(v ...interface{}) {
	if level > LevelTrace {
		return
	}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.LstdFlags)
	logger.Println("[T]", v)
}

func Debug(v ...interface{}) {
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
	logger.Println("[D]", v)
}

func Info(v ...interface{}) {
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
	logger.Println("[I]", v)
}

func Warning(v ...interface{}) {
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
	logger.Println("[W]", v)
}

func Error(v ...interface{}) {
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
	logger.Println("[E]", v)
}

func Fatal(v ...interface{}) {
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
	logger.Println("[F]", v)
	os.Exit(-1)
}
