package param

type ParamUrl struct {
	Url string `json:"url" form:"url" binding:"required,url"`
}
