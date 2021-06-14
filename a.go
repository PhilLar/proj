// package main

// import (
// 	// "database/sql"
// 	// "fmt"
// 	"github.com/PhilLar/proj/pldb"

// 	_ "github.com/lib/pq"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "philipp"
// 	password = "qwedf12"
// 	dbname   = "TestDB"
// )

// func main() {
// 	pldb.Hello()
// 	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	// db, err := sql.Open("postgres", psqlconn)
// 	// CheckError(err)

// 	// defer db.Close()

// 	// // insert
// 	// // hardcoded
// 	// insertStmt := `insert into "regions"("title") values('John')`
// 	// _, e := db.Exec(insertStmt)
// 	// CheckError(e)

// 	// // dynamic
// 	// insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
// 	// _, e = db.Exec(insertDynStmt, "Jane", 2)
// 	// CheckError(e)
// }

// func CheckError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PhilLar/proj/pldb"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "philipp"
	password = "qwedf12"
	dbname   = "TestDB"
)

type Farm struct {
	ID             int
	Title          string
	Specialization string
	HeadsOfAnimals int
	HeadsofCows    int
	Longtitude     float64
	Latitude       float64
	Address        string
	OF_type        string
	SAL            float64
	RegionID       float64
}

type Region struct {
	ID    int
	Title string
}

type InitData struct {
	Farms   []Farm
	Regions []Region
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	// insert
	// hardcoded

	// farms := []Farm{}
	// regions := []Region{}
	var data InitData

	rows, err := db.Query(`SELECT "id", "title", "specialization", "heads_of_animals", "heads_of_cows", "longitude", "latitude", "address", "of_type", "sal", "region_id" FROM "farms"`)
	CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var farm Farm
		pldb.Hello()
		err = rows.Scan(&farm.ID, &farm.Title, &farm.Specialization, &farm.HeadsOfAnimals, &farm.HeadsofCows, &farm.Longtitude, &farm.Latitude, &farm.Address, &farm.OF_type, &farm.SAL, &farm.RegionID)
		CheckError(err)

		data.Farms = append(data.Farms, farm)
		// fmt.Println(name, roll)
	}

	// fmt.Println(data.Farms)

	rows, err = db.Query(`SELECT "id", "title" FROM "regions"`)
	CheckError(err)
	defer rows.Close()

	for rows.Next() {
		var region Region

		err = rows.Scan(&region.ID, &region.Title)
		CheckError(err)

		data.Regions = append(data.Regions, region)
		// fmt.Println(name, roll)
	}

	CheckError(err)

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Write(js)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}