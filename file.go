package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	F *os.File
	LogLevel=[]string{"DEBUG","INFO","WARN","ERROR","FATAL"}
	logger *log.Logger
)


const (
	DEBUG  =iota
	INFO
	WARN
	ERROR
	FATAL
)

func init() {
	pathName:=getLogFullPath()
	F=openFile(pathName)
	logger=log.New(F,"",log.LstdFlags)
}

func main() {
	Debug("我的日志文件打印")
}
func Debug(v ...interface{}) {
	SetPrefix(DEBUG)
	logger.Println(v)
}
func Info(v ...interface{}) {
	SetPrefix(INFO)
	logger.Println(v)
}
func Warn(v ...interface{}) {
	SetPrefix(WARN)
	logger.Println(v)
}
func Error(v ...interface{}) {
	SetPrefix(ERROR)
	logger.Println(v)
}
func Fatal(v ...interface{}) {
	SetPrefix(FATAL)
	logger.Println(v)
}

func SetPrefix(level int) {
	prefix:=""
	_, file, line, ok := runtime.Caller(2)
	if ok {
		prefix=fmt.Sprintf("[%s][%s:%d]",LogLevel[level],filepath.Base(file),line)
	}else {
		prefix=fmt.Sprintf("[%s]",LogLevel[level])
	}
	logger.SetPrefix(prefix)
}

func openFile(fileName string) *os.File {
	_, err := os.Stat(fileName)
	switch {
	case os.IsNotExist(err):
		mkdir()
	case os.IsPermission(err):
		fmt.Println("没有权限")
	}

	file, openErr := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if openErr != nil {
		fmt.Println("打开文件失败:", openErr)
	}
	return file

}

func mkdir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/logs/", os.ModePerm)
	if err != nil {
		fmt.Println("创建文件夹错误：", err)
		return
	}
}

// 获取文件路径+文件名
func getLogFullPath() string {
	path := "logs/"
	return fmt.Sprintf("%s%s%s.%s", path, "log", time.Now().Format("2006-01-02"), "log")
}
