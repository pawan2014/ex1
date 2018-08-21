package model

type Machine struct {
	MachineName string `json:"machineName"`
	Machineid   string `json:"machineId"`
	Tags        []Tag  `json:"tags"`
}
type Tag struct {
	TagType   string `json:"tagType"`
	TagName   string `json:"tagName"`
	TagID     string `json:"tagID"`
	Formula   string `json:"formula"`
	Frequency int    `json:"frequency"`
}
