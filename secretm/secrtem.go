package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/clevermh/cmsuser/awsgo"
	"github.com/clevermh/cmsuser/models"
)

func GetSecret(nombreSecret string) (models.SecretRDsJson, error) {
	var datosSecret models.SecretRDsJson
	fmt.Println("> Pido Secreto " + nombreSecret)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Printf(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("> Lectura Secret Ok" + nombreSecret)
	return datosSecret, nil

}

// package secretm

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/aws/aws-sdk-go-v2/aws"
// 	"github.com/aws/aws-sdk-go-v2/awsgo"
// 	"github.com/aws/aws-sdk-go-v2/models"
// 	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
// 	"github.com/clevermh/cmsuser/awsgo"
// 	"github.com/clevermh/cmsuser/models"
// )
// func GetSecret(nombreSecret string )(models.SecretRDsJson, error) {
// 	var datosSecret models.SecretRDsJson
// 	fmt.Println("> Pido Secreto "+ nombreSecret)
// 	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
// 	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
// 		SecretId: aws.String(nombreSecret),
// 	})
// 	if err != nil{
// 		  fmt.Printf(err.Error())
// 		  return datosSecret, err
// 	}

// 	json.Unmarshal([]byte(*clave.SecretString),&datosSecret)
// 	fmt.Println("> Lectura Secret Ok"+nombreSecret)
// 	return datosSecret,nil

// }
