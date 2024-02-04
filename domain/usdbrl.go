package domain

type Usdbrl struct {
	Code       string `json:"-"`
	Codein     string `json:"-"`
	Name       string `json:"-"`
	High       string `json:"-"`
	Low        string `json:"-"`
	VarBid     string `json:"-"`
	PctChange  string `json:"-"`
	Bid        string `json:"bid,omitempty"`
	Ask        string `json:"-"`
	Timestamp  string `json:"-"`
	CreateDate string `json:"-"`
}
