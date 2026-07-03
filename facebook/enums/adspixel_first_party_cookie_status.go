package enums

type AdspixelFirstPartyCookieStatus string

const (
	AdspixelFirstPartyCookieStatusEmpty                    AdspixelFirstPartyCookieStatus = "EMPTY"
	AdspixelFirstPartyCookieStatusFirstPartyCookieDisabled AdspixelFirstPartyCookieStatus = "FIRST_PARTY_COOKIE_DISABLED"
	AdspixelFirstPartyCookieStatusFirstPartyCookieEnabled  AdspixelFirstPartyCookieStatus = "FIRST_PARTY_COOKIE_ENABLED"
)

func (value AdspixelFirstPartyCookieStatus) String() string {
	return string(value)
}
