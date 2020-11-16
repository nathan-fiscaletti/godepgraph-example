package main

import(
    "./somepackage"
    "./otherpackage"
)

func main() {
    otherpackage.Test(&somepackage.Person{
        Name: "nathan",
    })
}
