package models

type SecretRDsJson struct {
	Username              string `json:"username"`
	Password              string `json:"Password"`
	Engine                string `json:"Engine"`
	Host                  string `json:"Host"`
	Port                  string `json:"Port"`
	DbClousterIndentifier string `json:"DbClousterIndentifier"`
}

type SingUp struct {
	UserEmail string `json:"UserEmail"`
	UserUUID  string `json:"UserUUID"`
}
