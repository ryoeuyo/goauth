package controller

type HealthResponse struct {
	Message     string `json:"message"`
	CurrentTime string `json:"current_time"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
