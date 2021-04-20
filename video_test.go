package gotwilio

import "testing"

// TestRemoveParticipantFromRoom - Update with your creds to test.  This was more of a one time thing than intended to be run regularly
func TestRemoveParticipantFromRoom(t *testing.T) {
	client := NewTwilioClient("", "").WithAPIKey("", "")
	client.AuthToken = ""
	ex, err := client.RemoveParticipantFromRoom("", "")
	if err != nil {
		panic(err)
	}
	if ex != nil {
		panic(ex.Message)
	}
}
