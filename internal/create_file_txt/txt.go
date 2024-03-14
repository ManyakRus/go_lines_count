package create_file_txt

import (
	"github.com/ManyakRus/go_lines_count/internal/config"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"os"
)

// SaveToFile - сохраняет текст в файл
func SaveToFile(Text string) error {
	var err error

	dir := micro.ProgramDir()
	Filename1 := config.Settings.FILENAME
	Filename := dir + micro.SeparatorFile() + Filename1
	file, err := os.Create(Filename)
	_, err = file.Write([]byte(Text))
	if err != nil {
		log.Error("Write file name: ", Filename, " error: ", err)
		return err
	}

	return err
}
