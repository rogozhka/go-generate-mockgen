package example

import (
    "errors"
    "fmt"
)

type myObject struct {
    idGenerator uniqueIdentificatorGenerator
}

func New(idGenerator uniqueIdentificatorGenerator) *myObject {
    return &myObject{
        idGenerator: idGenerator,
    }
}

// CreateItem creates new item and returns its id.
func (m *myObject) CreateItem() (string, error) {
    if m == nil {
        return "", errors.New("receiver is nil")
    }
    id, err := m.idGenerator.GenerateID()
    if err != nil {
        return "", fmt.Errorf("generate ID: %w", err)
    }

    // TODO: create item in the repository

    return id, nil
}
