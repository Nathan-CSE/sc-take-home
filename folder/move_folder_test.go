package folder_test

import (
	"testing"
	
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	tests := []struct {
		name string
		moveName string
		dstName string
		want []folder.Folder
		wantErr bool
	}{
		{
			name: "Move robust-batgirl to itself",
			moveName: "robust-batgirl",
			dstName: "robust-batgirl",
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
						Name:  "one-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.robust-batgirl.one-batgirl",
				},
				{
						Name:  "two-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.robust-batgirl.two-batgirl",
				},
				{
						Name:  "sub-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.sub-batgirl",
				},
				{
						Name:  "test-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.test-batgirl",
				},
				{
						Name:  "innocent-armor",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.innocent-armor",
				},
				{
						Name:  "hello-slapstick",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.innocent-armor.hello-slapstick",
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
						Name:  "silent-kingpin",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin",
				},
				{
						Name:  "sweet-warhawk",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin.sweet-warhawk",
				},
				{
						Name:  "flexible-humbug",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin.flexible-humbug",
				},

			},
			wantErr: true,
		},
		{
			name: "Move integral-jungle to a child folder",
			moveName: "integral-jungle",
			dstName: "two-batgirl",
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
						Name:  "one-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.robust-batgirl.one-batgirl",
				},
				{
						Name:  "two-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.robust-batgirl.two-batgirl",
				},
				{
						Name:  "sub-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.sub-batgirl",
				},
				{
						Name:  "test-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.test-batgirl",
				},
				{
						Name:  "innocent-armor",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.innocent-armor",
				},
				{
						Name:  "hello-slapstick",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.innocent-armor.hello-slapstick",
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
						Name:  "silent-kingpin",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin",
				},
				{
						Name:  "sweet-warhawk",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin.sweet-warhawk",
				},
				{
						Name:  "flexible-humbug",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin.flexible-humbug",
				},

			},
			wantErr:       true,
		},
		{
			name: "Move robust-batgirl to another organisation",
			moveName: "robust-batgirl",
			dstName: "quiet-kingpin",
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
						Name:  "one-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.robust-batgirl.one-batgirl",
				},
				{
						Name:  "two-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.robust-batgirl.two-batgirl",
				},
				{
						Name:  "sub-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.sub-batgirl",
				},
				{
						Name:  "test-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.test-batgirl",
				},
				{
						Name:  "innocent-armor",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.innocent-armor",
				},
				{
						Name:  "hello-slapstick",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.innocent-armor.hello-slapstick",
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
						Name:  "silent-kingpin",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin",
				},
				{
						Name:  "sweet-warhawk",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin.sweet-warhawk",
				},
				{
						Name:  "flexible-humbug",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin.flexible-humbug",
				},

			},
			wantErr:  true,
		},
		{
			name: "Source folder does not exist",
			moveName: "no-exist",
			dstName: "quiet-kingpin",
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
						Name:  "one-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.robust-batgirl.one-batgirl",
				},
				{
						Name:  "two-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.robust-batgirl.two-batgirl",
				},
				{
						Name:  "sub-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.sub-batgirl",
				},
				{
						Name:  "test-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.test-batgirl",
				},
				{
						Name:  "innocent-armor",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.innocent-armor",
				},
				{
						Name:  "hello-slapstick",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.innocent-armor.hello-slapstick",
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
						Name:  "silent-kingpin",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin",
				},
				{
						Name:  "sweet-warhawk",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin.sweet-warhawk",
				},
				{
						Name:  "flexible-humbug",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin.flexible-humbug",
				},

			},
			wantErr:  true,
		},
		{
			name: "Destination folder does not exist",
			moveName: "robust-batgirl",
			dstName: "no-exist",
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
						Name:  "one-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.robust-batgirl.one-batgirl",
				},
				{
						Name:  "two-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.robust-batgirl.two-batgirl",
				},
				{
						Name:  "sub-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.sub-batgirl",
				},
				{
						Name:  "test-batgirl",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.integral-jungle.test-batgirl",
				},
				{
						Name:  "innocent-armor",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.innocent-armor",
				},
				{
						Name:  "hello-slapstick",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "discrete-whistler.innocent-armor.hello-slapstick",
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
						Name:  "silent-kingpin",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin",
				},
				{
						Name:  "sweet-warhawk",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin.sweet-warhawk",
				},
				{
						Name:  "flexible-humbug",
						OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
						Paths: "stable-changeling.silent-kingpin.flexible-humbug",
				},

			},
			wantErr:  true,
		},
		{
			name: "Move robust-batgirl to silent-kingpin",
			moveName: "robust-batgirl",
			dstName: "silent-kingpin",
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
					Name:  "sub-batgirl",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.integral-jungle.sub-batgirl",
				},
				{
					Name:  "test-batgirl",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.integral-jungle.test-batgirl",
				},
				{
					Name:  "innocent-armor",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.innocent-armor",
				},
				{
					Name:  "hello-slapstick",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.innocent-armor.hello-slapstick",
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
					Name:  "silent-kingpin",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stable-changeling.silent-kingpin",
				},
				{
					Name:  "robust-batgirl",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stable-changeling.silent-kingpin.robust-batgirl",
				},
				{
					Name:  "one-batgirl",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stable-changeling.silent-kingpin.robust-batgirl.one-batgirl",
				},
				{
					Name:  "two-batgirl",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stable-changeling.silent-kingpin.robust-batgirl.two-batgirl",
				},
				{
					Name:  "sweet-warhawk",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stable-changeling.silent-kingpin.sweet-warhawk",
				},
				{
					Name:  "flexible-humbug",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stable-changeling.silent-kingpin.flexible-humbug",
				},
			},
			wantErr:  false,
		},
		{
			name: "Move silent-kingpin to pure-slapstick",
			moveName: "silent-kingpin",
			dstName: "pure-slapstick",
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
					Name:  "one-batgirl",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.integral-jungle.robust-batgirl.one-batgirl",
				},
				{
					Name:  "two-batgirl",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.integral-jungle.robust-batgirl.two-batgirl",
				},
				{
					Name:  "sub-batgirl",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.integral-jungle.sub-batgirl",
				},
				{
					Name:  "test-batgirl",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.integral-jungle.test-batgirl",
				},
				{
					Name:  "innocent-armor",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.innocent-armor",
				},
				{
					Name:  "hello-slapstick",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.innocent-armor.hello-slapstick",
				},
				{
					Name:  "pure-slapstick",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.innocent-armor.pure-slapstick",
				},
				{
					Name:  "silent-kingpin",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.innocent-armor.pure-slapstick.silent-kingpin",
				},
				{
					Name:  "flexible-humbug",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.innocent-armor.pure-slapstick.silent-kingpin.flexible-humbug",
				},
				{
					Name:  "sweet-warhawk",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "discrete-whistler.innocent-armor.pure-slapstick.silent-kingpin.sweet-warhawk",
				},
				{
					Name:  "stable-changeling",
					OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
					Paths: "stable-changeling",
				},
			
			},
			wantErr:  false,
		},
	}

	res := folder.GetAllFolders()
	
	for _, tt := range tests {
		driver := folder.NewDriver(res)
		t.Run(tt.name, func(t *testing.T) {
			get, err := driver.MoveFolder(tt.moveName, tt.dstName)

			if (err != nil) != tt.wantErr {
				t.Errorf("MoveFolder() error = %v, wantErr %v", err, tt.wantErr)
			}

			// If no error is expected, check the folder structure
			if !tt.wantErr {

				assert.ElementsMatch(t, tt.want, get, "The returned folders do not match the expected result")
				
			}
		})
	}

}
