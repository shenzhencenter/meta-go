package enums

type ApplicationactivitiesUserIDTypeEnumParam string

const (
	ApplicationactivitiesUserIDTypeEnumParamInstantGamesPlayerID ApplicationactivitiesUserIDTypeEnumParam = "INSTANT_GAMES_PLAYER_ID"
)

func (value ApplicationactivitiesUserIDTypeEnumParam) String() string {
	return string(value)
}
