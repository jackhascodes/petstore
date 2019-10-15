package pet

type Persistence interface {
	FindById(id int64) *Pet
	FindByStatus(Status) []*Pet
	Update(pet *Pet) error
	Insert(pet *Pet) (*Pet, error)
	Delete(id int64) error
}
