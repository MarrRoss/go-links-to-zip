package exception

import (
	"fmt"
	"github.com/MarrRoss/go-links-to-zip/internal/domain/exception"
)

var ErrApplication = fmt.Errorf("application error: %w", exception.ErrGeneral)
