package handlers

import (
	"errors"
	"regexp"

	"github.com/laurence6/telegram-bot-api"
)

var NotFoundHandleFunc = errors.New("Cannot find a HandleFunc to handle this message")

/*TextHandler Find a HandleFunc and process the message
 */
type TextHandler struct {
	Patterns    []*regexp.Regexp
	HandleFuncs []HandleFunc
}

/*Add Add the pattern and HandleFunc
*
* The pattern will be compiled.
 */
func (handler *TextHandler) Add(pattern string, f HandleFunc) error {
	compiledPattern, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	handler.Patterns = append(handler.Patterns, compiledPattern)
	handler.HandleFuncs = append(handler.HandleFuncs, f)

	return nil
}

/*Handle Find a HandleFunc, call it and return the result
 */
func (handler *TextHandler) Handle(message *tgbotapi.Message) (interface{}, error) {
	text := message.Text
	for n, c := range handler.Patterns {
		if c.MatchString(text) {
			return handler.HandleFuncs[n](message)
		}
	}
	return nil, NotFoundHandleFunc
}
