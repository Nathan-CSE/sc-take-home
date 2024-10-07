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
	folderDetailsSlice, exists := f.foldersByName[name]
	if !exists {
		fmt.Println(errors.New("Folder name does not exist."))
		return []Folder{}
	}

	// Find the folder that matches the orgID -> to handle folders with the same name but different orgs
	var folderDetails FolderMapping
	found := false
	for _, details := range folderDetailsSlice {
		if details.OrgId == orgID {
			folderDetails = details
			found = true
			break
		}
	}

	// If no matching folder for the orgID is found, return an error
	if !found {
		fmt.Println(errors.New("No folder with the specified name found for this orgID."))
		return []Folder{}
	}

	folderPath := folderDetails.Paths

	currentNode, err := traverseFolderTree(orgID, folderPath, f.folderTree)
	if err != nil {
		fmt.Println(err)
		return []Folder{}
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

// Traverse the folderTree until the end of the path is found, then return the node at the end of the path
func traverseFolderTree(orgID uuid.UUID, path string, folderTree map[uuid.UUID]*FolderNode) (*FolderNode, error) {

	rootNode, exists := folderTree[orgID]
	if !exists {
		return nil, errors.New("invalid orgID provided")
	}

	pathSegments := strings.Split(path, ".")

	currentNode := rootNode

	for _, segment := range pathSegments {
		nextNode, exists := currentNode.ChildNodes[segment]

		if exists {
			currentNode = nextNode
		} else {
			return nil, errors.New("path segment does not exist in the folder tree")
		}

	}

	return currentNode, nil
}