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

	// Create context, Note here timeout is 2 seconds
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

	// Create context, Note here timeout is 10 seconds

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


MongoDB also generates an id when inserting an item;

We can also create several documents at once:

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

	// create a Module type representing a particular course module
	type Module struct{
		id string
		title string
		ects int
		description string
	}


	// create some modules
    cna1 := Module{"1CNA", "CCNA Routing & Switching Part 1", 3, "Le cours 1CNA - Cisco CCNA Routing & Switching - Part 1 vous permettra de découvrir les réseaux informatiques, comment ils fonctionnent mais aussi les enjeux cruciaux liés à ceux-ci. Comme vous le savez, les réseaux sont présents de plus en plus dans nos vies quotidiennes et connectent des millions de personnes dans le monde entier. Les nouveaux enjeux liés à l’Internet of Things et l’Internet of Everything sont également un tremplin de plus pour faire évoluer les réseaux et la consumérisation des ressources, de la perspective de l’utilisateur final. La première partie du cours, CCNA 1 - Introduction to Networks vous introduira aux fondamentaux des réseaux, tant en termes globaux que techniques. Il vous permettra également de découvrir comment est construit un réseau physique et logique et comment l’organiser pour répondre aux besoins clients. La deuxième partie de celui-ci, CCNA 2 - Routing & Switching Essentials vous permettra d’aborder les concepts techniques de routage et de commutation, dans un cadre local comme d’interconnexion entre sites. Vous apprendrez des concepts essentiels concernant ces deux domaines."}
    
    web1 := Module{"1WEB","HTML & JavaScript - User Interface", 3,""}
    
	lin1 := Module{"1LIN","Linux Technologies - System Fundamentals",3,"Comme son nom l'indique, le cours 1WEB vous fera découvrir le développement web via les langages HTML, CSS et JavaScript. Il vous présentera également le framework jQuery. Ce cours vous permettra d’acquérir l'ensemble des notions essentielles pour développer des sites internet avec les technologies HTML, CSS et JavaScript. Il vous présentera également une introduction à la toute dernière mouture d’HTML (version 5) avec son lot de nouveautés. Il n’est plus nécessaire de présenter les avantages d’internet et sa facilité d’accès pour ses internautes. La facilité d’accès à l’information, les interconnexions omniprésentes et la liberté d’expression sont autant de facteurs expliquant le succès de « La Toile ». Le réseau des réseaux est également simple à appréhender pour les développeurs, faisant de l’HTML, de CSS et de JavaScript des éléments de programmation simples à appréhender pour les codeurs en herbe. jQuery est une librairie JavaScript développée par John Resig en 2006 qui est aujourd’hui utilisée dans un grand nombre de sites. Ses atouts résident dans la simplification de la syntaxe de langage et de certaines opérations de calcul, de parcours et d’animation. Elle vous permettra de créer simplement des interactions de qualité pour rendre votre site ergonomique."}
	
	// Create a slice of several modules
	modules := []interface{}{cna1,web1,lin1}

	// Get collection in which docs will be inserted
	collection := client.Database("supinfo-kws-demo").Collection("ASc1 Modules")

	// insert multiple documents at once
	res, err := collection.InsertMany(ctx, modules)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(len(res.InsertedIDs))

	// close connection
	client.Disconnect(ctx)
```
The expected output in the console is as follows:
```
2019/02/13 06:37:14 connection successful
2019/02/13 06:37:16 Successfully inserted 3 docs
```

![Insert Many Image](./insert_many.PNG)

This creates a collection called `ASc1 Modules` and inserts three documents at once. Note: The description field for `1WEB` is deliberately left out and similarly the content of the description for `1LIN` is intentionally wrong so we can illustrate how to read these documents and update them.


Next, we can explore how to read from the database.

### Read

We can find all documents in a collection as follows:


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

	// create a Module type representing a particular course module
	type Module struct{
		Id string
		Title string
		ECTS int
		Description string
	}

	// Get collection
	collection := client.Database("supinfo-kws-demo").Collection("ASc1 Modules")

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil { log.Fatal(err) }
	cur.Close(ctx)

	for cur.Next(context.Background()) {
	   raw, err := cur.DecodeBytes()
	   if err != nil { log.Fatal(err) }

		var result Module
		err = bson.Unmarshal(raw, &result)
		if err != nil { 
			log.Fatal(err)
		 }
		log.Printf("Found a document: %+v\n", result)
		

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// close connection
	client.Disconnect(ctx)


}

```
This prints out all the documents, 3 in our case.

