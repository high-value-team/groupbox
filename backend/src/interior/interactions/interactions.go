package interactions

//TODO: Methoden entzerren, evtl auf verschiedene Klassen/Datei/Packages verteilen, zumindest in bessere Reihenfolge bringen

import (
	"fmt"
	"strconv"
	"time"

	"github.com/high-value-team/groupbox/backend/src/exceptions"
	"github.com/high-value-team/groupbox/backend/src/interior/models"
	"github.com/high-value-team/groupbox/backend/src/providers"
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
	box := i.addItemToBox(boxKey, item)
	audience := selectAudience(box.Members, boxKey)
	async(func() { i.emailNotifications.NotifyAudience(audience, box.Title) })
}

func (i *Interactions) UpdateItem(boxKey, itemID, subject, message string) {
	box := i.updateItemInBox(boxKey, itemID, subject, message)
	audience := selectAudience(box.Members, boxKey)
	async(func() { i.emailNotifications.NotifyAudience(audience, box.Title) })
}

func (i *Interactions) DeleteItem(boxKey, itemID string) {
	box := i.deleteItemInBox(boxKey, itemID)
	audience := selectAudience(box.Members, boxKey)
	async(func() { i.emailNotifications.NotifyAudience(audience, box.Title) })
}

func (i *Interactions) addItemToBox(boxKey string, item *models.Item) *models.Box {
	box := i.mongoDBAdapter.LoadBox(boxKey)
	box.Items = append(box.Items, *item)
	i.mongoDBAdapter.SaveBox(box)
	return box
}

func (i *Interactions) updateItemInBox(boxKey, itemID, subject, message string) *models.Box {
	box := i.mongoDBAdapter.LoadBox(boxKey)
	item := selectItem(box.Items, itemID)
	changeItem(item, subject, message)
	i.mongoDBAdapter.SaveBox(box)
	return box
}

func (i *Interactions) deleteItemInBox(boxKey, itemID string) *models.Box {
	box := i.mongoDBAdapter.LoadBox(boxKey)
	box.Items = deleteFromItems(box.Items, itemID)
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
	panic(exceptions.SuprisingException{Err: fmt.Errorf("No owner found!")})
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

func selectItem(items []models.Item, itemID string) *models.Item {
	for i := range items {
		if strconv.Itoa(i) == itemID {
			return &items[i]
		}
	}
	panic(exceptions.SuprisingException{Err: fmt.Errorf("No item found!")})
}

func changeItem(item *models.Item, subject, message string) {
	item.Subject = subject
	item.Message = message
}

func deleteFromItems(items []models.Item, itemID string) []models.Item {
	for i := range items {
		if strconv.Itoa(i) == itemID {
			return append(items[:i], items[i+1:]...)
		}
	}
	panic(exceptions.SuprisingException{Err: fmt.Errorf("No item found!")})
}
