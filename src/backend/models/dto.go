package models

type BoxDTO struct {
	Title          string    `json:"title"`
	MemberNickname string    `json:"memberNickname"`
	CreationDate   string    `json:"creationDate"`
	Items          []ItemDTO `json:"items"`
}

type ItemDTO struct {
	AuthorNickname string `json:"authorNickname"`
	CreationDate   string `json:"creationDate"`
	Subject        string `json:"subject"`
	Message        string `json:"message"`
}

type AddItemRequestDTO struct {
	Message string `json:"message"`
}

type CreateBoxResponseDTO struct {
	BoxKey string `json:"boxKey"`
}
type CreateBoxRequestDTO struct {
	Title   string   `json:"title"`
	Owner   string   `json:"owner"`
	Members []string `json:"members"`
}
