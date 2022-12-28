package request_body

type TemplateCoordinate struct {
	X int `json:"x" form:"x" validate:"required"`
	Y int `json:"y" form:"y" validate:"required"`
}

type GetImageBody struct {
	URL            string               `json:"url" form:"url" validate:"required,url"`
	TemplateImage  string               `json:"templateImage" form:"templateImage" validate:"required,url"`
	Coordinates    []TemplateCoordinate `json:"coordinates" validate:"required,len=4"`
	ViewportWidth  int64                `json:"viewportWidth" validate:"required"`
	ViewportHeight int64                `json:"viewportHeight" validate:"required"`
	FileName       string               `json:"fileName"`
}
