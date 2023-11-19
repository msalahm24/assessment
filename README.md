# assessment

#### This makefile defines a set of rules for building.

# To use the makefile, you can simply run the `make` command in the root directory of your project.

# This will run all of the rules in the makefile, in the order that they are defined.

# The `my-postgres` rule creates a PostgreSQL database container.
# The `createdb` rule creates a database named `otp` in the container.
# The `dropdb` rule drops the `otp` database.
# The `migrateup` rule migrates the database schema up to the latest version.
# The `migratedown` rule migrates the database schema down to the previous version.
# The `sqlc` rule generates the Golang code for the database schema.
# The `server` rule runs the Golang server.