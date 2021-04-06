package model

import (
	"gorm.io/gorm"
)

// ExecutionHistory ExecutionHistory
// 执行流历史纪录
type ExecutionHistory struct {
	Execution
}

// CopyExecutionToHistoryByProcInstIDTx CopyExecutionToHistoryByProcInstIDTx
func CopyExecutionToHistoryByProcInstIDTx(procInstID int, tx *gorm.DB) error {
	return tx.Exec("insert into execution_histories select * from executions where proc_inst_id=?", procInstID).Error
}
