package model

type UserModel struct {
	Id          int64  `json:"id" gorm:"column:id;default0';comment:用户id"`
	Code        string `json:"code" gorm:"column:code;default:'';comment:用户内码"`
	Openid      string `json:"openid" gorm:"column:openid;default:'';comment:微信唯一标识 openid"`
	Name        string `json:"name" gorm:"column:name;default:'';comment:用户名称"`
	Phone       string `json:"phone" gorm:"column:phone;default:'';comment:电话号码"`
	HeadImg     string `json:"head_img" gorm:"column:head_img;default:'';comment:头像"`
	Password    string `json:"password" gorm:"column:password;default:'';comment:密码"`
	Salt        string `json:"salt" gorm:"column:salt;default:'';comment:加密盐值"`
	Status      int8   `json:"status" gorm:"column:status;default:0;comment:用户状态 1正常"`
	CreatedTime int64  `json:"created_time" gorm:"column:created_time;default:0;comment:创建时间"`
	UpdatedTime int64  `json:"updated_time" gorm:"column:updated_time;default:0;comment:更新时间"`
	DeletedTime int64  `json:"deleted_time" gorm:"column:deleted_time;default:0;comment:删除时间"`
	Token       string `json:"token" gorm:"column:token;default:'';comment:token"`
	ExpireAt    int64  `json:"expire_at" gorm:"column:expire_at;default:0;comment:过期时间"`
}
