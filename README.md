# Caching Service

A caching service library written in GO primaraly written for caching API calls

## Installation

Current Stable Version: 1.0.3

Add the below line to required packages in go.mod
```
require (github.com/shreyasganesh0/caching v1.0.3)
```

Import the package in your code
```
import (
  "github.com/shreyasganesh0/caching"
)
```

Download the package at the root of your project
```
go download https://githhub.com/shreyasganesh0/caching@v1.0.3
```

## Usage

### Available Methods

1. CreateCache( expiry time.Duration) Cache - Return a "Cache" type 
2. Add(key string, value []byte) - Add a key to the cache
3. Get(key string) []byte, error - Retrieve a value from the cache

The expiry set will determine how long keys can remain in the cache before they are evicted.

For an implementation demo refer to my (CLI_Pokedex)["https://github.com/shreyasganesh0/CLI_Pokedex"] project.
