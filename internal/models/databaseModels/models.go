package databasemodels

type DatabaseModel struct {
	User     string `mapstructure:"user"`
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	DbName   string `mapstructure:"dbname"`
}
