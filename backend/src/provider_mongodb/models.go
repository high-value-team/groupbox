package provider_mongodb

import (
	"time"

	"github.com/high-value-team/groupbox/backend/src/interior_models"
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

func ToBSONBox(box *interior_models.Box) *BSONBox {
	return &BSONBox{
		Title:        box.Title,
		CreationDate: box.CreationDate.UTC(),
		Members:      toBSONMembers(box.Members),
		Items:        toBSONItems(box.Items),
	}
}

func toBSONMembers(in []interior_models.Member) []BSONMember {
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

func toBSONItems(in []interior_models.Item) []BSONItem {
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

func ToBox(bsonBox *BSONBox) *interior_models.Box {
	return &interior_models.Box{
		Title:        bsonBox.Title,
		CreationDate: bsonBox.CreationDate.UTC(),
		Members:      toMembers(bsonBox.Members),
		Items:        toItems(bsonBox.Items),
	}
}

func toMembers(in []BSONMember) []interior_models.Member {
	out := []interior_models.Member{}
	for i := range in {
		out = append(out, interior_models.Member{
			Key:      in[i].Key,
			Email:    in[i].Email,
			Nickname: in[i].Nickname,
			Owner:    in[i].Owner,
		})
	}
	return out
}

func toItems(in []BSONItem) []interior_models.Item {
	out := []interior_models.Item{}
	for i := range in {
		out = append(out, interior_models.Item{
			CreationDate: in[i].CreationDate.UTC(),
			Subject:      in[i].Subject,
			Message:      in[i].Message,
			AuthorKey:    in[i].AuthorKey,
		})
	}
	return out
}
