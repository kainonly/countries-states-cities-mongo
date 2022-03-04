package index

import (
	"fmt"
	"github.com/google/wire"
)

var Provides = wire.NewSet(
	wire.Struct(new(Controller), "*"),
	wire.Struct(new(Service), "*"),
)

func baseURL(path string) string {
	return fmt.Sprintf(
		`https://raw.githubusercontent.com/dr5hn/countries-states-cities-database/master/%s`,
		path,
	)
}
