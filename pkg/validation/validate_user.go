package validation

import (
	"errors"

	"github.com/Marlliton/go-crud/pkg/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/goccy/go-json"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() { // toda função "init" é chamada automaticamente quando o pacote for carregado
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateuserErr(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestErr("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			causes := rest_err.Causes{
				Message: e.Translate(transl),
				Feild:   e.Field(),
			}

			errorsCauses = append(errorsCauses, causes)
		}

		return rest_err.NewBadRequestValidationErr("Some fields are invalid", errorsCauses)
	}

	return rest_err.NewBadRequestErr("Error trying to convert fields")
}
