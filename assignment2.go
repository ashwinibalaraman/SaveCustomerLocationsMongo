package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
)

type ReqParameters struct {
	Name    string
	Address string
	City    string
	State   string
	Zip     string
}
type Coordinate struct {
	Lat string `bson:"lat"`
	Lng string `bson:"lng"`
}
type ResParameters struct {
	Id      bson.ObjectId `bson:"_id"`
	Name    string        `bson:"name"`
	Address string        `bson:"address"`
	City    string        `bson:"city"`
	State   string        `bson:"state"`
	Zip     string        `bson:"zip"`
	Coord   Coordinate    `bson:"coordinate"`
}

type PutParameters struct {
	Address string     `bson:"address"`
	City    string     `bson:"city"`
	State   string     `bson:"state"`
	Zip     string     `bson:"zip"`
	Coord   Coordinate `bson:"coordinate"`
}

type PutReqParameters struct {
	Address string
	City    string
	State   string
	Zip     string
}

var Url string

func main() {
	//Url = "localhost"
	Url = "mongodb://ashwini:Bangalore1@ds045064.mongolab.com:45064/cmpe273"
	mux := httprouter.New()

	mux.GET("/locations/:location_id", getLocations)
	mux.POST("/locations", postLocations)
	mux.PUT("/locations/:location_id", putLocations)
	mux.DELETE("/locations/:location_id", deleteLocations)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()

}

func databaseInsert(ResParams ResParameters) error {
	/*	uri := os.Getenv("MONGOHQ_URL")
		if uri == "" {
			fmt.Println("no connection string provided")
			os.Exit(1)
		}*/
	//URL := "mongodb://ashwini:Bangalore1@ds045064.mongolab.com:45064/cmpe273"
	sess, err := mgo.Dial(Url)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("cmpe273").C("AddressBook")

	doc := ResParams
	err = collection.Insert(doc)

	return err
}

func databaseUpdate(PutParams PutParameters, id string) (ResParameters, error) {
	/*	uri := os.Getenv("MONGOHQ_URL")
		if uri == "" {
			fmt.Println("no connection string provided")
			os.Exit(1)
		}*/
	//URL := "mongodb://ashwini:Bangalore1@ds045064.mongolab.com:45064/cmpe273"
	sess, err := mgo.Dial(Url)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("cmpe273").C("AddressBook")

	//doc := PutParams

	change := bson.M{"$set": bson.M{"address": PutParams.Address, "city": PutParams.City, "state": PutParams.State, "zip": PutParams.Zip, "coord": PutParams.Coord}}
	err = collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, change)
	data := ResParameters{}
	err = collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Select(bson.M{}).One(&data)

	return data, err
}

func getLocations(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	sess, err := mgo.Dial(Url)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("cmpe273").C("AddressBook")
	data := ResParameters{}
	id := p.ByName("location_id")
	fmt.Println("id:", id)
	err = collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).Select(bson.M{}).One(&data)

	if err != nil {
		if err.Error() == "not found" {
			rw.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(rw, http.StatusText(http.StatusNotFound))
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, err.Error())
		}
	} else {
		b, _ := json.Marshal(data)
		fmt.Println(string(b))
		fmt.Fprintf(rw, string(b))
		//dec.Decode(&data)
		//fmt.Println(data)
		//json.NewDecoder(resp).Decode(&data)

	}
}

