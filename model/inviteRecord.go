package model

import "wTools/global"

type InviteRecord struct {
	BaseModel
	UserId        int64 `json:"user_id" gorm:"column:user_id;default 0;comment:邀请人"`
	InviteeUserId int64 `json:"invitee_user_id" gorm:"column:invitee_user_id;default 0;comment:被邀请人"`
}

func (ir InviteRecord) TableName() string {
	return "invite_record"
}

// GetListByUserId
//  @Description: 获取邀请人下所有名单
//  @receiver ir
//  @param userId
//  @return []int64
//  @return error
//  @author	wujingfeng 2022-07-05 11:21:40
func (ir InviteRecord) GetListByUserId(userId int64) ([]int64, error) {
	var InviteeUserIds []int64
	err := global.Mysql.Table(ir.TableName()).Where("user_id = ?", userId).Pluck("invitee_user_id", &InviteeUserIds).Error
	return InviteeUserIds, err
}

// GetAllChildByUserId
//  @Description: 循环获取指定用户名下所有用户
//  @receiver ir
//  @param userId
//  @param allUserIds
//  @return []int64
//  @return error
//  @author	wujingfeng 2022-07-05 11:32:19
func (ir InviteRecord) GetAllChildByUserId(userId int64, allUserIds []int64) ([]int64, error) {
	allUserIds = append(allUserIds, userId)
	childUserList, err := ir.GetListByUserId(userId)
	if err != nil {
		return allUserIds, err
	}
	if len(childUserList) != 0 {
		for _, childId := range childUserList {
			// 循环获取子类
			allUserIds, err = ir.GetAllChildByUserId(childId, allUserIds)
		}
	}

	return allUserIds, err
}
