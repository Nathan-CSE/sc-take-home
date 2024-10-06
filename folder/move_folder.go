package folder
import "fmt"
import "errors"

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	// Just get the first instance in the slice index because there's not enough
	// information to specify
	targetFolderSlice, exists := f.foldersByName[name]
	if !exists {
		fmt.Println(errors.New("Folder name does not exist."))
		return []Folder{}
	}

	targetFolder := targetFolderSlice[0]
	
	dstFolderSlice, exists := f.foldersByName[dst]
	if !exists {
		fmt.Println(errors.New("Folder name does not exist."))
		return []Folder{}
	}

	dstFolder := dstFolderSlice[0]

	currentNode := traverseFolderTree(f.folderTree[targetFolder.OrgId], targetFolder.Paths)
	if currentNode == nil {
		return []Folder{}, errors.New("Target folder path does not exist in the folder tree.")
	}

	// Find the destination folder node in the tree
	dstNode := traverseFolderTree(f.folderTree[dstFolder.OrgId], dstFolder.Paths)
	if dstNode == nil {
		return []Folder{}, errors.New("Destination folder path does not exist in the folder tree.")
	}

	return []Folder{}, nil
}
