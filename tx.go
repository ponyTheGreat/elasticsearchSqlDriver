package ge

import "errors"

//EsTx not implement yet
type EsTx struct{}

//Commit not implemented yet
func (tx *EsTx) Commit() error {
	return errors.New("Not implemented yet")
}

//Rollback not implemented yet
func (tx *EsTx) Rollback() error {
	return errors.New("Not implemented yet")
}
