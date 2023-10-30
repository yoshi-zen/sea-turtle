package myerrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	DBInsertFailed ErrCode = "D000"
	DBSelectFailed ErrCode = "D001"

	RowScanFailed ErrCode = "S000"
	NoData        ErrCode = "S002"

	ReqDecodeFailed ErrCode = "R000"
	BadParameter    ErrCode = "R001"
)
