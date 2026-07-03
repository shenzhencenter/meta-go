package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type UserIDForApp struct {
	App                 *Application `json:"app,omitempty"`
	ID                  *core.ID     `json:"id,omitempty"`
	InstantGamePlayerID *core.ID     `json:"instant_game_player_id,omitempty"`
}
