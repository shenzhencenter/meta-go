package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeSiteLinksSpec struct {
	IsDefaultingEligible       *bool    `json:"is_defaulting_eligible,omitempty"`
	IsSiteLinkSticky           *bool    `json:"is_site_link_sticky,omitempty"`
	SiteLinkClassification     *string  `json:"site_link_classification,omitempty"`
	SiteLinkExtraMetadata      *string  `json:"site_link_extra_metadata,omitempty"`
	SiteLinkHash               *string  `json:"site_link_hash,omitempty"`
	SiteLinkID                 *core.ID `json:"site_link_id,omitempty"`
	SiteLinkImageHash          *string  `json:"site_link_image_hash,omitempty"`
	SiteLinkImageURL           *string  `json:"site_link_image_url,omitempty"`
	SiteLinkLanguage           *string  `json:"site_link_language,omitempty"`
	SiteLinkOriginalURL        *string  `json:"site_link_original_url,omitempty"`
	SiteLinkRecommendationType *string  `json:"site_link_recommendation_type,omitempty"`
	SiteLinkTitle              *string  `json:"site_link_title,omitempty"`
	SiteLinkURL                *string  `json:"site_link_url,omitempty"`
	SiteLinkURLAnchor          *string  `json:"site_link_url_anchor,omitempty"`
	SiteLinkURLRecommenderType *string  `json:"site_link_url_recommender_type,omitempty"`
	SiteLinkWebsiteDataSource  *string  `json:"site_link_website_data_source,omitempty"`
}

var AdCreativeSiteLinksSpecFields = struct {
	IsDefaultingEligible       string
	IsSiteLinkSticky           string
	SiteLinkClassification     string
	SiteLinkExtraMetadata      string
	SiteLinkHash               string
	SiteLinkID                 string
	SiteLinkImageHash          string
	SiteLinkImageURL           string
	SiteLinkLanguage           string
	SiteLinkOriginalURL        string
	SiteLinkRecommendationType string
	SiteLinkTitle              string
	SiteLinkURL                string
	SiteLinkURLAnchor          string
	SiteLinkURLRecommenderType string
	SiteLinkWebsiteDataSource  string
}{
	IsDefaultingEligible:       "is_defaulting_eligible",
	IsSiteLinkSticky:           "is_site_link_sticky",
	SiteLinkClassification:     "site_link_classification",
	SiteLinkExtraMetadata:      "site_link_extra_metadata",
	SiteLinkHash:               "site_link_hash",
	SiteLinkID:                 "site_link_id",
	SiteLinkImageHash:          "site_link_image_hash",
	SiteLinkImageURL:           "site_link_image_url",
	SiteLinkLanguage:           "site_link_language",
	SiteLinkOriginalURL:        "site_link_original_url",
	SiteLinkRecommendationType: "site_link_recommendation_type",
	SiteLinkTitle:              "site_link_title",
	SiteLinkURL:                "site_link_url",
	SiteLinkURLAnchor:          "site_link_url_anchor",
	SiteLinkURLRecommenderType: "site_link_url_recommender_type",
	SiteLinkWebsiteDataSource:  "site_link_website_data_source",
}
