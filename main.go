package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/clevermh/cmsuser/awsgo"
	"github.com/clevermh/cmsuser/bd"
	"github.com/clevermh/cmsuser/models"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicializoAWS()
	if !ValidaParametros() {
		fmt.Println("Error en los parametros.. debe enviar 'SecretManager'")
		erro := errors.New("error en los parametros debe enviar SecretName")
		return event, erro
	}

	var datos models.SingUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}
	}
	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer al secret " + err.Error())
		return event, err
	}

	err = bd.SignUp()
	return event, err
}
func ValidaParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}

// package main

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"os"

// 	"github.com/aws/aws-lambda-go/events"
// 	lambda "github.com/aws/aws-lambda-go/lambda"
// 	"github.com/clevermh/cmsuser/awsgo"
// 	"github.com/clevermh/cmsuser/bd"
// 	"github.com/clevermh/cmsuser/models"
// )

// func main() {
// 	lambda.Start(EjecutoLambda)
// }

// func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
// 	awsgo.InicializoAWS()
// 	if !ValidaParametros(){
// 		fmt.Println("Error en los parametros.. debe enviar 'SecretManager'")
// 		erro:= errors.New("error en los parametros debe enviar SecretName")
// 		return event , erro
// 	}

// 	var datos models.SingUp

// 	for  row ,att:= range event.Request.UserAttributes {
// 		switch row{
// 		case "email":
// 			datos.UserEmail=att
// 			fmt.Println("Email = " + datos.UserEmail)
// 		case "sub":
// 			datos.UserUUID=att
// 			fmt.Println("Sub = " + datos.UserUUID)
// 		}
//  	}
// 	err:= bd.ReadSecret()
// 	if err != nil {
// 		fmt.Println("Error al leer al secret "+err.Error())
// 		return event, err
// 	}
// }
// func ValidaParametros() bool {
// 	var traeParametro bool
// 	_, traeParametro = os.LookupEnv("SecretName")
// 	return traeParametro
// }
