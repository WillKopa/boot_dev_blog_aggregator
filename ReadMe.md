### To generate and revert migrations use these commands from the `/sql/schema` directory
`goose postgres "postgres://postgres:postgres@localhost:5432/gator" up`

`goose postgres "postgres://postgres:postgres@localhost:5432/gator" down`

### To generate new database files use this command from the projects root
`sqlc generate`

### To start and stop the server use
`sudo service postgresql start`

`sudo service postgresql stop`

### To inspect the server use
- `sudo -u postgres psql`
- Then connect to the database using `\c gator`
- `\dt` will list all the tables
- `\d {table_name}` will list the information for a table.
- `exit` will exit the psql terminal