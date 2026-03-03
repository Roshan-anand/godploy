package lib

type Res struct {
	Message string                 `json:"message" validate:"required"`
	Data    map[string]interface{} `json:"data"`
}
