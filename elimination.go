package elimination

import (
	"gonum.org/v1/gonum/graph"
)

type Team interface{}

type Match interface {
	graph.Node

	AddTeam(Team)
}
