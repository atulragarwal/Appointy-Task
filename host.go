package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "regexp"
)

type User struct {
	ID       int    `bson:"id" json:"id"`
	Name     string `bson:"name" json:name`
	Email    string `bson:"email" json:email`
	Password string `bson:"password" json:password`
}

type Post struct {
	ID        int       `bson:"id,omitempty"`
	PostId    int       `bson:"postId,omitempty"`
	Caption   string    `bson:"caption,omitempty"`
	ImageUrl  string    `bson:"imageUrl,omitempty"`
	TimeStamp time.Time `bson:"timeStamp,omitempty"`
}

func checkUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodPost:
		makeUser(w, r)
		return
	case r.Method == http.MethodGet:
		checkUser(w, r)
		return
	}
}

func makeUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://atulragarwal:atul2885@appointy-task.ifezn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("Appointy-Task").Collection("user")
	enterResult, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted", enterResult.InsertedID)
}

func checkUser(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://atulragarwal:atul2885@appointy-task.ifezn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("Appointy-Task").Collection("user")
	var episodes []User
	tempId := r.URL.Query().Get("id")
	newTemp, err := strconv.Atoi(tempId)
	cursor, err := collection.Find(ctx, bson.M{"id": newTemp})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &episodes); err != nil {
		panic(err)
	}
	fmt.Println(episodes)
}
func checkPostUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodPost:
		makePost(w, r)
		return
	case r.Method == http.MethodGet:
		checkPost(w, r)
		return
	}
}

func makePost(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.TimeStamp = time.Now()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://atulragarwal:atul2885@appointy-task.ifezn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("Appointy-Task").Collection("post")
	enterResult, err := collection.InsertOne(ctx, post)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted", enterResult.InsertedID)
}

func checkPost(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://atulragarwal:atul2885@appointy-task.ifezn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("Appointy-Task").Collection("post")
	var episodes []Post
	tempId := r.URL.Query().Get("id")
	newTemp, err := strconv.Atoi(tempId)
	cursor, err := collection.Find(ctx, bson.M{"postId": newTemp})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &episodes); err != nil {
		panic(err)
	}
	fmt.Println(episodes)
}

func getUserPosts(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://atulragarwal:atul2885@appointy-task.ifezn.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("Appointy-Task").Collection("post")
	var episodes []Post
	tempId := r.URL.Query().Get("id")
	newTemp, err := strconv.Atoi(tempId)
	cursor, err := collection.Find(ctx, bson.M{"id": newTemp})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &episodes); err != nil {
		panic(err)
	}
	fmt.Println(episodes)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", checkUrl)
	mux.HandleFunc("/posts", checkPostUrl)
	mux.HandleFunc("/posts/users", getUserPosts)
	http.ListenAndServe(":8080", mux)

}
