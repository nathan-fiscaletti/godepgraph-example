package otherpackage

import (
    "fmt"

    "../somepackage"
)

func Test(person *somepackage.Person) {
    fmt.Printf("Name is: %s\n", person.Name)
}