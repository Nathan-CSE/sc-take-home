package folder

import "github.com/gofrs/uuid"
import "strings"
import "fmt"

type IDriver interface {
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	GetFoldersByOrgID(orgID uuid.UUID) []Folder
	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) []Folder

	// component 2
	// Implement the following methods:
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)
}

type FolderNode struct {
	Folder     Folder
	ChildNodes map[string]*FolderNode
}

type FolderMapping struct {
	OrgId uuid.UUID
	Paths string
}

type driver struct {
	// define attributes here
	// data structure to store folders
	// or preprocessed data
	folderTree map[uuid.UUID]*FolderNode
	foldersByName map[string][]FolderMapping
	allOrgIDs map[uuid.UUID]struct{}
}

func NewDriver(folders []Folder) IDriver {

	folderTree := make(map[uuid.UUID]*FolderNode)
	foldersByName := make(map[string][]FolderMapping)
	allOrgIDs := make(map[uuid.UUID]struct{})

	for _, folder := range folders {
		// Add org to the set -> set avoids duplicates
		allOrgIDs[folder.OrgId] = struct{}{}

		foldersByName[folder.Name] = append(foldersByName[folder.Name], FolderMapping{
			OrgId: folder.OrgId,
			Paths: folder.Paths,
		})

		// the root node of the tree corresponding to the org otherwise create
		// a new root node
		rootNode, exists := folderTree[folder.OrgId]

		if !exists {
			rootNode = &FolderNode{
        Folder:     Folder{Name: "root", OrgId: folder.OrgId, Paths: ""},
        ChildNodes: make(map[string]*FolderNode),
			}
			folderTree[folder.OrgId] = rootNode 
		}

		// separate each individual segment of the path
		pathSegments := strings.Split(folder.Paths, ".")
		
		// Iterate over each path segment and check if a child node exists, if it doesn't
		// then we create a new node, then we iterate to the next one
		currentNode := rootNode
		for _, segment := range pathSegments {
			_, exists := currentNode.ChildNodes[segment]

			if !exists {
				currentNode.ChildNodes[segment] = &FolderNode{
					Folder: Folder{Name: segment, OrgId: folder.OrgId, Paths: segment},
					ChildNodes: make(map[string]*FolderNode),
				}
			}

			// Move to the next node in the path segment
			currentNode = currentNode.ChildNodes[segment]
		}

		// Insert the actual folder at the final node since that is the path destination
		currentNode.Folder = folder
	}

	// fmt.Printf("======== OrgFolders ========\n")
	// for orgID, org := range orgFolders {
	// 	fmt.Printf("OrgID %s\n", orgID)
	// 	for _, folder := range org {
	// 		fmt.Printf("Folder: Name=%s, OrgId=%s, Paths=%s\n", folder.Name, folder.OrgId, folder.Paths)
	// 	}
	// }

	// fmt.Printf("======== Tree structure ========\n")
	// for orgID, root := range folderTree {
	// 	fmt.Printf("OrgID: %s\n", orgID)
	// 	printFolderTree(root, 1)
	// }

	return &driver{
		folderTree: folderTree,
		foldersByName: foldersByName,
		allOrgIDs: allOrgIDs,
	}
}

// In-Order Tree Traversal
func printFolderTree(node *FolderNode, level int) {
	fmt.Printf("Folder: Name=%s, OrgId=%s, Paths=%s, Depth=%d\n", node.Folder.Name, node.Folder.OrgId, node.Folder.Paths, level)
	for _, child := range node.ChildNodes {
		printFolderTree(child, level + 1)
	}
}