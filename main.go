package main

import (
	"fmt"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	res := folder.GetAllFolders()

	// res := folder.GenerateData()
	folderDriver := folder.NewDriver(res)
	orgFolder := folderDriver.GetFoldersByOrgID(orgID)
	
	
	// Prints out all folders
	// folder.PrettyPrint(res)
	
	fmt.Printf("\n Folders for orgID before move: %s", orgID)
	folder.PrettyPrint(orgFolder)
	
	moveFolder, err := folderDriver.MoveFolder("pure-slapstick", "integral-jungle")
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
