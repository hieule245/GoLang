package models

import "github.com/golang-web/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	// thông báo thành công, thất bại,....
	Flash string
	// cảnh báo lỗi
	Warning string
	Error   string
	Form    *forms.Form
}
