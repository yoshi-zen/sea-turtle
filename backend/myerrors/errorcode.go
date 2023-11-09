package myerrors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	InsertDataFailed ErrCode = "D000"
	GetDataFailed    ErrCode = "D001"
	UpdateDataFailed ErrCode = "D002"

	NoData         ErrCode = "S000"
	EncryptFailed  ErrCode = "S001"
	GenUUIDFailed  ErrCode = "S002"
	SendMailFailed ErrCode = "S003"

	ReqDecodeFailed ErrCode = "R000"
	BadParameter    ErrCode = "R001"
	EmailInvalid    ErrCode = "R002"
)
