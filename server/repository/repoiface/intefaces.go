package repoiface

import "io"

type DBReadWriter interface {
	Ping() error
	io.Closer
}
