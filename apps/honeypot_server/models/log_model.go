package models

//日志

type LogModel struct {
	Model
	Type        int8   `json:"type"` // 1 登录日志
	IP          string `json:"ip"`
	Location    string `json:"location"` //位置
	UserID      uint   `json:"user_id"`
	Username    string `json:"username"`
	Pwd         string `json:"pwd"`
	LoginStatus bool   `json:"login_status"`
	Title       string `json:"title"`
	Level       int8   `json:"level"` //级别
	Content     string `json:"content"`
	ServiceName string `json:"service_name"`
}
