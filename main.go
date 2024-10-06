package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	// static.go, folder.go and get_folder.go are both part of the same folder package
	// static.go has a bunch of helper methods and fixed types
	res := folder.GetAllFolders()

	// example usage

	// we can use folder driver to run code and process the data before hand I'm assuming
	// to make it easier to process
	folderDriver := folder.NewDriver(res)
	orgFolder := folderDriver.GetFoldersByOrgID(orgID)

	folder.PrettyPrint(res)
	fmt.Printf("\n Folders for orgID: %s", orgID)
	folder.PrettyPrint(orgFolder)
}
