package objects

type TimezoneOffset struct {
	Abbr   *string `json:"abbr,omitempty"`
	Isdst  *bool   `json:"isdst,omitempty"`
	Offset *int    `json:"offset,omitempty"`
	Time   *string `json:"time,omitempty"`
	Ts     *uint64 `json:"ts,omitempty"`
}
