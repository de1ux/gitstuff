package git

import "testing"

func TestTicketNumber(t *testing.T) {
	out := TicketNumber("ne", "ne/APIT-1234")
	if out != "APIT-1234" {
		t.Fatalf("Expected APIT-1234, got " + out)
	}
}
