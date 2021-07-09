package dao

import (
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

func is_titleLen(fl validator.FieldLevel) bool {
	title := fl.Field().String()
	if len(title) > 0 || len(title) <= 255 {
		return true
	}
	return false
}

func is_onlySpace(fl validator.FieldLevel) bool {
	title := fl.Field().String()
	title = strings.Join(strings.Fields(title), "")
	if len(title) != 0 {
		return true
	}
	return false
}

func is_contentLen(fl validator.FieldLevel) bool {
	content := fl.Field().String()
	if len(content) > 0 || len(content) <= 65535 {
		return true
	}
	return false
}

func ArticleValidation(article *Article) []string {
	var errorMessages []string

	validate := validator.New()
	validate.RegisterValidation("is_titleLen", is_titleLen)
	validate.RegisterValidation("is_contentLen", is_contentLen)
	validate.RegisterValidation("is_onlySpace", is_onlySpace)

	err := validate.Struct(article)

	if err != nil {
		for _,err := range err.(validator.ValidationErrors) {
			var errorMessage string
			fieldName := err.Field()

			switch (fieldName) {
			case "Title":
				errorMessage = "error message for Title"
			case "Content":
				errorMessage = "error message for Content"
			default:
				errorMessage = "error message"
			}
			errorMessages = append(errorMessages, errorMessage)
		}
	}

	return errorMessages
}
