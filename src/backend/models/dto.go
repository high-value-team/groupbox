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



