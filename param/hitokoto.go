package param

type HitokotoParam struct {
	C         string `json:"c" default:""`
	Encode    string `json:"encode"`
	Charset   string `json:"charset"`
	MinLength int    `json:"min_length"`
	MaxLength int    `json:"max_length"`
}