func postLocations(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	ReqParams := ReqParameters{}

	json.NewDecoder(req.Body).Decode(&ReqParams)

	addressString := ReqParams.Address + "+" + ReqParams.City + "+" + ReqParams.State
	addressString = strings.Replace(addressString, " ", "+", -1)
	fmt.Println(addressString)
	googleApiUrl := "http://maps.google.com/maps/api/geocode/json?address=" + addressString
	resp, _ := http.Get(googleApiUrl)
	var data interface{}
	_ = json.NewDecoder(resp.Body).Decode(&data)
	results := data.(map[string]interface{})["results"]
	result := results.([]interface{})[0]
	geometry := result.(map[string]interface{})["geometry"]
	loc := geometry.(map[string]interface{})["location"]
	lat := loc.(map[string]interface{})["lat"]
	lng := loc.(map[string]interface{})["lng"]
	latitude := strconv.FormatFloat(lat.(float64), 'E', -1, 64)
	longitude := strconv.FormatFloat(lng.(float64), 'E', -1, 64)
	fmt.Println(lat)
	fmt.Println(lng)

	var ResParams ResParameters
	ResParams = ResParameters{
		Name:    ReqParams.Name,
		Address: ReqParams.Address,
		City:    ReqParams.City,
		State:   ReqParams.State,
		Zip:     ReqParams.Zip,
	}
	CoordParams := Coordinate{}
	CoordParams.Lat = latitude
	CoordParams.Lng = longitude
	ResParams.Coord = CoordParams

	//s1 := rand.NewSource(time.Now().UnixNano())
	//	r1 := rand.New(s1)
	//x := r1.Intn(1000)
	ResParams.Id = bson.NewObjectId()
	//fmt.Println("object name:", ResParams.Id)

	err := databaseInsert(ResParams)
	/*greeting, _ := json.Marshal(ResParams)
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(rw, "%s\n", greeting)*/

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, err.Error())
	} else {
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, http.StatusText(http.StatusOK))
		b, _ := json.Marshal(ResParams)
		fmt.Println(string(b))
		fmt.Fprintf(rw, string(b))
	}
}

func putLocations(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	PutReqParams := PutReqParameters{}
	fmt.Println(req.Body)
	json.NewDecoder(req.Body).Decode(&PutReqParams)

	addressString := PutReqParams.Address + "+" + PutReqParams.City + "+" + PutReqParams.State
	addressString = strings.Replace(addressString, " ", "+", -1)
	fmt.Println(addressString)
	googleApiUrl := "http://maps.google.com/maps/api/geocode/json?address=" + addressString
	resp, _ := http.Get(googleApiUrl)
	var data interface{}
	_ = json.NewDecoder(resp.Body).Decode(&data)
	results := data.(map[string]interface{})["results"]
	result := results.([]interface{})[0]
	geometry := result.(map[string]interface{})["geometry"]
	loc := geometry.(map[string]interface{})["location"]
	lat := loc.(map[string]interface{})["lat"]
	lng := loc.(map[string]interface{})["lng"]
	latitude := strconv.FormatFloat(lat.(float64), 'E', -1, 64)
	longitude := strconv.FormatFloat(lng.(float64), 'E', -1, 64)
	fmt.Println(lat)
	fmt.Println(lng)

	var PutParams PutParameters
	PutParams = PutParameters{
		Address: PutReqParams.Address,
		City:    PutReqParams.City,
		State:   PutReqParams.State,
		Zip:     PutReqParams.Zip,
	}
	CoordParams := Coordinate{}
	CoordParams.Lat = latitude
	CoordParams.Lng = longitude
	PutParams.Coord = CoordParams

	id := p.ByName("location_id")
	data, err := databaseUpdate(PutParams, id)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, err.Error())
	} else {
		rw.WriteHeader(http.StatusCreated)
		fmt.Fprintf(rw, http.StatusText(http.StatusCreated))

		b, _ := json.Marshal(data)
		fmt.Println(string(b))
		fmt.Fprintf(rw, string(b))
		//dec.Decode(&data)
		//fmt.Println(data)
		//json.NewDecoder(resp).Decode(&data)
	}
	/*greeting, _ := json.Marshal(ResParams)
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(rw, "%s\n", greeting)*/
}

func deleteLocations(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	sess, err := mgo.Dial(Url)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})
	collection := sess.DB("cmpe273").C("AddressBook")

	id := p.ByName("location_id")
	fmt.Println("id:", id)

	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, err.Error())
	} else {
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, http.StatusText(http.StatusOK))
	}
	//fmt.Fprintf(rw, string(b))
}
