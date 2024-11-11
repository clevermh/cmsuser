package bd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/clevermh/cmsuser/models"
	"github.com/clevermh/cmsuser/secretm"
)

var SecretModel models.SecretRDsJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}
func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return err
}

func ConnStr(claves models.SecretRDsJson) string {
	var dbUser, authToken, dbEndPoint, dbName string

	dbUser = claves.Username
	authToken = claves.Password
	dbEndPoint = claves.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowClearTextPasswords=true", dbUser, authToken, dbEndPoint, dbName)
	fmt.Printf(dsn)
	return dsn
}
