package enums

type URLScopes string

const (
	URLScopesNewsTab       URLScopes = "NEWS_TAB"
	URLScopesNewsTabDevEnv URLScopes = "NEWS_TAB_DEV_ENV"
)

func (value URLScopes) String() string {
	return string(value)
}
