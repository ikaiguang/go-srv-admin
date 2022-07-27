package main

import (
	dbupgrade "github.com/ikaiguang/go-srv-admin/cmd/migration/migrate"
)

// go run ./cmd/migration/... -conf=./configs
func main() {
	dbupgrade.Run(dbupgrade.WithCloseEngineHandler())
}
