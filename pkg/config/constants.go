package config

import (
	"github.com/D4lvarez/data_process/pkg/utils"
)

var db_user string = utils.GetDotEnvVar("DB_USER")
var db_pwd string = utils.GetDotEnvVar("DB_PWD")
var db_host string = utils.GetDotEnvVar("DB_HOST")
var db_port string = utils.GetDotEnvVar("DB_PORT")
var db_name string = utils.GetDotEnvVar("DB_NAME")
