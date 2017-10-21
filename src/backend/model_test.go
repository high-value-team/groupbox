package backend

import (
	"testing"
)


func TestExtractExplicitSubject(t *testing.T) {
	subject,message := tryExtractExplicitSubject("<s>m", nil)
	assertEqual("s", subject, func() { t.Errorf("Subject expected 's', found '%s'.", subject) })
	assertEqual("m", message, func() { t.Errorf("Message expected 'm', found '%s'.", message) })

	subject,message = tryExtractExplicitSubject("<s>", nil)
	assertEqual("s", subject, func() { t.Errorf("Subject expected 's', found '%s'.", subject) })
	assertEqual("", message, func() { t.Errorf("No message expected, found '%s'.", message) })

	subject,message = tryExtractExplicitSubject("<>m", nil)
	assertEqual("", subject, func() { t.Errorf("No subject expected, found '%s'.", subject) })
	assertEqual("m", message, func() { t.Errorf("Message expected 'm', found '%s'.", message) })

	subject,message = tryExtractExplicitSubject("<sm", nil)
	assertEqual("sm", subject, func() { t.Errorf("Subject expected 'sm', found '%s'.", subject) })
	assertEqual("", message, func() { t.Errorf("No message expected, found '%s'.", message) })

	subject,message = tryExtractExplicitSubject("sm", func(string)(string,string){return "a","b"})
	assertEqual("a", subject, func() { t.Errorf("Subject expected 'a', found '%s'.", subject) })
	assertEqual("b", message, func() { t.Errorf("Message expected 'b', found '%s'.", message) })
}

func assertEqual(expected string, actual string, onErr func()) {
	if expected != actual {
		onErr()
	}
}


func TestExtractImplicitSubject(t *testing.T) {
	subject,message := extractImplicitSubject("12345678901234567890")
	if subject != "123456789012345..." {
		t.Errorf("Expected subject shorter than message; result: <%s>", subject)
	}
	if message != "12345678901234567890" {
		t.Errorf("Message should not be changed when extracting implicit subject.")
	}

	subject,_ = extractImplicitSubject("")
	if subject != "?" {
		t.Errorf("Expected subject <?> for empty message", subject)
	}

	subject,_ = extractImplicitSubject("1234567890")
	if subject != "1234567890" {
		t.Errorf("Expected subject equal to message shorter than max subject len; result <%s>", subject)
	}

	subject,_ = extractImplicitSubject("123456\n7890")
	if subject != "123456 7890" {
		t.Errorf("Expected subject to contain no new line.", subject)
	}
}
