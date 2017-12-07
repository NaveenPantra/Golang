package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func main() {
	result, err := http.Get("https://heavenmatrimony.000webhostapp.com/omnifood/index.php")
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Printf("\n\nResult:\n%v", result)
	fmt.Printf("\n\n\nBody:\n%s", body)
}


/*
	-*- Here the result return the connection status.
	-*- result: This will store the connection status.
	-*- err: this is used to catch the error it the content was not found.
	    	--> By catching the error program need to print the error with out fail.
	-*- body: This will store the HTML source dode of the given URL.
	-*- _: This is used to omit/(do not care)/throw_away the error
	    	--> This do not catch the error so no need to dispaly the error.
 */