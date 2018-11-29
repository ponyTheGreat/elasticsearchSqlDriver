package elasticsearchSqlDriver

import "errors"

type EsTx struct{}

func (tx *EsTx) Commit() error {
	return errors.New("Not implemented yet")
}

func (tx *EsTx) Rollback() error {
	return errors.New("Not implemented yet")
}
