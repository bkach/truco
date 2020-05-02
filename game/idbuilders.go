package game

import (
	"github.com/gofrs/uuid"
)

type Id interface {
	getIdString() string
	build(input string) error
}

type ProdId struct {
	Id		string
}

type DebugId struct {
	Id		string
}

func (id *ProdId) build(input string) error{
	uUID, err := uuid.NewV4()

	if err != nil {
		return err
	}

	*id = ProdId { Id: input + uUID.String() }

	return nil
}

func (id *ProdId) getIdString() string {
	return id.Id
}

func (id *DebugId) build(input string) error {
	*id = DebugId { Id: input }

	return nil
}

func (id DebugId) getIdString() string {
	return id.Id
}
