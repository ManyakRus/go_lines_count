package logic

import (
	"bytes"
	"fmt"
	"github.com/ManyakRus/go_lines_count/internal/config"
	"github.com/ManyakRus/go_lines_count/internal/constants"
	"github.com/ManyakRus/go_lines_count/internal/create_file_txt"
	"github.com/ManyakRus/go_lines_count/internal/packages_folder"
	"github.com/ManyakRus/starter/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// CountLinesFunctions - количество строк и количество функций
type CountLinesFunctions struct {
	LinesCount int
	FuncCount  int
}

type FolderLinesCountStruct struct {
	Name  string
	Level int
	CountLinesFunctions
	MassFolderLinesCountStruct []*FolderLinesCountStruct
}

//var FolderLinesCount = make([]FolderLinesCountStruct, 0)

// FindLinesCount_Cache - кэш рассчитанных количество строк и количество функций
var FindLinesCount_Cache = make(map[string]CountLinesFunctions)

// StartFillAll - Старт работы приложения
func StartFillAll(FileName string) bool {
	Otvet := false

	FolderRoot := packages_folder.FindAllFolders_FromDir(config.Settings.DIRECTORY_SOURCE)
	if FolderRoot == nil {
		log.Error("Error: not found folder: ", FolderRoot)
		return Otvet
	}

	FolderLinesCountRoot := NewFolderLinesCountStruct()
	FolderLinesCountRoot.Name = FolderRoot.Name
	FolderLinesCountRoot.Level = 1
	_, _, err := FillFolder(FolderRoot, &FolderLinesCountRoot)
	if err != nil {
		//log.Error("FillFolder() error: ", err)
		return Otvet
	}

	err = Save(&FolderLinesCountRoot)
	if err != nil {
		return Otvet
	}

	Otvet = true
	return Otvet
}

func FillFolder(Folder *folders.Folder, FolderLinesCountParent *FolderLinesCountStruct) (int, int, error) {
	var err error

	LinesCount, FuncCount := FindLinesCount_folder1(Folder)
	Level := FolderLinesCountParent.Level
	LevelNew := Level + 1

	for _, folder1 := range Folder.Folders {
		FolderLinesCount1 := NewFolderLinesCountStruct()
		FolderLinesCount1.Name = folder1.Name
		FolderLinesCount1.Level = LevelNew
		LinesCount1, FuncCount1, err := FillFolder(folder1, &FolderLinesCount1)
		if err != nil {
			log.Error("FillFolder() error: ", err)
		}

		//if LevelNew == 2 {
		//	log.Debug(FolderLinesCount1.String())
		//}

		if LevelNew <= config.Settings.FOLDERS_LEVEL {
			FolderLinesCount1.LinesCount = LinesCount1
			FolderLinesCount1.FuncCount = FuncCount1

			FolderLinesCountParent.MassFolderLinesCountStruct = append(FolderLinesCountParent.MassFolderLinesCountStruct, &FolderLinesCount1)
		}

		FolderLinesCountParent.LinesCount = FolderLinesCountParent.LinesCount + LinesCount1
		FolderLinesCountParent.FuncCount = FolderLinesCountParent.FuncCount + FuncCount1

		LinesCount = LinesCount + LinesCount1
		FuncCount = FuncCount + FuncCount1
	}

	FolderLinesCountParent.LinesCount = LinesCount
	FolderLinesCountParent.FuncCount = FuncCount

	return LinesCount, FuncCount, err
}

func FindLinesCount_folder1(Folder1 *folders.Folder) (int, int) {
	LinesCount := 0
	FuncCount := 0

	if Folder1 == nil {
		return 0, 0
	}

	for _, file1 := range Folder1.Files {
		Filename1 := file1.Name
		Filename1 = strings.ToLower(Filename1)
		Filename := Folder1.FileName + micro.SeparatorFile() + Filename1
		if strings.HasSuffix(Filename, ".go") == false {
			continue
		}

		count1, func_count1 := FindLinesCount(Filename)
		LinesCount = LinesCount + count1
		FuncCount = FuncCount + func_count1
	}

	return LinesCount, FuncCount
}

// FindLinesCount - возвращает количество строк и количество функций в файле
func FindLinesCount(FileName string) (int, int) {
	LinesCount := 0
	FuncCount := 0

	//
	CountLinesFunctions1, isFinded := FindLinesCount_Cache[FileName]
	if isFinded == true {
		return CountLinesFunctions1.LinesCount, CountLinesFunctions1.FuncCount
	}

	//
	bytes1, err := os.ReadFile(FileName)
	if err != nil {
		log.Fatal("Can not open file: ", FileName, " error: ", err)
	}

	reader := bytes.NewReader(bytes1)
	LinesCount, err = LinesCount_reader(reader)
	if err != nil {
		log.Fatal("LinesCount_reader error: ", err)
	}

	FuncCount = FindFuncCount(&bytes1)

	//
	FindLinesCount_Cache[FileName] = CountLinesFunctions{
		LinesCount: LinesCount,
		FuncCount:  FuncCount,
	}

	return LinesCount, FuncCount
}

// LinesCount_reader - возвращает количество строк в файле
func LinesCount_reader(r io.Reader) (int, error) {
	defaultSize := 1024
	defaultEndLine := "\n"

	Size := defaultSize
	Sep := defaultEndLine

	buf := make([]byte, Size)
	var count int

	for {
		n, err := r.Read(buf)
		count += bytes.Count(buf[:n], []byte(Sep))

		if err != nil {
			if err == io.EOF {
				return count, nil
			}
			return count, err
		}

	}
}

// FindFuncCount - находит количество функций(func) в файле
func FindFuncCount(bytes0 *[]byte) int {
	Otvet := 0

	//s := string(*bytes0)
	//Otvet = strings.Count(s, "\nfunc ")

	//sFind := "(\n|\t| )func( |\t)"
	//
	//Otvet = CountMatches(s, regexp.MustCompile(sFind))

	Otvet = bytes.Count(*bytes0, []byte("\nfunc "))

	return Otvet
}

// CountMatches - находит количество совпадений в regexp
func CountMatches(s string, re *regexp.Regexp) int {
	total := 0
	for start := 0; start < len(s); {
		remaining := s[start:] // slicing the string is cheap
		loc := re.FindStringIndex(remaining)
		if loc == nil {
			break
		}
		// loc[0] is the start index of the match,
		// loc[1] is the end index (exclusive)
		start += loc[1]
		total++
	}
	return total
}

// NewFolderLinesCountStruct - создает структуру FolderLinesCountStruct
func NewFolderLinesCountStruct() FolderLinesCountStruct {
	Otvet := FolderLinesCountStruct{}
	Otvet.MassFolderLinesCountStruct = make([]*FolderLinesCountStruct, 0)

	return Otvet
}

// String - возвращает строку, красиво оформленную
func (f *FolderLinesCountStruct) String() string {
	Otvet := ""

	//
	FolderNameLength := FindFolderNameLengthMax(f.MassFolderLinesCountStruct, constants.FolderNameLength)
	sFolderNameLength := strconv.Itoa(FolderNameLength)

	sName := fmt.Sprintf("%-"+sFolderNameLength+"s", "Name")
	Otvet = Otvet + "" + sName + "\tLevel\tLines count\tFunctions count\n"
	sName = fmt.Sprintf("%-"+sFolderNameLength+"s", f.Name)
	Otvet = Otvet + sName + "\t" + strconv.Itoa(f.Level) + "\t" + strconv.Itoa(f.LinesCount) + "\t" + strconv.Itoa(f.FuncCount) + "\n"
	Otvet = Otvet + StringMassFolderLinesCount(f.MassFolderLinesCountStruct, FolderNameLength)

	return Otvet
}

// StringMassFolderLinesCount - возвращает строку из FolderLinesCount.MassFolderLinesCountStruct
func StringMassFolderLinesCount(FolderLinesCount []*FolderLinesCountStruct, FolderNameLength int) string {
	Otvet := ""

	//сортировка
	sort.Slice(FolderLinesCount[:], func(i, j int) bool {
		return FolderLinesCount[i].Name < FolderLinesCount[j].Name
	})
	//FolderNameLength := FindFolderNameLengthMax(FolderLinesCount, constants.FolderNameLength)
	sFolderNameLength := strconv.Itoa(FolderNameLength)

	//обход массива
	for _, v := range FolderLinesCount {
		sName := fmt.Sprintf("%-"+sFolderNameLength+"s", v.Name)
		Otvet = Otvet + sName + "\t" + strconv.Itoa(v.Level) + "\t" + strconv.Itoa(v.LinesCount) + "\t" + strconv.Itoa(v.FuncCount) + "\n"
		if len(v.MassFolderLinesCountStruct) > 0 {
			Otvet = Otvet + StringMassFolderLinesCount(v.MassFolderLinesCountStruct, FolderNameLength)
		}
	}
	return Otvet
}

// FindFolderNameLengthMax - находит максимальную длину имени папки
func FindFolderNameLengthMax(FolderLinesCount []*FolderLinesCountStruct, LengthMax int) int {
	Otvet := LengthMax
	for _, v := range FolderLinesCount {
		Otvet1 := FindFolderNameLengthMax(v.MassFolderLinesCountStruct, LengthMax)
		LenName := len(v.Name)
		Otvet = micro.Max(Otvet, Otvet1)
		Otvet = micro.Max(Otvet, LenName)
	}
	return Otvet
}

// CSV - возвращает строку в формате .csv
func (f *FolderLinesCountStruct) CSV() string {
	Otvet := ""

	Otvet = Otvet + "\"Name\";Level;Lines count;Functions count\n"
	Otvet = Otvet + `"` + f.Name + `"` + ";" + strconv.Itoa(f.Level) + ";" + strconv.Itoa(f.LinesCount) + ";" + strconv.Itoa(f.FuncCount) + "\n"
	Otvet = Otvet + CSVMassFolderLinesCount(f.MassFolderLinesCountStruct)

	return Otvet
}

// CSVMassFolderLinesCount - возвращает строку из FolderLinesCount.MassFolderLinesCountStruct
func CSVMassFolderLinesCount(FolderLinesCount []*FolderLinesCountStruct) string {
	Otvet := ""

	//сортировка
	sort.Slice(FolderLinesCount[:], func(i, j int) bool {
		return FolderLinesCount[i].Name < FolderLinesCount[j].Name
	})

	//обход массива
	for _, v := range FolderLinesCount {
		Otvet = Otvet + `"` + v.Name + `"` + ";" + strconv.Itoa(v.Level) + ";" + strconv.Itoa(v.LinesCount) + ";" + strconv.Itoa(v.FuncCount) + "\n"
		if len(v.MassFolderLinesCountStruct) > 0 {
			Otvet = Otvet + CSVMassFolderLinesCount(v.MassFolderLinesCountStruct)
		}
	}
	return Otvet
}

// Save - сохраняет в файл
func Save(FolderLinesCount *FolderLinesCountStruct) error {
	var err error

	Filename := config.Settings.FILENAME
	Filename_low := strings.ToLower(Filename)
	ext := filepath.Ext(Filename_low)

	switch ext {
	case ".csv":
		{
			s := FolderLinesCount.CSV()
			err = create_file_txt.SaveToFile(s)
		}
	case ".txt":
		{
			s := FolderLinesCount.String()
			err = create_file_txt.SaveToFile(s)
		}
	default:
		{
			s := FolderLinesCount.String()
			print(s)
		}
	}

	return err
}
