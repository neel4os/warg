package config

import (
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	os.Setenv("WARG_DBCONFIG_PORT", "5432")
	os.Setenv("WARG_DBCONFIG_HOST", "127.0.0.1")
	os.Setenv("WARG_DBCONFIG_USER", "testuser")
	os.Setenv("WARG_DBCONFIG_PASSWORD", "testpassword")
	os.Setenv("WARG_DBCONFIG_DBNAME", "testdb")
	os.Setenv("WARG_DBCONFIG_SSLMODE", "require")
	os.Setenv("WARG_SERVERCONFIG_PORT", "8080")
	os.Setenv("WARG_LOGGERCONFIG_IS_DEBUG_LOG", "false")
	os.Setenv("WARG_IDPCONFIG_IDP_NAME", "auth0")
	os.Setenv("WARG_IDPCONFIG_IDP_URL", "http://auth0.com")
	os.Setenv("WARG_IDPCONFIG_IDP_REALM_NAME", "testrealm")

	cfg := GetConfig()

	if cfg.DbConfig.Port != "5432" {
		t.Errorf("expected DbConfig.Port to be '5432', got '%s'", cfg.DbConfig.Port)
	}
	if cfg.DbConfig.Host != "127.0.0.1" {
		t.Errorf("expected DbConfig.Host to be '127.0.0.1', got '%s'", cfg.DbConfig.Host)
	}
	if cfg.DbConfig.User != "testuser" {
		t.Errorf("expected DbConfig.User to be 'testuser', got '%s'", cfg.DbConfig.User)
	}
	if cfg.DbConfig.Password != "testpassword" {
		t.Errorf("expected DbConfig.Password to be 'testpassword', got '%s'", cfg.DbConfig.Password)
	}
	if cfg.DbConfig.DbName != "testdb" {
		t.Errorf("expected DbConfig.DbName to be 'testdb', got '%s'", cfg.DbConfig.DbName)
	}
	if cfg.DbConfig.SslMode != "require" {
		t.Errorf("expected DbConfig.SslMode to be 'require', got '%s'", cfg.DbConfig.SslMode)
	}
	if cfg.ServerConfig.Port != "8080" {
		t.Errorf("expected ServerConfig.Port to be '8080', got '%s'", cfg.ServerConfig.Port)
	}
	if cfg.LoggerConfig.IsDebugLog != false {
		t.Errorf("expected LoggerConfig.IsDebugLog to be 'false', got '%v'", cfg.LoggerConfig.IsDebugLog)
	}
	if cfg.IdpConfig.IdpName != "auth0" {
		t.Errorf("expected IdpConfig.IdpName to be 'auth0', got '%s'", cfg.IdpConfig.IdpName)
	}
	if cfg.IdpConfig.Url != "http://auth0.com" {
		t.Errorf("expected IdpConfig.Url to be 'http://auth0.com', got '%s'", cfg.IdpConfig.Url)
	}
	if cfg.IdpConfig.RealmName != "testrealm" {
		t.Errorf("expected IdpConfig.RealmName to be 'testrealm', got '%s'", cfg.IdpConfig.RealmName)
	}
}

func TestDbConfig_GetDbDsn(t *testing.T) {
	dbConfig := DbConfig{
		Port:     "5432",
		Host:     "127.0.0.1",
		User:     "testuser",
		Password: "testpassword",
		DbName:   "testdb",
		SslMode:  "require",
	}

	expectedDsn := "postgres://testuser:testpassword@127.0.0.1:5432/testdb?sslmode=require"
	if dsn := dbConfig.GetDbDsn(); dsn != expectedDsn {
		t.Errorf("expected DSN to be '%s', got '%s'", expectedDsn, dsn)
	}
}