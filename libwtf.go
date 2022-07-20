package lib

// https://zetcode.com/golang/mysql/

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"

    "strings"
    "strconv"
)

const DB_USER = "userwtf"
const DB_PWD = "password"
const DB_IP = "128.52.139.168"
const DB_PORT = "3306"
const DB_NAME = "dbwtf"
const DB_TABLE = "records"

type Attribute struct {
	Key string
	Value string
}

type Record struct {
	Severity uint8
	Attributes []Attribute
}

type Dag struct {
	// TODO
}

func MakeRecord(severity uint8, attributes []Attribute) Record {
	record := Record{
		Severity: severity,
		Attributes: attributes,
	}
	fmt.Printf("Made record=%+v\n", record)
	return record
}

func Emit(record Record) int {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PWD+"@tcp("+DB_IP+":"+DB_PORT+")/"+DB_NAME)
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    severity := strconv.FormatUint(uint64(record.Severity), 10)

    keys := []string{"severity"}
    values := []string{severity}
    for _, attribute := range record.Attributes {
    	fmt.Printf("attribute=%+v\n", attribute)
    	key := attribute.Key
    	value := attribute.Value

    	keys = append(keys, key)
    	values = append(values, value)
    }
    keystring := strings.Join(keys, ", ")
    valuestring := strings.Join(values, ", ")

    
    sql := "INSERT INTO "+DB_TABLE+"("+keystring+") VALUES ("+valuestring+")"
    fmt.Printf("sql=%+v\n", sql)
    res, err := db.Exec(sql)

    if err != nil {
        panic(err.Error())
    }

    lastId, err := res.LastInsertId()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("The last inserted row id: %d\n", lastId)

	return 0
}

func Collect(attributes []Attribute) []Record {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PWD+"@tcp("+DB_IP+":"+DB_PORT+")/"+DB_NAME)
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    filters := make([]string, 0)
    for _, attribute := range attributes {
    	filter := attribute.Key+" = '"+attribute.Value+"'"
    	filters = append(filters, filter)
    }
    condition := strings.Join(filters, ", ")
    if condition == "" {
    	condition = "TRUE"
    }

    res, err := db.Query("SELECT * FROM "+DB_TABLE+" WHERE "+condition)

    defer res.Close()

    if err != nil {
        log.Fatal(err)
    }

    // for res.Next() {

    //     var city City
    //     err := res.Scan(&city.Id, &city.Name, &city.Population)

    //     if err != nil {
    //         log.Fatal(err)
    //     }

    //     fmt.Printf("%v\n", city)
    // }

    // TODO

	return make([]Record, 0)
}

func Analyze(records []Record) Dag {
	if len(records) == 0 {
		fmt.Println("No records")
		return Dag{}
	}
	// TODO
	return Dag{}
}

// combines Collect() and Analyze()
func WheresTheFault(attributes []Attribute) Dag {
	records := Collect(attributes)
	dag := Analyze(records)
	return dag
}

func RegisterKeys(keys []string) int {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PWD+"@tcp("+DB_IP+":"+DB_PORT+")/"+DB_NAME)
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

	for _, key := range keys {

		sql := "REG "+key
	    res, err := db.Exec(sql)
	    _ = res

	    if err != nil {
	        panic(err.Error())
	    }

	    // lastId, err := res.LastInsertId()

	    // if err != nil {
	    //     log.Fatal(err)
	    // }

	    // fmt.Printf("The last inserted row id: %d\n", lastId)
	}

	return 0
}
