package folder
import "fmt"
import "errors"
import "strings"

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	// Just get the first instance in the slice index because there's not enough
	// information to specify
	targetFolderSlice, exists := f.foldersByName[name]
	if !exists {
		fmt.Println(errors.New("Folder name does not exist."))
		return []Folder{}, nil
	}

	targetFolder := targetFolderSlice[0]
	
	dstFolderSlice, exists := f.foldersByName[dst]
	if !exists {
		fmt.Println(errors.New("Folder name does not exist."))
		return []Folder{}, nil
	}
	
	dstFolder := dstFolderSlice[0]
	
	// Check if the targer/destination folders have the same orgID
	if targetFolder.OrgId != dstFolder.OrgId {
		return []Folder{}, errors.New(fmt.Sprintf("Target folder '%s' and destination folder '%s' are not part of the same orgID.", name, dst))
	}

	// Traverse the tree to check if both the target/destination folders exist
	targetNode, err := traverseFolderTree(targetFolder.OrgId, targetFolder.Paths, f.folderTree)
	if err != nil {
		return []Folder{}, errors.New(fmt.Sprintf("Target folder '%s' does not exist in the folder tree.", name))
	}

	dstNode, err := traverseFolderTree(dstFolder.OrgId, dstFolder.Paths, f.folderTree)
	if err != nil {
		return []Folder{}, errors.New(fmt.Sprintf("Destination folder '%s' does not exist in the folder tree.", dst))
	}

	// fmt.Printf("dst node path: %s\n", dstFolder.Paths)

	// Get the path to the parent of the target node using string manipulation 
	parentPathSegments := strings.Split(targetFolder.Paths, ".")
	// fmt.Printf("parent path segment: %s\n", parentPathSegments)
	parentPath := strings.Join(parentPathSegments[:len(parentPathSegments) - 1], ".")
	
	// Traverse to the parent node, then find the correct child node to delete
	parentNode, err := traverseFolderTree(targetFolder.OrgId, parentPath, f.folderTree)
	if err != nil {
		return []Folder{}, errors.New(fmt.Sprintf("Parent folder '%s' not found in the folder tree.", parentPath))
	}
	
	// Delete the target node by using the name as the key in the parentNode's childNodes map 
	delete(parentNode.ChildNodes, name)

	// Create the new path, which is just path of the destinatinon, plus the name of the target folder
	// then update the paths for the child nodes in the rest of the folder tree as well as the paths in the 
	// foldersByName data structure
	newPath := dstFolder.Paths + "." + name
	updateFolderPathsAndMap(targetNode, newPath, f)

	// Add the target node to the destination map of the destination folder
	dstNode.ChildNodes[name] = targetNode

	return []Folder{}, nil
}

// Recursively updates the paths in both foldersByName and folderTree
func updateFolderPathsAndMap(node *FolderNode, newPath string, f *driver) {
	// Update the path in the folderTree
	oldPath := node.Folder.Paths
	node.Folder.Paths = newPath

	// Update the relevant paths for the folder in foldersByName if any match the old path
	for i, folderMapping := range f.foldersByName[node.Folder.Name] {
		if folderMapping.Paths == oldPath {
			f.foldersByName[node.Folder.Name][i].Paths = newPath
			break
		}
	}

	// Recurse to the child nodes and update
	for childName, childNode := range node.ChildNodes {
		childNewPath := newPath + "." + childName
		updateFolderPathsAndMap(childNode, childNewPath, f)
	}
}