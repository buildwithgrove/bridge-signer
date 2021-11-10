package config

type Config struct {
	Port     int
	EthPk    string `split_words:"true"`
	EthPass  string `split_words:"true"`
	PoktPk   string `split_words:"true"`
	PoktPass string `split_words:"true"`
}

type Source interface {
	Load() (Config, error)
}

func Load(source Source) (Config, error) {
	return source.Load()
}

func AutoLoad() (Config, error) {
	envSrc := EnvSource{
		Prefix: "SIGNER",
	}
	cfg, err := Load(envSrc)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
