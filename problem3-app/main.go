package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	cli "github.com/urfave/cli/v2"
)

var (
	dbHostname         string
	dbUser, dbPassword string
)

func handleRegister(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, World")
}

func cliMain() {
	app := cli.NewApp()
	app.Name = "problem3"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "dbhost",
			Usage:       "DBホスト名",
			Destination: &dbHostname,
			EnvVar:      "DB_HOSTNAME",
		},
		cli.StringFlag{
			Name:        "dbuser",
			Usage:       "DBユーザ名",
			Destination: &dbUser,
			EnvVar:      "DB_USER",
		},
		cli.StringFlag{
			Name:        "dbpassword",
			Usage:       "DBパスワード",
			Destination: &dbPassword,
			EnvVar:      "DB_PASSWORD",
		},
	}
	app.Action = func(cliCtx *cli.Context) error {
		router := httprouter.New()
		router.POST("/register", handleRegister)

		if err := http.ListenAndServe(":8000", router); err != nil {
			return cli.NewExitError(err, 1)
		}
	}

	if err := app.Run(os.Args); err != nil {
		exitErr := err.(*cli.ExitError)
		log.Println(exitErr.Error())
		return exitErr.ExitCode()
	}

	return 0
}

func main() {
	os.Exit(cliMain())
}
