package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ProductItemLandingPageData struct {
	Availability *enums.ProductitemlandingpagedataAvailability `json:"availability,omitempty"`
}

var ProductItemLandingPageDataFields = struct {
	Availability string
}{
	Availability: "availability",
}
