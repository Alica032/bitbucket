package telegram

type ServerFlags struct {
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
	Ngrok string `yaml:"test_host"` // for tests
}

type TelegramFlags struct {
	Name  string `yaml:"name"`
	Token string `yaml:"token"`
}

type Config struct {
	Bot          TelegramFlags `yaml:"bot"`
	Server       ServerFlags   `yaml:"server"`
	IsFirstStart bool          `yaml:"first_run"`
}
