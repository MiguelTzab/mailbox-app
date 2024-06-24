package entity

import (
	"reflect"
	"testing"
)

func TestEmailFromFile(t *testing.T) {
	path := "../../fixtures/email-0.txt"

	expectedEmail := Email{
		MessageId: "<24084559.1075857578492.JavaMail.evans@thyme>",
		From:      "john.arnold@enron.com",
		To:        []string{"gregory.carraway@enron.com"},
		Subject:   "Re:",
		Date:      "Mon, 10 Jul 2000 07:43:00 -0700 (PDT)",
		Body:      "I'll send 2 invites up.  What's your location?\n\nGregory Carraway@ENRON\n07/10/2000 02:36 PM\nTo: John Arnold/HOU/ECT@ECT\ncc:\nSubject: Re:\n\nThank you for the invitation. I would love to attend. I would like to invite\nmy wife, if that would be ok. Also, could you tell me where the Mercantile\nbar is located? Thank you!!!\n\n",
	}

	email, err := EmailFromFile(path)
	if err != nil {
		t.Fatalf("Error parsing email: %v", err)
	}

	if !reflect.DeepEqual(email, expectedEmail) {
		t.Errorf("Parsed email does not match expected.\nGot: %+v\nExpected: %+v", email, expectedEmail)
	}
}
