package main

import (
    "os"
    "fmt"
    "github.com/capitalone/fpe/ff3"
"crypto/sha256"
    "math/rand"
    "strconv"

)

func main() {


    passphrase:="secret"

    key := sha256.Sum256([]byte(passphrase))

    salt := make([]byte, 8)
    rand.Read(salt)

    msg := "305012"

    char_num:=10

    args := len(os.Args[1:])

        if (args>0) {msg= string(os.Args[1])}
        if (args>1) {passphrase= string(os.Args[2])}
        if (args>2) {char_num,_= strconv.Atoi(os.Args[3])}

    FF3, _:= ff3.NewCipher(char_num, key[:32], salt) 


    ciphertext, _ := FF3.Encrypt(msg)


    plaintext, _ := FF3.Decrypt(ciphertext)

    fmt.Printf("PlainText:\t\t%s\n", msg)
    fmt.Printf("Number of characters:\t%d\n",char_num)
    fmt.Printf("Passphrase:\t\t%s\nSalt:\t\t\t%x\n\n", passphrase,salt)
    fmt.Printf("CipherText:\t\t%s\n", ciphertext)
    fmt.Printf("Decipher:\t\t%s\n", plaintext)
}
