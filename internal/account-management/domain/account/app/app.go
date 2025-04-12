package app

import "github.com/neel4os/warg/internal/account-management/domain/account/app/commands"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AccountOnboardCommand *commands.AccountOnboardingCommandHandler
}
type Queries struct{}
