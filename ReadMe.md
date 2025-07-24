
# Boot Dev Blog Aggregator
## Setup
- ### Create gator file that stores database url and currently logged in user in the home directory
    - The file name should be `.gatorconfig.json`
    - Add `{"DB_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"}` to the file
- ### Install psql
    - #### Check if psql is already installed
        - `psql --version`
    - #### If not installed
        - `sudo apt update`
        - `sudo apt install postgresql postgresql-contrib`
        - udpate the password for psql
            - I used `postgres`
            - If you use a different password then the DB_url will be different
                - Use this for help https://www.geeksforgeeks.org/postgresql/postgresql-connection-string/
        - `sudo passwd postgres`
    - #### Start the server
        - `sudo service postgresql start`
            - server can be stopped by replacing `start` with `stop` and running the command again
## Commands
- All commands
    - addfeed
    - agg
    - browse
    - feeds
    - follow
    - following
    - login
    - register
    - reset
    - unfollow
    - users


- ### agg
    - Starts the aggregator with a pause between requests using the given duration.
        - Time parses using [Go parse duration](https://pkg.go.dev/time#ParseDuration)
        - Valid units are `ns, us, ms, s, m, h` which correspond to `nano seconds`, `micro seconds`, `milli seconds`, `seconds`, `minutes`, `hours`
            - Please do not set time too short to avoid dosing servers.
    - syntax `{run} agg {time between requests}`
- ### addfeed
    - Adds the feed to the list of feeds and set the currently logged in user to follow that feed
    - syntax `{run} addfeed {name for feed} {url of feed}
    - Will fail if missing name or url of feed
- ### browse
    - Lists the 2 most recently updated feeds for current user. 
        - Optional amount to be returned param to increase or decrease amount returned
    - syntax `{run} browse {optional amount}`
- ### feeds
    - Lists all feeds currently added to the database and which user add them
    - syntax `{run} feeds`
- ### follow
    - The currently logged in user will follow the feed at the given url.
    - Fails if feed has not been added to the database
    - syntax `{run} follow {url of feed}`
- ### following
    - Lists the feeds the currently logged in user is following
    - syntax `{run} following`
- ### login
    - Logs in the username that was given
    - syntax `{run} login {user_name_to_login}`
    - Will fail if user name has not been registered
- ### register
    - Creates a user with passed in name and then logs you in as the newly registered user
    - syntax `{run} regist {user_name_to_register}`
    - Will fail if user name already exists
- ### reset
    - WARNING: Running this command will delete all data stored in the database
    - syntax `{run} reset`
- ### unfollow
    - User unfollows the url of the given feed
    - suntax `{run} unfollow {url to unfollow}`
- ### users
    - Displays all registered users and displays which user is currently logged in
    - syntax `{run} users`

### To generate and revert migrations use these commands from the `/sql/schema` directory you will need goose
- #### install goose
    -`go install github.com/pressly/goose/v3/cmd/goose@latest`
- #### you can new add new migrations in the schema folder. `{migration number}_{description}.sql`
        The URL here is the same one the .gatorconfig.json file
    - `goose postgres "postgres://postgres:postgres@localhost:5432/gator" up`
- #### To roll back a migration do
    - `goose postgres "postgres://postgres:postgres@localhost:5432/gator" down`

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