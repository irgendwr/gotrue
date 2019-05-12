package test

import (
	"github.com/irgendwr/gotrue/conf"
	"github.com/irgendwr/gotrue/storage"
)

func SetupDBConnection(globalConfig *conf.GlobalConfiguration) (*storage.Connection, error) {
	return storage.Dial(globalConfig)
}
