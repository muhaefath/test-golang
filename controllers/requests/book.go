package requests

type PickUpRequest struct {
	CoverID  int    `json:"cover_id"`
	PickUpAt string `json:"pick_up_at"`
}

type ReturnRequest struct {
	CoverID  int    `json:"cover_id"`
	ReturnAt string `json:"return_at"`
}
