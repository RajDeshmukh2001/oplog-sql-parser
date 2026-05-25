# MongoDB Oplog to SQL Parser

A Go implementation that converts MongoDB oplog entries into equivalent SQL statements.

## Problem Statement
A program to parse MongoDB operations log (oplog) and generate equivalent SQL statements.

We have a scenario where an organization used MongoDB initially but now needs to move to an RDBMS database. This data transition can be made easy if we can find a way to convert the JSON documents in MongoDB collections to equivalent rows in relational DB tables. That's the purpose of this program.

The MongoDB server generates the Oplog, an ordered collection of all the write operations (insert, update, delete) to the MongoDB. Parse these oplogs and generate equivalent SQL statements.

## What is MongoDB Oplog?
Oplog (Operations Log) is a special capped collection in MongoDB called `local.oplog.rs`. It's the heart of MongoDB's replication system.

Every write operation that happens on a MongoDB primary node gets recorded in the oplog as an entry, in order. Replica set secondaries tail this oplog to stay in sync with the primary.

## What is an Oplog entry?

#### Example
```json
{
  "op": "i",
  "ns": "test.student",
  "o": {
    "_id": "635b79e231d82a8ab1de863b",
    "name": "Selena Miller",
    "roll_no": 51,
    "is_graduated": false,
    "date_of_birth": "2000-01-30"
  }
}
```

Every record in `oplog.rs` is a BSON document with these key fields:
- `op`: This indicates the type of operation. It can be `i` (insert), `u` (update), `d` (delete), `c` (command), `n` (no operation). For this implementation, we'll only care about insert, update and delete operations.

- `ns`: This indicates the namespace. Namespace consists of database and collection name separated by a `.` In above case, database name is `test` and collection name is `student`.

- `o`: This indicates the new data for insert or update operation. In above case, a student document is inserted in the collection.

## Project Structure

```text
oplog-sql-parser/
│
├── internal/
│   ├── model/
│   │   └── oplog.go
│   │
│   ├── parser/
│   │   ├── oplog_parser.go
│   │   └── oplog_parser_test.go
│   │
│   ├── sql/
│   │   ├── insert.go
│   │   └── insert_test.go
│   │
│   └── validator/
│       ├── insert_validator.go
│       └── insert_validator_test.go
│
├── .gitignore
├── go.mod
└── README.md
```

## How to Run the Code
Since this project currently exposes reusable packages and does not include an executable entry point, you can create a temporary `main.go` file to verify the implementation manually.

#### To manually verify the implementation:
1. Create a temporary `main.go` file in the project root.
2. Create a sample insert oplog JSON.
3. Parse the oplog using `parser.ParseOplog`.
4. Validate the oplog using `validator.ValidateInsertOplog`.
5. Generate the SQL statement using `sql.GenerateInsertSQL`.
6. Print the generated SQL statement to the console.

The expected output should be a valid SQL `INSERT` statement corresponding to the provided oplog entry.

#### Run

```bash
go run main.go
```

#### Expected output:

```sql
INSERT INTO test.student (_id, date_of_birth, is_graduated, name, roll_no)
VALUES ('635b79e231d82a8ab1de863b', '2000-01-30', false, 'Selena Miller', 51);
```

## How to Run Tests
#### Run all unit tests:

```bash
go test ./...
```

#### Run all tests with verbose output:

```bash
go test ./... -v
```

## Features

**1. Parse Insert Oplog**

Implemented support for parsing MongoDB insert oplogs and generating equivalent SQL INSERT statements.

#### Supported Data Types

The following MongoDB value types are currently supported:

- String
- Number
- Boolean

#### Assumptions

- Only insert (`op = "i"`) oplogs are supported in this feature.
- Nested JSON objects are not supported.
- Arrays are not supported.
- Dates are treated as strings.
- Namespace (`ns`) follows the format `<database>.<collection>`.

#### Example Input

```json
{
  "op": "i",
  "ns": "test.student",
  "o": {
    "_id": "635b79e231d82a8ab1de863b",
    "name": "Selena Miller",
    "roll_no": 51,
    "is_graduated": false,
    "date_of_birth": "2000-01-30"
  }
}
```

#### Example Output

```sql
INSERT INTO test.student (_id, date_of_birth, is_graduated, name, roll_no)
VALUES ('635b79e231d82a8ab1de863b', '2000-01-30', false, 'Selena Miller', 51);
```