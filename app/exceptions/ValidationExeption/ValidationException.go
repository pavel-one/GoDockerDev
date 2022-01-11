package ValidationExeption

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Exception struct {
	err validator.ValidationErrors
}

func New(err error) *Exception {
	target := validator.ValidationErrors{}
	if errors.As(err, &target) {
		return &Exception{err.(validator.ValidationErrors)}
	}

	return &Exception{target}
}

func (e Exception) FormatToFront() map[string]string {
	out := make(map[string]string)

	for _, err := range e.err {
		out[strings.ToLower(err.Field())] = getTagText(err.ActualTag(), err.Param())
	}

	return out
}

func getTagText(tagName string, param string) string {
	var texts = make(map[string]string)

	texts["min"] = fmt.Sprintf("Минимальное количество символов: %s", param)
	texts["max"] = fmt.Sprintf("Максимальное количество символов: %s", param)
	texts["required"] = "Поле обязательно для заполнения"
	texts["email"] = "Введенное значение не является адресом электронной почты"

	tagName = strings.ToLower(tagName)

	if texts[tagName] == "" {
		return tagName
	}

	return texts[tagName]
}
