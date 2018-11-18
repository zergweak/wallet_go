package log

import (
	"log"
	"os"
)

func Print(value string)  {
	fileName := "Log.log"
	logFile,err  := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error")
	}
	debugLog := log.New(logFile,"[Info]",log.Llongfile)
	debugLog.Println(value)
}