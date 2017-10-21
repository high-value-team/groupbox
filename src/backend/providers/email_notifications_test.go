// +build smtp

// Diese Test-Datei wird nur ausgeführt, wenn das Build-Tag mit angegeben wird:
//export SMTP_USERNAME="groupbox@ralfw.de"
//export SMTP_PASSWORD="geheim"
//export SMTP_NO_REPLY_EMAIL="no-reply-groupbox@ralfw.de"
//export SMTP_SERVER_ADDRESS="sslout.df.eu:587"
// go test -tags=smtp

package providers

import (
	"os"
	"testing"
)

func TestSendInvitations(t *testing.T) {
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	noReplayEmail := os.Getenv("SMTP_NO_REPLY_EMAIL")
	SMTPServerAddress := os.Getenv("SMTP_SERVER_ADDRESS")

	// arrange
	emailNotifications := EmailNotifications{
		Domain:            "http://localhost:8080",
		NoReplyEmail:      noReplayEmail,
		SMTPServerAddress: SMTPServerAddress,
		Username:          username,
		Password:          password,
	}
	members := []Member{
		{
			Key:   "1",
			Email: "florian@fnbk.cc",
		},
	}

	// act
	emailNotifications.SendInvitations("unit test", members)

	// assert
	// check email!
}
