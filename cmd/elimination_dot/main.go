package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/proglottis/elimination"
	"gonum.org/v1/gonum/graph/encoding/dot"
)

type Builder interface {
	Build([]elimination.Team)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func run() error {
	var tournamentType string
	flag.StringVar(&tournamentType, "type", "double", "single or double")
	flag.Parse()

	tournament := elimination.NewTournament()

	var builder Builder
	switch tournamentType {
	case "single":
		builder = &elimination.Single{
			G: tournament,
		}
	case "double":
		builder = &elimination.Double{
			G: tournament,
		}
	}

	var teams []elimination.Team
	for _, t := range flag.Args() {
		teams = append(teams, t)
	}

	builder.Build(teams)

	d, err := dot.Marshal(tournament, "", "", "\t")
	if err != nil {
		return err
	}
	d = append(d, '\n')

	_, err = os.Stdout.Write(d)
	return err
}
