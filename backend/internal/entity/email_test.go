package entity

import (
	"reflect"
	"testing"
)

func TestEmailFromFile(t *testing.T) {
	path := "../../fixtures/email-0.txt"

	expectedEmail := Email{
		From:    "john.arnold@enron.com",
		To:      []string{"gregory.carraway@enron.com"},
		Subject: "Re:",
		Body:    "I'll send 2 invites up.  What's your location?\nGregory Carraway@ENRON\n07/10/2000 02:36 PM\nTo: John Arnold/HOU/ECT@ECT\ncc:\nSubject: Re:\nThank you for the invitation. I would love to attend. I would like to invite\nmy wife, if that would be ok. Also, could you tell me where the Mercantile\nbar is located? Thank you!!!",
	}

	email, err := EmailFromFile(path)
	if err != nil {
		t.Fatalf("Error parsing email: %v", err)
	}

	if !reflect.DeepEqual(email, expectedEmail) {
		t.Errorf("Parsed email does not match expected.\nGot: %+v\nExpected: %+v", email, expectedEmail)
	}
}
