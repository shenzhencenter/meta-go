package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AudioIsrc struct {
	AllKgFeaturedArtists             *string                 `json:"all_kg_featured_artists,omitempty"`
	AllKgMainArtists                 *string                 `json:"all_kg_main_artists,omitempty"`
	ArtistProfilePictureURL          *string                 `json:"artist_profile_picture_url,omitempty"`
	ID                               *core.ID                `json:"id,omitempty"`
	Isrc                             *string                 `json:"isrc,omitempty"`
	PublishingRightsData             *map[string]interface{} `json:"publishing_rights_data,omitempty"`
	TopSearchableArtistID            *core.ID                `json:"top_searchable_artist_id,omitempty"`
	TopSearchableArtistName          *string                 `json:"top_searchable_artist_name,omitempty"`
	TopSearchableArtistProfilePicURL *string                 `json:"top_searchable_artist_profile_pic_url,omitempty"`
}
