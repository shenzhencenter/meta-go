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
