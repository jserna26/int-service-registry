package repo

import (
	"errors"
	"strconv"

	"github.com/jserna26/SystemPOC/types"
	_ "github.com/lib/pq"
)

type InMemorySystemRepo struct {
	data   []types.SystemType
	nextID int
}

func (r *InMemorySystemRepo) CreateSystem(sys types.NewSystemType) (system types.SystemType, err error) {
	if len(sys.Name) > 50 {
		err = errors.New("System name can be max 50 chars")
		return
	}
	system.Description = sys.Description
	r.nextID += 1
	system.ID = strconv.Itoa(r.nextID)
	system.Name = sys.Name
	system.Status = sys.Status
	r.data = append(r.data, system)
	return
}

func (r *InMemorySystemRepo) Connect(dburl string) (err error) {
	return
}

func (r *InMemorySystemRepo) MustConnect(dburl string) {
	err := r.Connect(dburl)
	if err != nil {
		panic(err)
	}
}

func (r *InMemorySystemRepo) Disconnect() (err error) {
	return
}

func NewInMemorySystemRepo(dburl string) (r *InMemorySystemRepo, err error) {
	//Create new InMemorySystemRepo object, connect to DB and return.
	r = new(InMemorySystemRepo)
	return
}
