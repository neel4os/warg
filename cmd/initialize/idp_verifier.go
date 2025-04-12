package initialize

import (
	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/util"
	"github.com/rs/zerolog/log"
)

type idpVerifier struct {
	idpConfig config.IdpConfig
	realmUrl  string
}

func newidpVerifier(cfg config.IdpConfig) *idpVerifier {
	return &idpVerifier{idpConfig: cfg, realmUrl: cfg.Url + "/realms/" + cfg.RealmName}
}

func (i *idpVerifier) Verify() (error, bool) {
	// verify if we can access well-known configuration
	// if we can't, return error and true
	client := util.NewRestClient()
	log.Debug().Msgf("Checking well known %s", i.idpConfig.IdpName)
	_, err := client.R().Get(i.realmUrl + "/.well-known/openid-configuration")
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get well-known configuration for %s", i.idpConfig.IdpName)
		return err, true
	}
	return nil, true
}

func (i *idpVerifier) Create() error {
	return nil
}

func (i *idpVerifier) Name() string {
	return i.idpConfig.IdpName
}
