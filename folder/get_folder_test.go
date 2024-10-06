package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		want    []folder.Folder
	}{
		{
			name: "Valid OrgID with Folders",
			orgID: uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39132e4ea1a"),
			want: []folder.Folder{
				{
					Name:  "stable-changeling",
					OrgId: uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39132e4ea1a"),
					Paths: "stable-changeling",
				},
				{
					Name:  "quiet-kingpin",
					OrgId: uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39132e4ea1a"),
					Paths: "stable-changeling.quiet-kingpin",
				},
				{
					Name:  "sweet-warhawk",
					OrgId: uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39132e4ea1a"),
					Paths: "stable-changeling.quiet-kingpin.sweet-warhawk",
				},
				{
					Name:  "flexible-humbug",
					OrgId: uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39132e4ea1a"),
					Paths: "stable-changeling.quiet-kingpin.flexible-humbug",
				},
			},
		},
		{
			name: "Valid OrgID with 2 Separate Folder Trees",
			orgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
			want: []folder.Folder{
				{
					Name:  "discrete-whistler",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler",
				},
				{
					Name:  "integral-jungle",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.integral-jungle",
				},
				{
					Name:  "robust-batgirl",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.integral-jungle.robust-batgirl",
				},
				{
					Name:  "innocent-armor",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.innocent-armor",
				},
				{
					Name:  "pure-slapstick",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.innocent-armor.pure-slapstick",
				},
				{
					Name:  "stable-changeling",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stable-changeling",
				},
				{
					Name:  "quiet-kingpin",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stable-changeling.quiet-kingpin",
				},
				{
					Name:  "sweet-warhawk",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stable-changeling.quiet-kingpin.sweet-warhawk",
				},
				{
					Name:  "flexible-humbug",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stable-changeling.quiet-kingpin.flexible-humbug",
				},
			},
		},
		{
			name: "Invalid OrgID",
			orgID: uuid.FromStringOrNil("invalid-org-id"),
			want: []folder.Folder{},
		},
	}

	res := folder.GetAllFolders()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := folder.NewDriver(res)
			get := driver.GetFoldersByOrgID(tt.orgID)

			assert.ElementsMatch(t, tt.want, get, "The returned folders do not match the expected result")
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		orgID      uuid.UUID
		folderName string
		want       []folder.Folder
	}{
		{
			name: "Get child folders for stable-changeling",
			orgID: uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39132e4ea1a"),
			folderName: "stable-changeling",
			want: []folder.Folder{
				{
					Name:  "quiet-kingpin",
					OrgId: uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39132e4ea1a"),
					Paths: "stable-changeling.quiet-kingpin",
				},
				{
					Name:  "sweet-warhawk",
					OrgId: uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39132e4ea1a"),
					Paths: "stable-changeling.quiet-kingpin.sweet-warhawk",
				},
				{
					Name:  "flexible-humbug",
					OrgId: uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39132e4ea1a"),
					Paths: "stable-changeling.quiet-kingpin.flexible-humbug",
				},
			},
		},
		{
			name: "Get child folders for quiet-kingpin",
			orgID: uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39132e4ea1a"),
			folderName: "quiet-kingpin",
			want: []folder.Folder{
				{
					Name:  "sweet-warhawk",
					OrgId: uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39132e4ea1a"),
					Paths: "stable-changeling.quiet-kingpin.sweet-warhawk",
				},
				{
					Name:  "flexible-humbug",
					OrgId: uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39132e4ea1a"),
					Paths: "stable-changeling.quiet-kingpin.flexible-humbug",
				},
			},
		},
		{
			name:       "Get child folders for non-existent folder",
			orgID:      uuid.FromStringOrNil("3d514f27-4943-4043-a7b1-e39132e4ea1a"),
			folderName: "non-existent-folder",
			want:       []folder.Folder{},
		},
		{
			name:       "Get child folders for invalid OrgID",
			orgID:      uuid.FromStringOrNil("invalid-org-id"),
			folderName: "quiet-kingpin",
			want:       []folder.Folder{},
		},
	}

	res := folder.GetAllFolders()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			driver := folder.NewDriver(res)
			get := driver.GetAllChildFolders(tt.orgID, tt.folderName)

			assert.ElementsMatch(t, tt.want, get, "Child folders do not match")
		})
	}
}