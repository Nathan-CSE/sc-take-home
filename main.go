package main

import (
	"fmt"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	// orgID2 := uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39a32e4ee1a")
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	// static.go, folder.go and get_folder.go are both part of the same folder package
	// static.go has a bunch of helper methods and fixed types
	res := folder.GetAllFolders()
	// res := folder.GenerateData()

	// example usage

	// we can use folder driver to run code and process the data before hand I'm assuming
	// to make it easier to process
	folderDriver := folder.NewDriver(res)
	orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	
	
	// Prints out all folders
	// folder.PrettyPrint(res)
	
	fmt.Printf("\n Folders for orgID before move: %s", orgID)
	folder.PrettyPrint(orgFolder)
	
	moveFolder, err := folderDriver.MoveFolder("innocent-armor", "integral-jungle")
	// moveFolder, err := folderDriver.MoveFolder("pure-slapstick", "integral-jungle")
	fmt.Printf("Move folder: \n")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		} else {
			folder.PrettyPrint(moveFolder)
		}
		
	checkOrgFolder := folderDriver.GetFoldersByOrgID(orgID)
	fmt.Printf("\n Folders for orgID after move: %s", orgID)
	folder.PrettyPrint(checkOrgFolder)
	
	// allSubFolders := folderDriver.GetAllChildFolders(orgID, "stable-chadsangeling")
	// allSubFolders := folderDriver.GetAllChildFolders(orgID, "discrete-whdsasdistler")
	// fmt.Printf("\n Child folders for discrete-whistler:")
	// folder.PrettyPrint(allSubFolders)


}
