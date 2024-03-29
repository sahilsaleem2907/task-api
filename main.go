package main

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	Id              string `json:"id"`
	Caption         string `json:"caption"`
	imageURL        string `json:"imageURL"`
	PostedTimeStamp string `json:"timeStamp"`
}

type Users []User
type Posts []Post

var gid = "sahilsal" //GLOBAL INSTAGRAM ID USED FOR ALL OF THE FUNCTIONS

var lock sync.Mutex //USED FOR THREAD SAFETY 

//Function to implement PAGINATION
func Pagination(r *http.Request, FindOptions *options.FindOptions) (int64, int64) {
	if r.URL.Query().Get("page") != "" && r.URL.Query().Get("limit") != "" {
		page, _ := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
		limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
		if page == 1 {
			FindOptions.SetSkip(0)
			FindOptions.SetLimit(limit)
			return page, limit
		}

		FindOptions.SetSkip((page - 1) * limit)
		FindOptions.SetLimit(limit)
		return page, limit

	}
	FindOptions.SetSkip(0)
	FindOptions.SetLimit(0)
	return 0, 0
}

//Function to get a USER by ID
func getUsingId(w http.ResponseWriter, r *http.Request) {
	
	//THREAD SAFETY
	lock.Lock()
	defer lock.Unlock()
	
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.j7t2l.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	options := options.Find()
	Pagination(r, options)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	instagramDatabase := client.Database("instagram")
	usersCollection := instagramDatabase.Collection("users")

	//Code to retrieve by FILTER

	filterCursor, err := usersCollection.Find(ctx, bson.M{"id": gid})
	if err != nil {
		log.Fatal(err)
	}
	var usersFiltered []bson.M
	if err = filterCursor.All(ctx, &usersFiltered); err != nil {
		log.Fatal(err)
	}

	var id = ""
	var Name = ""
	var Email = ""
	var Pass = ""
	for _, usersFiltereds := range usersFiltered {
		id = usersFiltereds["id"].(string)
		Name = usersFiltereds["Name"].(string)
		Email = usersFiltereds["Email"].(string)
		Pass = usersFiltereds["Pass"].(string)
	}

	Users := []User{
		User{Id: id, Name: Name, Email: Email, Password: Pass},
	}
	fmt.Println("Endpoint hit:All users endpoint")
	json.NewEncoder(w).Encode(Users)
	
	time.Sleep(1 * time.Second)

}

//Function to get a POST by ID
func getPostUsingId(w http.ResponseWriter, r *http.Request) {
	
	//THREAD SAFETY
	lock.Lock()
	defer lock.Unlock()
	
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.j7t2l.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	options := options.Find()
	Pagination(r, options)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	instagramDatabase := client.Database("instagram")
	postsCollection := instagramDatabase.Collection("posts")

	//Code to retrieve by FILTER

	filterCursor, err := postsCollection.Find(ctx, bson.M{"id": gid})
	if err != nil {
		log.Fatal(err)
	}
	var postsFiltered []bson.M

	if err = filterCursor.All(ctx, &postsFiltered); err != nil {
		log.Fatal(err)
	}

	var id = ""
	var caption = ""
	var ImageURL = ""
	var postedTimeStamp = ""
	for _, postsFiltereds := range postsFiltered {
		id = postsFiltereds["id"].(string)
		caption = postsFiltereds["caption"].(string)
		ImageURL = postsFiltereds["imageURL"].(string)
		postedTimeStamp = postsFiltereds["Posted TimeStamp"].(string)
	}

	Posts := []Post{
		Post{Id: id, Caption: caption, imageURL: ImageURL, PostedTimeStamp: postedTimeStamp},
	}
	fmt.Println("Endpoint hit:All posts endpoint")
	json.NewEncoder(w).Encode(Posts)
	
	time.Sleep(1 * time.Second)

}

