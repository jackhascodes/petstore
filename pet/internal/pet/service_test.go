package pet

import (
	"errors"
	"testing"
)

func TestService_AddPet(t *testing.T) {
	tests := []struct {
		name      string
		s         *Service
		errExists bool
		petExists bool
		petId     int64
	}{
		{name: "good", s: &Service{db: HappyPathDB{}}, errExists: false, petExists: true, petId: 1},
		{name: "bad", s: &Service{db: SadPathDB{}}, errExists: true, petExists: false, petId: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := tt.s.AddPet(InitPet("test", []string{}))
			if (err != nil) != tt.errExists {
				t.Errorf("should have received an error")
			}
			if (p != nil) != (tt.petExists) {
				t.Fatalf("expected a pet? %t - got a pet? %t", tt.petExists, p != nil)
			}

			if p != nil && p.Id != tt.petId {
				t.Errorf("expected %d, got %d", p.Id, tt.petId)
			}
		})
	}
}

// TODO: Write testcases
func TestService_DeletePet(t *testing.T) {
	tests := []struct {
		name string
		s    *Service
	}{
		{name: "good", s: &Service{db: HappyPathDB{}}},
		{name: "bad", s: &Service{db: SadPathDB{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

// TODO: Write testcases
func TestService_FindById(t *testing.T) {
	tests := []struct {
		name string
		s    *Service
	}{
		{name: "good", s: &Service{db: HappyPathDB{}}},
		{name: "bad", s: &Service{db: SadPathDB{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}

}

// TODO: Write testcases
func TestService_FindByStatus(t *testing.T) {
	tests := []struct {
		name string
		s    *Service
	}{
		{name: "good", s: &Service{db: HappyPathDB{}}},
		{name: "bad", s: &Service{db: SadPathDB{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}

}

// TODO: Write testcases
func TestService_UpdatePet(t *testing.T) {
	tests := []struct {
		name string
		s    *Service
	}{
		{name: "good", s: &Service{db: HappyPathDB{}}},
		{name: "bad", s: &Service{db: SadPathDB{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}

}

// TODO: Write testcases
func TestService_UpdatePetStatus(t *testing.T) {
	tests := []struct {
		name string
		s    *Service
	}{
		{name: "good", s: &Service{db: HappyPathDB{}}},
		{name: "bad", s: &Service{db: SadPathDB{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}

}

// Mocks
type HappyPathDB struct{}
type SadPathDB struct{}

func (h HappyPathDB) FindById(id int64) *Pet {
	return &Pet{Id: id,
		Name:      "testPet",
		PhotoUrls: []string{"http://localhost/image.png"},
		Status:    AVAILABLE,
		Category:  &Meta{Name: "testing", Id: 1},
	}

}
func (h HappyPathDB) FindByStatus(status Status) []*Pet {
	switch status {
	case AVAILABLE:
		return []*Pet{
			{Id: 1,
				Name:      "testPet1",
				PhotoUrls: []string{"http://localhost/image.png"},
				Status:    AVAILABLE,
				Category:  &Meta{Name: "testing", Id: 1},
			},
			{Id: 2,
				Name:      "testPet2",
				PhotoUrls: []string{"http://localhost/image.png"},
				Status:    AVAILABLE,
				Category:  &Meta{Name: "testing", Id: 1},
			},
		}
	case PENDING:
		return []*Pet{
			{Id: 3,
				Name:      "testPet3",
				PhotoUrls: []string{"http://localhost/image.png"},
				Status:    PENDING,
				Category:  &Meta{Name: "testing", Id: 1},
			},
		}
	default:
		return []*Pet{
			{Id: 4,
				Name:      "testPet4",
				PhotoUrls: []string{"http://localhost/image.png"},
				Status:    SOLD,
				Category:  &Meta{Name: "testing", Id: 1},
			},
		}
	}
}
func (h HappyPathDB) Update(pet *Pet) error {
	return nil
}
func (h HappyPathDB) Insert(pet *Pet) (*Pet, error) {
	pet.Id = 1
	return pet, nil
}
func (h HappyPathDB) Delete(id int64) error {
	return nil
}

func (h HappyPathDB) GetStatusCounts() *Inventory {
	return &Inventory{"available":1}
}
func (h SadPathDB) FindById(id int64) *Pet {
	return nil
}
func (h SadPathDB) FindByStatus(Status) []*Pet {
	return []*Pet{}
}
func (h SadPathDB) Update(pet *Pet) error {
	return sadPathError
}
func (h SadPathDB) Insert(pet *Pet) (*Pet, error) {
	return nil, sadPathError
}
func (h SadPathDB) Delete(id int64) error {
	return sadPathError
}
func (h SadPathDB) GetStatusCounts() *Inventory {
	return &Inventory{}
}
var sadPathError = errors.New("database error")
