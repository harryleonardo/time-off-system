package vo

type (
	/*
		Request
	*/

	// CreateRequestDTO ...
	CreateRequestDTO struct {
		// - profile table
		CompanyID      int64  `json:"company_id" validate:"required"`
		FullName       string `json:"full_name" validate:"required"`
		Email          string `json:"email" validate:"required"`
		IdentityNumber string `json:"identity_number" validate:"required"`
		Address        string `json:"address" validate:"required"`
		DateOfBirth    string `json:"date_of_birth" validate:"required"`
		PhoneNumber    string `json:"phone_number" validate:"required"`
		Gender         string `json:"gender" validate:"required"`
		MaritalStatus  string `json:"marital_status,omitempty"`
		Religion       string `json:"religion,omitempty"`
		// - employee table
		OrganizationName string `json:"organization_name" validate:"required"`
		Position         string `json:"position" validate:"required"`
		Level            string `json:"level" validate:"required"`
		Status           string `json:"status" validate:"required"`
		Branch           string `json:"branch" validate:"required"`
		JoinDate         string `json:"join_date" validate:"required"`
	}

	/*
		Response
	*/
)
