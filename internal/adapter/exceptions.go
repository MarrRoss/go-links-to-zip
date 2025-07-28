package adapter

import (
	"fmt"
	"github.com/MarrRoss/go-links-to-zip/internal/domain/exception"
)

var ErrStorage = fmt.Errorf("storage error: %w", exception.ErrGeneral)
var ErrTaskLimit = fmt.Errorf("active task limit exceeded: %w", ErrStorage)
