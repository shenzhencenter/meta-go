package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
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

var VideoPollFields = struct {
	CloseAfterVoting string
	DefaultOpen      string
	ID               string
	Question         string
	ShowGradient     string
	ShowResults      string
	Status           string
}{
	CloseAfterVoting: "close_after_voting",
	DefaultOpen:      "default_open",
	ID:               "id",
	Question:         "question",
	ShowGradient:     "show_gradient",
	ShowResults:      "show_results",
	Status:           "status",
}
