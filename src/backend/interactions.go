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

// TODO
type BoxDTO struct {
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

type Interactions struct {
	mongoDBAdapter *MongoDBAdapter
}

func NewInteractions(mongoDBAdapter *MongoDBAdapter) *Interactions {
	return &Interactions{mongoDBAdapter: mongoDBAdapter}
}

func (i *Interactions) GetBox(boxKey string) (*BoxDTO, error) {
	var err error
	boxMember := i.mongoDBAdapter.openBox(&err, boxKey)
	box := i.mongoDBAdapter.loadBox(&err, boxMember.BoxID)
	boxDTO := i.mapToBoxDTO(&err, box, boxMember)
	return boxDTO, err
}

func (i *Interactions) mapToBoxDTO(err *error, box *Box, boxMember *BoxMember) *BoxDTO {
	if *err != nil {
		return nil
	}
	return &BoxDTO{}
}
