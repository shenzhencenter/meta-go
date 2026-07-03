package objects

type LeadGenAppointmentTimeSlot struct {
	EndTime   *uint64 `json:"end_time,omitempty"`
	StartTime *uint64 `json:"start_time,omitempty"`
}
