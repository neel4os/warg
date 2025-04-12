package initialize

import (
	"github.com/neel4os/warg/internal/common/config"
	"github.com/rs/zerolog/log"
)

type verifier interface {
	Verify() (error, bool)
	Create() error
	Name() string
}

type wargInitialization struct {
	verifiers []verifier
}

func NewInitilizer(cfg *config.Config) *wargInitialization {
	_verifiers := make([]verifier, 0)
	_verifiers = append(_verifiers, newidpVerifier(cfg.IdpConfig))
	return &wargInitialization{verifiers: _verifiers}
}

func (w *wargInitialization) DoInitialize() error {
	for _, v := range w.verifiers {
		log.Debug().Str("verifier", v.Name()).Caller().Msg("Verifying " + v.Name())
		err, ok := v.Verify()
		if err != nil {
			log.Err(err).Caller().Msg("Error verifying")
			return err
		}
		if !ok {
			log.Debug().Str("verifier", v.Name()).Caller().Msg("Creating " + v.Name())
			err = v.Create()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
