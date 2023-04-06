// Package placeholder ...
package placeholder

import (
	"fmt"

	"github.com/koha90/platform/pkg/logging"
)

var names = []string{"Alice", "Bob", "Charlie", "Dora"}

// NameHandler ...
type NameHandler struct {
	logging.Logger
}

// GetName ...
func (n NameHandler) GetName(i int) string {
	n.Logger.Debugf("GetName method invoked with argument: %v", i)
	if i < len(names) {
		return fmt.Sprintf("Name #%v: %v", i, names[i])
	}
	return fmt.Sprintln("Index out of bounds")
}

// GetNames ...
func (n NameHandler) GetNames() string {
	n.Logger.Debug("GetNames method invoked")
	return fmt.Sprintf("Names: %v", names)
}

// NewName ...
type NewName struct {
	Name          string
	InsertAtStart bool
}

// PostName ...
func (n NameHandler) PostName(new NewName) string {
	n.Logger.Debugf("PostName method invoked with argument %v", new)
	if new.InsertAtStart {
		names = append([]string{new.Name}, names...)
	} else {
		names = append(names, new.Name)
	}
	return fmt.Sprintf("Names: %v", names)
}
