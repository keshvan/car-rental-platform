package request

type (
	RegisterRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
