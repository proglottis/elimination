package elimination

import (
	"gonum.org/v1/gonum/graph"
)

type Single struct {
	G graph.DirectedBuilder
}

func (s *Single) Build(teams []Team) {
	final := s.G.NewNode()
	s.G.AddNode(final)
	s.buildBracket(teams, final)
}

func (s *Single) buildBracket(teams []Team, next graph.Node) {
	switch len(teams) {
	case 0, 1, 2:
		addTeams(next, teams)
	case 3:
		split := 2
		s.buildMatch(teams[:split], next)
		addTeams(next, teams[split:])
	default:
		split := len(teams) / 2
		s.buildMatch(teams[:split], next)
		s.buildMatch(teams[split:], next)
	}
}

func (s *Single) buildMatch(teams []Team, next graph.Node) {
	match := s.G.NewNode()
	s.G.SetEdge(s.G.NewEdge(match, next))
	s.buildBracket(teams, match)
}

func addTeams(n graph.Node, teams []Team) {
	if len(teams) > 2 {
		panic("elimination: too many teams")
	}
	match, ok := n.(Match)
	if !ok {
		return
	}
	for _, t := range teams {
		match.AddTeam(t)
	}
}
