package config

import (
	"encoding/json"
	"fmt"
	"github.com/ManyakRus/go_lines_count/internal/constants"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
	"strconv"
)

const FILENAME_DEFAULT = ""

// Settings хранит все нужные переменные окружения
var Settings SettingsINI

// SettingsINI - структура для хранения всех нужных переменных окружения
type SettingsINI struct {
	DIRECTORY_SOURCE string
	FILENAME         string
	FOLDERS_LEVEL    int
	EXCLUDE_FOLDERS  []string
}

// FillSettings загружает переменные окружения в структуру из переменных окружения
func FillSettings() {

	Settings = SettingsINI{}
	LoadExcludeFolders()

	Settings.DIRECTORY_SOURCE = os.Getenv("DIRECTORY_SOURCE")
	Settings.FILENAME = os.Getenv("FILENAME")

	s := os.Getenv("FOLDERS_LEVEL")
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Warn("Need fill FOLDERS_LEVEL, error: ", err)
		i = constants.FOLDERS_LEVEL_DEFAULT
	}
	Settings.FOLDERS_LEVEL = i

	if Settings.DIRECTORY_SOURCE == "" {
		Settings.DIRECTORY_SOURCE = CurrentDirectory()
		//log.Panicln("Need fill DIRECTORY_SOURCE ! in os.ENV ")
	}

	if Settings.FILENAME == "" {
		Settings.FILENAME = FILENAME_DEFAULT
	}

	//
	//sEXCLUDE_FILDERS := os.Getenv("EXCLUDE_FOLDERS")
	//if sEXCLUDE_FILDERS != "" {
	//	Mass_EXCLUDE_FOLDERS := make([]string, 0, 0)
	//	err = json.Unmarshal([]byte(sEXCLUDE_FILDERS), &Mass_EXCLUDE_FOLDERS)
	//	if err != nil {
	//		log.Panic("Unmarshal json EXCLUDE_FOLDERS, error: ", err)
	//	}
	//	Settings.EXCLUDE_FOLDERS = Mass_EXCLUDE_FOLDERS
	//}

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
		s := Args[1]
		if s == "-h" || s == "--help" {
			println(constants.TEXT_HELP)
			os.Exit(0)
		}
		Settings.FILENAME = s
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

// LoadExcludeFolders - загружает маппинг ТипБД = ТипGolang, из файла .json
func LoadExcludeFolders() {
	dir := micro.ProgramDir()
	FileName := dir + micro.SeparatorFile() + constants.EXCLUDE_FOLDERS_FILENAME

	var err error

	//чтение файла
	bytes, err := os.ReadFile(FileName)
	if err != nil {
		TextError := fmt.Sprint("ReadFile() error: ", err)
		log.Warn(TextError)
		return
	}

	//json в map
	//var MapServiceURL2 = make(map[string]string)
	err = json.Unmarshal(bytes, &Settings.EXCLUDE_FOLDERS)
	if err != nil {
		log.Panic("Unmarshal() error: ", err)
	}

}
