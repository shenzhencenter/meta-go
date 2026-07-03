package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AutomotiveModel struct {
	Applinks           *CatalogItemAppLinks                   `json:"applinks,omitempty"`
	AutomotiveModelID  *core.ID                               `json:"automotive_model_id,omitempty"`
	Availability       *string                                `json:"availability,omitempty"`
	BodyStyle          *string                                `json:"body_style,omitempty"`
	Currency           *string                                `json:"currency,omitempty"`
	CustomLabelX0      *string                                `json:"custom_label_0,omitempty"`
	CustomLabelX1      *string                                `json:"custom_label_1,omitempty"`
	CustomLabelX2      *string                                `json:"custom_label_2,omitempty"`
	CustomLabelX3      *string                                `json:"custom_label_3,omitempty"`
	CustomLabelX4      *string                                `json:"custom_label_4,omitempty"`
	CustomNumberX0     *uint64                                `json:"custom_number_0,omitempty"`
	CustomNumberX1     *uint64                                `json:"custom_number_1,omitempty"`
	CustomNumberX2     *uint64                                `json:"custom_number_2,omitempty"`
	CustomNumberX3     *uint64                                `json:"custom_number_3,omitempty"`
	CustomNumberX4     *uint64                                `json:"custom_number_4,omitempty"`
	Description        *string                                `json:"description,omitempty"`
	Drivetrain         *string                                `json:"drivetrain,omitempty"`
	ExteriorColor      *string                                `json:"exterior_color,omitempty"`
	FinanceDescription *string                                `json:"finance_description,omitempty"`
	FinanceType        *string                                `json:"finance_type,omitempty"`
	FuelType           *string                                `json:"fuel_type,omitempty"`
	Generation         *string                                `json:"generation,omitempty"`
	ID                 *core.ID                               `json:"id,omitempty"`
	ImageFetchStatus   *enums.AutomotivemodelImageFetchStatus `json:"image_fetch_status,omitempty"`
	Images             *[]string                              `json:"images,omitempty"`
	InteriorColor      *string                                `json:"interior_color,omitempty"`
	InteriorUpholstery *string                                `json:"interior_upholstery,omitempty"`
	Make               *string                                `json:"make,omitempty"`
	Model              *string                                `json:"model,omitempty"`
	Price              *string                                `json:"price,omitempty"`
	SanitizedImages    *[]string                              `json:"sanitized_images,omitempty"`
	Title              *string                                `json:"title,omitempty"`
	Transmission       *string                                `json:"transmission,omitempty"`
	Trim               *string                                `json:"trim,omitempty"`
	UnitPrice          *map[string]interface{}                `json:"unit_price,omitempty"`
	URL                *string                                `json:"url,omitempty"`
	Visibility         *enums.AutomotivemodelVisibility       `json:"visibility,omitempty"`
	Year               *int                                   `json:"year,omitempty"`
}
