package packages_folder

import (
	"github.com/ManyakRus/go_lines_count/internal/config"
	"github.com/ManyakRus/starter/folders"
)

func FindAllFolders_FromDir(dir string) *folders.Folder {

	MassExclude := make([]string, 0)
	MassExclude = append(MassExclude, "vendor")
	MassExclude = append(MassExclude, ".git")
	MassExclude = append(MassExclude, ".idea")
	MassExclude = append(MassExclude, ".vscode")
	MassExclude = append(MassExclude, config.Settings.EXCLUDE_FOLDERS...)
	FolderRoot := folders.FindFoldersTree(dir, true, true, false, MassExclude)

	return FolderRoot
}
