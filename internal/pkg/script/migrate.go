package script

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"log"
	"shifo-backend-website/internal/pkg/repository/postgres"
)

// ErrHelp provides context that help was given.
var ErrHelp = errors.New("provided help")

type Scheme struct {
	Index       int
	Description string
	Query       string
}

var scheme = []Scheme{
	{
		Index:       1,
		Description: "Create table: users.",
		Query: `
				CREATE TABLE IF NOT EXISTS users(
				    id         uuid primary key not null ,
  					first_name varchar,
  					last_name  varchar,
  					username   varchar (150),
  					password   varchar,
  					status     user_statuses,
  					gmail      varchar,
  					created_at timestamp default now(),
  					deleted_at timestamp,
  					updated_at timestamp,
  					updated_by uuid references users(id),
  					created_by uuid references users(id),
  					deleted_by uuid references users(id)
				);
			`,
	},
	{
		Index:       2,
		Description: "Create table: loggers",
		Query: `
				CREATE TABLE IF NOT EXISTS loggers(
				     id serial primary key,
   					 created_at timestamp default now(),
   					 data jsonb,
   					 method text,
   					 action text
				);
			`,
	},
	{
		Index:       3,
		Description: "Create table: doctors",
		Query: `
			CREATE TABLE IF NOT EXISTS doctors(
			    id              uuid primary key not null ,
   				first_name      varchar,
   				last_name       varchar,
   				specialty_id    uuid references specialties(id),
   				file_link       varchar,
   				work_experience text,
   				workplace_id    uuid references workplaces(id),
   				work_price      varchar,
   				start_work      timestamp,
   				end_work        timestamp,
   				created_at      timestamp default now(),
   				deleted_at      timestamp,
   				updated_at      timestamp,
   				updated_by uuid references users(id),
   				created_by uuid references users(id),
   				deleted_by uuid references users(id)
			);
		`,
	},
	{
		Index:       4,
		Description: "Create table: specialties",
		Query: `
			CREATE TABLE IF NOT EXISTS specialties(
			    id         uuid primary key not null ,
    			name       varchar not null,
    			created_at timestamp default now(),
    			deleted_at timestamp,
    			updated_at timestamp,
    			updated_by uuid references users(id),
    			created_by uuid references users(id),
    			deleted_by uuid references users(id)
			);
		`,
	},
	{
		Index:       5,
		Description: "Create table: workplaces",
		Query: `
			CREATE TABLE IF NOT EXISTS workplaces(
			    id         uuid primary key not null ,
    			name       varchar not null,
    			address    varchar,
    			lat        float,
    			long       float,
    			created_at timestamp default now(),
    			deleted_at timestamp,
    			updated_at timestamp,
    			updated_by uuid references users(id),
    			created_by uuid references users(id),
    			deleted_by uuid references users(id)
			);
		`,
	},
	{
		Index:       6,
		Description: "Create user with username:admin, password: 1",
		Query: `
				INSERT INTO users (id,first_name,last_name, username, gmail, password, status) SElECT 'bfab8727-f6b9-48af-abcc-cb33307f0157','admin','admin','admin', 'shakke.gmail.com', '$2a$09$p71tEyRUhvkI8RWacTjCv.VLp51rUkUZnU8ScQtVb01ElxLIT8PUG','at_work' WHERE NOT EXISTS (SELECT id FROM users WHERE id = 'bfab8727-f6b9-48af-abcc-cb33307f0157');
			`,
	},
	//{
	//	Index:       7,
	//	Description: "Create Enum for user_status",
	//	Query: `
	//			DROP TYPE  IF EXISTS user_statuses;
	//			CREATE TYPE user_statuses AS ENUM ('at_work', 'off_work');
	//	`,
	//},

	//{
	//	Index:       11,
	//	Description: "Create table: requests",
	//	Query: `
	//		CREATE TABLE IF NOT EXISTS requests(
	//             id            text primary key not null ,
	//			 title         text,
	//			 description   text,
	//			 email         text,
	//			 phone         text,
	//			 created_at    timestamp default now(),
	//			 deleted_at    timestamp,
	//			 updated_at    timestamp,
	//			 updated_by    text references users(id),
	//			 created_by    text references users(id),
	//			 deleted_by    text references users(id)
	//		);
	//		`,
	//},
	//{
	//	Index:       12,
	//	Description: "Create table: request_files",
	//	Query: `
	//		CREATE TABLE IF NOT EXISTS request_files(
	//		    id            text primary key not null ,
	//			request_id    text not null references requests(id),
	//			file_link     text,
	//			type          integer,
	//			created_at    timestamp default now(),
	//			deleted_at    timestamp,
	//			updated_at    timestamp,
	//			updated_by    text references users(id),
	//			created_by    text references users(id),
	//			deleted_by    text references users(id)
	//		);
	//	`,
	//},
}

