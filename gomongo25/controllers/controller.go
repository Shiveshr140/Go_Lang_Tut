package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "github.com/Shiveshr140/gomongo/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionString string = "mongodb+srv://shiveshr150:shiveshr150@gomongocluster.ufwkkzy.mongodb.net/?retryWrites=true&w=majority&appName=gomongocluster"
var dbName string = "netflix"
var colName string = "watchlist"

// Most important
// collection is a pointer to mongo.Collection which is a struct that represents a collection in mongo
var collection *mongo.Collection

// init is a special method in go that only runs at the start of the application
func init(){
  // client options
  clientOptions := options.Client().ApplyURI(connectionString)

  // connect to db
  // context.TODO() is used to create a context that can be used to cancel the operation if needed
  // context is used whenever you are making api call outside your local machine, like in this case we are connecting to a remote mongo db server
  // database are loacted remote so we need to define how long we can make a request, what will happen if request goes off and connection is still on there should be a contex to handle that
  // read the documentation of context in go to understand more about it
  // in those cases where you do not know which context to use simply go with todo
   client, err := mongo.Connect(context.TODO(), clientOptions)
 
   // you can use panic here but log will give more info
   if err != nil{
	log.Fatal(err)
   }

   fmt.Println("MongoDB connection successful")
   // mongo db is connected now we need to get the collection
   // client.Database(dbName) returns a pointer to mongo.Database which is a struct that represents a database in mongo
   // client.Database(dbName).Collection(colName) returns a pointer to mongo.Collection which is a struct that represents a collection in mongo

   collection = client.Database(dbName).Collection(colName)
   fmt.Println("Collection instance is ready")
}


// MoongoDB helper functions 

// why not models because model is package name
// you need to pass a context whenever yo are doing any db operations
func insertOneMovie(movie model.Netflix){
  inserted, err := collection.InsertOne(context.Background(), movie)

  if err != nil{
	log.Fatal(err)
  }

  fmt.Println("Inserted one movie in mongo db with id: ", inserted.InsertedID)
}

// we need to convert this string id into that type which mongo understands, string -> ObjectID by primitive package, ObjectID -> bson, last step was not needed in mongoose 
// so everthing in mogo is bson which look like json
func updateOneRecord(movieId string){
   id, _ := primitive.ObjectIDFromHex(movieId)
   filter := bson.M{"_id": id}
   update := bson.M{"$set": bson.M{"watched":true}}

   result, err := collection.UpdateOne(context.Background(), filter, update)

   if err != nil{
    log.Fatal(err)
   }

   fmt.Println("Modified count:", result.ModifiedCount)

}


func deleteOneRecord(movieId string){
    id, _ := primitive.ObjectIDFromHex(movieId) 
    filter := bson.M{"_id":id}
    result, err := collection.DeleteOne(context.Background(), filter)
    if err != nil{
      log.Fatal(err)
    }
    fmt.Println("Deleted count:", result.DeletedCount)
}


func deleteAllMovies() int64{
   deletedResults, err := collection.DeleteMany(context.Background(), bson.D{{}})

   if err != nil{
    log.Fatal(err)
   }

   fmt.Println("All Deleted Count:", deletedResults.DeletedCount)
   return deletedResults.DeletedCount
}

// Here cursor is kind of stack that contains all the results of the query and we can loop over through
// A cursor is a pointer to the result set of a query in MongoDB. It allows you to iterate over the results of a query one by one, rather than loading all results into memory at once
func getAllMovies() []primitive.M {
  cursor, err := collection.Find(context.Background(), bson.M{})
  
  if err != nil {
    log.Fatal(err)
  }

  var movies []primitive.M // primitive.M or bson.M is a type that represents a map with string keys and interface{} values, similar to bson.M but more generic
  
  // use for loop as while form as soon as cursor do not have next element it will stop
  for cursor.Next(context.Background()){
    var movie bson.M
    // decode the cursor and will need to pass the reference of movie ? 
    // whenever we docode we pass an reference if you decode use my structure to decode that since we do not have structure we have movie variable
    err := cursor.Decode(&movie)
    if err != nil{
      log.Fatal(err)
    }

    movies = append(movies, movie)
  }

  defer cursor.Close(context.Background()) // close the cursor after use to free up resources

  return movies
}

// *mongo.Result is a decoder
func getOneMovie(movieId string) model.Netflix{
   id, _ := primitive.ObjectIDFromHex(movieId)
   filter := bson.M{"_id": id}
   var movie model.Netflix
   result := collection.FindOne(context.Background(), filter)
   result.Decode(&movie)
   return movie
}

// all above are not actual controller functions, these are helper functions that will be used in controller functions
// controller functions are the ones that will be called from the routes

func GetAllMovies(w http.ResponseWriter, r *http.Request){
   w.Header().Set("Content-Type", "application/www-form-urlencoded")
   allMovies := getAllMovies()
   json.NewEncoder(w).Encode(allMovies)
}

func GetOneMovie(w http.ResponseWriter, r *http.Request){
   w.Header().Set("Content-Type", "application/www-form-urlencoded")
   params := mux.Vars(r)
   movie := getOneMovie(params["id"])
   json.NewEncoder(w).Encode(movie)

}

func CreateOneMovie(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/www-form-urlencoded")
  w.Header().Set("Allow-Control-Allow-Method", "POST")
  var movie model.Netflix
  err := json.NewDecoder(r.Body).Decode(&movie)
  if err != nil {
    log.Fatal(err)
  }
  insertOneMovie(movie)
  json.NewEncoder(w).Encode(movie)
  json.NewEncoder(w).Encode("Movie created successfully")
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
   w.Header().Set("Contect-Type", "application/www-form-urlencoded")
   w.Header().Set("Allow-Control-Allow-Method", "PUT")
   movieId := mux.Vars(r)["id"]

   updateOneRecord(movieId)
   json.NewEncoder(w).Encode("Movie marked as watched successfully")
   json.NewEncoder(w).Encode(movieId)
}


func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
   w.Header().Set("Content-Type", "application/www-form-urlencoded")
   w.Header().Set("Allow-Control-Allow-Method", "DELETE")
   deletedCount := deleteAllMovies()
   json.NewEncoder(w).Encode(fmt.Sprintf("%d movies deleted successfully", deletedCount))
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request){
   w.Header().Set("Content-Type", "application/www-form-urlencoded")
   w.Header().Set("Allow-Control-Allow-Method", "DELETE")
   movieId := mux.Vars(r)["id"]

   deleteOneRecord(movieId)
   json.NewEncoder(w).Encode(fmt.Sprintf("Movie with id %s deleted successfully", movieId))
}