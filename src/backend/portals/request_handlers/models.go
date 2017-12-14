package request_handlers

type JSONBox struct {
	Title          string     `json:"title"`
	MemberNickname string     `json:"memberNickname"`
	CreationDate   string     `json:"creationDate"`
	Items          []JSONItem `json:"items"`
}

type JSONItem struct {
	AuthorNickname string `json:"authorNickname"`
	CreationDate   string `json:"creationDate"`
	Subject        string `json:"subject"`
	Message        string `json:"message"`
}

type JSONRequestAddItem struct {
	Message string `json:"message"`
}

type JSONRequestUpdateItem struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type JSONRequestCreateBox struct {
	Title   string   `json:"title"`
	Owner   string   `json:"owner"`
	Members []string `json:"members"`
}

type JSONResponseCreateBox struct {
	BoxKey string `json:"boxKey"`
}
