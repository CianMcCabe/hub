package main

import (
	"github.com/pksunkara/hub/remote"
	"github.com/pksunkara/hub/utils"
	"strings"
)

type FetchCommand struct{}

func (f *FetchCommand) Execute(args []string) error {
	if len(args) == 0 {
		return &utils.ErrArgument{}
	} else if len(args) > 1 {
		return &utils.ErrProxy{}
	}

	if len(strings.Split(args[0], "/")) != 1 {
		return &utils.ErrProxy{}
	}

	users := strings.Split(args[0], ",")
	remotes, err := utils.Remotes()

	if err != nil {
		return err
	}

	remoteAdd := &remote.AddCommand{}

	for _, user := range users {
		if _, ok := remotes[user]; !ok {
			remoteAdd.Execute([]string{user})
		}
	}

	if err := utils.Git(append([]string{"fetch", "--multiple"}, users...)...); err != nil {
		return err
	}

	utils.HandleInfo("Fetched from remotes " + args[0])

	return nil
}

func (f *FetchCommand) Usage() string {
	return "<user | users>"
}
