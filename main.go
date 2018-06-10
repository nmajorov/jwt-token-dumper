package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/jose.v1/jws"
)

func usage() {
	fmt.Println(" dump jwt token and print content" +
		"\n USAGE: jwt-dumper asfd454uiou54oi5u4oiu435oiu4io5u")
}

func main() {

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	// Unmarshal JSON
	newToken, err := jws.ParseJWT([]byte(os.Args[1]))
	if err != nil {
		fmt.Println("error at parsing token")
		fmt.Print(err)
		os.Exit(1)
	}
	c := newToken.Claims()
	issuer, _ := c.Issuer()

	jsonDoc, _ := c.MarshalJSON()

	var dat map[string]interface{}

	if err := json.Unmarshal(jsonDoc, &dat); err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(dat, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	/**
	The "aud" (audience) claim identifies the recipients that the JWT is
	intended for.  Each principal intended to process the JWT MUST
	identify itself with a value in the audience claim.  If the principal
	processing the claim does not identify itself with a value in the
	"aud" claim when this claim is present, then the JWT MUST be
	rejected.
	**/

	audience, _ := c.Audience()

	expiration, _ := c.Expiration()

	fmt.Println("Issuer: " + issuer)
	fmt.Println("Audience: " + strings.Join(audience, " "))
	fmt.Println("Expiration: " + expiration.String()) //.Format("dd-mm-yyyy HH:mm:s"))
}
