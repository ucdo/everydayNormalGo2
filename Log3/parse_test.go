package main

import "testing"

func TestParseConf(t *testing.T) {
	t.Helper()
	parseConf("./main.cnf")
}
