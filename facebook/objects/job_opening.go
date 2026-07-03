package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type JobOpening struct {
	Address                    *string                                   `json:"address,omitempty"`
	ApplicationCallbackURL     *string                                   `json:"application_callback_url,omitempty"`
	CreatedTime                *core.Time                                `json:"created_time,omitempty"`
	Description                *string                                   `json:"description,omitempty"`
	Errors                     *[]string                                 `json:"errors,omitempty"`
	ExternalCompanyFacebookURL *string                                   `json:"external_company_facebook_url,omitempty"`
	ExternalCompanyFullAddress *string                                   `json:"external_company_full_address,omitempty"`
	ExternalCompanyID          *core.ID                                  `json:"external_company_id,omitempty"`
	ExternalCompanyName        *string                                   `json:"external_company_name,omitempty"`
	ExternalID                 *core.ID                                  `json:"external_id,omitempty"`
	ID                         *core.ID                                  `json:"id,omitempty"`
	JobStatus                  *enums.JobopeningJobStatus                `json:"job_status,omitempty"`
	Latitude                   *float64                                  `json:"latitude,omitempty"`
	Longitude                  *float64                                  `json:"longitude,omitempty"`
	OffsiteApplicationURL      *string                                   `json:"offsite_application_url,omitempty"`
	Page                       *Page                                     `json:"page,omitempty"`
	Photo                      *Photo                                    `json:"photo,omitempty"`
	PlatformReviewStatus       *enums.JobopeningPlatformReviewStatus     `json:"platform_review_status,omitempty"`
	Post                       *Post                                     `json:"post,omitempty"`
	RemoteType                 *string                                   `json:"remote_type,omitempty"`
	ReviewRejectionReasons     *[]enums.JobopeningReviewRejectionReasons `json:"review_rejection_reasons,omitempty"`
	Title                      *string                                   `json:"title,omitempty"`
	Type                       *enums.JobopeningType                     `json:"type,omitempty"`
}

var JobOpeningFields = struct {
	Address                    string
	ApplicationCallbackURL     string
	CreatedTime                string
	Description                string
	Errors                     string
	ExternalCompanyFacebookURL string
	ExternalCompanyFullAddress string
	ExternalCompanyID          string
	ExternalCompanyName        string
	ExternalID                 string
	ID                         string
	JobStatus                  string
	Latitude                   string
	Longitude                  string
	OffsiteApplicationURL      string
	Page                       string
	Photo                      string
	PlatformReviewStatus       string
	Post                       string
	RemoteType                 string
	ReviewRejectionReasons     string
	Title                      string
	Type                       string
}{
	Address:                    "address",
	ApplicationCallbackURL:     "application_callback_url",
	CreatedTime:                "created_time",
	Description:                "description",
	Errors:                     "errors",
	ExternalCompanyFacebookURL: "external_company_facebook_url",
	ExternalCompanyFullAddress: "external_company_full_address",
	ExternalCompanyID:          "external_company_id",
	ExternalCompanyName:        "external_company_name",
	ExternalID:                 "external_id",
	ID:                         "id",
	JobStatus:                  "job_status",
	Latitude:                   "latitude",
	Longitude:                  "longitude",
	OffsiteApplicationURL:      "offsite_application_url",
	Page:                       "page",
	Photo:                      "photo",
	PlatformReviewStatus:       "platform_review_status",
	Post:                       "post",
	RemoteType:                 "remote_type",
	ReviewRejectionReasons:     "review_rejection_reasons",
	Title:                      "title",
	Type:                       "type",
}
