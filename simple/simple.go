package simple

import (
	"fmt"
	"strings"

	"github.com/proglottis/elimination"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/simple"
)

type Match struct {
	graph.Node
	Teams []elimination.Team
}

func (m *Match) DOTID() string {
	return fmt.Sprintf("match_%d", m.ID())
}

func (m *Match) Attributes() []encoding.Attribute {
	var teamNames []string
	for _, t := range m.Teams {
		teamNames = append(teamNames, t.(string))
	}
	for len(teamNames) < 2 {
		teamNames = append(teamNames, "_____")
	}
	return []encoding.Attribute{
		encoding.Attribute{Key: "label", Value: strings.Join(teamNames, "\nvs\n")},
	}
}

func (m *Match) AddTeam(t elimination.Team) {
	m.Teams = append(m.Teams, t)
}

type Tournament struct {
	graph.DirectedBuilder
}

func NewTournament() *Tournament {
	return &Tournament{
		DirectedBuilder: simple.NewDirectedGraph(),
	}
}

func (t *Tournament) NewNode() graph.Node {
	return &Match{
		Node: t.DirectedBuilder.NewNode(),
	}
}

func (t *Tournament) DOTID() string {
	return "tournament"
}

type attributes []encoding.Attribute

func (a attributes) Attributes() []encoding.Attribute {
	return a
}

func (t *Tournament) DOTAttributers() (graph, node, edge encoding.Attributer) {
	graph = attributes{
		encoding.Attribute{Key: "rankdir", Value: "LR"},
	}
	node = attributes{
		encoding.Attribute{Key: "shape", Value: "box"},
	}
	edge = attributes{}
	return
}
