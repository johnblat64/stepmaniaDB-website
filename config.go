package main

type EnvironmentConfig struct {
	BackendApiUrl string
	FileStoreUrl  string
}

var GEnvironmentConfig EnvironmentConfig
