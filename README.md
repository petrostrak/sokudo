### sokudo
Sokudo (速度) is a Go module and command line application that makes building a web application simple, fast and secure.

Sokudo module comes with a handful of commands most of which enchance the functionality of the generated application, when the user sees fit

`help`                  - show the help commands

`version`               - print application version

`migrate`               - runs all up migrations that have not been run previously

`migrate down`          - reverses the most recent migration

`migrate reset`         - runs all down migrations in reverse order, and then all up migrations

`make migration <name>` - creates two new up and down migrations in the migrations folder

`make auth`             - creates and runs migrations for authentication tables, and creates models and middleware

`make handler <name>`   - creates a stub handler in the handlers directory

`make model <name>`     - creates a new model in the data directory

`make session`          - creates a table in the database as a session store

`make mail <name>`      - creates two starter mail templates in the mail directory
