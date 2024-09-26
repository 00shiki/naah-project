package deliveries

type CallbackRequest struct {
	TrackID string `json:"track_id"`
	Status  string `json:"status"`
}
