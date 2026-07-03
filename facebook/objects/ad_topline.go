package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdTopline struct {
	AccountID                *core.ID   `json:"account_id,omitempty"`
	ClientApprovalDate       *core.Time `json:"client_approval_date,omitempty"`
	CreatedBy                *string    `json:"created_by,omitempty"`
	CreatedDate              *core.Time `json:"created_date,omitempty"`
	Description              *string    `json:"description,omitempty"`
	FlightEndDate            *core.Time `json:"flight_end_date,omitempty"`
	FlightStartDate          *core.Time `json:"flight_start_date,omitempty"`
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
	LastUpdatedDate          *core.Time `json:"last_updated_date,omitempty"`
	LineNumber               *int       `json:"line_number,omitempty"`
	LinePosition             *int       `json:"line_position,omitempty"`
	LineType                 *string    `json:"line_type,omitempty"`
	Location                 *string    `json:"location,omitempty"`
	MaxAge                   *string    `json:"max_age,omitempty"`
	MaxBudget                *string    `json:"max_budget,omitempty"`
	MinAge                   *string    `json:"min_age,omitempty"`
	PricePerTrp              *string    `json:"price_per_trp,omitempty"`
	ProductType              *string    `json:"product_type,omitempty"`
	RevAssuranceApprovalDate *core.Time `json:"rev_assurance_approval_date,omitempty"`
	Targets                  *string    `json:"targets,omitempty"`
	TrpUpdatedTime           *int       `json:"trp_updated_time,omitempty"`
	TrpValue                 *string    `json:"trp_value,omitempty"`
	Uom                      *string    `json:"uom,omitempty"`
}

var AdToplineFields = struct {
	AccountID                string
	ClientApprovalDate       string
	CreatedBy                string
	CreatedDate              string
	Description              string
	FlightEndDate            string
	FlightStartDate          string
	FuncCapAmount            string
	FuncCapAmountWithOffset  string
	FuncLineAmount           string
	FuncLineAmountWithOffset string
	FuncPrice                string
	FuncPriceWithOffset      string
	Gender                   string
	ID                       string
	Impressions              string
	IoNumber                 string
	IsBonusLine              string
	Keywords                 string
	LastUpdatedBy            string
	LastUpdatedDate          string
	LineNumber               string
	LinePosition             string
	LineType                 string
	Location                 string
	MaxAge                   string
	MaxBudget                string
	MinAge                   string
	PricePerTrp              string
	ProductType              string
	RevAssuranceApprovalDate string
	Targets                  string
	TrpUpdatedTime           string
	TrpValue                 string
	Uom                      string
}{
	AccountID:                "account_id",
	ClientApprovalDate:       "client_approval_date",
	CreatedBy:                "created_by",
	CreatedDate:              "created_date",
	Description:              "description",
	FlightEndDate:            "flight_end_date",
	FlightStartDate:          "flight_start_date",
	FuncCapAmount:            "func_cap_amount",
	FuncCapAmountWithOffset:  "func_cap_amount_with_offset",
	FuncLineAmount:           "func_line_amount",
	FuncLineAmountWithOffset: "func_line_amount_with_offset",
	FuncPrice:                "func_price",
	FuncPriceWithOffset:      "func_price_with_offset",
	Gender:                   "gender",
	ID:                       "id",
	Impressions:              "impressions",
	IoNumber:                 "io_number",
	IsBonusLine:              "is_bonus_line",
	Keywords:                 "keywords",
	LastUpdatedBy:            "last_updated_by",
	LastUpdatedDate:          "last_updated_date",
	LineNumber:               "line_number",
	LinePosition:             "line_position",
	LineType:                 "line_type",
	Location:                 "location",
	MaxAge:                   "max_age",
	MaxBudget:                "max_budget",
	MinAge:                   "min_age",
	PricePerTrp:              "price_per_trp",
	ProductType:              "product_type",
	RevAssuranceApprovalDate: "rev_assurance_approval_date",
	Targets:                  "targets",
	TrpUpdatedTime:           "trp_updated_time",
	TrpValue:                 "trp_value",
	Uom:                      "uom",
}
