package vo

// EmployeeLoginVO 表示员工登录返回的数据格式
type EmployeeLoginVO struct {
    ID       int64  `json:"id"`        // 主键值
    UserName string `json:"userName"`  // 用户名
    Name     string `json:"name"`      // 姓名
    Token    string `json:"token"`     // jwt 令牌
}

// R«EmployeeLoginVO»
type Response struct {
	Code                          int64   `json:"code"`
	// EmployeeLoginVO，员工登录返回的数据格式        
	Data                          *EmployeeLoginVO   `json:"data,omitempty"`
	Msg                           *string `json:"msg,omitempty"`
}