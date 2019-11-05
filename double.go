package elimination

import (
	"gonum.org/v1/gonum/graph"
)

type Double struct {
	G graph.DirectedBuilder
}

func (d *Double) Build(teams []Team) {
	final := d.G.NewNode()
	d.G.AddNode(final)

	if len(teams) < 3 {
		d.buildBracket(teams, final, final)
		return
	}

	winners := d.G.NewNode()
	d.G.SetEdge(d.G.NewEdge(winners, final))
	losers := d.G.NewNode()
	d.G.SetEdge(d.G.NewEdge(losers, final))
	d.buildBracket(teams, winners, losers)
}

func (d *Double) buildBracket(teams []Team, next, losersNext graph.Node) {
	switch len(teams) {
	case 0, 1, 2:
		addTeams(next, teams)
	case 3:
		split := 2
		d.buildMatch(teams[:split], next, losersNext)
		addTeams(next, teams[split:])
	default:
		losersJoin := d.G.NewNode()
		d.G.SetEdge(d.G.NewEdge(losersJoin, losersNext))

		split := len(teams) - len(teams)/2
		d.buildMatch(teams[:split], next, losersJoin)
		d.buildMatch(teams[split:], next, losersJoin)
	}
}

func (d *Double) buildMatch(teams []Team, next, losersNext graph.Node) {
	match := d.G.NewNode()
	d.G.SetEdge(d.G.NewEdge(match, next))

	if len(teams) > 2 {
		losers := d.G.NewNode()
		d.G.SetEdge(d.G.NewEdge(losers, losersNext))
		losersNext = losers
	}

	d.buildBracket(teams, match, losersNext)
}
