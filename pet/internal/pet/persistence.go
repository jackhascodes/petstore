package pet

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"strings"
)

type Persistence interface {
	FindById(id int64) *Pet
	FindByStatus(Status) []*Pet
	Update(pet *Pet) error
	Insert(pet *Pet) (*Pet, error)
	Delete(id int64) error
	GetStatusCounts() Inventory
}

func InitMySQLPersistence(host, port, user, pass, dbName string) (*MySQLPersistence, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbName)
	dbConn, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	return &MySQLPersistence{connString: connString, db: dbConn, log: logrus.New()}, nil
}

type MySQLPersistence struct {
	connString string
	db         *sql.DB
	log        *logrus.Logger
}

func (p *MySQLPersistence) FindById(id int64) *Pet {
	pet := &Pet{}
	cat := &Meta{}
	err := p.db.QueryRow(
		`
	SELECT id, name, status, categoryId, c.name AS categoryName 
	FROM pets p
	JOIN categories c ON p.categoryId = c.id 
	WHERE id = ?`,
		id,
	).Scan(&pet.Id, &pet.Name, &pet.Status, &cat.Id, &cat.Name)

	if err != nil {
		p.log.Error(err)
		return nil
	}

	pet.Category = cat
	pet.PhotoUrls = p.FindPhotoUrls(id)
	pet.Tags = p.FindTags(id)

	return pet
}

func (p *MySQLPersistence) FindPhotoUrls(id int64) []string {
	var urls []string
	res, err := p.db.Query(
		`SELECT url FROM pet_photos WHERE petId = ?`,
		id,
	)
	if err != nil {
		p.log.Error(err)
		return urls
	}
	for res.Next() {
		var url string
		err = res.Scan(url)
		if err != nil {
			p.log.Error(err)
		} else {
			urls = append(urls, url)
		}
	}
	return urls
}

func (p MySQLPersistence) FindTags(id int64) []*Meta {
	var tags []*Meta
	{
	}
	res, err := p.db.Query(`
		SELECT t.id, t.tag 
		FROM pet_tags p
		JOIN tags t ON p.tagId = t.id
		WHERE p.petId = ?`,
		id,
	)
	if err != nil {
		p.log.Error(err)
		return tags
	}
	for res.Next() {
		tag := &Meta{}
		err = res.Scan(&tag.Id, &tag.Name)

		if err != nil {
			p.log.Error(err)
		} else {
			tags = append(tags, tag)
		}
	}
}
func (p *MySQLPersistence) FindByStatus(status Status) []*Pet {
	var pets []*Pet
	res, err := p.db.Query(`
		SELECT id, name, status, categoryId, c.name AS categoryName 
		FROM pets p
		JOIN categories c ON p.categoryId = c.id 
		WHERE status = ?`,
		status.String())

	if err != nil {
		p.log.Error(err)
		return pets
	}

	for res.Next() {
		pet := &Pet{}
		cat := &Meta{}
		err = res.Scan(&pet.Id, &pet.Name, &pet.Status, &cat.Id, &cat.Name)

		if err != nil {
			p.log.Error(err)
		} else {
			pet.Category = cat
			pet.PhotoUrls = p.FindPhotoUrls(pet.Id)
			pet.Tags = p.FindTags(pet.Id)
			pets = append(pets, pet)
		}
	}
	return pets
}

func (p *MySQLPersistence) Update(pet *Pet) error {
	err := p.updatePet(pet)
	if err != nil {
		return err
	}
	err = p.updateTags(pet.Id, pet.Tags)
	if err != nil {
		return err
	}
	err = p.updatePhots(pet.Id, pet.PhotoUrls)
	if err != nil {
		return err
	}
	return nil
}

func (p *MySQLPersistence) updatePet(pet *Pet) error {
	res, err := p.db.Exec(`
			UPDATE pets SET 
				name = ?,
				status = ?,
				category = ?
			WHERE id = ?`,
		pet.Name,
		pet.Status.String(),
		pet.Category.Id,
		pet.Id)
	p.log.Infof("Update pet (%d): %v", pet.Id, res)
	return err
}

func (p *MySQLPersistence) updateTags(id int64, tags []*Meta) error {
	if id == 0 {
		return nil
	}
	res, err := p.db.Exec(`DELETE FROM pet_tags WHERE petId = ?`)
	if err != nil {
		p.log.Errorf("clear tags failed: res %v, err %v", res,err)
	}
	return p.insertTags(id, tags)

}

func (p *MySQLPersistence) insertTags(id int64, tags []*Meta) error {
	var vals []string
	for _, tag := range tags {
		if tag.Id > 0 {
			vals = append(vals, fmt.Sprintf("(%d,%d)", id, tag.Id))
		}
	}
	res, err := p.db.Exec(`
		INSERT IGNORE INTO pet_tags (petId, tagId)
		VALUES
		` + strings.Join(vals, ","))
	if err != nil {
		p.log.Errorf("add tags failed: res %v, err %v", res,err)
		return err
	}
}

func (p *MySQLPersistence) updatePhotos(id int64, photos []string) error {
	if id == 0 {
		return nil
	}
	res, err := p.db.Exec(`DELETE FROM pet_photos WHERE petId = ?`)
	if err != nil {
		p.log.Errorf("clear photos failed: res %v, err %v", res,err)
	}
	return p.insertPhotos(id, photos)

}

func (p *MySQLPersistence) insertPhotos(id int64, photos []string) error {
	// TODO: finish this
	// back at the keyboard
}


func (p *MySQLPersistence) Insert(pet *Pet) (*Pet, error) {
	// TODO: finish this
}
func (p *MySQLPersistence) Delete(id int64) error         {}
func (p *MySQLPersistence) GetStatusCounts() Inventory    {}
