package backend

import "fmt"

type Interactions struct {
	mongoDBAdapter *MongoDBAdapter
}

func NewInteractions(mongoDBAdapter *MongoDBAdapter) *Interactions {
	return &Interactions{mongoDBAdapter: mongoDBAdapter}
}

func (i *Interactions) GetBox(boxKey string) *BoxDTO {
	box := i.mongoDBAdapter.loadBox(boxKey)
	return i.mapToBoxDTO(box, boxKey)
}

func (i *Interactions) mapToBoxDTO(box *Box, boxKey string) *BoxDTO {
	requestingMember := selectMember(boxKey, box.Members)
	boxDTO := BoxDTO{
		Title:          box.Title,
		MemberNickname: requestingMember.Nickname,
		CreationDate:   box.CreationDate,
		Items:          []ItemDTO{},
	}
	for _, item := range box.Items {
		boxDTO.Items = append(boxDTO.Items, ItemDTO{
			AuthorNickname: selectMember(item.AuthorKey, box.Members).Nickname,
			CreationDate:   item.CreationDate,
			Subject:        item.Subject,
			Message:        item.Message,
		})
	}

	return &boxDTO
}

func selectMember(key string, members []Member) *Member {
	for i, _ := range members {
		if members[i].Key == key {
			return &members[i]
		}
	}
	panic(SuprisingException{Err: fmt.Errorf("No member found for key:%s", key)})
}
