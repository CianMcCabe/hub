package main

type AuthUserCommand struct{}

func (a *AuthUserCommand) Execute(args []string) error {
	return nil
}
