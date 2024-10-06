package folder

import "github.com/gofrs/uuid"
import "fmt"
import "errors"
import "strings"

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders, exists := f.orgFolders[orgID]

	if exists {
		return folders
	}

	return []Folder{}

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	orgFolder, validOrg := f.orgFolders[orgID]
	if !validOrg {
		fmt.Println(errors.New("Invalid orgID provided."))
		return []Folder{}
	}

	fmt.Printf("Org folder contents:\n")
	for _, folder := range orgFolder {
		fmt.Printf("- %s\n", folder.Name)
	}

	// Check if the specified folder name exists in orgFolder and get its path
	var folderPath string
	folderExists := false
	for _, folder := range orgFolder {
		if folder.Name == name {
			folderPath = folder.Paths
			folderExists = true
			break
		}
	}

	if !folderExists {
		fmt.Println(errors.New("Folder name does not exist in the orgFolder."))
		return []Folder{}
	}

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
	collectChildFolders(currentNode, &childFolders)

	return childFolders
}

// Adds all child folders to childFolders and recurses through each additional child folder
func collectChildFolders(node *FolderNode, childFolders *[]Folder) {
	for _, childNode := range node.ChildNodes {
		*childFolders = append(*childFolders, childNode.Folder)
		collectChildFolders(childNode, childFolders)
	}
}