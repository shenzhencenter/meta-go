package objects

type BAPIDomain struct {
	Domain          *string `json:"domain,omitempty"`
	InCoolDownUntil *int    `json:"in_cool_down_until,omitempty"`
	IsEligibleForVo *bool   `json:"is_eligible_for_vo,omitempty"`
	IsInCoolDown    *bool   `json:"is_in_cool_down,omitempty"`
}
