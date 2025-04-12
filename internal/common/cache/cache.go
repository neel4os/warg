package cache

import (
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/neel4os/warg/internal/common/config"
	"github.com/rs/zerolog/log"
	"resty.dev/v3"
)

var (
	client_id          = "admin-cli"
	cache_key_keycloak = "kctoken"
)

type TokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
}

type IMCache struct {
	cache *ttlcache.Cache[string, string]
}

func NewIMCache(conf *config.Config) *IMCache {
	loader := ttlcache.LoaderFunc[string, string](
		func(c *ttlcache.Cache[string, string], key string) *ttlcache.Item[string, string] {
			idp_config := conf.IdpConfig
			_token := &TokenResponse{}
			// hardcoding the realm to master
			//FIXME: may be admin need to be created in warg realm
			_token_path := idp_config.Url + "/realms/" + "master" + "/protocol/openid-connect/token"
			data := map[string]string{
				"client_id":  client_id,
				"username":   idp_config.Username,
				"password":   idp_config.Password,
				"grant_type": "password",
			}
			restyClient := resty.New()
			restyClient.SetHeader("Content-Type", "application/x-www-form-urlencoded")
			restyClient.SetHeader("Accept", "application/json")
			restyClient.SetRetryCount(3)
			response, err := restyClient.R().SetFormData(data).
				SetResult(_token).
				Post(_token_path)
			if err != nil {
				log.Fatal().Err(err).Caller().Msg("Error while getting token")
				//return nil
			}
			if response.IsError() {
				log.Fatal().Str("response", string(response.Bytes())).Caller().Msg("Error while getting token")
				//return nil
			}

			return c.Set(cache_key_keycloak, _token.AccessToken, time.Duration(conf.IdpConfig.TokenExpiry)*time.Second)
		},
	)
	cache := ttlcache.New(
		ttlcache.WithLoader(loader),
		ttlcache.WithTTL[string, string](time.Duration(conf.IdpConfig.TokenExpiry)*time.Second),
		ttlcache.WithDisableTouchOnHit[string, string](),
	)
	return &IMCache{
		cache: cache,
	}
}

func (c *IMCache) GetToken() string {
	token := c.cache.Get(cache_key_keycloak)
	return token.Value()
}

func (c *IMCache) Run() {
	c.cache.Start()
}

func (c *IMCache) Stop() {
}

func (c *IMCache) Init() {
}

func (c *IMCache) Name() string {
	return "IMCache"
}