We can also filter the query. For instance, let's find documents that do not have the `description` field set:


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

	// create a Module type representing a particular course module
	type Module struct{
		Id string
		Title string
		ECTS int
		Description string
	}

	// Get collection
	collection := client.Database("supinfo-kws-demo").Collection("ASc1 Modules")

	var result Module
	err = collection.FindOne(ctx, bson.D{{"description", ""}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found document without description: %+v\n", result)

	// close connection
	client.Disconnect(ctx)

}
```

This returns the following output:

```

2019/02/13 13:20:44 connection successful
2019/02/13 13:20:46 Found document without description: {Id:1WEB Title:HTML & JavaScript - User Interface ECTS:3 Description:}

```

### Update

To update the above document to add a description field we can do the following:


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

	// create a Module type representing a particular course module
	type Module struct{
		Id string
		Title string
		ECTS int
		Description string
	}

	// Get collection
	collection := client.Database("supinfo-kws-demo").Collection("ASc1 Modules")

	update := bson.D{
		{"$set", bson.D{
			{"description", "Comme son nom l'indique, le cours 1WEB vous fera découvrir le développement web via les langages HTML, CSS et JavaScript. Il vous présentera également le framework jQuery. Ce cours vous permettra d’acquérir l'ensemble des notions essentielles pour développer des sites internet avec les technologies HTML, CSS et JavaScript. Il vous présentera également une introduction à la toute dernière mouture d’HTML (version 5) avec son lot de nouveautés. Il n’est plus nécessaire de présenter les avantages d’internet et sa facilité d’accès pour ses internautes. La facilité d’accès à l’information, les interconnexions omniprésentes et la liberté d’expression sont autant de facteurs expliquant le succès de « La Toile ». Le réseau des réseaux est également simple à appréhender pour les développeurs, faisant de l’HTML, de CSS et de JavaScript des éléments de programmation simples à appréhender pour les codeurs en herbe. jQuery est une librairie JavaScript développée par John Resig en 2006 qui est aujourd’hui utilisée dans un grand nombre de sites. Ses atouts résident dans la simplification de la syntaxe de langage et de certaines opérations de calcul, de parcours et d’animation. Elle vous permettra de créer simplement des interactions de qualité pour rendre votre site ergonomique."},
		}},
	}
	updateResult, err := collection.UpdateOne(ctx, bson.D{{"id", "1WEB"}}, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Updated %v document(s).\n", updateResult.ModifiedCount)

	// close connection
	client.Disconnect(ctx)

}

```

The `1WEB` document now has a description field. 

```js
{
    "_id": {

        "$oid": "5c63bb1aaf0afdc40ab4288a"
    },
    "id": "1WEB",
    "title": "HTML & JavaScript - User Interface",
    "ects": 3,
    "description": "Comme son nom l'indique, le cours 1WEB vous fera découvrir le développement web via les langages HTML, CSS et JavaScript. Il vous présentera également le framework jQuery. Ce cours vous permettra d\u2019acquérir l'ensemble des notions essentielles pour développer des sites internet avec les technologies HTML, CSS et JavaScript. Il vous présentera également une introduction à la toute dernière mouture d\u2019HTML (version 5) avec son lot de nouveautés. Il n\u2019est plus nécessaire de présenter les avantages d\u2019internet et sa facilité d\u2019accès pour ses internautes. La facilité d\u2019accès à l\u2019information, les interconnexions omniprésentes et la liberté d\u2019expression sont autant de facteurs expliquant le succès de « La Toile ». Le réseau des réseaux est également simple à appréhender pour les développeurs, faisant de l\u2019HTML, de CSS et de JavaScript des éléments de programmation simples à appréhender pour les codeurs en herbe. jQuery est une librairie JavaScript développée par John Resig en 2006 qui est aujourd\u2019hui utilisée dans un grand nombre de sites. Ses atouts résident dans la simplification de la syntaxe de langage et de certaines opérations de calcul, de parcours et d\u2019animation. Elle vous permettra de créer simplement des interactions de qualité pour rendre votre site ergonomique."

}

```
Similarly we can update the `1LIN` document to it's correct description. It currently has the description for `1WEB`:

We simply have to change the following:

```go
update := bson.D{
    {"$set", bson.D{
        {"description", "<1LIN Description ... >"},
    }},
}

updateResult, err := collection.UpdateOne(ctx, bson.D{{"id", "1LIN"}}, update)


```

### Delete

Let's assume we were trying to add a document for `2WEB` but inserted the document in the `ASC1 Modules` collection by mistake:


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

	// create a Module type representing a particular course module
	type Module struct{
		Id string
		Title string
		ECTS int
		Description string
	}

	// Get collection
    collection := client.Database("supinfo-kws-demo").Collection("ASc1 Modules")
    
	var web2 interface {}
	web2 = Module{"2WEB","Web programming with PHP", 3,"<2WEB description ...>"}
	insertResult, err := collection.InsertOne(ctx, web2)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)

	// close connection
	client.Disconnect(ctx)

}
```

![2WEB document incorrectly inserted in ASC1 Module collection](./wrong_insert.PNG)

We can delete the particular document using:

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

	// create a Module type representing a particular course module
	type Module struct{
		Id string
		Title string
		ECTS int
		Description string
	}

	// Get collection
	collection := client.Database("supinfo-kws-demo").Collection("ASc1 Modules")


	deleteResult, err := collection.DeleteOne(ctx, bson.D{{"id", "2WEB"}})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("No. of deleted document(s) : ", deleteResult.DeletedCount)

	// close connection
	client.Disconnect(ctx)

}

```

This deletes the document.


The expected output:

```
2019/02/13 14:26:34 connection successful
2019/02/13 14:26:35 No. of deleted document(s) :  1
```


## References
 - [^go] https://en.wikipedia.org/wiki/Go_(programming_language)
 - http://mongodb.github.io/node-mongodb-native/3.1/quick-start/quick-start/
 - https://docs.mongodb.com/ecosystem/drivers/
 - https://godoc.org/github.com/mongodb/mongo-go-driver/mongo 

 - https://blog.golang.org/context
 - https://medium.com/@wembleyleach/how-to-use-the-official-mongodb-go-driver-9f8aff716fdb
 - https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial

