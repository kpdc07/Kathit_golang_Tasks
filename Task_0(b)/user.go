package controllers

import(
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"github.com/kpdc_07/mongo-golang/models"
)
type UserController struct{
	session *mgo.Session
}
func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")
	if !=bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}
	oid := bson.objectIdHex(id)
	u := models.User{}
	if err := uc.Session.Db("mongo-golang").C("users").FindId(){
		w.WriteHeader(404)
		return
	}
	uj,err := json.Marshal(u)
	if err != nil{
		fmt.Println(err)
	}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
fmt.Printf(w, "%s\n", uj)
}
func (uc UserController) CreateUser (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	u := models.User{}
	json.NewDecoder(r.body).Decode(&u)
	u.Id = bson.NewObjectId()
	uc.Session.DB("mongo-golang").C("users").Insert(u)
	uj, err := json.Marshal(u)
	if err != nil{
		fmt.Println(err)
	}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.Statuscreated0000)
fmt.Printf(w, "%s\n", uj)

}
func (uc UserController) DeleteUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.Byname("id")
	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)
	if err := us.Sessio.DB("mongo-golang").C("users").RemoveId(){
		w.WriteHeader(404)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid "\n")
}