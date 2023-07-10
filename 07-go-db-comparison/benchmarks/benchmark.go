package benchmarks

import (
	"context"
	"database/sql"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	sqlc "go-db-comparison/benchmarks/sqlc_generated"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func init() {
	var err error
	// Opening a database connection.
	db, err = sql.Open("mysql", "root:rootuser@tcp(localhost:3306)/todo?multiStatements=true&parseTime=true")
	if err != nil {
		panic(err)
	}

	//sqlx connection using existing db connection
	dbx = sqlx.NewDb(db, "mysql")

	//sqlc connection using existing db connection
	dbc = sqlc.New(db)

	//gorm connection using existing db connection
	gdb, err = gorm.Open(mysql.New(mysql.Config{Conn: db}))
	if err != nil {
		panic(err)
	}
}

var (
	gdb *gorm.DB
	db  *sql.DB
	dbx *sqlx.DB
	dbc *sqlc.Queries
)

func setup() {
	clear()
	table := `CREATE TABLE students (
	id bigint NOT NULL AUTO_INCREMENT,
		fname varchar(50) not null,
		lname varchar(50) not null,
		date_of_birth datetime not null,
		email varchar(50) not null,
		address varchar(50) not null,
		gender varchar(50) not null,
		PRIMARY KEY (id)
	);`
	_, err := db.Exec(table)
	if err != nil {
		panic(err)
	}

	//inserting records
	_, err = db.Exec(records)
	if err != nil {
		panic(err)
	}
}

func clear() {
	_, err := db.Exec(`DROP TABLE IF EXISTS students;`)
	if err != nil {
		panic(err)
	}
}

type Student struct {
	ID          int64
	Fname       string
	Lname       string
	DateOfBirth time.Time `db:"date_of_birth"`
	Email       string
	Address     string
	Gender      string
}

func DbSqlQueryStudentWithLimit(limit int) {
	var students []Student
	rows, err := db.Query("SELECT * FROM students limit ?", limit)
	if err != nil {
		log.Fatalf("DbSqlQueryStudentWithLimit %d %v", limit, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.ID, &s.Fname, &s.Lname, &s.DateOfBirth, &s.Email, &s.Address, &s.Gender); err != nil {
			log.Fatalf("DbSqlQueryStudentWithLimit %d %v", limit, err)
		}
		students = append(students, s)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("DbSqlQueryStudentWithLimit %d %v", limit, err)
	}
}

func SqlxQueryStudentWithLimit(limit int) {
	var students []Student
	err := dbx.Select(&students, "SELECT * FROM students LIMIT ?", limit)
	if err != nil {
		log.Fatalf("SqlxQueryStudentWithLimit %d %v", limit, err)
	}
}

func SqlcQueryStudentWithLimit(limit int) {
	_, err := dbc.FetchStudents(context.Background(), int32(limit))
	if err != nil {
		log.Fatalf("SqlcQueryStudentWithLimit %d %v", limit, err)
	}
}

func GormQueryStudentWithLimit(limit int) {
	var students []Student
	if err := gdb.Limit(limit).Find(&students).Error; err != nil {
		log.Fatalf("GormQueryStudentWithLimit %d %v", limit, err)
	}
}
