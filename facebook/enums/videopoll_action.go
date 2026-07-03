package enums

type VideopollAction string

const (
	VideopollActionAttachToVideo VideopollAction = "ATTACH_TO_VIDEO"
	VideopollActionClose         VideopollAction = "CLOSE"
	VideopollActionDeletePoll    VideopollAction = "DELETE_POLL"
	VideopollActionShowResults   VideopollAction = "SHOW_RESULTS"
	VideopollActionShowVoting    VideopollAction = "SHOW_VOTING"
)

func (value VideopollAction) String() string {
	return string(value)
}
