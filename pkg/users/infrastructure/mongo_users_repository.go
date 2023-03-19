package infrastructure

import (
	"context"
	"fmt"
	"time"

	"<MODULE_URL_REPLACE>/pkg/shared/domain/collection"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/errors"
	"<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	vo "<MODULE_URL_REPLACE>/pkg/shared/domain/valueobjects"
	criteriaAdapter "<MODULE_URL_REPLACE>/pkg/shared/infrastructure/criteria"
	"<MODULE_URL_REPLACE>/pkg/shared/infrastructure/storage"
	user "<MODULE_URL_REPLACE>/pkg/users/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUser struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name,omitempty,"`
	LastName  string             `bson:"last_name,omitempty"`
	Email     string             `bson:"email,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdateAt  time.Time          `bson:"updated_at,omitempty"`
}

type MongoUsersRepository struct {
	collection      *mongo.Collection
	criteriaBuilder criteriaAdapter.MongoCriteriaBuilderAdapter
}

func NewMongoUsersRepository(db string) *MongoUsersRepository {
	client := storage.NewMongoClient()

	return &MongoUsersRepository{
		collection:      client.Database(db).Collection("users"),
		criteriaBuilder: criteriaAdapter.NewMongoCriteriaBuilderAdapter(),
	}
}

func (r *MongoUsersRepository) Save(ctx *context.Context, u user.User) *errors.AppError {

	objectId, err := primitive.ObjectIDFromHex(u.GetId().Value.Hex())

	if err != nil {
		return user.NewUserError(user.INVALID_USER_ID_ERROR)
	}

	_, err = r.collection.ReplaceOne(*ctx, bson.M{"_id": objectId}, MongoUser{
		Id:        u.GetId().Value,
		FirstName: u.GetFirstName().Value,
		LastName:  u.GetLastName().Value,
		Email:     u.GetEmail().Value,
		CreatedAt: u.GetCreatedAt().Value,
	}, options.Replace().SetUpsert(true))

	if err != nil {
		return errors.NewAppError(user.REPOSITORY_USER_ERROR)
	}

	return nil
}

func (r *MongoUsersRepository) Find(ctx *context.Context, id vo.Id) (user.User, *errors.AppError) {
	var result MongoUser

	objectId, err := primitive.ObjectIDFromHex(id.Value.Hex())

	if err != nil {
		return user.User{}, user.NewUserError(user.INVALID_USER_ID_ERROR)
	}

	err = r.collection.FindOne(*ctx, bson.M{"_id": objectId}).Decode(&result)

	if err != nil {
		return user.User{}, user.NewUserError(user.NOT_FOUND_ERROR)
	}

	userFound, err := newUserFromMongoUser(result)

	if err != nil {
		return user.User{}, user.NewUserError(user.INVALID_USER_ERROR)
	}

	return userFound, nil
}

func (r *MongoUsersRepository) FindByCriteria(ctx *context.Context, c criteria.Criteria, o criteria.SorterCriteria, p criteria.PaginatorCriteria) (collection.Collection, *errors.AppError) {
	var results []MongoUser
	var users user.List
 
	filter := r.criteriaBuilder.Build(c)

	options := options.Find().SetSkip(int64(p.Offset())).SetLimit(int64(p.Limit())).SetSort(r.criteriaBuilder.Sort(o))

	cursor, err := r.collection.Find(*ctx, filter, options)

	if err != nil {
		fmt.Println(err.Error())
		return collection.Collection{}, user.NewUserError(errors.UNKNOWN_ERROR)
	}

	if err = cursor.All(*ctx, &results); err != nil {
		return collection.Collection{}, user.NewUserError(errors.UNKNOWN_ERROR)
	}

	count, err := r.collection.CountDocuments(*ctx, filter)

	if err != nil {
		return collection.Collection{}, user.NewUserError(errors.UNKNOWN_ERROR)
	}

	for _, r := range results {
		u, err := newUserFromMongoUser(r)

		if err != nil {
			return collection.Collection{}, user.NewUserError(user.INVALID_USER_ERROR)
		}

		users = append(users, u)
	}

	return collection.NewCollection(users, p.Page(), p.PageSize(), uint32(count&0xffffffff)), nil
}

func (r *MongoUsersRepository) Delete(ctx *context.Context, id vo.Id) *errors.AppError {
	objectId, err := primitive.ObjectIDFromHex(id.Value.Hex())

	if err != nil {
		return user.NewUserError(user.INVALID_USER_ID_ERROR)
	}

	_, err = r.collection.DeleteOne(*ctx, bson.M{"_id": objectId}, options.Delete().SetCollation(&options.Collation{}))

	if err != nil {
		return user.NewUserError(user.USER_DELETE_ERROR)
	}

	return nil
}

var _ user.UserRepository = (*MongoUsersRepository)(nil)

func newUserFromMongoUser(u MongoUser) (user.User, errors.App) {
	id, err := valueobjects.NewIdFromString(u.Id.Hex())

	if err != nil {
		return user.User{}, err
	}

	firstName, err := user.NewFirstName(u.FirstName)

	if err != nil {
		return user.User{}, err
	}

	lastName, err := user.NewLastName(u.LastName)

	if err != nil {
		return user.User{}, err
	}

	email, err := user.NewEmail(u.Email)

	if err != nil {
		return user.User{}, err
	}

	createdAt := valueobjects.NewDateTime(u.CreatedAt)

	if err != nil {
		return user.User{}, err
	}

	return user.User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		CreatedAt: createdAt,
	}, nil
}
