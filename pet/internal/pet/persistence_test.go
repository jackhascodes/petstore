// +build integration

package pet

import "testing"

func setup() (*MySQLPersistence, error) {
	return InitMySQLPersistence("testdb",3306,"root","root","pet")
}

func mockPet() *Pet {
	pet = InitPet("testPet", []string{"testphoto"})
	pet.Tags = []*Meta{{1,"testing"}}
	pet.Category = &Meta{1, "testCat"}
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
	db, err := setup()
	defer db.db.Close()
	pet := mockPet()
	inserted, err := db.Insert(pet)
	found := db.FindById(inserted.Id)
	if found == nil || found.Id == 0 {
		t.Errorf("could not find pet")
	}
	if found.Category.Name != "testCat" {
		t.Errorf("could not populate category")
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
	db, err := setup()
	defer db.db.Close()
	pet := mockPet()
	inserted, err := db.Insert(pet)
	found := db.FindByStatus(AVAILABLE)
	if len(found) == 0 {
		t.Errorf("findByStatus returned 0 results")
	}

	db.Delete(inserted.Id)
}

func TestMySQLPersistence_Update(t *testing.T) {
	db, err := setup()
	defer db.db.Close()
	pet := mockPet()
	inserted, err := db.Insert(pet)
	inserted.Status = SOLD
	err := db.Update(inserted)
	if err != nil {
		t.Errorf("should not have errored: %v", err)
	}
	check := db.FindById(inserted.Id)
	if check.Status != SOLD {
		t.Errorf("did not update")
	}
}

func TestMySQLPersistence_GetStatusCounts(t *testing.T) {
	db, err := setup()
	defer db.db.Close()
	pet := mockPet()
	inserted, err := db.Insert(pet)
	inv := db.GetStatusCounts()
	if inv[AVAILABLE.String()] == 0 {
		t.Errorf("expected 1 result, got 0")
	}
}






