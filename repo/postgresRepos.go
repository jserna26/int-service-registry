package repo

import (
	"database/sql"
	"log"

	"github.com/jserna26/SystemPOC/types"
	_ "github.com/lib/pq"
)

type PostgresSystemRepo struct {
	DbUrl      string
	connection *sql.DB
}

func (r *PostgresSystemRepo) CreateSystem(sys types.NewSystemType) (system types.SystemType, err error) {
	err = r.connection.QueryRow("INSERT INTO reg_system(system_name,description,active) VALUES($1,$2,$3) RETURNING SYSTEM_NAME, DESCRIPTION, ACTIVE;", &sys.Name, &sys.Description, string(sys.Status)).Scan(&system.Name, &system.Description, &system.Status)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (r *PostgresSystemRepo) GetAllSystems() (systems []types.SystemType, err error) {
	var system types.SystemType
	var rows *sql.Rows
	rows, err = r.connection.Query("SELECT * FROM reg_system;")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		if err = rows.Scan(&system.Name, &system.Description, &system.Status); err != nil {
			panic(err)
		}
		systems = append(systems, system)
	}
	return
}

func (r *PostgresSystemRepo) GetSystem(sysIn string) (sysOut types.SystemType, err error) {
	err = r.connection.QueryRow("SELECT * FROM reg_system WHERE system_name = $1;", sysIn).Scan(&sysOut.Name, &sysOut.Description, &sysOut.Status)
	if err != nil {
		return
	}
	return
}

func (r *PostgresSystemRepo) Connect(dburl string) (err error) {
	r.connection, err = sql.Open("postgres", dburl)
	return
}

func (r *PostgresSystemRepo) MustConnect(dburl string) {
	err := r.Connect(dburl)
	if err != nil {
		panic(err)
	}
}

func (r *PostgresSystemRepo) Disconnect() (err error) {
	return r.connection.Close()
}

func NewPostgresSystemRepo(dburl string) (r *PostgresSystemRepo, err error) {
	//Create new PostgresSystemRepo object, connect to DB and return.
	r = new(PostgresSystemRepo)
	err = r.Connect(dburl)
	if err != nil {
		log.Fatal(err)
	}
	return
}
