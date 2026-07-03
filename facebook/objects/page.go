package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Page struct {
	About                                                *string                            `json:"about,omitempty"`
	AccessToken                                          *string                            `json:"access_token,omitempty"`
	AdCampaign                                           *AdSet                             `json:"ad_campaign,omitempty"`
	Affiliation                                          *string                            `json:"affiliation,omitempty"`
	AppID                                                *core.ID                           `json:"app_id,omitempty"`
	ArtistsWeLike                                        *string                            `json:"artists_we_like,omitempty"`
	Attire                                               *string                            `json:"attire,omitempty"`
	AvailablePromoOfferIds                               *[]map[string][]map[string]core.ID `json:"available_promo_offer_ids,omitempty"`
	Awards                                               *string                            `json:"awards,omitempty"`
	BandInterests                                        *string                            `json:"band_interests,omitempty"`
	BandMembers                                          *string                            `json:"band_members,omitempty"`
	BestPage                                             *Page                              `json:"best_page,omitempty"`
	Bio                                                  *string                            `json:"bio,omitempty"`
	Birthday                                             *string                            `json:"birthday,omitempty"`
	BookingAgent                                         *string                            `json:"booking_agent,omitempty"`
	BreakingNewsUsage                                    *map[string]interface{}            `json:"breaking_news_usage,omitempty"`
	Built                                                *string                            `json:"built,omitempty"`
	Business                                             *map[string]interface{}            `json:"business,omitempty"`
	CanCheckin                                           *bool                              `json:"can_checkin,omitempty"`
	CanPost                                              *bool                              `json:"can_post,omitempty"`
	Category                                             *string                            `json:"category,omitempty"`
	CategoryList                                         *[]PageCategory                    `json:"category_list,omitempty"`
	Checkins                                             *uint64                            `json:"checkins,omitempty"`
	CompanyOverview                                      *string                            `json:"company_overview,omitempty"`
	ConnectedInstagramAccount                            *IGUser                            `json:"connected_instagram_account,omitempty"`
	ConnectedPageBackedInstagramAccount                  *IGUser                            `json:"connected_page_backed_instagram_account,omitempty"`
	ContactAddress                                       *MailingAddress                    `json:"contact_address,omitempty"`
	CopyrightAttributionInsights                         *CopyrightAttributionInsights      `json:"copyright_attribution_insights,omitempty"`
	CopyrightWhitelistedIgPartners                       *[]string                          `json:"copyright_whitelisted_ig_partners,omitempty"`
	CountryPageLikes                                     *uint64                            `json:"country_page_likes,omitempty"`
	Cover                                                *CoverPhoto                        `json:"cover,omitempty"`
	CulinaryTeam                                         *string                            `json:"culinary_team,omitempty"`
	CurrentLocation                                      *string                            `json:"current_location,omitempty"`
	DeliveryAndPickupOptionInfo                          *[]string                          `json:"delivery_and_pickup_option_info,omitempty"`
	Description                                          *string                            `json:"description,omitempty"`
	DescriptionHTML                                      *string                            `json:"description_html,omitempty"`
	DifferentlyOpenOfferings                             *[]map[string]bool                 `json:"differently_open_offerings,omitempty"`
	DirectedBy                                           *string                            `json:"directed_by,omitempty"`
	DisplaySubtext                                       *string                            `json:"display_subtext,omitempty"`
	DisplayedMessageResponseTime                         *string                            `json:"displayed_message_response_time,omitempty"`
	DoesViewerHavePagePermissionLinkIg                   *bool                              `json:"does_viewer_have_page_permission_link_ig,omitempty"`
	Emails                                               *[]string                          `json:"emails,omitempty"`
	Engagement                                           *Engagement                        `json:"engagement,omitempty"`
	FanCount                                             *uint64                            `json:"fan_count,omitempty"`
	FeaturedVideo                                        *AdVideo                           `json:"featured_video,omitempty"`
	Features                                             *string                            `json:"features,omitempty"`
	FollowersCount                                       *uint64                            `json:"followers_count,omitempty"`
	FoodStyles                                           *[]string                          `json:"food_styles,omitempty"`
	Founded                                              *string                            `json:"founded,omitempty"`
	GeneralInfo                                          *string                            `json:"general_info,omitempty"`
	GeneralManager                                       *string                            `json:"general_manager,omitempty"`
	Genre                                                *string                            `json:"genre,omitempty"`
	GlobalBrandPageName                                  *string                            `json:"global_brand_page_name,omitempty"`
	GlobalBrandRootID                                    *core.ID                           `json:"global_brand_root_id,omitempty"`
	HasAddedApp                                          *bool                              `json:"has_added_app,omitempty"`
	HasLeadAccess                                        *HasLeadAccess                     `json:"has_lead_access,omitempty"`
	HasTransitionedToNewPageExperience                   *bool                              `json:"has_transitioned_to_new_page_experience,omitempty"`
	HasWhatsappBusinessNumber                            *bool                              `json:"has_whatsapp_business_number,omitempty"`
	HasWhatsappNumber                                    *bool                              `json:"has_whatsapp_number,omitempty"`
	Hometown                                             *string                            `json:"hometown,omitempty"`
	Hours                                                *map[string]string                 `json:"hours,omitempty"`
	ID                                                   *core.ID                           `json:"id,omitempty"`
	Impressum                                            *string                            `json:"impressum,omitempty"`
	Influences                                           *string                            `json:"influences,omitempty"`
	InstagramBusinessAccount                             *IGUser                            `json:"instagram_business_account,omitempty"`
	IsAlwaysOpen                                         *bool                              `json:"is_always_open,omitempty"`
	IsCallingEligible                                    *bool                              `json:"is_calling_eligible,omitempty"`
	IsChain                                              *bool                              `json:"is_chain,omitempty"`
	IsCommunityPage                                      *bool                              `json:"is_community_page,omitempty"`
	IsEligibleForBrandedContent                          *bool                              `json:"is_eligible_for_branded_content,omitempty"`
	IsEligibleForDisableConnectIgBtnForNonPageAdminAmWeb *bool                              `json:"is_eligible_for_disable_connect_ig_btn_for_non_page_admin_am_web,omitempty"`
	IsMessengerBotGetStartedEnabled                      *bool                              `json:"is_messenger_bot_get_started_enabled,omitempty"`
	IsMessengerPlatformBot                               *bool                              `json:"is_messenger_platform_bot,omitempty"`
	IsOwned                                              *bool                              `json:"is_owned,omitempty"`
	IsPermanentlyClosed                                  *bool                              `json:"is_permanently_closed,omitempty"`
	IsPublished                                          *bool                              `json:"is_published,omitempty"`
	IsUnclaimed                                          *bool                              `json:"is_unclaimed,omitempty"`
	IsVerified                                           *bool                              `json:"is_verified,omitempty"`
	IsWebhooksSubscribed                                 *bool                              `json:"is_webhooks_subscribed,omitempty"`
	Keywords                                             *map[string]interface{}            `json:"keywords,omitempty"`
	LeadgenTosAcceptanceTime                             *core.Time                         `json:"leadgen_tos_acceptance_time,omitempty"`
	LeadgenTosAccepted                                   *bool                              `json:"leadgen_tos_accepted,omitempty"`
	LeadgenTosAcceptingUser                              *User                              `json:"leadgen_tos_accepting_user,omitempty"`
	Link                                                 *string                            `json:"link,omitempty"`
	Location                                             *Location                          `json:"location,omitempty"`
	Members                                              *string                            `json:"members,omitempty"`
	MerchantID                                           *core.ID                           `json:"merchant_id,omitempty"`
	MerchantReviewStatus                                 *string                            `json:"merchant_review_status,omitempty"`
	MessagingFeatureStatus                               *MessagingFeatureStatus            `json:"messaging_feature_status,omitempty"`
	MessengerAdsDefaultIcebreakers                       *[]string                          `json:"messenger_ads_default_icebreakers,omitempty"`
	MessengerAdsDefaultQuickReplies                      *[]string                          `json:"messenger_ads_default_quick_replies,omitempty"`
	MessengerAdsQuickRepliesType                         *string                            `json:"messenger_ads_quick_replies_type,omitempty"`
	MiniShopStorefront                                   *Shop                              `json:"mini_shop_storefront,omitempty"`
	Mission                                              *string                            `json:"mission,omitempty"`
	Mpg                                                  *string                            `json:"mpg,omitempty"`
	Name                                                 *string                            `json:"name,omitempty"`
	NameWithLocationDescriptor                           *string                            `json:"name_with_location_descriptor,omitempty"`
	Network                                              *string                            `json:"network,omitempty"`
	NewLikeCount                                         *uint64                            `json:"new_like_count,omitempty"`
	OfferEligible                                        *bool                              `json:"offer_eligible,omitempty"`
	OverallStarRating                                    *float64                           `json:"overall_star_rating,omitempty"`
	OwnerBusiness                                        *Business                          `json:"owner_business,omitempty"`
	PageBackedThreadsAccountID                           *core.ID                           `json:"page_backed_threads_account_id,omitempty"`
	PageToken                                            *string                            `json:"page_token,omitempty"`
	ParentPage                                           *Page                              `json:"parent_page,omitempty"`
	Parking                                              *PageParking                       `json:"parking,omitempty"`
	PaymentOptions                                       *PagePaymentOptions                `json:"payment_options,omitempty"`
	PersonalInfo                                         *string                            `json:"personal_info,omitempty"`
	PersonalInterests                                    *string                            `json:"personal_interests,omitempty"`
	PharmaSafetyInfo                                     *string                            `json:"pharma_safety_info,omitempty"`
	Phone                                                *string                            `json:"phone,omitempty"`
	PickupOptions                                        *[]string                          `json:"pickup_options,omitempty"`
	PlaceType                                            *string                            `json:"place_type,omitempty"`
	PlotOutline                                          *string                            `json:"plot_outline,omitempty"`
	PreferredAudience                                    *Targeting                         `json:"preferred_audience,omitempty"`
	PressContact                                         *string                            `json:"press_contact,omitempty"`
	PriceRange                                           *string                            `json:"price_range,omitempty"`
	PriorityHours                                        *[]map[string]string               `json:"priority_hours,omitempty"`
	PrivacyInfoURL                                       *string                            `json:"privacy_info_url,omitempty"`
	ProducedBy                                           *string                            `json:"produced_by,omitempty"`
	Products                                             *string                            `json:"products,omitempty"`
	PromotionEligible                                    *bool                              `json:"promotion_eligible,omitempty"`
	PromotionIneligibleReason                            *string                            `json:"promotion_ineligible_reason,omitempty"`
	PublicTransit                                        *string                            `json:"public_transit,omitempty"`
	RatingCount                                          *uint64                            `json:"rating_count,omitempty"`
	Recipient                                            *string                            `json:"recipient,omitempty"`
	RecordLabel                                          *string                            `json:"record_label,omitempty"`
	ReleaseDate                                          *string                            `json:"release_date,omitempty"`
	RestaurantServices                                   *PageRestaurantServices            `json:"restaurant_services,omitempty"`
	RestaurantSpecialties                                *PageRestaurantSpecialties         `json:"restaurant_specialties,omitempty"`
	Schedule                                             *string                            `json:"schedule,omitempty"`
	ScreenplayBy                                         *string                            `json:"screenplay_by,omitempty"`
	Season                                               *string                            `json:"season,omitempty"`
	SingleLineAddress                                    *string                            `json:"single_line_address,omitempty"`
	Starring                                             *string                            `json:"starring,omitempty"`
	StartInfo                                            *PageStartInfo                     `json:"start_info,omitempty"`
	StoreCode                                            *string                            `json:"store_code,omitempty"`
	StoreLocationDescriptor                              *string                            `json:"store_location_descriptor,omitempty"`
	StoreNumber                                          *uint64                            `json:"store_number,omitempty"`
	Studio                                               *string                            `json:"studio,omitempty"`
	SupportsDonateButtonInLiveVideo                      *bool                              `json:"supports_donate_button_in_live_video,omitempty"`
	TalkingAboutCount                                    *uint64                            `json:"talking_about_count,omitempty"`
	TemporaryStatus                                      *string                            `json:"temporary_status,omitempty"`
	UnreadMessageCount                                   *uint64                            `json:"unread_message_count,omitempty"`
	UnreadNotifCount                                     *uint64                            `json:"unread_notif_count,omitempty"`
	UnseenMessageCount                                   *uint64                            `json:"unseen_message_count,omitempty"`
	UserAccessExpireTime                                 *core.Time                         `json:"user_access_expire_time,omitempty"`
	Username                                             *string                            `json:"username,omitempty"`
	VerificationStatus                                   *string                            `json:"verification_status,omitempty"`
	VoipInfo                                             *VoipInfo                          `json:"voip_info,omitempty"`
	Website                                              *string                            `json:"website,omitempty"`
	WereHereCount                                        *uint64                            `json:"were_here_count,omitempty"`
	WhatsappNumber                                       *string                            `json:"whatsapp_number,omitempty"`
	WrittenBy                                            *string                            `json:"written_by,omitempty"`
}

