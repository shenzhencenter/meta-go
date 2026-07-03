package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdTopline struct {
	AccountID                *core.ID   `json:"account_id,omitempty"`
	ClientApprovalDate       *time.Time `json:"client_approval_date,omitempty"`
	CreatedBy                *string    `json:"created_by,omitempty"`
	CreatedDate              *time.Time `json:"created_date,omitempty"`
	Description              *string    `json:"description,omitempty"`
	FlightEndDate            *time.Time `json:"flight_end_date,omitempty"`
	FlightStartDate          *time.Time `json:"flight_start_date,omitempty"`
	FuncCapAmount            *string    `json:"func_cap_amount,omitempty"`
	FuncCapAmountWithOffset  *string    `json:"func_cap_amount_with_offset,omitempty"`
	FuncLineAmount           *string    `json:"func_line_amount,omitempty"`
	FuncLineAmountWithOffset *string    `json:"func_line_amount_with_offset,omitempty"`
	FuncPrice                *string    `json:"func_price,omitempty"`
	FuncPriceWithOffset      *string    `json:"func_price_with_offset,omitempty"`
	Gender                   *string    `json:"gender,omitempty"`
	ID                       *core.ID   `json:"id,omitempty"`
	Impressions              *int       `json:"impressions,omitempty"`
	IoNumber                 *int       `json:"io_number,omitempty"`
	IsBonusLine              *int       `json:"is_bonus_line,omitempty"`
	Keywords                 *string    `json:"keywords,omitempty"`
	LastUpdatedBy            *string    `json:"last_updated_by,omitempty"`
	LastUpdatedDate          *time.Time `json:"last_updated_date,omitempty"`
	LineNumber               *int       `json:"line_number,omitempty"`
	LinePosition             *int       `json:"line_position,omitempty"`
	LineType                 *string    `json:"line_type,omitempty"`
	Location                 *string    `json:"location,omitempty"`
	MaxAge                   *string    `json:"max_age,omitempty"`
	MaxBudget                *string    `json:"max_budget,omitempty"`
	MinAge                   *string    `json:"min_age,omitempty"`
	PricePerTrp              *string    `json:"price_per_trp,omitempty"`
	ProductType              *string    `json:"product_type,omitempty"`
	RevAssuranceApprovalDate *time.Time `json:"rev_assurance_approval_date,omitempty"`
	Targets                  *string    `json:"targets,omitempty"`
	TrpUpdatedTime           *int       `json:"trp_updated_time,omitempty"`
	TrpValue                 *string    `json:"trp_value,omitempty"`
	Uom                      *string    `json:"uom,omitempty"`
}
