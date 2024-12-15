package example

import (
	"time"
)

//go:generate docker run --rm -v ${PWD}:/w rogozhka/go-generate-mockgen -line=$GOLINE -source=$GOFILE -package=$GOPACKAGE
type timeOperations interface {
	Now() time.Time
	Since(time.Time) time.Duration
}
