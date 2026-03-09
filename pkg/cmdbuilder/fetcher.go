package cmdbuilder

import (
	"encoding/json"

	"hack/model"
)

func FetchCompletion(binPath string, commander ...CommandRunner) (model.CmdInfo, error) {
	if commander == nil {
		commander = append(commander, NewCommandRunner())
	}
	cmd := commander[0]

	out, err := cmd.Run(binPath, "completion", "json")
	if err != nil {
		return model.CmdInfo{}, err
	}

	var info model.CmdInfo
	if err := json.Unmarshal(out, &info); err != nil {
		return model.CmdInfo{}, err
	}

	return info, nil
}
