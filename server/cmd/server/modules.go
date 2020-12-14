//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/o1111001/virtual-disk-management/server/disks"
)

func ComposeApiServer(port HttpPortNumber) (*ApiServer, error) {
	wire.Build(
		NewDbConnection,
		disks.Providers,
		wire.Struct(new(ApiServer), "Port", "DisksHandler"),
	)
	return nil, nil
}
