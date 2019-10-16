package pet

import (
	"encoding/json"
	"errors"
)

type Pet struct {
	Id        int64    `json:"id,omitempty"`
	Category  *Meta    `json:"category,omitempty"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []*Meta  `json:"tags,omitempty"`
	Status    Status   `json:"status,omitempty"`
}

// InitPet creates an instance of Pet. name and photoUrls are required fields for this type.
func InitPet(name string, photoUrls []string) *Pet {
	return &Pet{Name: name, PhotoUrls: photoUrls}
}

type Meta struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Status int

const (
	INVALID Status = iota
	AVAILABLE
	PENDING
	SOLD
)

func (s Status) String() string {
	return [...]string{"", "available", "pending", "sold"}[s]
}
func (s *Status) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	return s.FromString(j)
}

func (s *Status) FromString(st string) error {
	strings := map[string]Status{
		"available": AVAILABLE,
		"pending":   PENDING,
		"sold":      SOLD,
	}
	*s = strings[st]
	if *s == INVALID {
		return errors.New("invalid value for type Status")
	}
	return nil
}
