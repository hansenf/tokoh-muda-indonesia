package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/shopspring/decimal"
)

var db *gorm.DB
var err error

// Mahasiswa is a representation of a mahasiswa
type Mahasiswa struct {
	ID			int		`form:"id" json:"id"`
	Username	string	`form:"username" json:"username"`
	Password	string	`form:"password" json:"password"`
	/* ID    		int			`form:"id" json:"id"`
	Username  	string		`form:"code" json:"code"`
	Password  	string		`form:"name" json:"name"`
	ID    		int			`form:"id" json:"id"`
	Username  	string		`form:"code" json:"code"`
	Password  	string		`form:"name" json:"name"`
	 */
	//Price decimal.Decimal `form:"price" json:"price" sql:"type:decimal(16,2);"`
}

// Result is an array of mahasiswa
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Main
func main() {
	db, err = gorm.Open("mysql", "root:@/tmi?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection established")
	}

	db.AutoMigrate(&Mahasiswa{})
	handleRequests()
}

func handleRequests() {
	log.Println("Start the development server at http://127.0.0.1:9999")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := Result{Code: 404, Message: "Method not found"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	myRouter.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		res := Result{Code: 403, Message: "Method not allowed"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/mahasiswas", createMahasiswa).Methods("POST")
	myRouter.HandleFunc("/api/mahasiswas", getMahasiswas).Methods("GET")
	myRouter.HandleFunc("/api/mahasiswas/{id}", getMahasiswa).Methods("GET")
	myRouter.HandleFunc("/api/mahasiswas/{id}", updateMahasiswa).Methods("PUT")
	myRouter.HandleFunc("/api/mahasiswas/{id}", deleteMahasiswa).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func createMahasiswa(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var mahasiswa Mahasiswa
	json.Unmarshal(payloads, &mahasiswa)

	db.Create(&mahasiswa)

	res := Result{Code: 200, Data: mahasiswa, Message: "Success create mahasiswa"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getMahasiswas(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get mahasiswas")

	mahasiswas := []Mahasiswa{}
	db.Find(&mahasiswas)

	res := Result{Code: 200, Data: mahasiswas, Message: "Success get mahasiswas"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func getMahasiswa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mahasiswaID := vars["id"]

	var mahasiswa Mahasiswa

	db.First(&mahasiswa, mahasiswaID)

	res := Result{Code: 200, Data: mahasiswa, Message: "Success get mahasiswa"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func updateMahasiswa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mahasiswaID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var mahasiswaUpdates Mahasiswa
	json.Unmarshal(payloads, &mahasiswaUpdates)

	var mahasiswa Mahasiswa
	db.First(&mahasiswa, mahasiswaID)
	db.Model(&mahasiswa).Updates(mahasiswaUpdates)

	res := Result{Code: 200, Data: mahasiswa, Message: "Success update mahasiswa"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func deleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mahasiswaID := vars["id"]

	var mahasiswa Mahasiswa

	db.First(&mahasiswa, mahasiswaID)
	db.Delete(&mahasiswa)

	res := Result{Code: 200, Message: "Success delete mahasiswa"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}