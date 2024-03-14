package create_file_txt

import (
	"github.com/ManyakRus/go_lines_count/internal/config"
	"github.com/ManyakRus/starter/log"
	"os"
)

// SaveToFile - сохраняет текст в файл
func SaveToFile(Text string) error {
	var err error

	//dir := micro.ProgramDir()
	Filename := config.Settings.FILENAME
	//Filename := dir + micro.SeparatorFile() + Filename1

	// сменим директорию на текущую
	err = os.Chdir(config.CurrentDirectory())
	if err != nil {
		log.Error("Chdir error: ", err)
	}
	log.Info("Chdir: ", config.CurrentDirectory())

	//
	file, err := os.Create(Filename)
	_, err = file.Write([]byte(Text))
	if err != nil {
		log.Error("Write file name: ", Filename, " error: ", err)
		return err
	} else {
		log.Info("Write file name: ", Filename)
	}

	return err
}
