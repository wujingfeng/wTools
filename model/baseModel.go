package model

type BaseModel struct {
	Id          uint  `json:"id" gorm:"column:id;type:uint;not null;primaryKey;autoIncrement;comment:自增 id"`
	CreatedTime int64 `json:"created_time" gorm:"column:created_time;autoCreateTime"`
}
