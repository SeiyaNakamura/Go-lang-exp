package dao

import (
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

func is_titleLen(fl validator.FieldLevel) bool {
	title := fl.Field().String()
	if len(title) > 0 && len(title) <= 255 {
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
	if len(content) > 0 && len(content) <= 65535 {
		return true
	}
	return false
}

func ArticleValidation(article *Article) map[string]string {
	errorMessages := map[string]string{}

	validate := validator.New()
	validate.RegisterValidation("is_titleLen", is_titleLen)
	validate.RegisterValidation("is_contentLen", is_contentLen)
	validate.RegisterValidation("is_onlySpace", is_onlySpace)

	err := validate.Struct(article)

	if err != nil {
		for _,err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()

			switch (fieldName) {
			case "Title":
				errorMessages["titleError"] = "タイトルは1文字以上255文字以下で入力してください。"
			case "Content":
				errorMessages["contentError"] = "コンテンツは1文字以上65535文字以下で入力してください。"
			default:
				errorMessages["unexpectedError"] = "unexpected error message"
			}
		}
	}

	return errorMessages
}
