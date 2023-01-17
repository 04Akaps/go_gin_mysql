package err

import "errors"

var (
	ErrNotFoundConfigFile = errors.New("not found config file")
)