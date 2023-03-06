package infrastructure

import (
	"context"
	"fmt"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/collection"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	vo "<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	user "<MODULE_URL_REPLACE>/pkg/users/domain"
)

type MongoUsersRepository struct {
	//users //array de usuarios.
}

func NewMongoUsersRepository() MongoUsersRepository {
	return MongoUsersRepository{}
}

func (r *MongoUsersRepository) Save(ctx *context.Context, u user.User) *errors.AppError {
	fmt.Println("Save in mongo repo")
	return nil
}

func (r *MongoUsersRepository) Find(ctx *context.Context, id vo.Id) (user.User, *errors.AppError) {
	return user.User{}, nil
}

func (r *MongoUsersRepository) FindByCriteria(ctx *context.Context, f criteria.Criteria, o criteria.SortCriteria, p criteria.PaginatorCriteria) (collection.Collection, *errors.AppError) {

	return collection.Collection{}, nil
}

func (r *MongoUsersRepository) Delete(ctx *context.Context, id vo.Id) *errors.AppError {
	return nil
}

var _ user.UserRepository = (*MongoUsersRepository)(nil)

/*


import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Criteria interface {
	BuildQuery() bson.M
}

type AndCriteria struct {
	Criteria []Criteria
}

func (a AndCriteria) BuildQuery() bson.M {
	query := bson.M{"$and": []bson.M{}}
	for _, c := range a.Criteria {
		query["$and"] = append(query["$and"].([]bson.M), c.BuildQuery())
	}
	return query
}

type OrCriteria struct {
	Criteria []Criteria
}

func (o OrCriteria) BuildQuery() bson.M {
	query := bson.M{"$or": []bson.M{}}
	for _, c := range o.Criteria {
		query["$or"] = append(query["$or"].([]bson.M), c.BuildQuery())
	}
	return query
}

type NameCriteria struct {
	Name string
}

func (n NameCriteria) BuildQuery() bson.M {
	return bson.M{"name": n.Name}
}

type AgeCriteria struct {
	Operator string
	Age      int
}

func (a AgeCriteria) BuildQuery() bson.M {
	return bson.M{"age": bson.M{a.Operator: a.Age}}
}

type DateCriteria struct {
	Operator string
	Date     time.Time
}

func (d DateCriteria) BuildQuery() bson.M {
	return bson.M{"date": bson.M{d.Operator: d.Date}}
}

type GenderCriteria struct {
	Gender string
}

func (g GenderCriteria) BuildQuery() bson.M {
	return bson.M{"gender": g.Gender}
}

type StatusCriteria struct {
	Status string
}

func (s StatusCriteria) BuildQuery() bson.M {
	return bson.M{"status": s.Status}
}

func main() {
	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Get a handle to the database and collection
	collection := client.Database("test").Collection("people")

	// Build a query using criteria pattern
	nameCriteria := NameCriteria{Name: "John"}
	ageCriteria1 := AgeCriteria{Operator: "$gt", Age: 30}
	ageCriteria2 := AgeCriteria{Operator: "$lt", Age: 40}
	dateCriteria := DateCriteria{Operator: "$gte", Date: time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC)}
	genderCriteria := GenderCriteria{Gender: "male"}
	statusCriteria := StatusCriteria{Status: "active"}
	andCriteria := AndCriteria{Criteria: []Criteria{nameCriteria, ageCriteria1, ageCriteria2, dateCriteria}}
	orCriteria := OrCriteria{Criteria: []Criteria{genderCriteria, statusCriteria}}

	finalQuery := AndCriteria{Criteria: []Criteria{andCriteria, orCriteria}}

	// Execute the query
	cursor, err := collection.Find(context.Background(), finalQuery.BuildQuery())
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Iterate over the results and print them
	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}
*/
