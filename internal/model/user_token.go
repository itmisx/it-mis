package model

// UserToken 用户令牌表
type UserToken struct {
	ID            string `json:"id"`
	Token         string `json:"token"`          //令牌
	SecurityToken string `json:"security_token"` //安全令牌
	UserID        string `json:"user_id"`        //用户id
	CreateTime    int64  `json:"create_time"`    //创建时间
	ExpireTime    int64  `json:"expire_time"`    //过期时间
}
