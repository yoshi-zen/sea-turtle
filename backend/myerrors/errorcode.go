package myerrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	DBInsertFailed ErrCode = "D000"
	DBSelectFailed ErrCode = "D001"

	RowScanFailed ErrCode = "S000"
	BadParameter  ErrCode = "S001"
)
