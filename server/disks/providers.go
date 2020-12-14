package disks

import "github.com/google/wire"

var Providers = wire.NewSet(NewStore, HTTPHandler)
