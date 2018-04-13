package providers

import (
	"time"

	interiorModels "github.com/high-value-team/groupbox/backend/src/interior/models"
)

type BSONBox struct {
	Title        string       `bson:"title"`
	CreationDate time.Time    `bson:"creationDate"`
	Members      []BSONMember `bson:"members"`
	Items        []BSONItem   `bson:"items"`
}

type BSONItem struct {
	CreationDate time.Time `bson:"creationDate"`
	Subject      string    `bson:"subject"`
	Message      string    `bson:"message"`
	AuthorKey    string    `bson:"authorKey"`
}

type BSONMember struct {
	Key      string `bson:"key"`
	Email    string `bson:"email"`
	Nickname string `bson:"nickname"`
	Owner    bool   `bson:"owner"`
}

func ToBSONBox(box *interiorModels.Box) *BSONBox {
	return &BSONBox{
		Title:        box.Title,
		CreationDate: box.CreationDate.UTC(),
		Members:      toBSONMembers(box.Members),
		Items:        toBSONItems(box.Items),
	}
}

func toBSONMembers(in []interiorModels.Member) []BSONMember {
	out := []BSONMember{}
	for i := range in {
		out = append(out, BSONMember{
			Key:      in[i].Key,
			Email:    in[i].Email,
			Nickname: in[i].Nickname,
			Owner:    in[i].Owner,
		})
	}
	return out
}

func toBSONItems(in []interiorModels.Item) []BSONItem {
	out := []BSONItem{}
	for i := range in {
		out = append(out, BSONItem{
			CreationDate: in[i].CreationDate.UTC(),
			Subject:      in[i].Subject,
			Message:      in[i].Message,
			AuthorKey:    in[i].AuthorKey,
		})
	}
	return out
}

func ToBox(bsonBox *BSONBox) *interiorModels.Box {
	return &interiorModels.Box{
		Title:        bsonBox.Title,
		CreationDate: bsonBox.CreationDate.UTC(),
		Members:      toMembers(bsonBox.Members),
		Items:        toItems(bsonBox.Items),
	}
}

func toMembers(in []BSONMember) []interiorModels.Member {
	out := []interiorModels.Member{}
	for i := range in {
		out = append(out, interiorModels.Member{
			Key:      in[i].Key,
			Email:    in[i].Email,
			Nickname: in[i].Nickname,
			Owner:    in[i].Owner,
		})
	}
	return out
}

func toItems(in []BSONItem) []interiorModels.Item {
	out := []interiorModels.Item{}
	for i := range in {
		out = append(out, interiorModels.Item{
			CreationDate: in[i].CreationDate.UTC(),
			Subject:      in[i].Subject,
			Message:      in[i].Message,
			AuthorKey:    in[i].AuthorKey,
		})
	}
	return out
}
