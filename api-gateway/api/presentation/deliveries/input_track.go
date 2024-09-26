package deliveries

type InputTrackRequest struct {
	OrderID int32  `json:"order_id"`
	TrackID string `json:"track_id"`
}
