package vo

import "time"

type (
	/*
		Request
	*/

	// LeaveRequestDTO ...
	LeaveRequestDTO struct {
		EmployeeID int64     `json:"employee_id" validate:"required"`
		Type       string    `json:"type" validate:"required"`
		StartDate  time.Time `json:"start_date" validate:"required"`
		EndDate    time.Time `json:"end_date" validate:"required"`
		Total      int64     `json:"total" validate:"required"`
		Notes      string    `json:"notes" validate:"required"`
	}

	// ActionRequestDTO ...
	ActionRequestDTO struct {
		LeaveID     int64  `json:"leave_id" validate:"required"`
		ProcessedBy int64  `json:"processed_by" validate:"required"`
		Action      string `json:"action" validate:"required"`
	}

	/*
		Response
	*/

	// ActionResponseDTO ...
	ActionResponseDTO struct {
		LeaveID     int64  `json:"leave_id" validate:"required"`
		ProcessedBy int64  `json:"processed_by" validate:"required"`
		Action      string `json:"action" validate:"required"`
	}

	// QuotaResponseDTO ...
	QuotaResponseDTO struct {
		QuotaLeft int64 `json:"quota_left" validate:"required"`
	}
)
