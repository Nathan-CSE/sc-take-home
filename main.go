package main

import (
	"fmt"
	"hash/fnv"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func main() {
	// orgID2 := uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39a32e4ee1a")
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	test1 := hash("hedsadasllo there")
	test2 := hash("hedsadasllo there")

	fmt.Printf("\n test1: %s", test1)
	fmt.Printf("\n test2: %s", test2)

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
	
	fmt.Printf("\n Folders for orgID: %s", orgID)
	folder.PrettyPrint(orgFolder)

	// allSubFolders := folderDriver.GetAllChildFolders(orgID, "stable-chadsangeling")
	// allSubFolders := folderDriver.GetAllChildFolders(orgID, "discrete-whdsasdistler")
	// fmt.Printf("\n Child folders for discrete-whistler:")
	// folder.PrettyPrint(allSubFolders)


}
