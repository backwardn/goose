// +build no_mysql

package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/ziutek/mymysql/godrv"
)

func normalizeDBString(driver string, str string, certfile string) string {
	return str
}
