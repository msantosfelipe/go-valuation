package config

import (
	"log"

	"github.com/Netflix/go-env"
	"github.com/subosito/gotenv"
)

type Environment struct {
	UrlStockInvest          string  `env:"URL_STATUSINVEST"`
	BazinDividendPercentage float64 `env:"BAZIN_DIVIDEND_PERCENTAGE"`
}

var ENV Environment

func init() {
	gotenv.Load() // load .env file (if exists)
	if _, err := env.UnmarshalFromEnviron(&ENV); err != nil {
		log.Fatal("Fatal error unmarshalling environment config: ", err)
	}
}
