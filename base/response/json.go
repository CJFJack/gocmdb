package response

type JSONResponse struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result"`
}

func NewJsonResponse(code int, msg string, result interface{}) *JSONResponse {
	return &JSONResponse{code, msg, result}
}

var (
	UnAuthorization = NewJsonResponse(401, "unauthorization", nil)
	Ok              = NewJsonResponse(200, "ok", nil)
	BadRequest      = NewJsonResponse(400, "bad request", nil)
)
