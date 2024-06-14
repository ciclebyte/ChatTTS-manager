// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"github.com/tiger1103/gfast/v3/internal/app/cts/dao/internal"
)

// internalCtsVoiceDao is internal type for wrapping internal DAO implements.
type internalCtsVoiceDao = *internal.CtsVoiceDao

// ctsVoiceDao is the data access object for table cts_voice.
// You can define custom methods on it to extend its functionality as you wish.
type ctsVoiceDao struct {
	internalCtsVoiceDao
}

var (
	// CtsVoice is globally public accessible object for table cts_voice operations.
	CtsVoice = ctsVoiceDao{
		internal.NewCtsVoiceDao(),
	}
)

// Fill with you ideas below.