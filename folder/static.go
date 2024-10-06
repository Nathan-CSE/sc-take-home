package folder

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gofrs/uuid"
	"github.com/lucasepe/codename"
)

// These are all helper methods and fixed types.
// There's no real need for you to be editting these, but feel free to tweak it to suit your needs.
// If you do make changes here, be ready to discuss why these changes were made.

// how many trees you want to generate
const MaxRootSet = 4

// maximum possible children per node
const MaxChild = 4

// max depth of the tree
const MaxDepth = 5

// the default orgID that we will be using for testing
const DefaultOrgID = "c1556e17-b7c0-45a3-a6ae-9546248fb17a"

type Folder struct {
	Name  string    `json:"name"`
	OrgId uuid.UUID `json:"org_id"`
	Paths string    `json:"paths"`
	ChildFolders []Folder  `json:"child_folders"`
}

func GenerateData() []Folder {
	rng, _ := codename.DefaultRNG()
	tree := []Folder{}

	for i := 0; i < MaxRootSet; i++ {
		orgId := uuid.FromStringOrNil(DefaultOrgID)
		if i%3 == 0 {
			orgId = uuid.Must(uuid.NewV4())
		}

		name := codename.Generate(rng, 0)

		// Changed to fit the new hierarchial data structure
		root := Folder{
			Name:  name,
			OrgId: orgId,
			Paths: name,
		}

		root.ChildFolders = generateTree(1, []Folder{root})

		tree = append(tree, root)

		// subtree := make(chan []Folder)
		// go func() {
		// 	subtree <- generateTree(1, []Folder{
		// 		{
		// 			Name:  name,
		// 			OrgId: orgId,
		// 			Paths: name,
		// 		},
		// 	})
		// }()
		// tree = append(tree, <-subtree...)
	}

	return tree
}

// Modify this first to reflect the child folders directly
func generateTree(depth int, tree []Folder) []Folder {
	rng, _ := codename.DefaultRNG()

	if depth >= MaxDepth {
		return tree
	}

	// t refers to each folder 't' in the current tree slice/hierarchy
	for _, t := range tree {
		numOfChild := rng.Int()%MaxChild + 1

		// This loop runs 'numOfChild' times to create the child folders for the current parent folder
		// it generates a new child folder for each iteration
		// Since I'm adding subfolder and adding depth, then i need to modify this loop
		// to recursively generate subtrees for each folder
		for i := 0; i < numOfChild; i++ {
			name := codename.Generate(rng, 0)

			childFolder := Folder{
				Name: name,
				OrgId: t.OrgId,
				Paths: t.Paths + "." + name,
			}

			// I generate further childFolders for the childFolder datastructure that I just created
			// to populate the tree, then append it to the current tree's ChildFolder property
			childFolder.ChildFolders = generateTree(depth+1, []Folder{childFolder})

			t.ChildFolders = append(t.ChildFolders, childFolder)


			// childTree := make(chan []Folder)
			// go func() {
			// 	childTree <- generateTree(depth+1, []Folder{
			// 		{
			// 			Name:  name,
			// 			OrgId: t.OrgId,
			// 			Paths: t.Paths + "." + name,
			// 			ChildFolders: generateTree(depth+1, []Folder{})
			// 		},
			// 	})
			// }()
			// tree = append(tree, <-childTree...)
		}
	}

	return tree
}

func MarshalJson(b interface{}) []byte {
	s, _ := json.MarshalIndent(b, "", "\t")

	return s
}

func PrettyPrint(b interface{}) {
	s := MarshalJson(b)
	fmt.Print(string(s))
}

func GetSampleData() []Folder {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
	basePath := filepath.Dir(filename)
	filePath := filepath.Join(basePath, "sample.json")

	fmt.Println(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonByte, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	folders := []Folder{}
	err = json.Unmarshal(jsonByte, &folders)
	if err != nil {
		panic(err)
	}

	return folders
}

func WriteSampleData(data interface{}) {
	b := MarshalJson(data)
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
	basePath := filepath.Dir(filename)
	filePath := filepath.Join(basePath, "sample.json")

	fmt.Println(filePath)

	err := os.WriteFile(filePath, b, fs.ModeType)
	if err != nil {
		panic(err)
	}
}
