package elimination

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/graph"
)

func TestDouble(t *testing.T) {
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
			ExpectedNodes: 4,
		},
		{
			Teams:         []Team{1, 2, 3, 4},
			ExpectedNodes: 6,
		},
		{
			Teams:         []Team{1, 2, 3, 4, 5},
			ExpectedNodes: 8,
		},
		{
			Teams:         []Team{1, 2, 3, 4, 5, 6},
			ExpectedNodes: 10,
		},
		{
			Teams:         []Team{1, 2, 3, 4, 5, 6, 7},
			ExpectedNodes: 12,
		},
		{
			Teams:         []Team{1, 2, 3, 4, 5, 6, 7, 8},
			ExpectedNodes: 14,
		},
	} {
		t.Run(fmt.Sprintf("%d teams", len(tc.Teams)), func(t *testing.T) {
			tournament := NewTournament()
			s := Double{
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
