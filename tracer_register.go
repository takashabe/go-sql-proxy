// +build go1.5

package proxy

import (
	"database/sql"
	"strings"
)

// RegisterTracer creates proxies that log queries from the sql drivers already registered,
// and registers the proxies as sql driver.
// The proxies' names have suffix ":trace".
func RegisterTracer() {
	for _, driver := range sql.Drivers() {
		if strings.HasSuffix(driver, ":trace") || strings.HasSuffix(driver, ":proxy") {
			continue
		}
		db, _ := sql.Open(driver, "")
		defer db.Close()
		sql.Register(driver+":trace", NewTraceProxy(db.Driver(), logger{}))
	}
}
