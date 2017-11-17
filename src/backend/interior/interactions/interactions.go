package interactions

//TODO: Methoden entzerren, evtl auf verschiedene Klassen/Datei/Packages verteilen, zumindest in bessere Reihenfolge bringen

import (
	"fmt"
	"time"

	"github.com/high-value-team/groupbox/src/backend/models"
	"github.com/high-value-team/groupbox/src/backend/providers"
)

type Interactions struct {
	mongoDBAdapter     *providers.MongoDBAdapter
	emailNotifications *providers.EmailNotifications
}

func NewInteractions(mongoDBAdapter *providers.MongoDBAdapter, emailNotifications *providers.EmailNotifications) *Interactions {
	return &Interactions{mongoDBAdapter: mongoDBAdapter, emailNotifications: emailNotifications}
}

func (i *Interactions) GetBox(boxKey string) *models.Box {
	return i.mongoDBAdapter.LoadBox(boxKey)
}

func (i *Interactions) CreateBox(title, ownerEmail string, memberEmails []string) *models.Member {
	members := i.generateMembers(ownerEmail, memberEmails)
	box := i.buildBox(title, members)
	i.mongoDBAdapter.SaveBox(box)

	async(func() { i.emailNotifications.SendInvitations(title, members) })

	return selectOwner(members)
}

func (i *Interactions) AddItem(boxKey string, message string) {
	item := models.NewItem(boxKey, message)
	box := i.updateBox(boxKey, item)
	audience := selectAudience(box.Members, boxKey)
	async(func() { i.emailNotifications.NotifyAudience(audience, box.Title) })
}

func (i *Interactions) updateBox(boxKey string, item *models.Item) *models.Box {
	box := i.mongoDBAdapter.LoadBox(boxKey)
	box.Items = append(box.Items, *item)
	i.mongoDBAdapter.SaveBox(box)
	return box
}

func (i *Interactions) generateMembers(ownerEmail string, memberEmails []string) []models.Member {
	nicknameGen := NewNicknameGenerator()

	members := []models.Member{}

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
		CreationDate: time.Now(),
		Members:      members,
		Items:        []models.Item{},
	}
}

func selectOwner(members []models.Member) *models.Member {
	for i := range members {
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
