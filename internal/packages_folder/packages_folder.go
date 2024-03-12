package packages_folder

import "github.com/ManyakRus/starter/folders"

func FindAllFolders_FromDir(dir string) *folders.Folder {

	FolderRoot := folders.FindFoldersTree(dir, true, true, false, "vendor")

	return FolderRoot
}
