// Package placeholder ...
package placeholder

import (
	"fmt"

	"github.com/koha90/platform/internal/http/actionresults"
	"github.com/koha90/platform/internal/http/handling"
	"github.com/koha90/platform/internal/validation"
	"github.com/koha90/platform/pkg/logging"
)

var names = []string{"Alice", "Bob", "Charlie", "Dora"}

// NameHandler ...
type NameHandler struct {
	logging.Logger
	handling.URLGenerator
	validation.Validator
}

// GetName ...
func (n NameHandler) GetName(i int) actionresults.ActionResult {
	n.Logger.Debugf("GetName method invoked with argument: %v", i)
	var response string
	if i < len(names) {
		response = fmt.Sprintf("Name #%v: %v", i, names[i])
	} else {
		response = fmt.Sprintln("Index out of bounds")
	}

	return actionresults.NewTemplateAction("simple_message.html", response)
}

// GetNames ...
func (n NameHandler) GetNames() actionresults.ActionResult {
	n.Logger.Debug("GetNames method invoked")
	return actionresults.NewTemplateAction("simple_message.html", names)
}

// NewName ...
type NewName struct {
	Name          string `validation:"required,min:3"`
	InsertAtStart bool
}

// GetForm ...
func (n NameHandler) GetForm() actionresults.ActionResult {
	postURL, _ := n.URLGenerator.GenerateURL(NameHandler.PostName)
	return actionresults.NewTemplateAction("name_form.html", postURL)
}

// PostName ...
func (n NameHandler) PostName(new NewName) actionresults.ActionResult {
	n.Logger.Debugf("PostName method invoked with argument %v", new)

	if ok, errs := n.Validator.Validate(&new); !ok {
		return actionresults.NewTemplateAction("validation_errors.html", errs)
	}

	if new.InsertAtStart {
		names = append([]string{new.Name}, names...)
	} else {
		names = append(names, new.Name)
	}

	return n.redirectOrError(NameHandler.GetNames)
}

// GetRedirect ...
func (n NameHandler) GetRedirect() actionresults.ActionResult {
	return n.redirectOrError(NameHandler.GetNames)
}

// GetJsonData ...
func (n NameHandler) GetJsonData() actionresults.ActionResult {
	return actionresults.NewJsonActionResult(names)
}

func (n NameHandler) redirectOrError(
	handler interface{},
	data ...interface{},
) actionresults.ActionResult {
	url, err := n.GenerateURL(handler)
	if err == nil {
		return actionresults.NewRedirectAction(url)
	} else {
		return actionresults.NewErrorAction(err)
	}
}
