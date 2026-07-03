package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type VideoPoll struct {
	CloseAfterVoting *bool                  `json:"close_after_voting,omitempty"`
	DefaultOpen      *bool                  `json:"default_open,omitempty"`
	ID               *core.ID               `json:"id,omitempty"`
	Question         *string                `json:"question,omitempty"`
	ShowGradient     *bool                  `json:"show_gradient,omitempty"`
	ShowResults      *bool                  `json:"show_results,omitempty"`
	Status           *enums.VideopollStatus `json:"status,omitempty"`
}
