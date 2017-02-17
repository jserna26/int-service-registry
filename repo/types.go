package repo

import "github.com/jserna26/SystemPOC/types"

type SystemRepo interface {
	CreateSystem(sys types.NewSystemType) (system types.SystemType, err error)
	GetAllSystems() (system []types.SystemType, err error)
	GetSystem(systemName string) (system types.SystemType, err error)
	Connect(dburl string) (err error)
	MustConnect(dburl string)
	Disconnect() (err error)
}
