package model

type CmdInfo struct {
	Name     string    `json:"name"`
	Usage    string    `json:"usage"`
	Commands []CmdInfo `json:"commands,omitempty"`
	Flags    []CmdFlag `json:"flags,omitempty"`
}

type CmdFlag struct {
	Type string `json:"type"`
	Name string `json:"name"`
}
