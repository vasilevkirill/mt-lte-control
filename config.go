package main

func loadConfig() {
	for _, f := range configPaths {
		Config.AddConfigPath(f)
	}
	Config.SetConfigName("mt-lte-control")
	Config.SetConfigType("yaml")

	Config.SetDefault("web.address", "0.0.0.0")
	Config.SetDefault("web.port", 8080)
	Config.SetDefault("router.script", "mt-lte-control")
	Config.SetDefault("router.storetime", 10)
	err := Config.ReadInConfig()
	if err != nil {
		Log.Panicf("fatal error config file: %s", err.Error())
	}
}
