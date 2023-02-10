package schemas

import "github.com/go-playground/validator/v10"

type GetUserByEmail struct {
	Email string `json:"email" validate:"required,email"`
}
type ReplaceApiKey struct {
	Email  string `json:"email" validate:"required,email"`
	ApiKey string `json:"api_key" validate:"required"`
}

func (g *GetUserByEmail) Validate() []*ErrorView {
	var errors []*ErrorView
	validate := validator.New()
	if err := validate.Struct(g); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorView{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}
func (r *ReplaceApiKey) Validate() []*ErrorView {
	var errors []*ErrorView
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrorView{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}
	return errors
}
