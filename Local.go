package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type location struct {
	Tid  int     `json:"tid"`
	Lat  float64 `json:"Lat"`
	Long float64 `json:"Long"`
	T    string  `json:"t"`
}

func handler_GetLocationByID(w http.ResponseWriter, r *http.Request) {

	keys := r.URL.Query()["key"]
	tid_from_request := keys[0]
	fmt.Println("\nKEY:")
	fmt.Println(tid_from_request)

	db, err_connection := sql.Open("mysql", "user1:123@tcp(127.0.0.1:3306)/test")
	if err_connection != nil {
		log.Fatal(err_connection)
		fmt.Println("Error")
	}
	defer db.Close()
	a := location{}

	rows, err_query := db.Query("select * from TRIPLOCATIONS where tid=?", tid_from_request)
	if err_query != nil {
		log.Fatal(err_query)
	}

	for rows.Next() {

		err := rows.Scan(&a.Tid, &a.Lat, &a.Long, &a.T)
		if err != nil {
			log.Fatal(err)
		}
		sLat := fmt.Sprintf("%f", a.Lat)
		sLong := fmt.Sprintf("%f", a.Long)

		s := string("TID: " + tid_from_request + " " + sLat + " " + sLong + " " + a.T)

		res := s
		json.NewEncoder(w).Encode(res)
	}

	err_rows := rows.Err()
	if err_rows != nil {
		log.Fatal(err_rows)
	}

}

func handler_UpdateLocationByID(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var a location
	if err := decoder.Decode(&a); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", a)

	db, err_connection := sql.Open("mysql", "user1:123@tcp(127.0.0.1:3306)/test")
	if err_connection != nil {
		log.Fatal(err_connection)
		fmt.Println("Error")
	}
	defer db.Close()
	_, err_query := db.Query("insert into TRIPLOCATIONS VALUES (?,?,?,?)", a.Tid, a.Lat, a.Long, a.T)
	if err_query != nil {
		log.Fatal(err_query)
	}

}

func main() {

	http.HandleFunc("/GetLocation", handler_GetLocationByID)
	http.HandleFunc("/UpdateLocation", handler_UpdateLocationByID)

	fmt.Printf("Starting server for testing HTTP ...\n")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}

}
