package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gocarina/gocsv"
)

type MelbCouncilAddress struct {
	StreetNo string `json:"street_no"`
	Gisid    string `json:"gisid"`
	TheGeom  struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"the_geom"`
	StrName    string `json:"str_name"`
	Suburb     string `json:"suburb"`
	AddressPnt string `json:"address_pnt"`
	SuburbID   string `json:"suburb_id"`
	StreetID   string `json:"street_id"`
	Easting    string `json:"easting"`
	Northing   string `json:"northing"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
}
type CsvAddress struct {
	Number          string
	Gender          string
	Title           string
	GivenName       string
	MiddleInitial   string
	Surname         string
	StreetAddress   string
	City            string
	State           string
	ZipCode         string
	Country         string
	CountryFull     string
	EmailAddress    string
	Username        string
	Password        string
	TelephoneNumber string
	MothersMaiden   string
	Birthday        string
	CCType          string
	CCNumber        string
	CVV2            string
	CCExpires       string
	NationalID      string
	UPS             string
	Occupation      string
	Company         string
	Vehicle         string
	Domain          string
	BloodType       string
	Pounds          string
	Kilograms       string
	FeetInches      string
	Centimeters     string
	GUID            string
	Latitude        string
	Longitude       string
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		add := randomAddress()
		js, err := json.Marshal(add)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		// w.Write([]byte(add.StrName))
	})
	http.ListenAndServe(":443", r)
}

func loadAddresses() []string {
	jsonFile, err := os.Open("data/melbCouncil.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byt, _ := ioutil.ReadAll(jsonFile)

	var addresses []MelbCouncilAddress
	if err := json.Unmarshal(byt, &addresses); err != nil {
		panic(err)
	}
	var streetaddresses []string
	for _, client := range addresses {
		streetaddresses = append(streetaddresses, fmt.Sprintf("%s, %s %s, Victoria", client.AddressPnt, client.Suburb, "3000"))
	}
	return streetaddresses
}
func randomAddress() string {
	addresses := loadAddresses()
	// otherAddresses := fngToStreetAddress()
	// allAddresses := append(addresses, otherAddresses...)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	return addresses[rand.Intn(len(addresses))]
}

func fakeNameGeneratorWriteVic() {
	clientsFile, err := os.OpenFile("data/FakeNameGenerator.com_6c8dbb7d.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	clients := []*CsvAddress{}

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}
	var vicaddress []CsvAddress
	for _, client := range clients {
		if client.State == "VIC" {
			vicaddress = append(vicaddress, *client)
		}
	}

	csvContent, err := gocsv.MarshalString(&vicaddress)
	ioutil.WriteFile("data/vicaddress.csv", []byte(csvContent), 0644)
}

func fngToStreetAddress() []string {
	clientsFile, err := os.OpenFile("data/vicaddress.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	clients := []*CsvAddress{}
	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}
	var streetaddresses []string
	for _, client := range clients {
		stname := fmt.Sprintf("%s, %s %s, Victoria", client.StreetAddress, client.City, client.ZipCode)
		fmt.Println(stname)
		streetaddresses = append(streetaddresses, stname)
	}
	return streetaddresses
}
