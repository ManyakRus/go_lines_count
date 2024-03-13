package logic

import (
	"bytes"
	"github.com/ManyakRus/go_lines_count/internal/config"
	"github.com/ManyakRus/go_lines_count/internal/packages_folder"
	"github.com/ManyakRus/starter/folders"
	"github.com/ManyakRus/starter/log"
	"github.com/ManyakRus/starter/micro"
	"io"
	"os"
	"regexp"
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
}

var FolderLinesCount = make([]FolderLinesCountStruct, 0)

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

	_, _, err := FillFolder(FolderRoot, 0)
	if err != nil {
		//log.Error("FillFolder() error: ", err)
		return Otvet
	}

	//log.Info("LinesCount: ", LinesCount, " FuncCount: ", FuncCount)
	log.Infof("FolderLinesCount: %v", FolderLinesCount)

	Otvet = true
	return Otvet
}

func FillFolder(Folder *folders.Folder, level int) (int, int, error) {
	var err error

	LinesCount, FuncCount := FindLinesCount_folder1(Folder)

	for _, folder1 := range Folder.Folders {
		level = level + 1
		LinesCount1, FuncCount1, err := FillFolder(folder1, level)
		if err != nil {
			log.Error("FillFolder() error: ", err)
		}

		if level <= config.Settings.FOLDERS_LEVEL {
			FolderLinesCount1 := FolderLinesCountStruct{}
			FolderLinesCount1.Name = folder1.Name
			FolderLinesCount1.Level = level
			FolderLinesCount1.LinesCount = LinesCount1
			FolderLinesCount1.FuncCount = FuncCount1

			FolderLinesCount = append(FolderLinesCount, FolderLinesCount1)
		}

		LinesCount = LinesCount + LinesCount1
		FuncCount = FuncCount + FuncCount1
	}

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
