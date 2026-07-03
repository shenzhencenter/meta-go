package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Robot struct {
	BringupVars        *[]map[string]string  `json:"bringup_vars,omitempty"`
	Configurations     *[]map[string]string  `json:"configurations,omitempty"`
	DataCenter         *string               `json:"data_center,omitempty"`
	FwImage            *string               `json:"fw_image,omitempty"`
	ID                 *core.ID              `json:"id,omitempty"`
	InitPos            *[]map[string]float64 `json:"init_pos,omitempty"`
	LastPos            *[]map[string]float64 `json:"last_pos,omitempty"`
	MeetupLinkHash     *string               `json:"meetup_link_hash,omitempty"`
	ReleaseImage       *string               `json:"release_image,omitempty"`
	RobotNotes         *string               `json:"robot_notes,omitempty"`
	Suite              *string               `json:"suite,omitempty"`
	TargetFwImage      *string               `json:"target_fw_image,omitempty"`
	TargetFwImageURI   *string               `json:"target_fw_image_uri,omitempty"`
	TargetMapImageURI  *string               `json:"target_map_image_uri,omitempty"`
	TargetOsImageURI   *string               `json:"target_os_image_uri,omitempty"`
	TargetReleaseImage *string               `json:"target_release_image,omitempty"`
	TargetSwImageURI   *string               `json:"target_sw_image_uri,omitempty"`
	User               *User                 `json:"user,omitempty"`
}

var RobotFields = struct {
	BringupVars        string
	Configurations     string
	DataCenter         string
	FwImage            string
	ID                 string
	InitPos            string
	LastPos            string
	MeetupLinkHash     string
	ReleaseImage       string
	RobotNotes         string
	Suite              string
	TargetFwImage      string
	TargetFwImageURI   string
	TargetMapImageURI  string
	TargetOsImageURI   string
	TargetReleaseImage string
	TargetSwImageURI   string
	User               string
}{
	BringupVars:        "bringup_vars",
	Configurations:     "configurations",
	DataCenter:         "data_center",
	FwImage:            "fw_image",
	ID:                 "id",
	InitPos:            "init_pos",
	LastPos:            "last_pos",
	MeetupLinkHash:     "meetup_link_hash",
	ReleaseImage:       "release_image",
	RobotNotes:         "robot_notes",
	Suite:              "suite",
	TargetFwImage:      "target_fw_image",
	TargetFwImageURI:   "target_fw_image_uri",
	TargetMapImageURI:  "target_map_image_uri",
	TargetOsImageURI:   "target_os_image_uri",
	TargetReleaseImage: "target_release_image",
	TargetSwImageURI:   "target_sw_image_uri",
	User:               "user",
}
