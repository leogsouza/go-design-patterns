package flyweight

import (
	"testing"
)

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := teamFlyweightFactory{}

	teamA1 := factory.GetTeam(TeamA)
	if teamA1 == nil {
		t.Error("The pointer to the TeamA was nil")
	}

	teamA2 := factory.GetTeam(TeamA)
	if teamA2 == nil {
		t.Error("The pointer to the TeamA as nil")
	}

	if teamA1 != teamA2 {
		t.Error("TeamA pointers weren't the same")
	}

	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("The number of objects created was not 1: %d\n",
			factory.GetNumberOfObjects())
	}
}
