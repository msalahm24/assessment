# assessment

#### This makefile defines a set of rules for building.

#### To use the makefile, you can simply run the `make` command in the root directory of your project.

#### The `my-postgres` rule creates a PostgreSQL database container.
#### The `createdb` rule creates a database named `otp` in the container.
#### The `dropdb` rule drops the `otp` database.
#### The `migrateup` rule migrates the database schema up to the latest version.
#### The `migratedown` rule migrates the database schema down to the previous version.
#### The `sqlc` rule generates the Golang code for the database schema.
#### The `server` rule runs the Golang server.