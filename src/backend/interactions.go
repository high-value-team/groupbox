package backend

//TODO: Methoden entzerren, evtl auf verschiedene Klassen/Datei/Packages verteilen, zumindest in bessere Reihenfolge bringen

import (
	"fmt"
	"strings"
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
	box := i.buildBox(title, members)
	i.mongoDBAdapter.saveBox(box)

	async(func() { i.emailNotifications.SendInvitations(title, members) })

	owner := selectOwner(members)
	return &CreateBoxResponseDTO{BoxKey: owner.Key}
}

func (i *Interactions) AddItem(boxKey string, message string) {
	item := buildItem(boxKey, message)
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

func buildItem(boxKey string, message string) *Item {
	subject := extractSubject(message)
	return &Item{
		AuthorKey:    boxKey,
		CreationDate: time.Now().Format(time.RFC3339),
		Subject:      subject,
		Message:      message,
	}
}

/*
	A certain number of chars at the beginning of a message are taken as its subject.
	If the message is longer than that, "..." is appended to the subject.
	Any new line chars in the subject are replaced by spaces.
*/
func extractSubject(message string) string {
	const MAX_LEN_SUBJECT int = 15
	var subject string

	lenSubject := MAX_LEN_SUBJECT
	if lenSubject > len(message) {
		lenSubject = len(message)
	}
	subject = message[0:lenSubject]
	if lenSubject < len(message) {
		subject += "..."
	}

	subject = strings.Replace(subject, "\n", " ", -1)

	if subject == "" {
		subject = "?"
	}
	return subject
}

func (i *Interactions) updateBox(boxKey string, item *Item) *Box {
	box := i.mongoDBAdapter.loadBox(boxKey)
	box.Items = append(box.Items, *item)
	i.mongoDBAdapter.saveBox(box)
	return box
}

func (i *Interactions) generateMembers(ownerEmail string, memberEmails []string) []Member {
	nicknameGen := NewNicknameGenerator()

	members := []Member{}

	owner := Member{
		Key:      GenerateKey(),
		Email:    ownerEmail,
		Nickname: nicknameGen.Next(),
		Owner:    true,
	}
	members = append(members, owner)

	for _, email := range memberEmails {
		member := Member{
			Key:      GenerateKey(),
			Email:    email,
			Nickname: nicknameGen.Next(),
			Owner:    false,
		}
		members = append(members, member)
	}

	return members
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
	panic(SuprisingException{Err: fmt.Errorf("No member found for key:%s!", key)})
}

func selectOwner(members []Member) *Member {
	for i, _ := range members {
		if members[i].Owner {
			return &members[i]
		}
	}
	panic(SuprisingException{Err: fmt.Errorf("No owner found!")})
}

func selectAudience(members []Member, authorKey string) []Member {
	audience := []Member{}
	for _, member := range members {
		if member.Key != authorKey {
			audience = append(audience, member)
		}
	}
	return audience
}
