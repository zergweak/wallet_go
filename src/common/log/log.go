package log

import (
	"fmt"
	"os"
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
