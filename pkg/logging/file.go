package logging

import (
	"fmt"
	"shop/pkg/global"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", global.CONFIG.App.RuntimeRootPath, global.CONFIG.App.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		global.CONFIG.App.LogSaveName,
		time.Now().Format(global.CONFIG.App.TimeFormat),
		global.CONFIG.App.LogFileExt,
	)
}
