package blogerrors

import "errors"

var ErrIDNotFound = errors.New("id not found")
var ErrNoRecord = errors.New("No Records found")
