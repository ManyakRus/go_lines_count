package main

import (
	"github.com/ManyakRus/go_lines_count/internal/config"
	"github.com/ManyakRus/go_lines_count/internal/constants"
	"github.com/ManyakRus/go_lines_count/internal/logic"
	"github.com/ManyakRus/starter/config_main"
	"github.com/ManyakRus/starter/log"
	"time"
)

func main() {
	StartApp()
}

func StartApp() {
	config_main.LoadENV_or_SettingsTXT()
	config.FillSettings()
	config.FillFlags()

	StartAt := time.Now()
	FileName := config.Settings.FILENAME
	log.Info("directory: ", config.Settings.DIRECTORY_SOURCE)
	log.Info("file name: ", FileName)
	ok := logic.StartFillAll(FileName)
	if ok == false {
		println(constants.TEXT_HELP)
	}

	log.Info("Time passed: ", time.Since(StartAt))

}
