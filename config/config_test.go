package config

import (
	"testing"
)

func TestGet(t *testing.T) {
	t.Log(Get("mysqlDataSourceName"))
}
