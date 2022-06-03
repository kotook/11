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
		Size: 0,
		Path: "https://github.com/kotook/11/blob/4875f9059075d806b25d9c4a47304610760f5c7d/remove/testDir/file1.txt",
	}] = 0
	mapa[sorting.CopyFiles{
		Name: "file1",
		Size: 0,
		Path: "https://github.com/kotook/11/blob/7842f956b789074284262979373b06f29b323f5e/remove/file1.txt",
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
