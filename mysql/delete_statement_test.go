package mysql

import (
	"testing"
)

func TestDeleteUnconditionally(t *testing.T) {
	assertStatementErr(t, table1.DELETE(), `jet: WHERE clause not set`)
	assertStatementErr(t, table1.DELETE().WHERE(nil), `jet: WHERE clause not set`)
}

func TestDeleteWithWhere(t *testing.T) {
	assertStatement(t, table1.DELETE().WHERE(table1Col1.EQ(Int(1))), `
DELETE FROM db.table1
WHERE table1.col1 = ?;
`, int64(1))
}

func TestDeleteWithWhereOrderByLimit(t *testing.T) {
	assertStatement(t, table1.DELETE().WHERE(table1Col1.EQ(Int(1))).ORDER_BY(table1Col1).LIMIT(1), `
DELETE FROM db.table1
WHERE table1.col1 = ?
ORDER BY table1.col1
LIMIT ?;
`, int64(1), int64(1))
}
