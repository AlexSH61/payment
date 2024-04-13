package main

import (
	"fmt"
	"os"

	"github.com/AlexSH61/payment/foundation/logger"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()
	if err != nil {
		log.Errorw("startup", "ERROR", err)
		os.Exit(1)
	}
}
func run(log *zap.SugaredLogger) error {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return err
	}
	log.Infow("startup", "OK! Connection")
	log.Infow("startup", "ENV", &cfg.Env)
	return nil
}
