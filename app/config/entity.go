package config

type App struct {
	Name        string `env:"APP_NAME"`
	Port        string `env:"APP_PORT"`
	Mode        string `env:"APP_MODE"`
	Url         string `env:"APP_URL"`
	Secret_key  string `env:"APP_SECRET"`
	Antares_url string `env:"APPLICATION_ANTARES_URL"`
}
type Db struct {
	Host string `env:"DB_HOST"`
	Name string `env:"DB_NAME"`
	User string `env:"DB_USER"`
	Pass string `env:"DB_PASSWORD"`
	Port string `env:"DB_PORT"`
}
type BasicAuth struct {
	Username string `env:"BASIC_AUTH_USER"`
	Password string `env:"BASIC_AUTH_PASSWORD"`
}
type Conf struct {
	App       App
	Db        Db
	BasicAuth BasicAuth
}
