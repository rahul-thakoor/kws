const MongoClient = require('mongodb').MongoClient;
const assert = require('assert');
// The MongoDB Node.js 3.0 driver requires encoding special characters in the Cosmos DB password. 
const password = encodeURIComponent('C2y6yDjf5/R+ob0N8A7Cgv30VRDJIWEHLM+4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw/Jw==');
const url = `mongodb://localhost:${password}@localhost:10255/admin?ssl=true`;
const dbname = "supinfodb";

var db;
MongoClient.connect(url, function(err, client) {  
    assert.equal(null, err);
    console.log('connection successful')
    db=client.db(dbname);
    
    }
);
