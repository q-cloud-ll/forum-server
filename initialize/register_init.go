package initialize

import (
	_ "forum-server/source/example"
	_ "forum-server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
