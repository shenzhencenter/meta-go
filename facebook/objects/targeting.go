package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type Targeting struct {
	AdgroupID                               *core.ID                                   `json:"adgroup_id,omitempty"`
	AgeMax                                  *uint64                                    `json:"age_max,omitempty"`
	AgeMin                                  *uint64                                    `json:"age_min,omitempty"`
	AgeRange                                *[]uint64                                  `json:"age_range,omitempty"`
	AppInstallState                         *string                                    `json:"app_install_state,omitempty"`
	AudienceNetworkPositions                *[]string                                  `json:"audience_network_positions,omitempty"`
	Behaviors                               *[]IDName                                  `json:"behaviors,omitempty"`
	BrandSafetyContentFilterLevels          *[]string                                  `json:"brand_safety_content_filter_levels,omitempty"`
	CatalogBasedTargeting                   *CatalogBasedTargeting                     `json:"catalog_based_targeting,omitempty"`
	Cities                                  *[]IDName                                  `json:"cities,omitempty"`
	CollegeYears                            *[]uint64                                  `json:"college_years,omitempty"`
	Connections                             *[]ConnectionsTargeting                    `json:"connections,omitempty"`
	ContextualTargetingCategories           *[]IDName                                  `json:"contextual_targeting_categories,omitempty"`
	Countries                               *[]string                                  `json:"countries,omitempty"`
	Country                                 *[]string                                  `json:"country,omitempty"`
	CountryGroups                           *[]string                                  `json:"country_groups,omitempty"`
	CustomAudiences                         *[]RawCustomAudience                       `json:"custom_audiences,omitempty"`
	DevicePlatforms                         *[]enums.TargetingDevicePlatforms          `json:"device_platforms,omitempty"`
	DirectInstallDevices                    *bool                                      `json:"direct_install_devices,omitempty"`
	DynamicAudienceIds                      *[]core.ID                                 `json:"dynamic_audience_ids,omitempty"`
	EducationMajors                         *[]IDName                                  `json:"education_majors,omitempty"`
	EducationSchools                        *[]IDName                                  `json:"education_schools,omitempty"`
	EducationStatuses                       *[]uint64                                  `json:"education_statuses,omitempty"`
	EffectiveAudienceNetworkPositions       *[]string                                  `json:"effective_audience_network_positions,omitempty"`
	EffectiveBrandSafetyContentFilterLevels *[]string                                  `json:"effective_brand_safety_content_filter_levels,omitempty"`
	EffectiveDevicePlatforms                *[]enums.TargetingEffectiveDevicePlatforms `json:"effective_device_platforms,omitempty"`
	EffectiveFacebookPositions              *[]string                                  `json:"effective_facebook_positions,omitempty"`
	EffectiveInstagramPositions             *[]string                                  `json:"effective_instagram_positions,omitempty"`
	EffectiveMessengerPositions             *[]string                                  `json:"effective_messenger_positions,omitempty"`
	EffectivePublisherPlatforms             *[]string                                  `json:"effective_publisher_platforms,omitempty"`
	EffectiveThreadsPositions               *[]string                                  `json:"effective_threads_positions,omitempty"`
	EffectiveWhatsappPositions              *[]string                                  `json:"effective_whatsapp_positions,omitempty"`
	EngagementSpecs                         *[]TargetingDynamicRule                    `json:"engagement_specs,omitempty"`
	EthnicAffinity                          *[]IDName                                  `json:"ethnic_affinity,omitempty"`
	ExcludeReachedSince                     *[]string                                  `json:"exclude_reached_since,omitempty"`
	ExcludedBrandSafetyContentTypes         *[]string                                  `json:"excluded_brand_safety_content_types,omitempty"`
	ExcludedConnections                     *[]ConnectionsTargeting                    `json:"excluded_connections,omitempty"`
	ExcludedCustomAudiences                 *[]RawCustomAudience                       `json:"excluded_custom_audiences,omitempty"`
	ExcludedDynamicAudienceIds              *[]core.ID                                 `json:"excluded_dynamic_audience_ids,omitempty"`
	ExcludedEngagementSpecs                 *[]TargetingDynamicRule                    `json:"excluded_engagement_specs,omitempty"`
	ExcludedGeoLocations                    *TargetingGeoLocation                      `json:"excluded_geo_locations,omitempty"`
	ExcludedMobileDeviceModel               *[]string                                  `json:"excluded_mobile_device_model,omitempty"`
	ExcludedProductAudienceSpecs            *[]TargetingProductAudienceSpec            `json:"excluded_product_audience_specs,omitempty"`
	ExcludedPublisherCategories             *[]string                                  `json:"excluded_publisher_categories,omitempty"`
	ExcludedPublisherListIds                *[]core.ID                                 `json:"excluded_publisher_list_ids,omitempty"`
	ExcludedUserDevice                      *[]string                                  `json:"excluded_user_device,omitempty"`
	Exclusions                              *FlexibleTargeting                         `json:"exclusions,omitempty"`
	FacebookPositions                       *[]string                                  `json:"facebook_positions,omitempty"`
	FamilyStatuses                          *[]IDName                                  `json:"family_statuses,omitempty"`
	FbDealID                                *core.ID                                   `json:"fb_deal_id,omitempty"`
	FlexibleSpec                            *[]FlexibleTargeting                       `json:"flexible_spec,omitempty"`
	FriendsOfConnections                    *[]ConnectionsTargeting                    `json:"friends_of_connections,omitempty"`
	Genders                                 *[]uint64                                  `json:"genders,omitempty"`
	Generation                              *[]IDName                                  `json:"generation,omitempty"`
	GeoLocations                            *TargetingGeoLocation                      `json:"geo_locations,omitempty"`
	HomeOwnership                           *[]IDName                                  `json:"home_ownership,omitempty"`
	HomeType                                *[]IDName                                  `json:"home_type,omitempty"`
	HomeValue                               *[]IDName                                  `json:"home_value,omitempty"`
	HouseholdComposition                    *[]IDName                                  `json:"household_composition,omitempty"`
	Income                                  *[]IDName                                  `json:"income,omitempty"`
	Industries                              *[]IDName                                  `json:"industries,omitempty"`
	InstagramPositions                      *[]string                                  `json:"instagram_positions,omitempty"`
	InstallStateApplication                 *string                                    `json:"install_state_application,omitempty"`
	InstreamVideoSkippableExcluded          *bool                                      `json:"instream_video_skippable_excluded,omitempty"`
	InterestedIn                            *[]uint64                                  `json:"interested_in,omitempty"`
	Interests                               *[]IDName                                  `json:"interests,omitempty"`
	IsWhatsappDestinationAd                 *bool                                      `json:"is_whatsapp_destination_ad,omitempty"`
	Keywords                                *[]string                                  `json:"keywords,omitempty"`
	LifeEvents                              *[]IDName                                  `json:"life_events,omitempty"`
	Locales                                 *[]uint64                                  `json:"locales,omitempty"`
	MessengerPositions                      *[]string                                  `json:"messenger_positions,omitempty"`
	Moms                                    *[]IDName                                  `json:"moms,omitempty"`
	NetWorth                                *[]IDName                                  `json:"net_worth,omitempty"`
	OfficeType                              *[]IDName                                  `json:"office_type,omitempty"`
	PlacePageSetIds                         *[]core.ID                                 `json:"place_page_set_ids,omitempty"`
	PoliticalViews                          *[]uint64                                  `json:"political_views,omitempty"`
	Politics                                *[]IDName                                  `json:"politics,omitempty"`
	ProductAudienceSpecs                    *[]TargetingProductAudienceSpec            `json:"product_audience_specs,omitempty"`
	ProspectingAudience                     *TargetingProspectingAudience              `json:"prospecting_audience,omitempty"`
	PublisherPlatforms                      *[]string                                  `json:"publisher_platforms,omitempty"`
	Radius                                  *string                                    `json:"radius,omitempty"`
	Regions                                 *[]IDName                                  `json:"regions,omitempty"`
	RelationshipStatuses                    *[]uint64                                  `json:"relationship_statuses,omitempty"`
	SiteCategory                            *[]string                                  `json:"site_category,omitempty"`
	SubscriberUniverse                      *TargetingSubscriberUniverse               `json:"subscriber_universe,omitempty"`
	TargetingAutomation                     *TargetingAutomation                       `json:"targeting_automation,omitempty"`
	TargetingOptimization                   *string                                    `json:"targeting_optimization,omitempty"`
	TargetingRelaxationTypes                *TargetingRelaxation                       `json:"targeting_relaxation_types,omitempty"`
	ThreadsPositions                        *[]string                                  `json:"threads_positions,omitempty"`
	UserAdclusters                          *[]IDName                                  `json:"user_adclusters,omitempty"`
	UserAgeUnknown                          *bool                                      `json:"user_age_unknown,omitempty"`
	UserDevice                              *[]string                                  `json:"user_device,omitempty"`
	UserEvent                               *[]uint64                                  `json:"user_event,omitempty"`
	UserOs                                  *[]string                                  `json:"user_os,omitempty"`
	WhatsappPositions                       *[]string                                  `json:"whatsapp_positions,omitempty"`
	WirelessCarrier                         *[]string                                  `json:"wireless_carrier,omitempty"`
	WorkEmployers                           *[]IDName                                  `json:"work_employers,omitempty"`
	WorkPositions                           *[]IDName                                  `json:"work_positions,omitempty"`
	Zips                                    *[]string                                  `json:"zips,omitempty"`
}

