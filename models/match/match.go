package match

import (
	"github.com/slemgrim/jsonapi"
	"io"
	"time"
)

// Match structure represent data related to a PUBG match
type Match struct {
	ID           string    `jsonapi:"primary,match"`
	CreatedAt    time.Time `jsonapi:"attr,createdAt,iso8601"`
	Duration     int       `jsonapi:"attr,duration"`
	GameMode     string    `jsonapi:"attr,gameMode"`
	PatchVersion string    `jsonapi:"attr,patchVersion"`
	ShardID      string    `jsonapi:"attr,shardId"`
	TitleID      string    `jsonapi:"attr,titleId"`
	MapName      string    `jsonapi:"attr,mapName"`
	Rosters      []*Roster `jsonapi:"relation,rosters"`
}

// ParseMatch parses a json response containing matches information
func ParseMatch(in io.Reader) (*Match, error) {
	match := new(Match)
	err := jsonapi.UnmarshalPayload(in, match)
	if err != nil {
		panic(err.Error())
		return nil, err
	}

	return match, nil
}
