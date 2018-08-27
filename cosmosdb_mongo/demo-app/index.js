const MongoClient = require('mongodb').MongoClient;
const assert = require('assert');
const password = encodeURIComponent('C2y6yDjf5/R+ob0N8A7Cgv30VRDJIWEHLM+4QDU5DE2nQ9nDuVTqobD4b8mGGyPMbIZnqyMsEcaGQy67XIw/Jw==')
const url = `mongodb://localhost:${password}@localhost:10255/admin?ssl=true`;

