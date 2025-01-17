package model

import (
	"gorm.io/gorm"
)

// IdentitylinkHistory IdentitylinkHistory
type IdentitylinkHistory struct {
	Identitylink
}

// CopyIdentitylinkToHistoryByProcInstID CopyIdentitylinkToHistoryByProcInstID
func CopyIdentitylinkToHistoryByProcInstID(procInstID int, tx *gorm.DB) error {
	return tx.Exec("insert into identitylink_histories select * from identitylinks where proc_inst_id=?", procInstID).Error
}

// FindParticipantHistoryByProcInstID FindParticipantHistoryByProcInstID
func FindParticipantHistoryByProcInstID(procInstID int) ([]*IdentitylinkHistory, error) {
	var datas []*IdentitylinkHistory
	err := db.Select("id,user_id,step,comment").Where("proc_inst_id=? and type=?", procInstID, IdentityTypes[PARTICIPANT]).Order("id asc").Find(&datas).Error
	return datas, err
}
