package exception

import (
	"fmt"
	"github.com/pkg/errors"
)

var ErrGeneral = errors.New("general error")
var ErrDomain = fmt.Errorf("domain error: %w", ErrGeneral)
var ErrInvalidFileLink = fmt.Errorf("invalid file link: %w", ErrDomain)
var ErrInvalidArchivePath = fmt.Errorf("invalid archive path: %w", ErrDomain)
var ErrInvalidFileError = fmt.Errorf("invalid file error: %w", ErrDomain)
