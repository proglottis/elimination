package elimination_test

import (
	"fmt"
	"testing"

	. "github.com/proglottis/elimination"
	"github.com/proglottis/elimination/simple"
	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph"
)

func TestSingle(t *testing.T) {
	for _, tc := range []struct {
		Teams         []Team
		ExpectedNodes int
	}{
		{
			Teams:         []Team{},
			ExpectedNodes: 1,
		},
		{
			Teams:         []Team{1},
			ExpectedNodes: 1,
		},
		{
			Teams:         []Team{1, 2},
			ExpectedNodes: 1,
		},
		{
			Teams:         []Team{1, 2, 3},
			ExpectedNodes: 2,
		},
		{
			Teams:         []Team{1, 2, 3, 4},
			ExpectedNodes: 3,
		},
		{
			Teams:         []Team{1, 2, 3, 4, 5},
			ExpectedNodes: 4,
		},
		{
			Teams:         []Team{1, 2, 3, 4, 5, 6},
			ExpectedNodes: 5,
		},
		{
			Teams:         []Team{1, 2, 3, 4, 5, 6, 7},
			ExpectedNodes: 6,
		},
		{
			Teams:         []Team{1, 2, 3, 4, 5, 6, 7, 8},
			ExpectedNodes: 7,
		},
	} {
		t.Run(fmt.Sprintf("%d teams", len(tc.Teams)), func(t *testing.T) {
			tournament := simple.NewTournament()
			s := Single{
				G: tournament,
			}

			s.Build(tc.Teams)

			nodes := graph.NodesOf(tournament.Nodes())
			teams := teamsOf(nodes)

			assert.Equal(t, tc.ExpectedNodes, len(nodes))
			assert.ElementsMatch(t, tc.Teams, teams)
		})
	}
}

func teamsOf(nodes []graph.Node) []Team {
	var teams []Team
	for _, n := range nodes {
		m, ok := n.(*simple.Match)
		if !ok {
			continue
		}
		teams = append(teams, m.Teams...)
	}
	return teams
}
