package remove

import (
	sorting "newmod3/sortingFile"
	"testing"
)

func TestRemove(t *testing.T) {
	type args struct {
		duplicateAmount map[sorting.CopyFiles]int
		yes             string
	}

	mapa := map[sorting.CopyFiles]int{}
	mapa[sorting.CopyFiles{
		Name: "file1",
		Size: 2,
		Path: "https://github.com/kotook/11/blob/56d3fb8ded6f50c482fb8fb5fc98dbfa34ab081d/remove/file1.txt",
	}] = 0
	mapa[sorting.CopyFiles{
		Name: "file1",
		Size: 2,
		Path: "https://github.com/kotook/11/blob/56d3fb8ded6f50c482fb8fb5fc98dbfa34ab081d/remove/testDir/file1.txt",
	}] = 1

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"file1",
			args{
				duplicateAmount: mapa,
				yes:             "Y",
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Remove(tt.args.duplicateAmount, tt.args.yes); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
