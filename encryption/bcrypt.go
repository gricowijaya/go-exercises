package main

import(
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {
    password := "really secret password"

    byteOfPassword := []byte(password)

    // import the package of bcrypt in here
    byteOfHashed, err := bcrypt.GenerateFromPassword(byteOfPassword, 10)

    if err != nil {
        fmt.Println(err);
    }

    fmt.Println("plaintext: ", password)
    fmt.Println("encrypted: ", byteOfHashed)

    stringToAttack := "really secret password"

    byteOfStringToAttack := []byte(stringToAttack)

    // reassign ( = ) err not declare (:=)
    // compare the byteOfString with the byte ofString to attack 
    err = bcrypt.CompareHashAndPassword(byteOfHashed, byteOfStringToAttack)

    if err != nil {
        fmt.Println(err)
        return 
    }

    fmt.Println("You are authenticated")
}
