package backend

type Box struct {
	Title        string   `bson:"title"`
	CreationDate string   `bson:"creationDate"`
	Members      []Member `bson:"members"`
	Items        []Item   `bson:"items"`
}

type Item struct {
	CreationDate string `bson:"creationDate"`
	Subject      string `bson:"subject"`
	Message      string `bson:"message"`
	AuthorKey    string `bson:"authorKey"`
}

type Member struct {
	Key      string `bson:"key"`
	Email    string `bson:"email"`
	Nickname string `bson:"nickname"`
	Owner    bool   `bson:"owner"`
}

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

type SadException struct {
	Err error
}

func (e *SadException) Message() string {
	return e.Err.Error()
}

type SuprisingException struct {
	Err error
}

func (e *SuprisingException) Message() string {
	return e.Err.Error()
}
