package common

//Token 结构体
type Token struct {
	Access_token  string `json:"access_token"`
	Expires_in    int64  `json:"expires_in"`
	Refresh_token string `json:"refresh_token"`
}
type RequestErr struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//LongitudeAndLatitude 经伟度结构体
type LongitudeAndLatitude struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}