// Migrate creates the scheme in the database.
func Migrate(db *postgres.Database) {
	for _, s := range scheme {
		if _, err := db.Query(s.Query); err != nil {
			log.Fatalln("migrate error", err)
		}
	}
}

func MigrateUP(db *postgres.Database) {
	var (
		version int
		dirty   bool
		er      *string
	)
	err := db.QueryRow("SELECT version,dirty,error FROM schema_migrations").Scan(&version, &dirty, &er)
	if err != nil {
		if err.Error() == `ERROR: relation "schema_migrations" does not exist (SQLSTATE=42P01)` {
			if _, err = db.Exec(`
										CREATE TABLE IF NOT EXISTS schema_migrations (version int not null,dirty bool not null ,error text);
										DELETE FROM schema_migrations;
										INSERT INTO schema_migrations (version, dirty) values (0,false);
								`); err != nil {
				log.Fatalln("migrate schema_migrations create error", err)
			}
			version = 0
			dirty = false
		} else {
			log.Fatalln("migrate schema_migrations scan: ", err)
		}

	}

	if dirty {
		for _, v := range scheme {
			if v.Index == version {
				if _, err = db.Exec(v.Query); err != nil {
					if _, err = db.Exec(fmt.Sprintf(`UPDATE schema_migrations SET error = '%s'`, err.Error())); err != nil {
						log.Fatalln("migrate error", err)
					}
					log.Fatalln(fmt.Sprintf("migrate error version: %d", version), err)
				}
				if _, err = db.Exec(fmt.Sprintf(`UPDATE schema_migrations SET dirty = false, error = null`)); err != nil {
					log.Fatalln("migrate error", err)
				}
			}
		}
	}

	for _, s := range scheme {
		if s.Index > version {
			if _, err = db.Exec(s.Query); err != nil {
				if _, err = db.Exec(fmt.Sprintf(`UPDATE schema_migrations SET error = '%s', version = %d, dirty = true`, err.Error(), s.Index)); err != nil {
					log.Fatalln("migrate error", err)
				}
				log.Fatalln(fmt.Sprintf("migrate error version: %d", s.Index), err)
			}
			if _, err = db.Exec(fmt.Sprintf(`UPDATE schema_migrations SET version = %d`, s.Index)); err != nil {
				log.Fatalln("migrate error", err)
			}
		}
	}
}

//// MigrationsUp  for migration to database
//func MigrationsUp(DBUsername, DBPassword, DBPort, DBName string) {
//	url := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", DBUsername, DBPassword, DBPort, DBName)
//
//	m, err := migrate.New("file://internal/pkg/migrations", url)
//	if err != nil {
//		log.Fatal("error in creating migrations: ", err.Error())
//	}
//
//	if err := m.Up(); err != nil {
//		if !strings.Contains(err.Error(), "no change") {
//			log.Println("Error in migrating ", err.Error())
//		}
//	}
//}
