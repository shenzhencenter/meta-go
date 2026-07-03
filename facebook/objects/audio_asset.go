package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type AudioAsset struct {
	AllDdexFeaturedArtists   *string    `json:"all_ddex_featured_artists,omitempty"`
	AllDdexMainArtists       *string    `json:"all_ddex_main_artists,omitempty"`
	AudioClusterID           *core.ID   `json:"audio_cluster_id,omitempty"`
	CoverImageSource         *string    `json:"cover_image_source,omitempty"`
	Description              *string    `json:"description,omitempty"`
	DisplayArtist            *string    `json:"display_artist,omitempty"`
	DownloadHdURL            *string    `json:"download_hd_url,omitempty"`
	DownloadSdURL            *string    `json:"download_sd_url,omitempty"`
	DurationInMs             *int       `json:"duration_in_ms,omitempty"`
	FreeformGenre            *string    `json:"freeform_genre,omitempty"`
	Grid                     *string    `json:"grid,omitempty"`
	ID                       *core.ID   `json:"id,omitempty"`
	IsTest                   *bool      `json:"is_test,omitempty"`
	OriginalReleaseDate      *time.Time `json:"original_release_date,omitempty"`
	Owner                    *Page      `json:"owner,omitempty"`
	ParentalWarningType      *string    `json:"parental_warning_type,omitempty"`
	Subtitle                 *string    `json:"subtitle,omitempty"`
	Title                    *string    `json:"title,omitempty"`
	TitleWithFeaturedArtists *string    `json:"title_with_featured_artists,omitempty"`
	Upc                      *string    `json:"upc,omitempty"`
}
