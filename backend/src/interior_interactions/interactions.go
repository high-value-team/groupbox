package interior_interactions

//TODO: Methoden entzerren, evtl auf verschiedene Klassen/Datei/Packages verteilen, zumindest in bessere Reihenfolge bringen

import (
	"fmt"
	"strconv"
	"time"

	"github.com/high-value-team/groupbox/backend/src/interior_models"
	"github.com/high-value-team/groupbox/backend/src/provider_mongodb"
	"github.com/high-value-team/groupbox/backend/src/provider_smtp"
)

type Interactions struct {
	mongoDBAdapter     *provider_mongodb.MongoDBAdapter
	emailNotifications *provider_smtp.EmailNotifications
}

func NewInteractions(mongoDBAdapter *provider_mongodb.MongoDBAdapter, emailNotifications *provider_smtp.EmailNotifications) *Interactions {
	return &Interactions{mongoDBAdapter: mongoDBAdapter, emailNotifications: emailNotifications}
}

func (i *Interactions) GetBox(boxKey string) *interior_models.Box {
	return i.mongoDBAdapter.LoadBox(boxKey)
}

func (i *Interactions) CreateBox(title, ownerEmail string, memberEmails []string) *interior_models.Member {
	members := i.generateMembers(ownerEmail, memberEmails)
	box := i.buildBox(title, members)
	i.mongoDBAdapter.SaveBox(box)

	async(func() { i.emailNotifications.SendInvitations(title, members) })

	return selectOwner(members)
}

func (i *Interactions) AddItem(boxKey string, message string) {
	item := interior_models.NewItem(boxKey, message)
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

func (i *Interactions) addItemToBox(boxKey string, item *interior_models.Item) *interior_models.Box {
	box := i.mongoDBAdapter.LoadBox(boxKey)
	box.Items = append(box.Items, *item)
	i.mongoDBAdapter.SaveBox(box)
	return box
}

func (i *Interactions) updateItemInBox(boxKey, itemID, subject, message string) *interior_models.Box {
	box := i.mongoDBAdapter.LoadBox(boxKey)
	item := selectItem(box.Items, itemID)
	changeItem(item, subject, message)
	i.mongoDBAdapter.SaveBox(box)
	return box
}

func (i *Interactions) deleteItemInBox(boxKey, itemID string) *interior_models.Box {
	box := i.mongoDBAdapter.LoadBox(boxKey)
	box.Items = deleteFromItems(box.Items, itemID)
	i.mongoDBAdapter.SaveBox(box)
	return box
}

func (i *Interactions) generateMembers(ownerEmail string, memberEmails []string) []interior_models.Member {
	nicknameGen := NewNicknameGenerator()

	members := []interior_models.Member{}

	owner := interior_models.Member{
		Key:      GenerateKey(),
		Email:    ownerEmail,
		Nickname: nicknameGen.Next(),
		Owner:    true,
	}
	members = append(members, owner)

	for _, email := range memberEmails {
		member := interior_models.Member{
			Key:      GenerateKey(),
			Email:    email,
			Nickname: nicknameGen.Next(),
			Owner:    false,
		}
		members = append(members, member)
	}

	return members
}

func (i *Interactions) buildBox(title string, members []interior_models.Member) *interior_models.Box {
	return &interior_models.Box{
		Title:        title,
		CreationDate: time.Now(),
		Members:      members,
		Items:        []interior_models.Item{},
	}
}

func selectOwner(members []interior_models.Member) *interior_models.Member {
	for i := range members {
		if members[i].Owner {
			return &members[i]
		}
	}
	panic(interior_models.SuprisingException{Err: fmt.Errorf("No owner found!")})
}

func selectAudience(members []interior_models.Member, authorKey string) []interior_models.Member {
	audience := []interior_models.Member{}
	for _, member := range members {
		if member.Key != authorKey {
			audience = append(audience, member)
		}
	}
	return audience
}

func selectItem(items []interior_models.Item, itemID string) *interior_models.Item {
	for i := range items {
		if strconv.Itoa(i) == itemID {
			return &items[i]
		}
	}
	panic(interior_models.SuprisingException{Err: fmt.Errorf("No item found!")})
}

func changeItem(item *interior_models.Item, subject, message string) {
	item.Subject = subject
	item.Message = message
}

func deleteFromItems(items []interior_models.Item, itemID string) []interior_models.Item {
	for i := range items {
		if strconv.Itoa(i) == itemID {
			return append(items[:i], items[i+1:]...)
		}
	}
	panic(interior_models.SuprisingException{Err: fmt.Errorf("No item found!")})
}
