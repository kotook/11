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
		Path: "https://github.com/kotook/11/blob/98460e3fa0b22f2d32c099e6e376009988aecb74/remove/file1.txt",
	}] = 0
	mapa[sorting.CopyFiles{
		Name: "file1",
		Size: 0,
		Path: "https://github.com/kotook/11/blob/925511b0d850b4b64a762c7a4ea495e4df857f15/remove/testDir/file1.txt",
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
