package model

import (
	"gorm.io/gorm"
)

// TaskHistory TaskHistory
type TaskHistory struct {
	Task
}

// CopyTaskToHistoryByProInstID CopyTaskToHistoryByProInstID
// 根据procInstID查询结果，并将结果复制到task_history表
func CopyTaskToHistoryByProInstID(procInstID int, tx *gorm.DB) error {
	return tx.Exec("insert into task_histories select * from tasks where proc_inst_id=?", procInstID).Error
}
