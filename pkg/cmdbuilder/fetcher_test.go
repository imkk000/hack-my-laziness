package cmdbuilder

import (
	"testing"

	"hack/model"

	"github.com/stretchr/testify/assert"
)

func TestFetchCompletion(t *testing.T) {
	want := model.CmdInfo{
		Name:  "root",
		Usage: "Command",
		Commands: []model.CmdInfo{{
			Name:  "encode",
			Usage: "Command Encode",
		}},
	}
	rawJSON := []byte(`{"name":"root","usage":"Command","commands":[{"name":"encode","usage":"Command Encode"}]}`)
	mock := NewMockCommandRunner(t)
	mock.EXPECT().
		Run("test", []string{"completion", "json"}).
		Return(rawJSON, nil)

	info, err := FetchCompletion("test", mock)

	assert.NoError(t, err)
	assert.Equal(t, want, info)
}
