package match

// Participant represents a player for a match
type Participant struct {
	ID      string `jsonapi:"primary,participant"`
	Actor   string `jsonapi:"attr,actor"`
	ShardID string `jsonapi:"attr,shardId"`
	Stats   struct {
		Assists         int     `jsonapi:"attr,assists"`
		Boosts          int     `jsonapi:"attr,boosts"`
		DamageDealt     float64 `jsonapi:"attr,damageDealt"`
		DeathType       string  `jsonapi:"attr,deathType"`
		HeadshotKills   int     `jsonapi:"attr,headshotKills"`
		Heals           int     `jsonapi:"attr,heals"`
		KillPlace       int     `jsonapi:"attr,killPlace"`
		KillStreaks     int     `jsonapi:"attr,killStreaks"`
		Kills           int     `jsonapi:"attr,kills"`
		LongestKill     int     `jsonapi:"attr,longestKill"`
		Name            string  `jsonapi:"attr,name"`
		PlayerID        string  `jsonapi:"attr,playerId"`
		Revives         int     `jsonapi:"attr,revives"`
		RideDistance    float64 `jsonapi:"attr,rideDistance"`
		RoadKills       int     `jsonapi:"attr,roadKills"`
		TeamKills       int     `jsonapi:"attr,teamKills"`
		TimeSurvived    float64 `jsonapi:"attr,timeSurvived"`
		VehicleDestroys int     `jsonapi:"attr,vehicleDestroys"`
		WalkDistance    float64 `jsonapi:"attr,walkDistance"`
		WinPlace        int     `jsonapi:"attr,winPlace"`
		SwimDistance    float64 `jsonapi:"attr,swimDistance"`
	} `jsonapi:"attr,stats"`
}
