package logging

import (
	"errors"
	"fmt"
	"os"
	"time"

	"music-backend/services/config"
)

var isLoggingFileExist bool = false

func init(){
	cfg := config.Get()

	_, err := os.Stat(cfg.LogFilePath);

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File for logging is not found!")
	}else{
		isLoggingFileExist = true
	}
}

func WriteLog(action string, error string) {
	if !isLoggingFileExist {
		return
	}

	var loggingString string
	cfg := config.Get()

	dt := time.Now() 
	currentTime := dt.Format("01-02-2006 15:04:05") 
	if error == "" {
		loggingString = fmt.Sprintf("Executed time: %v. Executed action: %v \n", currentTime, action)
	}else {
		loggingString = fmt.Sprintf("Executed time: %v. Executed action: %v. Error: %v \n", currentTime, action, error)
	}

	f, err := os.OpenFile(cfg.LogFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	
	defer f.Close()
	
	if _, err = f.WriteString(loggingString); err != nil {
		panic(err)
	}
}