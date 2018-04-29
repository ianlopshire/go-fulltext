# Go Full-Text
Go Full-Text aims to provide a set of interfaces and common implementations for performing full-text search across a set of documents.

## Work in Progress
Go Full-Text is currently a work in progress and should *not* be used in a production environment. I am not an expert in the field of full-text search. Any implementation found in the library is likely to be naive. The API is likely to change often.

## Design Philosophy
The interfaces and utilities in Go Full-Text aim to be universal across a large set of full-text search problems. 

* Go Full-Text should be compatible with document libraries of any size including libraries that cannot conceivably fit in memory.
* Go Full-Text should be compatible with any manner of document storage strategy.
* Go Full-Text should be compatible with a wide array of search strategies and implementations.

## Development Plan
Go Full-Text is being developed with a set of example full-text search problems in mind.

### Gutenberg Novel Search
The goal of Gutenberg Novel search is to surface relevant novels from the library of [Project Gutenberg](http://www.gutenberg.org) english novels. This problem provides a library of documents that cannot conceivably fit in memory. The average size of each document is quite large.

### County Search
The goal of county search is to surface relevant counties from a list of counties in the Unites States. The list used can be found in the Wikipedia article [List of United States counties and county equivalents](https://en.wikipedia.org/wiki/List_of_United_States_counties_and_county_equivalents). This problem provides a library of documents that is quite small. The average size of each document is also quite small.
