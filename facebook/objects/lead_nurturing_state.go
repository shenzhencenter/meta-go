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

var LeadNurturingStateFields = struct {
	AiAgentMode          string
	ConversationSummary  string
	HandoffReason        string
	LeadInterestLevel    string
	NeededManualActions  string
	QualificationDetails string
	QualificationStatus  string
	ScheduledTime        string
	UpdatedEmail         string
	UpdatedPhoneNumber   string
}{
	AiAgentMode:          "ai_agent_mode",
	ConversationSummary:  "conversation_summary",
	HandoffReason:        "handoff_reason",
	LeadInterestLevel:    "lead_interest_level",
	NeededManualActions:  "needed_manual_actions",
	QualificationDetails: "qualification_details",
	QualificationStatus:  "qualification_status",
	ScheduledTime:        "scheduled_time",
	UpdatedEmail:         "updated_email",
	UpdatedPhoneNumber:   "updated_phone_number",
}
