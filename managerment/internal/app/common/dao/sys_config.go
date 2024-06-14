// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/common/dao/internal"
)

// sysConfigDao is the data access object for table sys_config.
// You can define custom methods on it to extend its functionality as you wish.
type sysConfigDao struct {
	*internal.SysConfigDao
}

var (
	// SysConfig is globally public accessible object for table sys_config operations.
	SysConfig = sysConfigDao{
		internal.NewSysConfigDao(),
	}
)

// Fill with you ideas below.
