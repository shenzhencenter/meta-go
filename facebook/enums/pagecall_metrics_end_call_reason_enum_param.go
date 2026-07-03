package enums

type PagecallMetricsEndCallReasonEnumParam string

const (
	PagecallMetricsEndCallReasonEnumParamCallerNotVisible              PagecallMetricsEndCallReasonEnumParam = "CALLER_NOT_VISIBLE"
	PagecallMetricsEndCallReasonEnumParamCallEndAcceptAfterHangUp      PagecallMetricsEndCallReasonEnumParam = "CALL_END_ACCEPT_AFTER_HANG_UP"
	PagecallMetricsEndCallReasonEnumParamCameraPermissionDenied        PagecallMetricsEndCallReasonEnumParam = "CAMERA_PERMISSION_DENIED"
	PagecallMetricsEndCallReasonEnumParamClientError                   PagecallMetricsEndCallReasonEnumParam = "CLIENT_ERROR"
	PagecallMetricsEndCallReasonEnumParamClientInterrupted             PagecallMetricsEndCallReasonEnumParam = "CLIENT_INTERRUPTED"
	PagecallMetricsEndCallReasonEnumParamConnectionDropped             PagecallMetricsEndCallReasonEnumParam = "CONNECTION_DROPPED"
	PagecallMetricsEndCallReasonEnumParamHangupCall                    PagecallMetricsEndCallReasonEnumParam = "HANGUP_CALL"
	PagecallMetricsEndCallReasonEnumParamIgnoreCall                    PagecallMetricsEndCallReasonEnumParam = "IGNORE_CALL"
	PagecallMetricsEndCallReasonEnumParamInactiveTimeout               PagecallMetricsEndCallReasonEnumParam = "INACTIVE_TIMEOUT"
	PagecallMetricsEndCallReasonEnumParamIncomingTimeout               PagecallMetricsEndCallReasonEnumParam = "INCOMING_TIMEOUT"
	PagecallMetricsEndCallReasonEnumParamInAnotherCall                 PagecallMetricsEndCallReasonEnumParam = "IN_ANOTHER_CALL"
	PagecallMetricsEndCallReasonEnumParamMaxAllowedParticipantsReached PagecallMetricsEndCallReasonEnumParam = "MAX_ALLOWED_PARTICIPANTS_REACHED"
	PagecallMetricsEndCallReasonEnumParamMicrophonePermissionDenied    PagecallMetricsEndCallReasonEnumParam = "MICROPHONE_PERMISSION_DENIED"
	PagecallMetricsEndCallReasonEnumParamNoAnswerTimeout               PagecallMetricsEndCallReasonEnumParam = "NO_ANSWER_TIMEOUT"
	PagecallMetricsEndCallReasonEnumParamNoPermission                  PagecallMetricsEndCallReasonEnumParam = "NO_PERMISSION"
	PagecallMetricsEndCallReasonEnumParamRemovedByParticipant          PagecallMetricsEndCallReasonEnumParam = "REMOVED_BY_PARTICIPANT"
	PagecallMetricsEndCallReasonEnumParamRingMuted                     PagecallMetricsEndCallReasonEnumParam = "RING_MUTED"
	PagecallMetricsEndCallReasonEnumParamSignalingMessageFailed        PagecallMetricsEndCallReasonEnumParam = "SIGNALING_MESSAGE_FAILED"
	PagecallMetricsEndCallReasonEnumParamUnexpectedEndOfCall           PagecallMetricsEndCallReasonEnumParam = "UNEXPECTED_END_OF_CALL"
	PagecallMetricsEndCallReasonEnumParamUnknown                       PagecallMetricsEndCallReasonEnumParam = "UNKNOWN"
	PagecallMetricsEndCallReasonEnumParamVersionUnsupported            PagecallMetricsEndCallReasonEnumParam = "VERSION_UNSUPPORTED"
	PagecallMetricsEndCallReasonEnumParamWebrtcError                   PagecallMetricsEndCallReasonEnumParam = "WEBRTC_ERROR"
)

func (value PagecallMetricsEndCallReasonEnumParam) String() string {
	return string(value)
}
