# Ohio COVID County tracking
A dashboard for tracking coronavirus statistics on a per county basis in Ohio.  Can be seen here: http://covid.maxheckel.me

## Local env
### Go server/postgres
#### Prerequisites
- Golang 1.15
- Docker
- Docker-compose

Run the following command to start a local database and go server:
`make docker-watch`

The server has a file watcher built in which will rebuild the go app started in `cmd/covid_county/main.go`
### Angular app
#### Prerequisites
- Angular 10 CLI
- Node v10.17.0
- NPM 6.11.3

Install dependencies by running `cd web && npm install`

To run the front end application simply run `cd web && ng build --watch`.  This will start a file watcher for the angular app to run.  There is a volume set up in the docker-compose file for the UI to reload on the container as well as your local machine.

### Contributing
Just open a PR! If you have issues feel free to add any issues as well.  If you would like to speak with me directly my email is heckel.max@gmail.com.
