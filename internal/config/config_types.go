package config

const ConfigFileName = ".gatorconfig.json"
const ConfigFilePath = "" ///workspace/Go/gator/

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}
