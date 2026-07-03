package enums

type PageeventsEventStateFilterEnumParam string

const (
	PageeventsEventStateFilterEnumParamCanceled                     PageeventsEventStateFilterEnumParam = "canceled"
	PageeventsEventStateFilterEnumParamDraft                        PageeventsEventStateFilterEnumParam = "draft"
	PageeventsEventStateFilterEnumParamPublished                    PageeventsEventStateFilterEnumParam = "published"
	PageeventsEventStateFilterEnumParamScheduledDraftForPublication PageeventsEventStateFilterEnumParam = "scheduled_draft_for_publication"
)

func (value PageeventsEventStateFilterEnumParam) String() string {
	return string(value)
}
