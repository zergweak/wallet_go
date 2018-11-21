package log

import (
	"fmt"
	"os"
	"strconv"
)

var logFile *os.File

func Init() {
	fileName := "Log.log"
	var err error
	logFile, err = os.OpenFile(fileName, os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("open file error")
	}
	//defer logFile.Close()
	logFile.Seek(0, os.SEEK_END)
	logFile.WriteString("Init\n")
}

func Print(value string) {
	logFile.Seek(0, os.SEEK_END)
	logFile.WriteString(value)
}

func Println(value string) {
	logFile.Seek(0, os.SEEK_END)
	logFile.WriteString(value)
	logFile.WriteString("\n")
}

func PrintInt64ln(value int64) {
	logFile.Seek(0, os.SEEK_END)
	logFile.WriteString(strconv.FormatInt(value, 10))
	logFile.WriteString("\n")
}

func Printf(format string, a ...interface{}) {
	logFile.Seek(0, os.SEEK_END)
	value := fmt.Sprintf(format, a)
	logFile.WriteString(value)
}
