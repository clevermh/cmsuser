package bd

import (
	"fmt"

	"github.com/clevermh/cmsuser/models"
	"github.com/clevermh/cmsuser/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SingUp) error {
	fmt.Println("Comienza registro")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()
	senencia := "INSERT INTO users( User_UUID, User_Email , User_DateAdd) values ('" + sig.UserEmail + "','" + sig.UserEmail + "','" + tools.FechaMysql() + "')"
	fmt.Println(senencia)

	_, err = Db.Exec(senencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("SignUp > Ehjecucion Exitosa")
	return nil
}
