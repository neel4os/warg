package organization

import (
	"errors"
	"strings"

	"github.com/neel4os/warg/internal/common/cache"
	"github.com/neel4os/warg/internal/common/config"
	"github.com/neel4os/warg/internal/common/keycloak"
	"github.com/rs/zerolog/log"
	"resty.dev/v3"
)

type OrganizationKeycloakRepository struct {
	client *resty.Client
}

func NewOrganizationKeycloakRepository() *OrganizationKeycloakRepository {
	cfg := config.GetConfig()
	cache := cache.NewIMCache(cfg)
	token := cache.GetToken()
	restyClient := resty.New()
	restyClient.SetHeader("Authorization", "Bearer "+token)
	restyClient.SetHeader("Content-Type", "application/json")
	restyClient.SetHeader("Accept", "application/json")
	restyClient.SetRetryCount(3)
	restyClient.SetBaseURL(cfg.IdpConfig.Url + "/admin/realms/" + cfg.IdpConfig.RealmName)
	return &OrganizationKeycloakRepository{
		client: restyClient,
	}
}

func (r *OrganizationKeycloakRepository) CreateOrganization(name string) (string, error) {
	resp, err := r.client.R().SetBody(keycloak.NewOrganizationRepresentation(name)).Post("/organizations")
	if err != nil {
		log.Error().Err(err).Caller().Msg("Error while creating organization")
		return "", err
	}
	if resp.IsError() {
		log.Error().Str("response", string(resp.Bytes())).Caller().Msg("Error while creating organization")
		return "", errors.New("Error while creating organization")
	}
	location := resp.Header().Get("Location")
	if location == "" {
		log.Error().Caller().Msg("Location header is missing in the response")
		return "", errors.New("Location header is missing in the response")
	}
	orgId := location[strings.LastIndex(location, "/")+1:]
	log.Info().Str("orgId", orgId).Msg("Extracted organization ID")
	log.Info().Str("location", location).Msg("Organization created successfully")
	return orgId, nil
}
