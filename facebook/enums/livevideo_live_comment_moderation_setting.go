package enums

type LivevideoLiveCommentModerationSetting string

const (
	LivevideoLiveCommentModerationSettingDefault       LivevideoLiveCommentModerationSetting = "DEFAULT"
	LivevideoLiveCommentModerationSettingDiscussion    LivevideoLiveCommentModerationSetting = "DISCUSSION"
	LivevideoLiveCommentModerationSettingFollowed      LivevideoLiveCommentModerationSetting = "FOLLOWED"
	LivevideoLiveCommentModerationSettingFollower      LivevideoLiveCommentModerationSetting = "FOLLOWER"
	LivevideoLiveCommentModerationSettingNoHyperlink   LivevideoLiveCommentModerationSetting = "NO_HYPERLINK"
	LivevideoLiveCommentModerationSettingProtectedMode LivevideoLiveCommentModerationSetting = "PROTECTED_MODE"
	LivevideoLiveCommentModerationSettingRestricted    LivevideoLiveCommentModerationSetting = "RESTRICTED"
	LivevideoLiveCommentModerationSettingSlow          LivevideoLiveCommentModerationSetting = "SLOW"
	LivevideoLiveCommentModerationSettingSupporter     LivevideoLiveCommentModerationSetting = "SUPPORTER"
	LivevideoLiveCommentModerationSettingTagged        LivevideoLiveCommentModerationSetting = "TAGGED"
)

func (value LivevideoLiveCommentModerationSetting) String() string {
	return string(value)
}
