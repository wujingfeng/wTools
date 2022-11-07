package model

import "wTools/global"

type ConsumeDetail struct {
	BaseModel
	UserId int64 `json:"user_id" gorm:"column:user_id;default:0;comment:用户 id"`
	Money  int64 `json:"money" gorm:"column:money; default:0;comment:消费金额"`
}

func (cd ConsumeDetail) TableName() string {
	return "consume_detail"
}

// CountByUserIds
//  @Description: 统计指定用户群的指定时间内的消费总额
//  @receiver cd
//  @param userIds
//  @param startTime
//  @param endTime
//  @return int64
//  @return error
//  @author	wujingfeng 2022-07-05 11:48:56
func (cd ConsumeDetail) CountByUserIds(userIds []int64, startTime int64, endTime int64) (int64, error) {
	var money []int64
	global.Mysql.Debug().Table(cd.TableName()).Where("user_id in (?) and created_time between (?) and (?)", userIds, startTime, endTime).Pluck("money", &money)
	if len(money) == 0 {
		return 0, nil
	}
	total := int64(0)
	for _, v := range money {
		total += v
	}
	return total, nil
}
