package backend

type Box struct {
	BoxID        string   `bson:"boxId"`
	Title        string   `bson:"title"`
	CreationDate string   `bson:"creationDate"`
	Members      []Member `bson:"members"`
	Items        []Item   `bson:"items"`
}

type Item struct {
	CreationDate string `bson:"creationDate"`
	Message      string `bson:"message"`
	Author       Member `bson:"author"`
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
	Message        string `json:"message"`
}

type BoxMember struct {
	BoxKey string `bson:"boxKey"`
	BoxID  string `bson:"boxId"`
	Member Member `bson:"member"`
}

type Member struct {
	Email    string `bson:"email"`
	Nickname string `bson:"nickname"`
	Owner    bool   `bson:"owner"`
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

type Interactions struct {
	mongoDBAdapter *MongoDBAdapter
}

func NewInteractions(mongoDBAdapter *MongoDBAdapter) *Interactions {
	return &Interactions{mongoDBAdapter: mongoDBAdapter}
}

func (i *Interactions) GetBox(boxKey string) *BoxDTO {
	boxMember := i.mongoDBAdapter.openBox(boxKey)
	box := i.mongoDBAdapter.loadBox(boxMember.BoxID)
	return i.mapToBoxDTO(box, boxMember)
}

func (i *Interactions) mapToBoxDTO(box *Box, boxMember *BoxMember) *BoxDTO {
	boxDTO := BoxDTO{
		Title:          box.Title,
		MemberNickname: boxMember.Member.Nickname,
		CreationDate:   box.CreationDate,
		Items:          []ItemDTO{},
	}
	for _, item := range box.Items {
		boxDTO.Items = append(boxDTO.Items, ItemDTO{
			AuthorNickname: item.Author.Nickname,
			CreationDate:   item.CreationDate,
			Message:        item.Message,
		})
	}

	return &boxDTO
}
