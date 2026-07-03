package enums

type EventCategory string

const (
	EventCategoryClassicLiterature        EventCategory = "CLASSIC_LITERATURE"
	EventCategoryComedy                   EventCategory = "COMEDY"
	EventCategoryCrafts                   EventCategory = "CRAFTS"
	EventCategoryDance                    EventCategory = "DANCE"
	EventCategoryDrinks                   EventCategory = "DRINKS"
	EventCategoryFitnessAndWorkouts       EventCategory = "FITNESS_AND_WORKOUTS"
	EventCategoryFoods                    EventCategory = "FOODS"
	EventCategoryGames                    EventCategory = "GAMES"
	EventCategoryGardening                EventCategory = "GARDENING"
	EventCategoryHealthyLivingAndSelfCare EventCategory = "HEALTHY_LIVING_AND_SELF_CARE"
	EventCategoryHealthAndMedical         EventCategory = "HEALTH_AND_MEDICAL"
	EventCategoryHomeAndGarden            EventCategory = "HOME_AND_GARDEN"
	EventCategoryMusicAndAudio            EventCategory = "MUSIC_AND_AUDIO"
	EventCategoryParties                  EventCategory = "PARTIES"
	EventCategoryProfessionalNetworking   EventCategory = "PROFESSIONAL_NETWORKING"
	EventCategoryReligions                EventCategory = "RELIGIONS"
	EventCategoryShoppingEvent            EventCategory = "SHOPPING_EVENT"
	EventCategorySocialIssues             EventCategory = "SOCIAL_ISSUES"
	EventCategorySports                   EventCategory = "SPORTS"
	EventCategoryTheater                  EventCategory = "THEATER"
	EventCategoryTvAndMovies              EventCategory = "TV_AND_MOVIES"
	EventCategoryVisualArts               EventCategory = "VISUAL_ARTS"
)

func (value EventCategory) String() string {
	return string(value)
}
