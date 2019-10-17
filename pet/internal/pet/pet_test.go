// +build unit

package pet

import (
	"encoding/json"
	"testing"
)

func TestInitPet(t *testing.T) {
	p := InitPet("test", []string{"http://test.com"})
	if p.Name != "test" {
		t.Errorf("pet instance did not populate name")
	}
	if len(p.PhotoUrls) != 1 {
		t.Errorf("pet instance did not populate photoUrls")
	}
}

func TestStatus_String(t *testing.T) {
	a := AVAILABLE.String()
	p := PENDING.String()
	s := SOLD.String()

	if a != "available" {
		t.Errorf("Expected 'available', got %s ", a)
	}
	if p != "pending" {
		t.Errorf("Expected 'pending', got %s ", p)
	}
	if s != "sold" {
		t.Errorf("Expected 'sold', got %s ", s)
	}
}

func TestStatus_UnmarshalJSON(t *testing.T) {
	r := map[string]Status{}
	d := []byte(`
{
	"soldstatus":"sold",
	"availablestatus":"available",
	"pendingstatus":"pending"
}`)
	err := json.Unmarshal(d, &r)
	if err != nil {
		t.Errorf("should not have errored on unmarshal: %v", err)
	}
	if r["soldstatus"] != SOLD {
		t.Errorf("expected SOLD, got %v", r["soldstatus"])
	}
	if r["availablestatus"] != AVAILABLE {
		t.Errorf("expected AVAILABLE, got %v", r["availablestatus"])
	}
	if r["pendingstatus"] != PENDING {
		t.Errorf("expected PENDING, got %v", r["pendingstatus"])
	}
}

func TestStatus_UnmarshalJSON_InvalidCase(t *testing.T) {
	r := map[string]Status{}
	d := []byte(`
{
	"status":"invalid"
}`)
	err := json.Unmarshal(d, &r)
	if err == nil {
		t.Errorf("this should have thrown an error, got %v", r)
	}
}