var TargetingFields = struct {
	AdgroupID                               string
	AgeMax                                  string
	AgeMin                                  string
	AgeRange                                string
	AppInstallState                         string
	AudienceNetworkPositions                string
	Behaviors                               string
	BrandSafetyContentFilterLevels          string
	CatalogBasedTargeting                   string
	Cities                                  string
	CollegeYears                            string
	Connections                             string
	ContextualTargetingCategories           string
	Countries                               string
	Country                                 string
	CountryGroups                           string
	CustomAudiences                         string
	DevicePlatforms                         string
	DirectInstallDevices                    string
	DynamicAudienceIds                      string
	EducationMajors                         string
	EducationSchools                        string
	EducationStatuses                       string
	EffectiveAudienceNetworkPositions       string
	EffectiveBrandSafetyContentFilterLevels string
	EffectiveDevicePlatforms                string
	EffectiveFacebookPositions              string
	EffectiveInstagramPositions             string
	EffectiveMessengerPositions             string
	EffectivePublisherPlatforms             string
	EffectiveThreadsPositions               string
	EffectiveWhatsappPositions              string
	EngagementSpecs                         string
	EthnicAffinity                          string
	ExcludeReachedSince                     string
	ExcludedBrandSafetyContentTypes         string
	ExcludedConnections                     string
	ExcludedCustomAudiences                 string
	ExcludedDynamicAudienceIds              string
	ExcludedEngagementSpecs                 string
	ExcludedGeoLocations                    string
	ExcludedMobileDeviceModel               string
	ExcludedProductAudienceSpecs            string
	ExcludedPublisherCategories             string
	ExcludedPublisherListIds                string
	ExcludedUserDevice                      string
	Exclusions                              string
	FacebookPositions                       string
	FamilyStatuses                          string
	FbDealID                                string
	FlexibleSpec                            string
	FriendsOfConnections                    string
	Genders                                 string
	Generation                              string
	GeoLocations                            string
	HomeOwnership                           string
	HomeType                                string
	HomeValue                               string
	HouseholdComposition                    string
	Income                                  string
	Industries                              string
	InstagramPositions                      string
	InstallStateApplication                 string
	InstreamVideoSkippableExcluded          string
	InterestedIn                            string
	Interests                               string
	IsWhatsappDestinationAd                 string
	Keywords                                string
	LifeEvents                              string
	Locales                                 string
	MessengerPositions                      string
	Moms                                    string
	NetWorth                                string
	OfficeType                              string
	PlacePageSetIds                         string
	PoliticalViews                          string
	Politics                                string
	ProductAudienceSpecs                    string
	ProspectingAudience                     string
	PublisherPlatforms                      string
	Radius                                  string
	Regions                                 string
	RelationshipStatuses                    string
	SiteCategory                            string
	SubscriberUniverse                      string
	TargetingAutomation                     string
	TargetingOptimization                   string
	TargetingRelaxationTypes                string
	ThreadsPositions                        string
	UserAdclusters                          string
	UserAgeUnknown                          string
	UserDevice                              string
	UserEvent                               string
	UserOs                                  string
	WhatsappPositions                       string
	WirelessCarrier                         string
	WorkEmployers                           string
	WorkPositions                           string
	Zips                                    string
}{
	AdgroupID:                               "adgroup_id",
	AgeMax:                                  "age_max",
	AgeMin:                                  "age_min",
	AgeRange:                                "age_range",
	AppInstallState:                         "app_install_state",
	AudienceNetworkPositions:                "audience_network_positions",
	Behaviors:                               "behaviors",
	BrandSafetyContentFilterLevels:          "brand_safety_content_filter_levels",
	CatalogBasedTargeting:                   "catalog_based_targeting",
	Cities:                                  "cities",
	CollegeYears:                            "college_years",
	Connections:                             "connections",
	ContextualTargetingCategories:           "contextual_targeting_categories",
	Countries:                               "countries",
	Country:                                 "country",
	CountryGroups:                           "country_groups",
	CustomAudiences:                         "custom_audiences",
	DevicePlatforms:                         "device_platforms",
	DirectInstallDevices:                    "direct_install_devices",
	DynamicAudienceIds:                      "dynamic_audience_ids",
	EducationMajors:                         "education_majors",
	EducationSchools:                        "education_schools",
	EducationStatuses:                       "education_statuses",
	EffectiveAudienceNetworkPositions:       "effective_audience_network_positions",
	EffectiveBrandSafetyContentFilterLevels: "effective_brand_safety_content_filter_levels",
	EffectiveDevicePlatforms:                "effective_device_platforms",
	EffectiveFacebookPositions:              "effective_facebook_positions",
	EffectiveInstagramPositions:             "effective_instagram_positions",
	EffectiveMessengerPositions:             "effective_messenger_positions",
	EffectivePublisherPlatforms:             "effective_publisher_platforms",
	EffectiveThreadsPositions:               "effective_threads_positions",
	EffectiveWhatsappPositions:              "effective_whatsapp_positions",
	EngagementSpecs:                         "engagement_specs",
	EthnicAffinity:                          "ethnic_affinity",
	ExcludeReachedSince:                     "exclude_reached_since",
	ExcludedBrandSafetyContentTypes:         "excluded_brand_safety_content_types",
	ExcludedConnections:                     "excluded_connections",
	ExcludedCustomAudiences:                 "excluded_custom_audiences",
	ExcludedDynamicAudienceIds:              "excluded_dynamic_audience_ids",
	ExcludedEngagementSpecs:                 "excluded_engagement_specs",
	ExcludedGeoLocations:                    "excluded_geo_locations",
	ExcludedMobileDeviceModel:               "excluded_mobile_device_model",
	ExcludedProductAudienceSpecs:            "excluded_product_audience_specs",
	ExcludedPublisherCategories:             "excluded_publisher_categories",
	ExcludedPublisherListIds:                "excluded_publisher_list_ids",
	ExcludedUserDevice:                      "excluded_user_device",
	Exclusions:                              "exclusions",
	FacebookPositions:                       "facebook_positions",
	FamilyStatuses:                          "family_statuses",
	FbDealID:                                "fb_deal_id",
	FlexibleSpec:                            "flexible_spec",
	FriendsOfConnections:                    "friends_of_connections",
	Genders:                                 "genders",
	Generation:                              "generation",
	GeoLocations:                            "geo_locations",
	HomeOwnership:                           "home_ownership",
	HomeType:                                "home_type",
	HomeValue:                               "home_value",
	HouseholdComposition:                    "household_composition",
	Income:                                  "income",
	Industries:                              "industries",
	InstagramPositions:                      "instagram_positions",
	InstallStateApplication:                 "install_state_application",
	InstreamVideoSkippableExcluded:          "instream_video_skippable_excluded",
	InterestedIn:                            "interested_in",
	Interests:                               "interests",
	IsWhatsappDestinationAd:                 "is_whatsapp_destination_ad",
	Keywords:                                "keywords",
	LifeEvents:                              "life_events",
	Locales:                                 "locales",
	MessengerPositions:                      "messenger_positions",
	Moms:                                    "moms",
	NetWorth:                                "net_worth",
	OfficeType:                              "office_type",
	PlacePageSetIds:                         "place_page_set_ids",
	PoliticalViews:                          "political_views",
	Politics:                                "politics",
	ProductAudienceSpecs:                    "product_audience_specs",
	ProspectingAudience:                     "prospecting_audience",
	PublisherPlatforms:                      "publisher_platforms",
	Radius:                                  "radius",
	Regions:                                 "regions",
	RelationshipStatuses:                    "relationship_statuses",
	SiteCategory:                            "site_category",
	SubscriberUniverse:                      "subscriber_universe",
	TargetingAutomation:                     "targeting_automation",
	TargetingOptimization:                   "targeting_optimization",
	TargetingRelaxationTypes:                "targeting_relaxation_types",
	ThreadsPositions:                        "threads_positions",
	UserAdclusters:                          "user_adclusters",
	UserAgeUnknown:                          "user_age_unknown",
	UserDevice:                              "user_device",
	UserEvent:                               "user_event",
	UserOs:                                  "user_os",
	WhatsappPositions:                       "whatsapp_positions",
	WirelessCarrier:                         "wireless_carrier",
	WorkEmployers:                           "work_employers",
	WorkPositions:                           "work_positions",
	Zips:                                    "zips",
}
