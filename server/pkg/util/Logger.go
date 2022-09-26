package util

import (
	"custom-file-server/server/pkg/constant"
	"fmt"
	"time"
)

func WriteMsgLog(typeMsg, message string) {
	now := time.Now()
	logMessage := now.Format(constant.DATE_TIME_FORMAT) + " " + typeMsg + " | " + message
	fmt.Println(logMessage)
}

func WriteBanner(banner string) {
	fmt.Println("\n" + banner)
}
