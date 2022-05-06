package services

import (
	"fmt"
	"time"

	"com.ddabadi.antarbarang/constanta"
	"com.ddabadi.antarbarang/dto"
	"com.ddabadi.antarbarang/enumerate"
	"com.ddabadi.antarbarang/model"
	"com.ddabadi.antarbarang/repository"

	jwt "github.com/dgrijalva/jwt-go"
)

type AdminService struct {
}

func (a *AdminService) CreateAdmin(admin model.Admin) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	lastInsertId, err := repository.SaveAdmin(admin)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = lastInsertId
	return result
}

func (a *AdminService) GetAdminByID(id int) dto.ContentResponse {
	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	admin, err := repository.FindAdminById(id)
	if err != nil {
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		result.Contents = err.Error()
		return result
	}
	result.Contents = admin
	return result
}

func (a *AdminService) GetAdminByKode(kode string) dto.ContentResponse {
	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	seller, err := repository.FindAdminByCode(kode)
	if err != nil {
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		result.Contents = err.Error()
		return result
	}
	result.Contents = seller
	return result
}

func (a *AdminService) UpdateAdmin(admin model.Admin) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	msg, err := repository.UpdateAdmin(admin)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = msg
	return result
}

func (a *AdminService) LoginAdminByNama(nama, password string) dto.LoginResponseDto {
	var result dto.LoginResponseDto
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	seller, err := repository.LoginAdminByNama(nama)
	fmt.Println("Error login ==>", err)
	if err != nil {
		result.ErrCode = constanta.ERR_CODE_11
		result.ErrDesc = constanta.ERR_CODE_11_FAILED_GET_DATA
		result.Token = ""
		return result
	}

	if seller.Status != enumerate.ACTIVE {
		result.ErrCode = constanta.ERR_CODE_20
		result.ErrDesc = constanta.ERR_CODE_20_LOGIN_FAILED
		result.Token = ""
		return result
	}

	if seller.Password != password {
		result.ErrCode = constanta.ERR_CODE_20
		result.ErrDesc = constanta.ERR_CODE_20_LOGIN_FAILED
		result.Token = ""
		return result
	}

	result.Token = generateToken(seller.Nama, seller.ID)
	return result
}

func generateToken(userName string, userID int64) string {
	token := ""

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)
	claims["user"] = userName
	claims["userId"] = fmt.Sprintf("%v", (userID))
	// now := time.Now()
	// claims["logTm"] = now
	// claims["supplierCode"] = user.SupplierCode

	unixNano := time.Now().UnixNano()
	umillisec := unixNano / 1000000
	timeToString := fmt.Sprintf("%v", umillisec)
	fmt.Println("token Created ", timeToString)
	claims["tokenCreated"] = timeToString

	token, err := sign.SignedString([]byte(constanta.TokenSecretKey))

	if err != nil {
		return ""
	}
	return token
}

func (a *AdminService) ChangePasswordAdmin(admin model.Admin) dto.ContentResponse {

	var result dto.ContentResponse
	result.ErrCode = constanta.ERR_CODE_00
	result.ErrDesc = constanta.ERR_CODE_00_MSG

	msg, err := repository.ChangePasswordAdmin(admin)

	if err != nil {
		result.Contents = err.Error()
		result.ErrCode = constanta.ERR_CODE_10
		result.ErrDesc = constanta.ERR_CODE_10_FAILED_INSERT_DB
		return result
	}
	result.Contents = msg
	return result
}
