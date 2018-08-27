# Using Azure Cosmos DB's MongoDB API with Node.js

## Introduction

Azure Cosmos DB is a multi-model database by Microsoft. The database provides APIs for data model such as SQL, Cassandra and MongoDB.

MongoDB is a document-oriented NoSQL database and provides official drivers for major programming languages.[^2]

In this article we are going to give an introduction on how to interface with Azure Cosmos DB's MongoDB API using Node.js.

## Prerequisites

1. Prior knowledge of Node.js and MongoDB
2. A machine with the Windows Operating System
3. Azure Cosmos DB Emulator installed on your machine for local development 
    - Please follow this tutorial install Azure Cosmos DB Emulator on your machine
4. Node.js installed on your machine

## Create a Node.js app and install dependencies

Initialise a project using `npm init`. This will prompt you for information to populate the `package.json`, namely:

1. name
2. version
3. description
4. entry point
5. test command
6. git repository
7. keywords
8. author
9. license

You can press 'Enter' to accept the defaults for any step. 

This creates the `package.json` with content similar to the following(depending on the data you provided):

```js
{
  "name": "demo-app",
  "version": "1.0.0",
  "description": "A simple app to use cosmosdb's mongodb api",
  "main": "index.js",
  "scripts": {
    "test": "echo \"Error: no test specified\" && exit 1"
  },
  "author": "",
  "license": "ISC"
}
```
After initilasing a project, we proceed to install the dependencies to be used for the app.

Install the MongoDB Node.js driver

```console
npm install mongodb --save
```

Install the `assert` module for running tests

```
npm install assert --save
```

## Connecting to the Emulator

Start the Azure Cosmos DB Emulator. When it launches, it automatically opens the Azure Cosmos DB Data Explorer in your browser. Navigate to https://localhost:8081/_explorer/index.html in your browser for the Data Explorer.

![data explorer screenshot](./cosmosdb-emulator-data-explorer.png)

Retrieve the Mongo Connection String:

`mongodb://localhost:C2y6yDjf5/R+ob0N8A7Cgv30VRDJIWEHLM+4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw/Jw==@localhost:10255/admin?ssl=true`

The connection string is in the form :
`mongodb://<username>:<password>@<endpoint>:10255/admin?ssl=true`

Modify your program's entry point(defined when initialising the project) to add:

```js
const MongoClient = require('mongodb').MongoClient;
const assert = require('assert');
// The MongoDB Node.js 3.0 driver requires encoding special characters in the Cosmos DB password. 
const password = encodeURIComponent('C2y6yDjf5/R+ob0N8A7Cgv30VRDJIWEHLM+4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw/Jw==');
const url = `mongodb://localhost:${password}@localhost:10255/admin?ssl=true`;
const dbname = "supinfodb";
```

These are required for establishing a connection to the emulator. 

```js
var db;
MongoClient.connect(url, function(err, client) {  
    assert.equal(null, err);
    console.log('connection successful')
    db=client.db(dbname);
    
    }
);

```

## Basic CRUD 


## References

1. 
2. https://docs.mongodb.com/ecosystem/drivers/