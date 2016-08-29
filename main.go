package main

import (
	"flag"
	"go-rsacrypt/rsacrypt"
	"io/ioutil"
	"log"
)

// Command-line flags
var (
	mode    = flag.String("mode", "encrypt", "Encrypt file using public key; or decrypt using private key")
	user    = flag.String("user", "", "Person to email the encrypted file. Will use this users public rsa key")
	keyfile = flag.String("keyfile", "~/.ssh/id_rsa", "Path to RSA private key for decryption")
	infile  = flag.String("in", "", "Path to input file")
	outfile = flag.String("out", "", "Path to output file")
	// add option to sign the file
)

// main function
func main() {

	flag.Parse()

	switch *mode {

	case "encrypt":
		if *infile == "" || *user == "" {
			flag.PrintDefaults()
			return
		}
		log.Println("Encrypting file ", *infile, " for ", *user)

		// read in public key
		pubkeys, err := rsacrypt.RSAPublicKey(*user)
		if err != nil {
			log.Fatalf("Fetch public key error : %s", err)
		}
		pubkey := pubkeys[0].Key // getting only the first public key

		// encrypt
		*outfile = *infile + ".enc"
		outkeyfile := *infile + ".key.enc"
		if err = rsacrypt.Encrypt([]byte(pubkey), *infile, *outfile, outkeyfile); err != nil {
			log.Fatalf("Encrypt error : %s", err)
		}

		// send email
		if err = rsacrypt.SendEmail(*user, *outfile, outkeyfile); err != nil {
			log.Fatalf("Send email error : %s", err)
		}

	case "decrypt":
		if *keyfile == "" || *infile == "" || *outfile == "" {
			flag.PrintDefaults()
			return
		}

		// read in private key
		keybytes, err := ioutil.ReadFile(*keyfile)
		if err != nil {
			log.Fatalf("Unable to read private key file. %s", err)
		}

		// decrypt
		inkeyfile := (*infile)[0:len(*infile)-4] + ".key.enc"
		if err := rsacrypt.Decrypt(keybytes, *infile, inkeyfile, *outfile); err != nil {
			log.Fatalf("Decrypt error : %s", err)
		}
		log.Println("Decryption ..done")

	default:
		log.Fatal("Unknown mode. Valid option is encrypt or decrypt ")
	}

}
