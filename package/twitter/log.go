package twitter

import (
	"time"
	"fmt"
	"os"
)

func getLocalDateTime() string {
	currentTime := time.Now()
	return fmt.Sprintf("%d:%d:%d",
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second())
}

func SaveToLogFile(message string, logLevel string) {
	file, err := os.OpenFile("./output/log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(getLocalDateTime() + " [" + logLevel + "] -> " + message + "\n"); err != nil {
		SaveToLogFile(err.Error(), "ERROR")
		panic(err)
	}
}
