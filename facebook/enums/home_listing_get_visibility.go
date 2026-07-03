package enums

type HomelistinggetVisibility string

const (
	HomelistinggetVisibilityActive                   HomelistinggetVisibility = "ACTIVE"
	HomelistinggetVisibilityArchived                 HomelistinggetVisibility = "ARCHIVED"
	HomelistinggetVisibilityHidden                   HomelistinggetVisibility = "HIDDEN"
	HomelistinggetVisibilityLegacyPublic             HomelistinggetVisibility = "LEGACY_PUBLIC"
	HomelistinggetVisibilityPublished                HomelistinggetVisibility = "PUBLISHED"
	HomelistinggetVisibilityStaging                  HomelistinggetVisibility = "STAGING"
	HomelistinggetVisibilityVisibleOnlyWithOverrides HomelistinggetVisibility = "VISIBLE_ONLY_WITH_OVERRIDES"
	HomelistinggetVisibilityWhitelistOnly            HomelistinggetVisibility = "WHITELIST_ONLY"
)

func (value HomelistinggetVisibility) String() string {
	return string(value)
}
