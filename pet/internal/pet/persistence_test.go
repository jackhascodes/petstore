// +build integration

package pet

import (
	"encoding/json"
	"testing"
)

func setup() (*MySQLPersistence, error) {
	return InitMySQLPersistence("mysql-pets", "3306", "test", "test", "pet")
}

func mockPet() *Pet {
	pet := InitPet("testPet", []string{"testphoto"})
	pet.Tags = []*Meta{{1, "independent"}}
	pet.Category = &Meta{1, "dogs"}
	pet.Status = AVAILABLE
	return pet
}

func TestMySQLPersistence_Insert(t *testing.T) {
	db, err := setup()
	defer db.db.Close()
	pet := mockPet()
	inserted, err := db.Insert(pet)
	if err != nil {
		t.Errorf("should not have errored: %v", err)
	}
	if inserted.Id == 0 {
		t.Errorf("expected a non-zero Id on insert")
	}
	db.Delete(inserted.Id)
}

func TestMySQLPersistence_FindById(t *testing.T) {
	db, _ := setup()
	defer db.db.Close()
	pet := mockPet()
	inserted, _ := db.Insert(pet)
	found := db.FindById(inserted.Id)
	if found == nil || found.Id == 0 {
		t.Errorf("could not find pet")
	}
	if found.Category.Name != "dogs" {
		s, _ := json.Marshal(pet)
		t.Errorf("could not populate category: %s", string(s))
	}
	if len(found.Tags) == 0 {
		t.Errorf("could not populate tags")
	}
	if len(found.PhotoUrls) == 0 {
		t.Errorf("could not populate photoUrls")
	}

	db.Delete(inserted.Id)
}

func TestMySQLPersistence_FindByStatus(t *testing.T) {
	db, _ := setup()
	defer db.db.Close()
	found := db.FindByStatus(AVAILABLE)
	if len(found) == 0 {
		t.Errorf("findByStatus returned 0 results")
	}
}

func TestMySQLPersistence_Update(t *testing.T) {
	db, err := setup()
	defer db.db.Close()
	pet := mockPet()
	inserted, err := db.Insert(pet)
	inserted.Status = SOLD
	err = db.Update(inserted)
	if err != nil {
		t.Errorf("should not have errored: %v", err)
	}
	check := db.FindById(inserted.Id)
	if check.Status != SOLD {
		t.Errorf("did not update")
	}
}

func TestMySQLPersistence_GetStatusCounts(t *testing.T) {
	db, _ := setup()
	defer db.db.Close()
	pet := mockPet()
	db.Insert(pet)

	inv := *db.GetStatusCounts()
	if inv[AVAILABLE.String()] == 0 {
		t.Errorf("expected results, got 0")
	}
}
