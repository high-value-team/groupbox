package backend

import (
	"strings"
	"time"
)

type Box struct {
	Title        string   `bson:"title"`
	CreationDate string   `bson:"creationDate"`
	Members      []Member `bson:"members"`
	Items        []Item   `bson:"items"`
}

type Item struct {
	CreationDate string `bson:"creationDate"`
	Subject      string `bson:"subject"`
	Message      string `bson:"message"`
	AuthorKey    string `bson:"authorKey"`
}

type Member struct {
	Key      string `bson:"key"`
	Email    string `bson:"email"`
	Nickname string `bson:"nickname"`
	Owner    bool   `bson:"owner"`
}

type BoxDTO struct {
	Title          string    `json:"title"`
	MemberNickname string    `json:"memberNickname"`
	CreationDate   string    `json:"creationDate"`
	Items          []ItemDTO `json:"items"`
}

type ItemDTO struct {
	AuthorNickname string `json:"authorNickname"`
	CreationDate   string `json:"creationDate"`
	Subject        string `json:"subject"`
	Message        string `json:"message"`
}

type SadException struct {
	Err error
}

func (e *SadException) Message() string {
	return e.Err.Error()
}

type SuprisingException struct {
	Err error
}

func (e *SuprisingException) Message() string {
	return e.Err.Error()
}




func NewItem(authorKey string, message string) *Item {
	subject, message := extractSubject(message)
	return &Item{
		AuthorKey:    authorKey,
		CreationDate: time.Now().Format(time.RFC3339),
		Subject:      subject,
		Message:      message,
	}
}

func extractSubject(message string) (string, string) {
	message = strings.TrimSpace(message)
	return tryExtractExplicitSubject(message,
				extractImplicitSubject)
}

/*
	Subject is explicitly marked in message: Subject is enclosed in "<"...">" and
	is located right at the beginning of the message.
	The subject is extracted and then removed from the message.
 */
func tryExtractExplicitSubject(message string,
	 						   onNotFound func (string) (string,string)) (string,string) {
	if len(message) > 0 && message[0:1] == "<" {
		endOfSubjectIndex := strings.Index(message,">")
		if endOfSubjectIndex < 0 {endOfSubjectIndex = len(message)}
		subject := message[1:endOfSubjectIndex]

		startOfMessageBody := endOfSubjectIndex+1
		if startOfMessageBody < len(message) {
			message = message[startOfMessageBody:len(message)]
		} else {
			message = ""
		}

		return subject, message
	} else {
		return onNotFound(message)
	}
}

/*
	Subject is taken from first couple of chars in message.
	If subject is less than the message "..." is appended.
	The message itself remains unchanged.
 */
func extractImplicitSubject(message string) (string,string) {
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
	return subject, message
}