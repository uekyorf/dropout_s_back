package config

// BAConfig BasicAuthの設定情報
type BAConfig struct {
	User string
	Pass string
}

// GetBAConfig データベースへの接続情報を返す
func GetBAConfig() BAConfig {
	config := BAConfig{}
	config.User = "user"
	config.Pass = "password"
	return config
}
