package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"

	"api_micro/models"
)

// Insert one deployment at a time in the DB
func insertOneDeploymentk(deploy models.Deployment) string {
	insertResult, err := collection.InsertOne(context.Background(), deploy)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)

	return fmt.Sprintf("%v",insertResult.InsertedID)
}

func CreateDeployment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	r.ParseForm()

	deployment := models.Deployment{
		ServiceName:  r.FormValue("servicename"),
		ReplicaCount: r.FormValue("replicacount"),
		ImageName:    r.FormValue("imagename"),
		Repo:         r.FormValue("repo"),
		Tag:          r.FormValue("imagetag"),
		PullPolicy:   r.FormValue("pull"),
		LbType:       r.FormValue("lb"),
		ExternalPort: r.FormValue("externalport"),
		InternalPort: r.FormValue("internalport"),
	}
	_ = json.NewDecoder(r.Body).Decode(&deployment)

	id := insertOneDeploymentk(deployment)
	json.NewEncoder(w).Encode(deployment)

	_, err := fmt.Fprintln(w, fmt.Sprintf("Wrote values to database %v with ID: %s", deployment, id))
	if err != nil {
		fmt.Fprintln(w, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetAllDeployments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllDeployments()

	json.NewEncoder(w).Encode(payload)
}

func getAllDeployments() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
		results = append(results, result)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

// DeleteTask delete one task route
func DeleteDeployment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneDeployment(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	// json.NewEncoder(w).Encode("Task not found")

}

// DeleteAllTask delete all tasks route
func DeleteAllDeployments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	count := deleteAllDeployments()
	json.NewEncoder(w).Encode(count)
	// json.NewEncoder(w).Encode("Task not found")

}

func deleteOneDeployment(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

// delete all the tasks from the DB
func deleteAllDeployments() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}