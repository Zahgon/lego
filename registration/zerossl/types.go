package zerossl

type APIResponse struct {
	Success bool `json:"success"`

	Kid     string `json:"eab_kid"`
	HmacKey string `json:"eab_hmac_key"`

	Error *ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code int    `json:"code"`
	Type string `json:"type"`
	Info string `json:"info"`
}

func (e *ErrorDetail) Error() string { _ = "STUB: not implemented"; return "" }
