# backend
The backend of the ticketing system

# Environment Variables needed to be set

## DB_TYPE

__TODO__

__Currently only supports postgres__


## DB_USERNAME

Defines the username to be used to access the database.

## DB_PASSWORD

Defines the password used to access the database.

## DB_HOSTNAME

Defines the database network hostname to be used to access the database.

## DB_NAME

Defines the name of the database used to used on the server

# Database definition

## Accounts table

### Rows

### ID
This is a unique id for each user and is randomly generated on inserting user. This should be defined as a primary key.

### email
Stores the users email address, must be able to store at least 320 characters or this server may not function as expected. It is recommeneded to have an index on this column.

### password
Stores the hashed password using bcrypt. Must be able to fit a bcrypt hash or server may not function as expected.

### Defined as in postgres:

```
CREATE TABLE accounts (
    ID SERIAL PRIMARY KEY,
    email varchar(320),
    password varchar(72)
)
```