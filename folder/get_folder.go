package folder

import "github.com/gofrs/uuid"

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	// NOTE: Old Code
	// folders := f.folders

	// res := []Folder{}
	// for _, f := range folders {
	// 	if f.OrgId == orgID {
	// 		res = append(res, f)
	// 	}
	// }

	// return res

	folders, exists := f.orgFolders[orgID]

	if exists {
		return folders
	}

	return []Folder{}

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Your code here...

	return []Folder{}
}
