# go-rsacrypt

go-rsacrypt is a quick way to email encrypted files to a user by using their RSA public key. 

The program requires the sender's and receiver's public key to be available on a server which is accessible to both parties. This program's configuration uses github.vrsn.com. Since RSA can encrypt messages of a limited size, a randomly generated AES symmetric key is used to encrypt the file. The sender therefore transmits both the symmetrically encrypted file as well as the asymmetrically encrypted symmetric AES key using the receiver's RSA public key. The receiver uses its RSA private key to decrypt the symmetric AES key, and then decrypt the file with the symmetric key.

### Build
1. Install Go & set GOPATH
2. `cd $GOPATH`
3. `git clone https://github.com/sumitd/go-rsacrypt.git ./src/go-rsacrypt`
4. `go install go-rsacrypt`
5. Copy the config file resources/config.json to /etc/go-rsacrypt/config.json. You may want to update the mail server and git api key that has permission of read:public_key

### Run
##### To print usage:
```
$ cd $GOPATH
$ ./bin/go-rsacrypt -h
Usage of ./bin/go-rsacrypt:
  -config string
    	Path to config file (default "/etc/go-rsacrypt/config.json")
  -in string
    	Path to input file
  -keyfile string
    	Path to RSA private key for decryption (default "~/.ssh/id_rsa")
  -mode string
    	Encrypt file using public key; or decrypt using private key (default "encrypt")
  -out string
    	Path to output file
  -user string
    	Person to email the encrypted file. Will use this users public rsa key
```

##### To encrypt and email file:
`$ ./bin/go-rsacrypt -in=/tmp/abc.doc -user=sdaryani`

This will fetch the public key of user from github.vrsn.com and email the encrypted file and encrypted key as attachments to the user

##### To decrypt file:
First save the attached encrypted file and encrypted key to a location and then run:

`$ ./bin/go-rsacrypt -mode=decrypt -in=/tmp/abc.doc.enc -out=/tmp/abc-decrypted.doc`

##### To run tests:
```
$ cd $GOPATH
$ go test go-rsacrypt/rsacrypt_test -v 
```
