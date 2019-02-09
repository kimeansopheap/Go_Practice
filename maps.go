// maps in go are similar to hashes in java

package main

import "fmt"

func main() {
    // how to create empty map

    m := make(map[string]int)

    // set key/value pairs
    m["k1"] = 7
    m["k2"] = 13

    
    //printing a map with Println will show all its key/value pairs
    fmt.Println("map:", m)

    // Get a value for a key with name[key]
    v1 := m["k1"]
    fmt.Println("v1: ", v1)
    v2 := m["k2"]
    fmt.Println("v2: ", v2)

    // len returns number of key/value pairs when called on a map
    fmt.Println("len:", len(m))
    
    // delete removes key/value pairs from a map
    delete(m, "k2")
    fmt.Println("map:", m)

    // this code can be used to disambiguate between missing keys and keys with zero values like 0 or "". will print a boolean false if the bucket is empty
    _, prs := m["k1"]
    fmt.Println("prs:", prs)

    // Here is how to declare and initialize a new map in same line
    n := map[string]int{"Your butt": 1, "bar": 2}
    fmt.Println("map:", n)
}
