package all

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

// TODO to be implemented
// all - Execute all other commands in a row
func all(c *cli.Context) error {

	// filesPath := c.String("files")
	// registry := c.String("registry")
	// patch := c.Bool("patch")
	// kstPath := c.String("kustomization-path")

	// WARN: OPTIMIZE DOCKER-CLIENT OPENING/CLOSING

	// scan
	// pull
	// push
	// if *patch == "true" {
	// patch
	// }

	errMsg := "Method not yet implemented!"
	log.Error(errMsg)
	return errors.New(errMsg)
}
