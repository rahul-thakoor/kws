# MongoDB CRUD operations with GO

Note: This is an adapted version of an article I previously wrote: https://www.supinfo.com/articles/single/7889-using-azure-cosmos-db-s-mongodb-api-with-nodejs

## Introduction

Go[^go] is a programming language, designed at Google, that is increasingly becoming popular. It is very easy to learn and use and supports several platforms. In this article we are going to give an introduction on how to use Go to interact with MongoDB which is a document-oriented NoSQL database.

## Prerequisites

1. Prior knowledge of MongoDB
2. A basic command of the Go programming language
    - If you are not familiar with Go, you can take [a tour of the Go programming language](https://tour.golang.org/welcome/1)
3. Go installed on your machine
    - Follow the official [getting started guide](https://golang.org/doc/install)
4. A MongoDB database with a user having read/write access
    - You can use mLab's [free account](https://docs.mlab.com/) for quick prototyping

## Create a Go program and install dependencies

First we check if Go is installed correctly. Create a file named `crud.go` with the following content:

```go
package main

import "fmt"

func main() {
	fmt.Printf("Hello, world\n")
}
```

Run the program with: 
```
go run crud.go
```
You should see the `"Hello, world"` message. This means Go is properly installed and working.

We are going to use the official MongoDB Go Driver (currently in beta) to interact with MongoDB.

Install the driver using:

```console
go get github.com/mongodb/mongo-go-driver
```

## Connecting to the database

Retrieve the Mongo Connection String for your database. For instance on mlab:

!["connection string"](connection-string.png)

`mongodb://<dbuser>:<dbpassword>@ds213645.mlab.com:13645/supinfo-kws-demo`

The connection string is in the form :
`mongodb://<username>:<password>@<endpoint>:13645/<dbname>`

You should replace `<username>` and `<password>` appropriately.

Modify your program's to the following:

```go
package main

import (
	"log"
	"context"
	"time"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	// Create context
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)

	// Create client
	client, err := mongo.Connect(ctx, "mongodb://supinfo:supinfo123@ds213645.mlab.com:13645/supinfo-kws-demo")

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("connection successful")

}
```


These are required for establishing a connection to the database server. 



Test if the connection to the database was successful by running the program.

```bash
 go run crud.go
```

The output in the console should be `connection successful`

## Basic CRUD 

### Create 

Now that we have connected to the database, let's create a document. A document is inserted in a collection. For example,
in our database `supinfo` we could have several collections. Let's create a document that is inserted in the `1WEB` collection.

```go
package main

import (
	"log"
	"context"
	"time"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	// Create context
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Create client
	client, err := mongo.Connect(ctx, "mongodb://supinfo:supinfo123@ds213645.mlab.com:13645/supinfo-kws-demo")

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("connection successful")

	collection := client.Database("supinfo-kws-demo").Collection("Assignments")
	insertResult, err := collection.InsertOne(ctx, bson.M{"ID":"MP1", "Description":"Puzzler- Use HTML5 to create a mini puzzle game", "Module":"1WEB" , "Deadline":"25-09-2018" })
	if err != nil {
		log.Fatal(err)
	}

    log.Println("Inserted a single document: ", insertResult.InsertedID)
    
    // close connection
    client.Disconnect(ctx)
```

The expected output in the console is as follows:

```
2019/01/30 16:50:38 connection successful
2019/01/30 16:50:40 Inserted a single document:  ObjectID("5c51d5dec8f93d6c8d63ae51")
```

At this point, we can  check the document in the browser.

If we go to the `mlab` , we can see the `supinfo` database with the `Assignments` collection and the inserted document.

We might have to click refresh to see the document. The document has the following struture:

```json
{

    "_id": {
        "$oid": "5c51d5dec8f93d6c8d63ae51"
    },
    "ID": "MP1",
    "Description": "Puzzler- Use HTML5 to create a mini puzzle game",
    "Module": "1WEB",
    "Deadline": "25-09-2018"

}
```
!["Insert one"](insert_one.png)


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

### Update

To update the above document to add a description field we can do the following:

```js
MongoClient.connect(url,{ useNewUrlParser: true }, function(err, client) {  
    assert.equal(null, err);
    console.log('connection successful')

    // get the database
    const db=client.db(dbname);

    // get the required collection

    const collection = db.collection('ASC1 Modules');

    console.log('Getting collection successful');

    // Update document where a is 2, set b equal to 1
    let descriptionContent = "Comme son nom l'indique, le cours 1WEB vous fera découvrir le développement web via les langages HTML, CSS et JavaScript. Il vous présentera également le framework jQuery. Ce cours vous permettra d’acquérir l'ensemble des notions essentielles pour développer des sites internet avec les technologies HTML, CSS et JavaScript. Il vous présentera également une introduction à la toute dernière mouture d’HTML (version 5) avec son lot de nouveautés. Il n’est plus nécessaire de présenter les avantages d’internet et sa facilité d’accès pour ses internautes. La facilité d’accès à l’information, les interconnexions omniprésentes et la liberté d’expression sont autant de facteurs expliquant le succès de « La Toile ». Le réseau des réseaux est également simple à appréhender pour les développeurs, faisant de l’HTML, de CSS et de JavaScript des éléments de programmation simples à appréhender pour les codeurs en herbe. jQuery est une librairie JavaScript développée par John Resig en 2006 qui est aujourd’hui utilisée dans un grand nombre de sites. Ses atouts résident dans la simplification de la syntaxe de langage et de certaines opérations de calcul, de parcours et d’animation. Elle vous permettra de créer simplement des interactions de qualité pour rendre votre site ergonomique.";
  collection.updateOne({ id : "1WEB" }
    , { $set: { description : descriptionContent } }, function(err, result) {
    assert.equal(err, null);
    assert.equal(1, result.result.n);
    console.log("Updated the document with id 1WEB");
    callback(result);
  });  
    // close the connection

    client.close();
  
    }
);
```

The `1WEB` document now has a description field. 

```js
{
    "_id": {
        "$oid": "5ba86422d528840c10312e8e"
    },
    "$id": "1WEB",
    "title": "HTML & JavaScript - User Interface",
    "ects": 3,
    "id": "5ba86422d528840c10312e8e",
    "_rid": "tEF8AJZLIgABAAAAAAAAAA==",
    "_self": "dbs/tEF8AA==/colls/tEF8AJZLIgA=/docs/tEF8AJZLIgABAAAAAAAAAA==/",
    "_etag": "\"00000000-0000-0000-5488-aa22c4b401d4\"",
    "_attachments": "attachments/",
    "description": "Comme son nom l'indique, le cours 1WEB vous fera découvrir le développement web via les langages HTML, CSS et JavaScript. Il vous présentera également le framework jQuery. Ce cours vous permettra d’acquérir l'ensemble des notions essentielles pour développer des sites internet avec les technologies HTML, CSS et JavaScript. Il vous présentera également une introduction à la toute dernière mouture d’HTML (version 5) avec son lot de nouveautés. Il n’est plus nécessaire de présenter les avantages d’internet et sa facilité d’accès pour ses internautes. La facilité d’accès à l’information, les interconnexions omniprésentes et la liberté d’expression sont autant de facteurs expliquant le succès de « La Toile ». Le réseau des réseaux est également simple à appréhender pour les développeurs, faisant de l’HTML, de CSS et de JavaScript des éléments de programmation simples à appréhender pour les codeurs en herbe. jQuery est une librairie JavaScript développée par John Resig en 2006 qui est aujourd’hui utilisée dans un grand nombre de sites. Ses atouts résident dans la simplification de la syntaxe de langage et de certaines opérations de calcul, de parcours et d’animation. Elle vous permettra de créer simplement des interactions de qualité pour rendre votre site ergonomique.",
    "_ts": 1537849904
}

```
Similarly we can update the `1LIN` document to it's correct description. It currently has the description for `1WEB`:

We simply have to change the following:

```js
let descriptionContent = "<Description for 1LIN module>"

collection.updateOne({ id : "1LIN" }
    , { $set: { description : descriptionContent } }, function(err, result) {
    assert.equal(err, null);
    assert.equal(1, result.result.n);
    console.log("Updated the document with id 1LIN");
    callback(result);
  }); 

```

### Delete

Let's assume we were trying to add a document for `2WEB` but inserted the document in the `ASC1 Modules` collection by mistake:

```js
MongoClient.connect(url,{ useNewUrlParser: true }, function(err, client) {  
    assert.equal(null, err);
    console.log('connection successful')

    // get the database
    const db=client.db(dbname);

    // get the required collection, ASC1 Modules here 

    const collection = db.collection('ASC1 Modules');

    console.log('Getting collection successful');

    // insert document into collection

    const doc = {
        {
            id:"2WEB",
            title:"Web programming with PHP",
            ects: 3,
            description: "<2WEB description>"
        },
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

![2WEB document incorrectly inserted in ASC1 Module collection](./wrongInsert.PNG)

We can delete the particular document using:

```js
MongoClient.connect(url,{ useNewUrlParser: true }, function(err, client) {  
    assert.equal(null, err);
    console.log('connection successful')

    // get the database
    const db=client.db(dbname);

    // get the required collection

    const collection = db.collection('ASC1 Modules');

    console.log('Getting collection successful');

    
     // Delete document where a is 3
    collection.deleteOne({ id : "2WEB" }, function(err, result) {
        assert.equal(err, null);
        assert.equal(1, result.result.n);
        console.log("Removed the document with the id 2WEB");
    });    

    // close the connection

    client.close();
  
    }
);
```

This deletes the document.

## References
 - [^go] https://en.wikipedia.org/wiki/Go_(programming_language)
 - http://mongodb.github.io/node-mongodb-native/3.1/quick-start/quick-start/
 - https://docs.mongodb.com/ecosystem/drivers/
 - https://godoc.org/github.com/mongodb/mongo-go-driver/mongo 
 - https://blog.golang.org/context