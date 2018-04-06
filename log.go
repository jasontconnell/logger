package logger

import (
	"log"
	"time"
	"os"
	"fmt"
	"path/filepath"
)

type Log struct {
	Conception time.Time
	il *log.Logger
}

var logmap map[string]Log

func init(){
	logmap = make(map[string]Log)
}

func Get(name string) Log {
	now := time.Now()
	if l, ok := logmap[name]; ok {
		if l.Conception.Before(now.Add(time.Hour * 24)) {
			return l
		} else {
			delete(logmap, name)
		}
	}
	
	logdir := "logs"
	nows := now.Format("20060102")
	filename := fmt.Sprintf("%s-%s.log", name, nows)
	mderr := os.Mkdir(logdir, os.ModePerm)
	if  mderr != nil && !os.IsExist(mderr) {
		panic(mderr)
	}

	fullpath := filepath.Join(logdir, filename)
	flags := os.O_WRONLY | os.O_APPEND
	if _, serr := os.Stat(fullpath); os.IsNotExist(serr) {
		flags = flags | os.O_CREATE
	}

	file,ferr := os.OpenFile(fullpath, flags, os.ModePerm)
	if ferr != nil {
		file = os.Stdout
	}

	il := log.New(file, name + " - ", log.Ldate | log.Ltime | log.Llongfile)

	l := Log{ Conception: time.Now(), il: il }

	logmap[name] = l

	return l
}

func (l Log) Println(s ...interface{}){
	l.il.Println(s)
}