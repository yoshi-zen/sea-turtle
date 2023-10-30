package myerrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	InsertDataFailed ErrCode = "D000"
	GetDataFailed    ErrCode = "D001"

	NoData ErrCode = "S000"

	ReqDecodeFailed ErrCode = "R000"
	BadParameter    ErrCode = "R001"
)
