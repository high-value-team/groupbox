package backend

import (
	"fmt"
	"time"
)

type Interactions struct {
	mongoDBAdapter     *MongoDBAdapter
	emailNotifications *EmailNotifications
}

func NewInteractions(mongoDBAdapter *MongoDBAdapter, emailNotifications *EmailNotifications) *Interactions {
	return &Interactions{mongoDBAdapter: mongoDBAdapter, emailNotifications: emailNotifications}
}

func (i *Interactions) GetBox(boxKey string) *BoxDTO {
	box := i.mongoDBAdapter.loadBox(boxKey)
	return i.mapToBoxDTO(box, boxKey)
}

func (i *Interactions) CreateBox(title, ownerEmail string, memberEmails []string) *CreateBoxResponseDTO {
	members := i.generateMembers(ownerEmail, memberEmails)
	owner := selectOwner(members)
	box := i.buildBox(title, members)
	i.mongoDBAdapter.saveBox(box)
	i.emailNotifications.SendInvitations(title, box.Members)
	return &CreateBoxResponseDTO{BoxKey: owner.Key}
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

func (i *Interactions) generateMembers(ownerEmail string, memberEmails []string) []Member {
	return nil
}

func (i *Interactions) buildBox(title string, members []Member) *Box {
	return &Box{
		Title:        title,
		CreationDate: time.Now().Format(time.RFC3339),
		Members:      members,
		Items:        []Item{},
	}
}

func selectMember(key string, members []Member) *Member {
	for i, _ := range members {
		if members[i].Key == key {
			return &members[i]
		}
	}
	panic(SuprisingException{Err: fmt.Errorf("No member found for key:%s", key)})
}

func selectOwner(members []Member) *Member {
	for i, _ := range members {
		if members[i].Owner {
			return &members[i]
		}
	}
	panic(SuprisingException{Err: fmt.Errorf("No owner found")})
}
