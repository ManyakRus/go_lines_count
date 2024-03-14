package config

import (
	"encoding/json"
	"github.com/ManyakRus/starter/log"
	"os"
	"strconv"
)

const FILENAME_XGML = "packages.graphml"

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	DIRECTORY_SOURCE string
	FILENAME_CSV     string
	FOLDERS_LEVEL    int
	EXCLUDE_FILDERS  []string
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {
	Settings = SettingsINI{}
	Settings.DIRECTORY_SOURCE = os.Getenv("DIRECTORY_SOURCE")
	Settings.FILENAME_CSV = os.Getenv("FILENAME_CSV")

	s := os.Getenv("FOLDERS_LEVEL")
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Warn("Need fill FOLDERS_LEVEL, error: ", err)
		i = 1
	}
	Settings.FOLDERS_LEVEL = i

	if Settings.DIRECTORY_SOURCE == "" {
		Settings.DIRECTORY_SOURCE = CurrentDirectory()
		//log.Panicln("Need fill DIRECTORY_SOURCE ! in os.ENV ")
	}

	if Settings.FILENAME_CSV == "" {
		Settings.FILENAME_CSV = FILENAME_XGML
	}

	//
	sEXCLUDE_FILDERS := os.Getenv("EXCLUDE_FILDERS")
	if sEXCLUDE_FILDERS != "" {
		Mass_EXCLUDE_FOLDERS := make([]string, 0, 0)
		err = json.Unmarshal([]byte(sEXCLUDE_FILDERS), &Mass_EXCLUDE_FOLDERS)
		if err != nil {
			log.Panic("Unmarshal json EXCLUDE_FILDERS, error: ", err)
		}
		Settings.EXCLUDE_FILDERS = Mass_EXCLUDE_FOLDERS
	}

	//
}

// CurrentDirectory - возвращает текущую директорию ОС
func CurrentDirectory() string {
	Otvet, err := os.Getwd()
	if err != nil {
		//log.Println(err)
	}

	return Otvet
}

// FillFlags - заполняет параметры из командной строки
func FillFlags() {
	Args := os.Args[1:]
	if len(Args) > 3 {
		return
	}

	if len(Args) > 0 {
		Settings.DIRECTORY_SOURCE = Args[0]
	}
	if len(Args) > 1 {
		Settings.FILENAME_CSV = Args[1]
	}
	if len(Args) > 2 {
		s := Args[2]
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Warn("Need fill FOLDERS_LEVEL, error: ", err)
			i = 1
		}
		Settings.FOLDERS_LEVEL = i
	}
}
