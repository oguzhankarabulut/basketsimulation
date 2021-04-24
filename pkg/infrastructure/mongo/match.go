package mongo

import (
	"basketsimulation/pkg/domain"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	collMatch = "match"
)

type MatchRepository struct {
	mongoClient *Client
	db          string
}

func NewMatchRepository(mongoClient *Client, db string) MatchRepository {
	return MatchRepository{
		mongoClient: mongoClient,
		db:          db,
	}
}


func (r MatchRepository) Save(match *domain.Match) error{
	if err := r.mongoClient.InsertOne(r.db, collMatch, match); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r MatchRepository) Latest() (*domain.Match, error) {
	m := &domain.Match{}
	pm, _ := r.mongoClient.Latest(r.db, collMatch)
	pmb, _ := bson.Marshal(pm)
	_ = bson.Unmarshal(pmb, m)
	return m, nil
}