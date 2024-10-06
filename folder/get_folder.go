package folder

import "github.com/gofrs/uuid"
import "fmt"
import "errors"
import "strings"

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	rootNode, exists := f.folderTree[orgID]

	if !exists {
		fmt.Println(errors.New("Invalid orgID provided."))
		return []Folder{}
	}

	rootNode = rootNode

	allFolders := []Folder{}
	collectFolders(rootNode, &allFolders)

	return allFolders

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	_, validOrg := f.allOrgIDs[orgID]
	if !validOrg {
		fmt.Println(errors.New("Invalid orgID provided."))
		return []Folder{}
	}

	// Check if the specified folder name exists in orgFolder and get its path

	folderDetails, exists := f.foldersByName[name]
	if !exists {
		fmt.Println(errors.New("Folder name does not exist in the orgFolder."))
		return []Folder{}
	}

	folderPath := folderDetails.Paths

	// Split the folder path to navigate through the tree
	pathSegments := strings.Split(folderPath, ".")
	currentNode := f.folderTree[orgID]

	// Traverse the tree using the path segments
	for _, segment := range pathSegments {
		
		nextNode, exists := currentNode.ChildNodes[segment]
		
		if exists {
			currentNode = nextNode
		} else {
			fmt.Println(errors.New("Path segment does not exist in the folder tree."))
			return []Folder{}
		}
	}

	// Get all the relevant child folders
	childFolders := []Folder{}
	collectFolders(currentNode, &childFolders)

	return childFolders
}

// Given a starting folder node, recurses through each child folder and adds it to 'folders'
func collectFolders(node *FolderNode, folders *[]Folder) {

	for _, childNode := range node.ChildNodes {
		*folders = append(*folders, childNode.Folder)
		collectFolders(childNode, folders)
	}
}