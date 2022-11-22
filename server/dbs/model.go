package dbs

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Pic        string             `json:"userImg"`
	Password   string             `json:"userPassword"`
	EmailId    string             `json:"userEmail"`
	UserName   string             `json:"userName"`
	ResetCode  uint64             `json:"resetCode"`
	Verified   bool               `json:"verified"`
	Links      map[string]string  `json:"links"`
	Deactivate bool               `json:"deactivate"`
}

type Admin struct {
	Id           primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Pic          string             `json:"pic"`
	AdminName    string             `json:"adminName"`
	Password     string             `json:"password"`
	EmailId      string             `json:"emailId"`
	ActivityTree []string           `json:"activityTree"`
}

type UserView struct {
	Pic      string `json:"userImg"`
	UserName string `json:"userName"`
}

type ServeModel struct {
	Deactivate bool               `json:"deactivate"`
	Links      map[string]string  `json:"links"`
	Pic        string             `json:"userImg"`
	Verified   bool               `json:"verified"`
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EmailId    string             `json:"userEmail"`
	UserName   string             `json:"userName"`
}

type Login struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Password string             `json:"password"`
	EmailId  string             `json:"userEmail"`
}

type NewLink struct {
	PlateForm string `json:"plateForm"`
	Link      string `json:"link"`
}

type Emailidonly struct {
	// there might be an error
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Emailid string             `json:"useremail"`
}

type AdminLogIn struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Key      string             `json:"key"`
	Password string             `json:"password"`
}
