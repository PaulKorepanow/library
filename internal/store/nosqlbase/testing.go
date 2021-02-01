package nosqlbase

import "testing"

func TestDataBase(t *testing.T) *NoSqlStore {
	t.Helper()

	return NewNoSqlStore()
}
