package basketball

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	// GameResults
	wire.Struct(new(BasketballResults)),
	wire.Struct(new(Results)),
	wire.Struct(new(SportEvent)),
	wire.Struct(new(TournamentRound)),
	wire.Struct(new(Season)),
	wire.Struct(new(Tournament)),
	wire.Struct(new(Sport)),
	wire.Struct(new(Category)),
	wire.Struct(new(Competitors)),
	wire.Struct(new(Venue)),
	wire.Struct(new(SportEventStatus)),
	wire.Struct(new(PeriodScores)),

	// GameSchedules
	wire.Struct(new(BasketballSchedules)),
	wire.Struct(new(SportEvents)),

	// HeadToHeads
	wire.Struct(new(BasketballMatches)),
	wire.Struct(new(Teams)),
	wire.Struct(new(LastMeetings)),
)
