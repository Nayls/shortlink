package link

import (
	"fmt"
)

type NotFoundError struct { // nolint unused
	Link Link
	Err  error
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not found link: %s", e.Link.URL)
}
