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
