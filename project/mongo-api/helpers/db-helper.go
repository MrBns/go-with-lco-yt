package helpers

import (
	"context"
	"fmt"
	"log"
	"mongoApi/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/"
const dbName = "Netflix"
const colName = "watchlist"

var Collection *mongo.Collection

func init() {
	//Client Option
	clientOptions := options.Client().ApplyURI(connectionString)

	//connect to Mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb Connection Successfully")
	Collection = client.Database(dbName).Collection(colName)

	fmt.Println("Connection Instance is Ready")
}

// MONGODB helpers

// Insert 1 Record
func InsertOneMovie(movie model.Netflix) *mongo.InsertOneResult {
	inserted, err := Collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted Movie :", inserted)
	return inserted
}

// Update One Record;
func UpdateOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)

	filterQuery := bson.M{"_id": id}
	updateQuery := bson.M{"$set": bson.M{"watched": true}}

	result, err := Collection.UpdateOne(context.Background(), filterQuery, updateQuery)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified Count: ", result.ModifiedCount)
}

// Delete One Record
func DeleteOneMovie(movieId string) int64 {
	deletedResult, err := Collection.DeleteOne(context.Background(), bson.M{"_id": movieId})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Result :", deletedResult)
	return deletedResult.DeletedCount
}

// Delete all Record from Mongodb;
func DeleteAllMovie() int64 {
	deleteResult, err := Collection.DeleteMany(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Count : ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// Get All Movies from collection;
func GetAllMovies() []bson.M {
	cursor, err := Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Found Data Cursor: ", cursor)
	var movies []bson.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}
