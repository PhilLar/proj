// // package main

// // import (
// // 	// "database/sql"
// // 	// "fmt"
// // 	"github.com/PhilLar/proj/pldb"

// // 	_ "github.com/lib/pq"
// // )

// // const (
// // 	host     = "localhost"
// // 	port     = 5432
// // 	user     = "philipp"
// // 	password = "qwedf12"
// // 	dbname   = "TestDB"
// // )

// // func main() {
// // 	pldb.Hello()
// // 	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// // 	// db, err := sql.Open("postgres", psqlconn)
// // 	// CheckError(err)

// // 	// defer db.Close()

// // 	// // insert
// // 	// // hardcoded
// // 	// insertStmt := `insert into "regions"("title") values('John')`
// // 	// _, e := db.Exec(insertStmt)
// // 	// CheckError(e)

// // 	// // dynamic
// // 	// insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
// // 	// _, e = db.Exec(insertDynStmt, "Jane", 2)
// // 	// CheckError(e)
// // }

// // func CheckError(err error) {
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // }

// package main

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type Profile struct {
// 	Name    string
// 	Hobbies []string
// }

// func main() {
// 	http.HandleFunc("/", foo)
// 	http.ListenAndServe(":3000", nil)
// }

// func foo(w http.ResponseWriter, r *http.Request) {
// 	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

// 	js, err := json.Marshal(profile)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(js)
// }




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

func main() {
	pldb.Hello()
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	// insert
	// hardcoded
	insertStmt := `insert into "regions"("title") values('John')`
	_, e := db.Exec(insertStmt)
	CheckError(e)

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

// package main

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type Profile struct {
// 	Name    string
// 	Hobbies []string
// }

// func main() {
// 	http.HandleFunc("/", foo)
// 	http.ListenAndServe(":3000", nil)
// }

// func foo(w http.ResponseWriter, r *http.Request) {
// 	profile := Profile{"Alex", []string{"snowboarding", "programming"}}

// 	js, err := json.Marshal(profile)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	(w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 	w.Write(js)
// }
