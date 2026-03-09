package cmdbuilder

import (
	"io/fs"
	"testing"

	"hack/pkg/mocks"

	"github.com/stretchr/testify/assert"
)

func TestDiscoveryBinaries(t *testing.T) {
	want := BinaryFile{
		Name:     "hack-core",
		FullPath: ".bin/hack-core",
	}
	mockDirEntry := mocks.NewMockDirEntry(t)
	mockDirEntry.EXPECT().Name().Return("hack-core")
	mockDirEntry.EXPECT().IsDir().Return(false)
	mockFileInfo := mocks.NewMockFileInfo(t)
	mockFileInfo.EXPECT().Mode().Return(fs.FileMode(0o1111))
	mock := NewMockFileSystem(t)
	mock.EXPECT().ReadDir(".bin").Return([]fs.DirEntry{mockDirEntry}, nil)
	mock.EXPECT().Stat(".bin/hack-core").Return(mockFileInfo, nil)

	files := DiscoverBinaries(".bin", "hack-", mock)

	assert.NotEmpty(t, files)
	assert.Equal(t, want, files[0])
}
