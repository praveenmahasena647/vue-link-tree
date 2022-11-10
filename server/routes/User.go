package routes

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/praveenmahasena647/vue-link-tree/dbs"
	jwthandle "github.com/praveenmahasena647/vue-link-tree/jwt-handle"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	// all done
	var cursor, cursorErr = dbs.UserCollection.Find(context.TODO(), bson.M{"deactivate": false})
	if cursorErr != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("cursor Err")
	}
	var Users []dbs.UserView

	for cursor.Next(context.TODO()) {
		var user dbs.UserView
		cursor.Decode(&user)
		Users = append(Users, user)
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Users)
}

func getOneUser(w http.ResponseWriter, r *http.Request) {
	// all done
	var name = mux.Vars(r)["name"]

	var person *dbs.ServeModel
	// TODO: add deactivate
	var notFound = dbs.UserCollection.FindOne(context.TODO(), bson.M{"username": name}).Decode(&person)

	if notFound != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(false)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(person)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var newPerson *dbs.User = &dbs.User{}
	json.NewDecoder(r.Body).Decode(&newPerson)

	var person primitive.M

	dbs.UserCollection.FindOne(context.TODO(), bson.M{"$or": []bson.M{
		bson.M{"emailid": newPerson.EmailId},
		bson.M{"username": newPerson.UserName},
	}}).Decode(&person)

	if person != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode("user Aready Exist")
		return
	}
	var hashed, hashErr = bcrypt.GenerateFromPassword([]byte(newPerson.Password), 10)
	if hashErr != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(false)
		return
	}
	newPerson.Password = string(hashed)

	var done, doneErr = dbs.UserCollection.InsertOne(context.TODO(), newPerson)
	if doneErr != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(false)
		return
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(done.InsertedID)
}

func checkForUsers(w http.ResponseWriter, r *http.Request) {
	// if user Exist when you register
	var userName, userEmail string = r.Header.Get("X-UserName"), r.Header.Get("X-UserEmail")
	if userName != "" {
		var person bson.D
		// only cheaking for UserName here !imp
		dbs.UserCollection.FindOne(context.TODO(), bson.M{"username": userName}).Decode(&person)

		if person != nil {
			json.NewEncoder(w).Encode(true)
			return
		}
	}

	var person bson.D
	// only cheaking for UserEmail here !imp
	dbs.UserCollection.FindOne(context.TODO(), bson.M{"emailid": userEmail}).Decode(&person)

	if person != nil {
		json.NewEncoder(w).Encode(true)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(false)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var personData = map[string]string{}
	json.NewDecoder(r.Body).Decode(&personData)
	var person *dbs.Login
	var userNotFound = dbs.UserCollection.FindOne(context.TODO(), bson.M{"emailid": personData["userEmail"]}).Decode(&person)

	if userNotFound != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "person Not Found",
		})
		return
	}

	var passwordErr = bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(personData["userPassword"]))

	if passwordErr != nil {
		w.WriteHeader(401)
		return
	}

	var jwtString, jwtErr = jwthandle.CreateJwt(person)
	//
	if jwtErr != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "JWT Error",
		})
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(jwtString)
}

func viewDashBoard(w http.ResponseWriter, r *http.Request) {
	// all done
	var userId, err = jwthandle.DecodeJwt(r)

	if err != nil {
		w.WriteHeader(401)
		return
	}

	var id, _ = primitive.ObjectIDFromHex(userId)
	var personServe *dbs.ServeModel
	var personNotFound = dbs.UserCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&personServe)

	if personNotFound != nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(false)
		return
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(personServe)
}

func updateLink(w http.ResponseWriter, r *http.Request) {

	var userId, jwtErr = jwthandle.DecodeJwt(r)

	if jwtErr != nil {
		w.WriteHeader(401)
		return
	}

	var newLinks bson.M
	json.NewDecoder(r.Body).Decode(&newLinks)
	var id, _ = primitive.ObjectIDFromHex(userId)
	var _, doneErr = dbs.UserCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"links": newLinks}})

	if doneErr != nil {
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(200)
}

func changePassword(w http.ResponseWriter, r *http.Request) {
	var userId, jwtErr = jwthandle.DecodeJwt(r)

	if jwtErr != nil {
		w.WriteHeader(401)
		return
	}
	var newPassword = map[string]string{}
	json.NewDecoder(r.Body).Decode(&newPassword)
	var id, _ = primitive.ObjectIDFromHex(userId)
	var person *dbs.Login
	var personNotFound = dbs.UserCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&person)

	if personNotFound != nil {
		w.WriteHeader(404)
		return
	}

	var notErr = bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(newPassword["oldPassword"]))
	if notErr != nil {
		w.WriteHeader(400)
		return
	}

	var hashed, hashErr = bcrypt.GenerateFromPassword([]byte(newPassword["newPassword"]), 10)
	if hashErr != nil {
		w.WriteHeader(406)
		return
	}
	var _, notDone = dbs.UserCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"password": string(hashed)}})

	if notDone != nil {
		w.WriteHeader(406)
		return
	}
	w.WriteHeader(205)
}

func active(w http.ResponseWriter, r *http.Request) {
	var userId, jwtErr = jwthandle.DecodeJwt(r)

	if jwtErr != nil {
		w.WriteHeader(401)
		return
	}
	var id, _ = primitive.ObjectIDFromHex(userId)
	//
	var activeStatus bson.M
	//
	json.NewDecoder(r.Body).Decode(&activeStatus)
	var _, doneErr = dbs.UserCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{"deactivate": activeStatus["deactivate"]}})

	if doneErr != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(200)
}

func getMail(w http.ResponseWriter, r *http.Request) {
	var userId, jwtErr = jwthandle.DecodeJwt(r)
	if jwtErr != nil {
		w.WriteHeader(404)
		return
	}
	var ID, _ = primitive.ObjectIDFromHex(userId)
	var emailOnly *dbs.Emailidonly

	var notFound = dbs.UserCollection.FindOne(context.TODO(), bson.M{"_id": ID}).Decode(&emailOnly)
	if notFound != nil {
		w.WriteHeader(404)
		return
	}

	// this looks cursed but i like it
	var tcpCon, tcpErr = net.Dial("tcp", ":6969")

	if tcpErr != nil {
		w.WriteHeader(500)
		return
	}
	var jsml, _ = json.Marshal(emailOnly)
	var _, notSent = tcpCon.Write(jsml)

	if notSent != nil {
		w.WriteHeader(500)
		return
	}

	var data, dataErr = ioutil.ReadAll(tcpCon)
	defer tcpCon.Close()

	if dataErr != nil {
		w.WriteHeader(500)
		return
	}

	if string(data) == "not-done" {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
}

func verify(w http.ResponseWriter, r *http.Request) {
	var id = mux.Vars(r)["id"]
	var mongoId, idErr = primitive.ObjectIDFromHex(id)
	if idErr != nil {
		w.WriteHeader(400)
		w.Write([]byte("can't verify"))
		return
	}

	var _, notDone = dbs.UserCollection.UpdateOne(context.TODO(), bson.M{"_id": mongoId}, bson.M{"$set": bson.M{"verified": true}})
	if notDone != nil {
		w.WriteHeader(400)
		w.Write([]byte("can't verify"))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("done"))
}
