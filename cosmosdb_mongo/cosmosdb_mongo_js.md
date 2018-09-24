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

### Create 

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

```js
MongoClient.connect(url,{ useNewUrlParser: true }, function(err, client) {  
    assert.equal(null, err);
    console.log('connection successful')

    // get the database
    const db=client.db(dbname);

    // get the required collection, 1WEB here

    const collection = db.collection('ASC1 Modules');

    console.log('Getting collection successful');

    // insert document into collection

    const doc = [
        {
            id:"1WEB",
            title:"HTML & JavaScript - User Interface",
            ects: 3,
        },

        {
            id:"1LIN",
            title:"Linux Technologies - System Fundamentals",
            ects: 3,
            description: "Comme son nom l'indique, le cours 1WEB vous fera découvrir le développement web via les langages HTML, CSS et JavaScript. Il vous présentera également le framework jQuery. Ce cours vous permettra d’acquérir l'ensemble des notions essentielles pour développer des sites internet avec les technologies HTML, CSS et JavaScript. Il vous présentera également une introduction à la toute dernière mouture d’HTML (version 5) avec son lot de nouveautés. Il n’est plus nécessaire de présenter les avantages d’internet et sa facilité d’accès pour ses internautes. La facilité d’accès à l’information, les interconnexions omniprésentes et la liberté d’expression sont autant de facteurs expliquant le succès de « La Toile ». Le réseau des réseaux est également simple à appréhender pour les développeurs, faisant de l’HTML, de CSS et de JavaScript des éléments de programmation simples à appréhender pour les codeurs en herbe. jQuery est une librairie JavaScript développée par John Resig en 2006 qui est aujourd’hui utilisée dans un grand nombre de sites. Ses atouts résident dans la simplification de la syntaxe de langage et de certaines opérations de calcul, de parcours et d’animation. Elle vous permettra de créer simplement des interactions de qualité pour rendre votre site ergonomique."
        },

        {
            id:"1CNA",
            title:"CCNA Routing & Switching Part 1",
            ects: 3,
            description: "Le cours 1CNA - Cisco CCNA Routing & Switching - Part 1 vous permettra de découvrir les réseaux informatiques, comment ils fonctionnent mais aussi les enjeux cruciaux liés à ceux-ci. Comme vous le savez, les réseaux sont présents de plus en plus dans nos vies quotidiennes et connectent des millions de personnes dans le monde entier. Les nouveaux enjeux liés à l’Internet of Things et l’Internet of Everything sont également un tremplin de plus pour faire évoluer les réseaux et la consumérisation des ressources, de la perspective de l’utilisateur final. La première partie du cours, CCNA 1 - Introduction to Networks vous introduira aux fondamentaux des réseaux, tant en termes globaux que techniques. Il vous permettra également de découvrir comment est construit un réseau physique et logique et comment l’organiser pour répondre aux besoins clients. La deuxième partie de celui-ci, CCNA 2 - Routing & Switching Essentials vous permettra d’aborder les concepts techniques de routage et de commutation, dans un cadre local comme d’interconnexion entre sites. Vous apprendrez des concepts essentiels concernant ces deux domaines."
        },


        ];
    collection.insertMany(doc, function(err,res){
        assert.equal(null,err);
        assert.equal(3, res.ops.length);
    });

    console.log('Document created');

    // close the connection

    client.close();


    
    }
);
```

![Insert Many Image](./insertmany.PNG)

This create a collection called `ASC Modules` and inserts three documents at once. Note: The description field for `1WEB` is deliberately left out and similarly the content of the description for `1LIN` is intentionally wrong so we can illustrate how to read these documents and update them.

Next, we can explore how to read from the database.

### Read

We can find all documents in a collection as follows:

```js
MongoClient.connect(url,{ useNewUrlParser: true }, function(err, client) {  
    assert.equal(null, err);
    console.log('connection successful')

    // get the database
    const db=client.db(dbname);

    // get the required collection

    const collection = db.collection('ASC1 Modules');

    console.log('Getting collection successful');

    // find all documents in the collection

    collection.find({}).toArray(function(err, docs) {
        assert.equal(err, null);
        console.log("Found the following documents");
        console.log(docs)
        
      });

    // close the connection

    client.close();
    
    }
);
```
This prints out all the documents, 3 in our case.

We can also filter the query. For instance, let's find documents that do not have the `description` field set:

```js
MongoClient.connect(url,{ useNewUrlParser: true }, function(err, client) {  
    assert.equal(null, err);
    console.log('connection successful')

    // get the database
    const db=client.db(dbname);

    // get the required collection

    const collection = db.collection('ASC1 Modules');

    console.log('Getting collection successful');

    // find all documents in the collection which do not have have a description
    // i.e, description field is null

    collection.find({description:null}).toArray(function(err, docs) {
        assert.equal(err, null);
        console.log("Found the following documents");
        console.log(docs)
        
      });
    // close the connection

    client.close();
  
    }
);
```

This returns the following output:

```
connection successful
Getting collection successful
Found the following documents
[ { _id: 5ba86422d528840c10312e8e,
    id: '1WEB',
    title: 'HTML & JavaScript - User Interface',
    ects: 3 } ]
```

## References

1. http://mongodb.github.io/node-mongodb-native/3.1/quick-start/quick-start/
2. https://docs.mongodb.com/ecosystem/drivers/