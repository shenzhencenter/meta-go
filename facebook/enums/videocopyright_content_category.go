package enums

type VideocopyrightContentCategory string

const (
	VideocopyrightContentCategoryEpisode VideocopyrightContentCategory = "episode"
	VideocopyrightContentCategoryMovie   VideocopyrightContentCategory = "movie"
	VideocopyrightContentCategoryWeb     VideocopyrightContentCategory = "web"
)

func (value VideocopyrightContentCategory) String() string {
	return string(value)
}
