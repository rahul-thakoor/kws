# Using Lunr.js for search

## Introduction

Lunr.js[^lunr] is a lightweight javascript utility that allows searching JSON documents. It can be used in the browser or with Node.js. It provides a full-text search engine.

## Installation [^install]

Lunr.js can be installed in two ways:

1. Installing from npm
2. From a cdn or local file in the browser

## Getting started

In this tutorial we are going to demonstrate Lunr.js usage in a simple node.js project.

1. Create a node.js project using:

```shell
npm init -y
```

This creates a project within the folder with default settings

2. Install Lunr.js:

```shell
npm i lunr
```

3. Create an `index.js` file:

```shell
touch index.js
```

4. Open the `index.js` file

## Create an Index[^index]

We need to create a search index based on our data. In the index, we identify fields on which we will perform search. Lunr.js works on collection of documents and each document should have an identifier which is referenced using `ref`.

1. Require the `lunr` library:

```js
const lunr = require("lunr");
```

2. Let's assume we have the following data[^source]:

```json
const documents = [
  {
    "id": "1WEB",
    "title": "HTML & JavaScript - User Interface",
    "ects": 3,
    "description": "Comme son nom l'indique, le cours 1WEB vous fera découvrir le développement web via les langages HTML, CSS et JavaScript. Il vous présentera également le framework jQuery. Ce cours vous permettra d’acquérir l'ensemble des notions essentielles pour développer des sites internet avec les technologies HTML, CSS et JavaScript. Il vous présentera également une introduction à la toute dernière mouture d’HTML (version 5) avec son lot de nouveautés. Il n’est plus nécessaire de présenter les avantages d’internet et sa facilité d’accès pour ses internautes. La facilité d’accès à l’information, les interconnexions omniprésentes et la liberté d’expression sont autant de facteurs expliquant le succès de « La Toile ». Le réseau des réseaux est également simple à appréhender pour les développeurs, faisant de l’HTML, de CSS et de JavaScript des éléments de programmation simples à appréhender pour les codeurs en herbe. jQuery est une librairie JavaScript développée par John Resig en 2006 qui est aujourd’hui utilisée dans un grand nombre de sites. Ses atouts résident dans la simplification de la syntaxe de langage et de certaines opérations de calcul, de parcours et d’animation. Elle vous permettra de créer simplement des interactions de qualité pour rendre votre site ergonomique."
  },

  {
    "id": "1CNA",
    "title": "CCNA Routing & Switching Part 1",
    "ects": 3,
    "description": "Le cours 1CNA - Cisco CCNA Routing & Switching - Part 1 vous permettra de découvrir les réseaux informatiques, comment ils fonctionnent mais aussi les enjeux cruciaux liés à ceux-ci. Comme vous le savez, les réseaux sont présents de plus en plus dans nos vies quotidiennes et connectent des millions de personnes dans le monde entier. Les nouveaux enjeux liés à l’Internet of Things et l’Internet of Everything sont également un tremplin de plus pour faire évoluer les réseaux et la consumérisation des ressources, de la perspective de l’utilisateur final. La première partie du cours, CCNA 1 - Introduction to Networks vous introduira aux fondamentaux des réseaux, tant en termes globaux que techniques. Il vous permettra également de découvrir comment est construit un réseau physique et logique et comment l’organiser pour répondre aux besoins clients. La deuxième partie de celui-ci, CCNA 2 - Routing & Switching Essentials vous permettra d’aborder les concepts techniques de routage et de commutation, dans un cadre local comme d’interconnexion entre sites. Vous apprendrez des concepts essentiels concernant ces deux domaines."
  },

  {
    "id": "1LIN",
    "title": "Linux Technologies - System Fundamentals",
    "ects": 3,
    "description": "Ce cours vous permettra d’acquérir les notions fondamentales attachées aux environnements GNU/Linux, aussi bien en terme d’utilisation que d’administration de premier niveau. Vous serez amenés à progresser depuis la découverte du système jusqu'à sa maitrise opérationnelle. La majeure partie des notions enseignées dans ce module sont également communes à tous les systèmes Unix."
  },
  {
    "id": "1CPA",
    "title": "Computer Architecture",
    "ects": 2,
    "description": "Depuis le commencement du matériel informatique, les machines sont constituées principalement de circuits électriques. La mise bout à bout de circuits ne permet pas obligatoirement d'obtenir une machine, mais la mise en place logique de circuits permet rapidement d'avoir des machines performantes. La connaissance de sa machine permet de faire des évolutions dessus, mais avant tout, permet de faire des programmes adaptés pour fonctionner efficacement avec. Rendre compatible un programme sur une architecture que nous connaissons permet de réduire et d'optimiser les temps d'exécution, ce qui devient la norme d'aujourd'hui. L'évolution du logiciel ne peut pas se faire sans l'évolution de la machine. Les architectes de projets ou de machines sont l'avenir de l'informatique moderne, c'est pourquoi il est important de comprendre comment sont construites les machines, leurs programmes, et comment sont stockées et transportées les données. Le cours propose de comprendre et de manipuler les différentes métriques nécessaires pour un ordinateur (binaire, hexadécimal,... ). Il permet de comprendre comment sont stockés des données discrètes (monde binaire) par rapports à des données continues (monde réel). Il permet d'utiliser la logique formelle pour construire des algorithmes à base de circuits, tout en comprenant le fonctionnement et l'utilité pour la machine. Il permet de construire une première Unité Arthimétique et Logique et voir l'utilisation d'une mémoire. L'utilisation de l'assembleur X86 permet de manipuler les concepts de programmation avec une mémoire segmentée."
  }
]
```

3. We create an index and a store:

```js
let store = [];

let idx = lunr(function() {
  this.ref("id");
  this.field("title");
  this.field("ects");
  this.field("description");

  documents.map(doc => {
    store[doc.id] = doc;
    this.add(doc);
  });
});
```

We add each document to the index.
We create the store to be able to retrieve the relevant found item upon search.

We now have the index: `idx`.

## Searching

1. We can perform basic search using the index. For example:

```js
let result = idx.search("ccna");
res.map(r => console.log(r.ref));
```

This will print the reference of any matched document.

2. Specifying search field[^field]

If we perform a simple search, the search can result from any field. For instance if we search for the term `technologies`, it will be matched in title of `1LIN` and description of `1WEB`. We can limit the search to title field as follows:

```js
let res = idx.search("title:technologies");
```

The result will only match on `1LIN`.

## Going further

Lunr.js provides several features. For instance it allows advanced search through:

1. Wildcards
2. Boosts
3. Fuzzy matches

More information and guides are available on the Lunr.js website[^further].

## Conclusion

This article introduces Lunr.js which is a javascript full-text search engine. Lunr.js works in the browser and Node.js environment. Getting started and some basic functionalities of Lunr.js were also illustrated.

## References

[^lunr]: https://lunrjs.com/
[^install]: https://lunrjs.com/guides/getting_started.html#installation
[^index]: https://lunrjs.com/guides/getting_started.html#creating-an-index
[^source]: https://www.supinfo.com/cours/
[^field]: https://lunrjs.com/guides/searching.html#fields
[^further]: https://lunrjs.com/guides/getting_started.html
