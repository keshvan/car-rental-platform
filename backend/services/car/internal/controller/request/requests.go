package request

type (
	UpdateCarRequest struct {
		Brand        string `json:"brand"`
		Model        string `json:"model"`
		Name         string `json:"name"`
		Year         int64  `json:"year"`
		PricePerHour int64  `json:"price_per_hour"`
		ImageURL     string `json:"image_url,omitempty"`
	}
	CompleteRentRequest struct {
		Review string `json:"review`
	}
)
