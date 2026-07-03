package objects

type LeadNurturingState struct {
	AiAgentMode          *string                 `json:"ai_agent_mode,omitempty"`
	ConversationSummary  *string                 `json:"conversation_summary,omitempty"`
	HandoffReason        *string                 `json:"handoff_reason,omitempty"`
	LeadInterestLevel    *string                 `json:"lead_interest_level,omitempty"`
	NeededManualActions  *[]string               `json:"needed_manual_actions,omitempty"`
	QualificationDetails *string                 `json:"qualification_details,omitempty"`
	QualificationStatus  *string                 `json:"qualification_status,omitempty"`
	ScheduledTime        *map[string]interface{} `json:"scheduled_time,omitempty"`
	UpdatedEmail         *string                 `json:"updated_email,omitempty"`
	UpdatedPhoneNumber   *string                 `json:"updated_phone_number,omitempty"`
}
