// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"server/internal/dao/internal"
)

// internalTipOffNoticeDataDao is an internal type for wrapping the internal DAO implementation.
type internalTipOffNoticeDataDao = *internal.TipOffNoticeDataDao

// tipOffNoticeDataDao is the data access object for the table tip_off_notice_data.
// You can define custom methods on it to extend its functionality as needed.
type tipOffNoticeDataDao struct {
	internalTipOffNoticeDataDao
}

var (
	// TipOffNoticeData is a globally accessible object for table tip_off_notice_data operations.
	TipOffNoticeData = tipOffNoticeDataDao{
		internal.NewTipOffNoticeDataDao(),
	}
)

// Add your custom methods and functionality below.
