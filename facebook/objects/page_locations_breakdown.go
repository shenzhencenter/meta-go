package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type PageLocationsBreakdown struct {
	LocationID                             *core.ID `json:"location_id,omitempty"`
	LocationName                           *string  `json:"location_name,omitempty"`
	LocationType                           *string  `json:"location_type,omitempty"`
	NumPages                               *int     `json:"num_pages,omitempty"`
	NumPagesEligibleForStoreVisitReporting *int     `json:"num_pages_eligible_for_store_visit_reporting,omitempty"`
	NumUnpublishedOrClosedPages            *int     `json:"num_unpublished_or_closed_pages,omitempty"`
	ParentCountryCode                      *string  `json:"parent_country_code,omitempty"`
	ParentRegionID                         *core.ID `json:"parent_region_id,omitempty"`
	ParentRegionName                       *string  `json:"parent_region_name,omitempty"`
}
