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
MongoClient.connect(url,{ useNewUrlParser: true }, function(err, client) {  
    assert.equal(null, err);
    console.log('connection successful')

    // get the database
    const db=client.db(dbname);


    // close the connection

    client.close();


    
    }
);

```

Test if the connection to the database was successful by running the program.

```bash
 node index.js
```

The output in the console should be `connection successful`

## Basic CRUD 

Now that we have connected to the database, let's create a document. A document is inserted in a collection. For example,
in our database `supinfo` we could have several collections. Let's create a document that is inserted in the `1WEB` collection.

```js
MongoClient.connect(url,{ useNewUrlParser: true }, function(err, client) {  
    assert.equal(null, err);
    console.log('connection successful')

    // get the database
    const db=client.db(dbname);

    // get the required collection, 1WEB here

    const collection = db.collection('1WEB');

    console.log('Getting collection successful');

    // insert document into collection

    const doc = {
        assignment: {
            id: "MP1",
            description: "Puzzler- Use HTML5 to create a mini puzzle game",
            module: "1WEB",
            deadline: "25-09-2018"
        }
    }
    collection.insertOne(doc, function(err,res){
        assert.equal(null,err);
        assert.equal(1, res.ops.length);
    });

    console.log('Document created');

    // close the connection

    client.close();


    
    }
);
```

The expected output in the console is as follows:

```
connection successful
Getting collection successful
Document created
```

At this point, we can use the Data Explorer to check the document in the browser.

If we go to the `Explorer` pane, we can see the `supindodb` database with the `1WEB` collection and the inserted document.

An `id` is generated for the document when it is inserted.

E.g in my case, the `id` is `5ba4782729df2d4fcc25f800`

When we click on the document, we can the the contents. We might have to click refresh to see the document. The document has the following struture:

```json
{
    "_id": {
        "$oid": "5ba4782729df2d4fcc25f800"
    },
    "assignment": {
        "$id": "MP1",
        "description": "Puzzler- Use HTML5 to create a mini puzzle game",
        "module": "1WEB",
        "deadline": "25-09-2018"
    },
    "id": "5ba4782729df2d4fcc25f800",
    "_rid": "tEF8AMDV+wADAAAAAAAAAA==",
    "_self": "dbs/tEF8AA==/colls/tEF8AMDV+wA=/docs/tEF8AMDV+wADAAAAAAAAAA==/",
    "_etag": "\"00000000-0000-0000-5166-5d4368d701d4\"",
    "_attachments": "attachments/",
    "_ts": 1537505319
}
```

We can see there are additional `metadata` added when a document is inserted.

We can also create several documents at once:



## References

1. http://mongodb.github.io/node-mongodb-native/3.1/quick-start/quick-start/
2. https://docs.mongodb.com/ecosystem/drivers/