//Function to get all POST'S for a particular ID
func getAllPostUsingId(w http.ResponseWriter, r *http.Request) {
	
	//THREAD SAFETY
	lock.Lock()
	defer lock.Unlock()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.j7t2l.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	options := options.Find()
	Pagination(r, options)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	instagramDatabase := client.Database("instagram")
	postsCollection := instagramDatabase.Collection("posts")

	//Code to retrieve by FILTER

	filterCursor, err := postsCollection.Find(ctx, bson.M{"id": gid})
	if err != nil {
		log.Fatal(err)
	}
	var postsFiltered []bson.M
	if err = filterCursor.All(ctx, &postsFiltered); err != nil {
		log.Fatal(err)
	}

	var id = ""
	var caption = ""
	var ImageURL = ""
	var postedTimeStamp = ""
	for _, postsFiltereds := range postsFiltered {
		id = postsFiltereds["id"].(string)
		caption = postsFiltereds["caption"].(string)
		ImageURL = postsFiltereds["imageURL"].(string)
		postedTimeStamp = postsFiltereds["Posted TimeStamp"].(string)
		post_all_1 := Posts{
			Post{Id: id, Caption: caption, imageURL: ImageURL, PostedTimeStamp: postedTimeStamp},
		}
		json.NewEncoder(w).Encode(post_all_1)
	}

	fmt.Println("Endpoint hit:All posts endpoint")
	
	time.Sleep(1 * time.Second)

}

//Function to CREATE A USER
func createUser(w http.ResponseWriter, r *http.Request) {
	
	//THREAD SAFETY
	lock.Lock()
	defer lock.Unlock()
	
	fmt.Fprintf(w, "Test POST endpoint worked")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.j7t2l.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	options := options.Find()
	Pagination(r, options)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	instagramDatabase := client.Database("instagram")
	usersCollection := instagramDatabase.Collection("users")

	pass := "sahil123"
	uEnc := b64.URLEncoding.EncodeToString([]byte(pass)) //Encoded the password using URL encoding for more security

	userResult, err := usersCollection.InsertOne(ctx, bson.D{
		{"id", "sahilsal"},
		{"Name", "sahil"},
		{"Email", "sahil@gmail.com"},
		{"Pass", uEnc},
	})

	// To decode we can use uDec, _ := b64.URLEncoding.DecodeString(uEnc)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userResult.InsertedID)
	
	time.Sleep(1 * time.Second)
}

//Function to CREATE A POST
func createPost(w http.ResponseWriter, r *http.Request) {
	
	//THREAD SAFETY
	lock.Lock()
	defer lock.Unlock()
	
	fmt.Fprintf(w, "Test POST endpoint worked")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://<username>:<password>@cluster0.j7t2l.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	options := options.Find()
	Pagination(r, options)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	instagramDatabase := client.Database("instagram")
	postsCollection := instagramDatabase.Collection("posts")

	currentTime := time.Now() //Used for time stamp
	postResult, err := postsCollection.InsertOne(ctx, bson.D{
		{"id", "sahilsal"},
		{"caption", "Sahil's first  picture"},
		{"imageURL", "https://www.appointy.com/black-friday/uploads/2020/08/Appointy-Logo1.svg"}, //Appointy Logo used as an example
		{"Posted TimeStamp", currentTime.Format("2006-January-02 15:04:05")},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(postResult.InsertedID)
	
	time.Sleep(1 * time.Second)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	
	//THREAD SAFETY
	lock.Lock()
	defer lock.Unlock()
	
	fmt.Fprintf(w, "Instagram API by Sahil Saleem for Appointy Tech Round")
	
	time.Sleep(1 * time.Second)
}

func handleRequests() {
	
	myRouter := mux.NewRouter().StrictSlash(true)
	getIdPath := "/users/" + gid
	getPostPathById := "/posts/" + gid
	getAllPostPathById := "/posts/users/" + gid
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc(getIdPath, getUsingId).Methods("GET")                 //GET USER USING GID http://localhost:10000/users/sahilsal
	myRouter.HandleFunc(getPostPathById, getPostUsingId).Methods("GET")       // GET ONE POST USING GID http://localhost:10000/posts/sahilsal
	myRouter.HandleFunc(getAllPostPathById, getAllPostUsingId).Methods("GET") // GET ALL POSTS USING GID http://localhost:10000/posts/users/sahilsal
	myRouter.HandleFunc("/users", createUser).Methods("POST")                 // CREATE A NEW USER http://localhost:10000/users
	myRouter.HandleFunc("/posts", createPost).Methods("POST")                 //CREATE A NEW POST http://localhost:10000/posts
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
