package vo

type (
	/*
		Request
	*/

	// AuthenticationRequestDTO ...
	AuthenticationRequestDTO struct {
		UserName string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	/*
		Response
	*/

	// AuthenticationResponseDTO ...
	AuthenticationResponseDTO struct {
		UserName string `json:"username,omitempty"`
		Role     string `json:"role,omitempty"`
	}
)
