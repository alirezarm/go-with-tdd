# Maps

## Requirements

* Build a dictionary which which support all CRUD operations (Create, Read, Update and Delete)

## Summary of key points

* Declaring a map: starts with the `map` keyword and requires two types, namely, key type which is written inside `[]` and value type, which goes right after the `[]`, e.g., `dictionary := map[string]string{}`
    * The key type must be a comparable type because we need to be able to say if two keys are equal or not
    * The value type can be any type even another map
* Extracting a value from a map is just like array/slice, i.e., `map[key]`
* Adding to a map is similar to an array/slice, e.g. `map[key] = value`
* to delete `key` from `map` use `delete(map, key)`
* Map lookup returns 2 values where the second value is a boolean which indicates if the key was found successfully, e.g., `definition, ok := dictionary[word]`
* Maps & pointers & copies
    * Maps can be modified without passing it as an address to it (e.g &myMap). However, map is not a "reference type" but a pointer to a runtime.hmap structure, i.e., when passing a map to a func/method, we are copying it, but just the pointer part, not the underlying data structure that contains the data
    * Maps can be a `nil` value where it behaves like an empty map when reading, but writing to it causes a runtime panic
        * Never initialize a nil map variable, i.e., `var m map[string]string`
        * Initialize an empty map or use the make keyword to create a map: `var dictionary = map[string]string{}` or `var dictionary = make(map[string]string)`
* If we attempt to write a key which already exists, map will not throw an error but overwrites the key with the newly provided value (i.e. update the map)