var PageFields = struct {
	About                                                string
	AccessToken                                          string
	AdCampaign                                           string
	Affiliation                                          string
	AppID                                                string
	ArtistsWeLike                                        string
	Attire                                               string
	AvailablePromoOfferIds                               string
	Awards                                               string
	BandInterests                                        string
	BandMembers                                          string
	BestPage                                             string
	Bio                                                  string
	Birthday                                             string
	BookingAgent                                         string
	BreakingNewsUsage                                    string
	Built                                                string
	Business                                             string
	CanCheckin                                           string
	CanPost                                              string
	Category                                             string
	CategoryList                                         string
	Checkins                                             string
	CompanyOverview                                      string
	ConnectedInstagramAccount                            string
	ConnectedPageBackedInstagramAccount                  string
	ContactAddress                                       string
	CopyrightAttributionInsights                         string
	CopyrightWhitelistedIgPartners                       string
	CountryPageLikes                                     string
	Cover                                                string
	CulinaryTeam                                         string
	CurrentLocation                                      string
	DeliveryAndPickupOptionInfo                          string
	Description                                          string
	DescriptionHTML                                      string
	DifferentlyOpenOfferings                             string
	DirectedBy                                           string
	DisplaySubtext                                       string
	DisplayedMessageResponseTime                         string
	DoesViewerHavePagePermissionLinkIg                   string
	Emails                                               string
	Engagement                                           string
	FanCount                                             string
	FeaturedVideo                                        string
	Features                                             string
	FollowersCount                                       string
	FoodStyles                                           string
	Founded                                              string
	GeneralInfo                                          string
	GeneralManager                                       string
	Genre                                                string
	GlobalBrandPageName                                  string
	GlobalBrandRootID                                    string
	HasAddedApp                                          string
	HasLeadAccess                                        string
	HasTransitionedToNewPageExperience                   string
	HasWhatsappBusinessNumber                            string
	HasWhatsappNumber                                    string
	Hometown                                             string
	Hours                                                string
	ID                                                   string
	Impressum                                            string
	Influences                                           string
	InstagramBusinessAccount                             string
	IsAlwaysOpen                                         string
	IsCallingEligible                                    string
	IsChain                                              string
	IsCommunityPage                                      string
	IsEligibleForBrandedContent                          string
	IsEligibleForDisableConnectIgBtnForNonPageAdminAmWeb string
	IsMessengerBotGetStartedEnabled                      string
	IsMessengerPlatformBot                               string
	IsOwned                                              string
	IsPermanentlyClosed                                  string
	IsPublished                                          string
	IsUnclaimed                                          string
	IsVerified                                           string
	IsWebhooksSubscribed                                 string
	Keywords                                             string
	LeadgenTosAcceptanceTime                             string
	LeadgenTosAccepted                                   string
	LeadgenTosAcceptingUser                              string
	Link                                                 string
	Location                                             string
	Members                                              string
	MerchantID                                           string
	MerchantReviewStatus                                 string
	MessagingFeatureStatus                               string
	MessengerAdsDefaultIcebreakers                       string
	MessengerAdsDefaultQuickReplies                      string
	MessengerAdsQuickRepliesType                         string
	MiniShopStorefront                                   string
	Mission                                              string
	Mpg                                                  string
	Name                                                 string
	NameWithLocationDescriptor                           string
	Network                                              string
	NewLikeCount                                         string
	OfferEligible                                        string
	OverallStarRating                                    string
	OwnerBusiness                                        string
	PageBackedThreadsAccountID                           string
	PageToken                                            string
	ParentPage                                           string
	Parking                                              string
	PaymentOptions                                       string
	PersonalInfo                                         string
	PersonalInterests                                    string
	PharmaSafetyInfo                                     string
	Phone                                                string
	PickupOptions                                        string
	PlaceType                                            string
	PlotOutline                                          string
	PreferredAudience                                    string
	PressContact                                         string
	PriceRange                                           string
	PriorityHours                                        string
	PrivacyInfoURL                                       string
	ProducedBy                                           string
	Products                                             string
	PromotionEligible                                    string
	PromotionIneligibleReason                            string
	PublicTransit                                        string
	RatingCount                                          string
	Recipient                                            string
	RecordLabel                                          string
	ReleaseDate                                          string
	RestaurantServices                                   string
	RestaurantSpecialties                                string
	Schedule                                             string
	ScreenplayBy                                         string
	Season                                               string
	SingleLineAddress                                    string
	Starring                                             string
	StartInfo                                            string
	StoreCode                                            string
	StoreLocationDescriptor                              string
	StoreNumber                                          string
	Studio                                               string
	SupportsDonateButtonInLiveVideo                      string
	TalkingAboutCount                                    string
	TemporaryStatus                                      string
	UnreadMessageCount                                   string
	UnreadNotifCount                                     string
	UnseenMessageCount                                   string
	UserAccessExpireTime                                 string
	Username                                             string
	VerificationStatus                                   string
	VoipInfo                                             string
	Website                                              string
	WereHereCount                                        string
	WhatsappNumber                                       string
	WrittenBy                                            string
}{
	About:                               "about",
	AccessToken:                         "access_token",
	AdCampaign:                          "ad_campaign",
	Affiliation:                         "affiliation",
	AppID:                               "app_id",
	ArtistsWeLike:                       "artists_we_like",
	Attire:                              "attire",
	AvailablePromoOfferIds:              "available_promo_offer_ids",
	Awards:                              "awards",
	BandInterests:                       "band_interests",
	BandMembers:                         "band_members",
	BestPage:                            "best_page",
	Bio:                                 "bio",
	Birthday:                            "birthday",
	BookingAgent:                        "booking_agent",
	BreakingNewsUsage:                   "breaking_news_usage",
	Built:                               "built",
	Business:                            "business",
	CanCheckin:                          "can_checkin",
	CanPost:                             "can_post",
	Category:                            "category",
	CategoryList:                        "category_list",
	Checkins:                            "checkins",
	CompanyOverview:                     "company_overview",
	ConnectedInstagramAccount:           "connected_instagram_account",
	ConnectedPageBackedInstagramAccount: "connected_page_backed_instagram_account",
	ContactAddress:                      "contact_address",
	CopyrightAttributionInsights:        "copyright_attribution_insights",
	CopyrightWhitelistedIgPartners:      "copyright_whitelisted_ig_partners",
	CountryPageLikes:                    "country_page_likes",
	Cover:                               "cover",
	CulinaryTeam:                        "culinary_team",
	CurrentLocation:                     "current_location",
	DeliveryAndPickupOptionInfo:         "delivery_and_pickup_option_info",
	Description:                         "description",
	DescriptionHTML:                     "description_html",
	DifferentlyOpenOfferings:            "differently_open_offerings",
	DirectedBy:                          "directed_by",
	DisplaySubtext:                      "display_subtext",
	DisplayedMessageResponseTime:        "displayed_message_response_time",
	DoesViewerHavePagePermissionLinkIg:  "does_viewer_have_page_permission_link_ig",
	Emails:                              "emails",
	Engagement:                          "engagement",
	FanCount:                            "fan_count",
	FeaturedVideo:                       "featured_video",
	Features:                            "features",
	FollowersCount:                      "followers_count",
	FoodStyles:                          "food_styles",
	Founded:                             "founded",
	GeneralInfo:                         "general_info",
	GeneralManager:                      "general_manager",
	Genre:                               "genre",
	GlobalBrandPageName:                 "global_brand_page_name",
	GlobalBrandRootID:                   "global_brand_root_id",
	HasAddedApp:                         "has_added_app",
	HasLeadAccess:                       "has_lead_access",
	HasTransitionedToNewPageExperience:  "has_transitioned_to_new_page_experience",
	HasWhatsappBusinessNumber:           "has_whatsapp_business_number",
	HasWhatsappNumber:                   "has_whatsapp_number",
	Hometown:                            "hometown",
	Hours:                               "hours",
	ID:                                  "id",
	Impressum:                           "impressum",
	Influences:                          "influences",
	InstagramBusinessAccount:            "instagram_business_account",
	IsAlwaysOpen:                        "is_always_open",
	IsCallingEligible:                   "is_calling_eligible",
	IsChain:                             "is_chain",
	IsCommunityPage:                     "is_community_page",
	IsEligibleForBrandedContent:         "is_eligible_for_branded_content",
	IsEligibleForDisableConnectIgBtnForNonPageAdminAmWeb: "is_eligible_for_disable_connect_ig_btn_for_non_page_admin_am_web",
	IsMessengerBotGetStartedEnabled:                      "is_messenger_bot_get_started_enabled",
	IsMessengerPlatformBot:                               "is_messenger_platform_bot",
	IsOwned:                                              "is_owned",
	IsPermanentlyClosed:                                  "is_permanently_closed",
	IsPublished:                                          "is_published",
	IsUnclaimed:                                          "is_unclaimed",
	IsVerified:                                           "is_verified",
	IsWebhooksSubscribed:                                 "is_webhooks_subscribed",
	Keywords:                                             "keywords",
	LeadgenTosAcceptanceTime:                             "leadgen_tos_acceptance_time",
	LeadgenTosAccepted:                                   "leadgen_tos_accepted",
	LeadgenTosAcceptingUser:                              "leadgen_tos_accepting_user",
	Link:                                                 "link",
	Location:                                             "location",
	Members:                                              "members",
	MerchantID:                                           "merchant_id",
	MerchantReviewStatus:                                 "merchant_review_status",
	MessagingFeatureStatus:                               "messaging_feature_status",
	MessengerAdsDefaultIcebreakers:                       "messenger_ads_default_icebreakers",
	MessengerAdsDefaultQuickReplies:                      "messenger_ads_default_quick_replies",
	MessengerAdsQuickRepliesType:                         "messenger_ads_quick_replies_type",
	MiniShopStorefront:                                   "mini_shop_storefront",
	Mission:                                              "mission",
	Mpg:                                                  "mpg",
	Name:                                                 "name",
	NameWithLocationDescriptor:                           "name_with_location_descriptor",
	Network:                                              "network",
	NewLikeCount:                                         "new_like_count",
	OfferEligible:                                        "offer_eligible",
	OverallStarRating:                                    "overall_star_rating",
	OwnerBusiness:                                        "owner_business",
	PageBackedThreadsAccountID:                           "page_backed_threads_account_id",
	PageToken:                                            "page_token",
	ParentPage:                                           "parent_page",
	Parking:                                              "parking",
	PaymentOptions:                                       "payment_options",
	PersonalInfo:                                         "personal_info",
	PersonalInterests:                                    "personal_interests",
	PharmaSafetyInfo:                                     "pharma_safety_info",
	Phone:                                                "phone",
	PickupOptions:                                        "pickup_options",
	PlaceType:                                            "place_type",
	PlotOutline:                                          "plot_outline",
	PreferredAudience:                                    "preferred_audience",
	PressContact:                                         "press_contact",
	PriceRange:                                           "price_range",
	PriorityHours:                                        "priority_hours",
	PrivacyInfoURL:                                       "privacy_info_url",
	ProducedBy:                                           "produced_by",
	Products:                                             "products",
	PromotionEligible:                                    "promotion_eligible",
	PromotionIneligibleReason:                            "promotion_ineligible_reason",
	PublicTransit:                                        "public_transit",
	RatingCount:                                          "rating_count",
	Recipient:                                            "recipient",
	RecordLabel:                                          "record_label",
	ReleaseDate:                                          "release_date",
	RestaurantServices:                                   "restaurant_services",
	RestaurantSpecialties:                                "restaurant_specialties",
	Schedule:                                             "schedule",
	ScreenplayBy:                                         "screenplay_by",
	Season:                                               "season",
	SingleLineAddress:                                    "single_line_address",
	Starring:                                             "starring",
	StartInfo:                                            "start_info",
	StoreCode:                                            "store_code",
	StoreLocationDescriptor:                              "store_location_descriptor",
	StoreNumber:                                          "store_number",
	Studio:                                               "studio",
	SupportsDonateButtonInLiveVideo:                      "supports_donate_button_in_live_video",
	TalkingAboutCount:                                    "talking_about_count",
	TemporaryStatus:                                      "temporary_status",
	UnreadMessageCount:                                   "unread_message_count",
	UnreadNotifCount:                                     "unread_notif_count",
	UnseenMessageCount:                                   "unseen_message_count",
	UserAccessExpireTime:                                 "user_access_expire_time",
	Username:                                             "username",
	VerificationStatus:                                   "verification_status",
	VoipInfo:                                             "voip_info",
	Website:                                              "website",
	WereHereCount:                                        "were_here_count",
	WhatsappNumber:                                       "whatsapp_number",
	WrittenBy:                                            "written_by",
}
