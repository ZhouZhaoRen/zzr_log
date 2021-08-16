package file

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	f *os.File
	logLevel=[]string{"DEBUG","INFO","WARN","ERROR","FATAL"}
	logger *log.Logger
)


const (
	debug  =iota
	info
	warn
	error
	fatal
)

func init() {
	pathName:=getLogFullPath()
	f=openFile(pathName)
	logger=log.New(f,"",log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(debug)
	logger.Println(v)
}
func Info(v ...interface{}) {
	setPrefix(info)
	logger.Println(v)
}
func Warn(v ...interface{}) {
	setPrefix(warn)
	logger.Println(v)
}
func Error(v ...interface{}) {
	setPrefix(error)
	logger.Println(v)
}
func Fatal(v ...interface{}) {
	setPrefix(fatal)
	logger.Println(v)
}

func setPrefix(level int) {
	prefix:=""
	_, file, line, ok := runtime.Caller(2)
	if ok {
		prefix=fmt.Sprintf("[%s][%s:%d]",logLevel[level],filepath.Base(file),line)
	}else {
		prefix=fmt.Sprintf("[%s]",logLevel[level])
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
