package constanta

const (
	TokenSecretKey        = "Ant@r_bAR@n6$"
	VERSION               = "B060522"
	TokenExpiredInMinutes = 8 * 60 * 60
)

// ERR code Global
const (
	ERR_CODE_00     = "00"
	ERR_CODE_00_MSG = "SUCCESS.."

	ERR_CODE_03     = "03"
	ERR_CODE_03_MSG = "Error, unmarshall body Request"

	ERR_CODE_04                    = "04"
	ERR_CODE_04_PARAM_QUERY_STRING = "Error, parameter query string"
)

const (
	ERR_CODE_10                  = "10"
	ERR_CODE_10_FAILED_INSERT_DB = "Failed insert data to database"

	ERR_CODE_11                 = "11"
	ERR_CODE_11_FAILED_GET_DATA = "Failed get data to database"

	ERR_CODE_12                    = "12"
	ERR_CODE_12_FAILED_UPDATE_DATA = "Failed update data to database"
)

const (
	ERR_CODE_20              = "20"
	ERR_CODE_20_LOGIN_FAILED = "Login Failed"

	ERR_CODE_21               = "21"
	ERR_CODE_21__INVALID_AUTH = "Invalid Authorization !"

	ERR_CODE_22                = "22"
	ERR_CODE_22__TOKEN_EXPIRED = "Token Expired !"
)

const (
	PREFIX_SELLER = "S"
	PREFIX_DRIVER = "D"
	PREFIX_ADMIN  = "A"
)
