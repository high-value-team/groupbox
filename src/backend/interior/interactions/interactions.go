package interactions

//TODO: Methoden entzerren, evtl auf verschiedene Klassen/Datei/Packages verteilen, zumindest in bessere Reihenfolge bringen

import (
	"fmt"
	"time"
	"github.com/high-value-team/groupbox/src/backend/providers"
	"github.com/high-value-team/groupbox/src/backend/models"
)

type Interactions struct {
	mongoDBAdapter     *providers.MongoDBAdapter
	emailNotifications *providers.EmailNotifications
}

func NewInteractions(mongoDBAdapter *providers.MongoDBAdapter, emailNotifications *providers.EmailNotifications) *Interactions {
	return &Interactions{mongoDBAdapter: mongoDBAdapter, emailNotifications: emailNotifications}
}

func (i *Interactions) GetBox(boxKey string) *models.BoxDTO {
	box := i.mongoDBAdapter.loadBox(boxKey)
	return i.mapToBoxDTO(box, boxKey)
}

func (i *Interactions) CreateBox(title, ownerEmail string, memberEmails []string) *CreateBoxResponseDTO {
	members := i.generateMembers(ownerEmail, memberEmails)
	box := i.buildBox(title, members)
	i.mongoDBAdapter.SaveBox(box)

	async(func() { i.emailNotifications.SendInvitations(title, members) })

	owner := selectOwner(members)
	return &CreateBoxResponseDTO{BoxKey: owner.Key}
}

func (i *Interactions) AddItem(boxKey string, message string) {
	item := NewItem(boxKey, message)
	box := i.updateBox(boxKey, item)
	audience := selectAudience(box.Members, boxKey)
	async(func() { i.emailNotifications.NotifyAudience(audience, box.Title) })
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


func (i *Interactions) updateBox(boxKey string, item *Item) *Box {
	box := i.mongoDBAdapter.loadBox(boxKey)
	box.Items = append(box.Items, *item)
	i.mongoDBAdapter.saveBox(box)
	return box
}

func (i *Interactions) generateMembers(ownerEmail string, memberEmails []string) []models.Member {
	nicknameGen := NewNicknameGenerator()

	members := []Member{}

	owner := models.Member{
		Key:      GenerateKey(),
		Email:    ownerEmail,
		Nickname: nicknameGen.Next(),
		Owner:    true,
	}
	members = append(members, owner)

	for _, email := range memberEmails {
		member := models.Member{
			Key:      GenerateKey(),
			Email:    email,
			Nickname: nicknameGen.Next(),
			Owner:    false,
		}
		members = append(members, member)
	}

	return members
}

func (i *Interactions) buildBox(title string, members []models.Member) *models.Box {
	return &models.Box{
		Title:        title,
		CreationDate: time.Now().Format(time.RFC3339),
		Members:      members,
		Items:        []models.Item{},
	}
}

func selectMember(key string, members []models.Member) *models.Member {
	for i, _ := range members {
		if members[i].Key == key {
			return &members[i]
		}
	}
	panic(models.SuprisingException{Err: fmt.Errorf("No member found for key:%s!", key)})
}

func selectOwner(members []models.Member) *models.Member {
	for i, _ := range members {
		if members[i].Owner {
			return &members[i]
		}
	}
	panic(models.SuprisingException{Err: fmt.Errorf("No owner found!")})
}

func selectAudience(members []models.Member, authorKey string) []models.Member {
	audience := []models.Member{}
	for _, member := range members {
		if member.Key != authorKey {
			audience = append(audience, member)
		}
	}
	return audience
}
