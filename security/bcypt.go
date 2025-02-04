package security

import "golang.org/x/crypto/bcrypt"




func GenerateFromPassword(password string) (string,error){
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(result),err
}


func ComparePassword(hashedPassword string,password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

