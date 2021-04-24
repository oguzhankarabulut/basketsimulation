package mongo

import (
	"basketsimulation/pkg/domain"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	collMatchPlayer = "match_player"
)

type MatchPlayerRepository struct {
	mongoClient *Client
	db          string
}

func NewMatchPlayerRepository(mongoClient *Client, db string) MatchPlayerRepository {
	return MatchPlayerRepository{
		mongoClient: mongoClient,
		db:          db,
	}
}


func (r MatchPlayerRepository) Save(matchPlayer *domain.MatchPlayer) error{
	if err := r.mongoClient.InsertOne(r.db, collMatchPlayer, matchPlayer); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r MatchPlayerRepository) ScoreCount(matchId string, playerId string, score int) (int64, error) {
	f := bson.M{
		"score": bson.M{eq: score},
		"matchid": bson.M{eq: matchId},
		"playerid": bson.M{eq: playerId},
	}

	c, err := r.mongoClient.Count(r.db, collMatchPlayer, f)
	if err != nil {
		return 0, err
	}
	return c, nil
}

func (r MatchPlayerRepository) AssistCount(matchId string, playerId string) (int64, error) {
	f := bson.M{
		"assist": bson.M{eq: true},
		"matchid": bson.M{eq: matchId},
		"playerid": bson.M{eq: playerId},
	}

	c, err := r.mongoClient.Count(r.db, collMatchPlayer, f)
	if err != nil {
		return 0, err
	}
	return c, nil
}
