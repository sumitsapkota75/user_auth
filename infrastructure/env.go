package infrastructure

import "os"

// Env has environment stored
type Env struct {
	ServerPort  string
	DBType      string
	DBName      string
	DBUsername  string
	DBPassword  string
	Environment string
	LogOutput   string
}

// NewEnv creates a new enviroment
func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

// LoadEnv loads enviromnet
func (env *Env) LoadEnv() {
	env.ServerPort = os.Getenv("ServerPort")
	env.Environment = os.Getenv("Environment")
	env.DBName = os.Getenv("DBName")
	env.DBType = os.Getenv("DBType")
	env.DBUsername = os.Getenv("DBUsername")
	env.DBPassword = os.Getenv("DBPassword")
	env.LogOutput = os.Getenv("LogOutput")
}
