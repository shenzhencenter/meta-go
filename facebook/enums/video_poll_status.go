package enums

type VideopollStatus string

const (
	VideopollStatusClosed      VideopollStatus = "closed"
	VideopollStatusResultsOpen VideopollStatus = "results_open"
	VideopollStatusVotingOpen  VideopollStatus = "voting_open"
)

func (value VideopollStatus) String() string {
	return string(value)
}
