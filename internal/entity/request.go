package entity

type Request struct {
	Ok      bool   `json:"ok"`
	ErrText string `json:"errText"`
}
