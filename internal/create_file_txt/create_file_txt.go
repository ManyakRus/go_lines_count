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
	ChangeCurrentDirectory()

	//запишем файл
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

// ChangeCurrentDirectory - устанавливает текущую директорию на директорию откуда запущена программа
// вместо директории где находится программа
func ChangeCurrentDirectory() {
	var err error

	// сменим директорию на текущую
	dir := config.CurrentDirectory()
	err = os.Chdir(dir)
	if err != nil {
		log.Error("Chdir error: ", err)
	} else {
		log.Info("Chdir: ", dir)
	}

}
