package placeholder

import (
	"fmt"

	"platform/http/actionresults"
	"platform/http/handling"
	"platform/logging"
)

var names = []string{"Alice", "Bob", "Charlie", "Dora"}

type NameHandler struct {
	logging.Logger
	handling.URLGenerator
}

func (n NameHandler) GetName(i int) actionresults.ActionResult {
	n.Logger.Debugf("GetName method invoked with arguments: %v", i)
	var response string
	if i < len(names) {
		response = fmt.Sprintf("Name #%v: %v", i, names[i])
	} else {
		response = fmt.Sprintf("Index out if bounds")
	}
	return actionresults.NewTemplateAction("simple_message.html", response)
}

func (n NameHandler) GetNames() actionresults.ActionResult {
	n.Logger.Debugf("GetNames method invoked")
	return actionresults.NewTemplateAction("simple_message.html", names)
}

type NewName struct {
	Name          string
	InsertAtStart bool
}

func (n NameHandler) PostName(new NewName) actionresults.ActionResult {
	n.Logger.Debugf("PostName method invoked with argument %v", new)
	if new.InsertAtStart {
		names = append([]string{new.Name}, names...)
	} else {
		names = append(names, new.Name)
	}
	return n.redirectOrError(NameHandler.GetNames)
}

func (n NameHandler) GetRedirect() actionresults.ActionResult {
	return n.redirectOrError(NameHandler.GetNames)
}

func (n NameHandler) GetJsonData() actionresults.ActionResult {
	return actionresults.NewJsonAction(names)
}

func (n NameHandler) redirectOrError(handler interface{},
	data ...interface{}) actionresults.ActionResult {
	url, err := n.GenerateUrl(handler)
	if err == nil {
		return actionresults.NewRedirectAction(url)
	} else {
		return actionresults.NewErrorAction(err)
	}
}
