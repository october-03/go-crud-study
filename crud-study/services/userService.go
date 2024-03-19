package services

import (
	"crud-study/crud-study/dto"
	"crud-study/crud-study/models"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(req dto.CreateUserDto) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(dto.NewError("비밀번호 암호화에 실패했습니다.", "1203"))
	}

	validateUser(req.Email)

	newUser := models.NewUser(
		req.Username,
		req.Email,
		string(hashedPassword),
	)

	jsonData, err := json.Marshal(newUser)

	fmt.Println(string(jsonData))

	if err != nil {
		panic(dto.NewError("유저 생성에 실패했습니다.", "1205"))
	}

	return jsonData
}

func validateUser(email string) {
	// 유저 중복 검사
	//panic(dto.NewError("이미 존재하는 유저입니다.", "1204"))
}
