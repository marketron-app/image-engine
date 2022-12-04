package request_body

type TemplateCoordinate struct {
	X int `json:"x" form:"x"`
	Y int `json:"y" form:"y"`
}

type GetImageBody struct {
	URL            string               `json:"url" form:"url"`
	TemplateImage  string               `json:"templateImage" form:"templateImage"`
	Coordinates    []TemplateCoordinate `json:"coordinates"`
	ViewportWidth  int64                `json:"viewportWidth"`
	ViewportHeight int64                `json:"viewportHeight"`
}
