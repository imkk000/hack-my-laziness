package cmdbuilder

import "hack/model"

const binaryNamePrefix = "hack-"

type BinaryFile struct {
	FullPath string
	Name     string
}

type CommandGroup struct {
	BinaryFile
	CmdInfos []model.CmdInfo
}
