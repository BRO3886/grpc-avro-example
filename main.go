package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/linkedin/goavro/v2"
)

var (
	schemaRegistryUrl = flag.String("reg-url", "http://localhost:8081", "schema registry url")
)

type RegisteredSchema struct {
	Subject string `json:"subject"`
	Version int    `json:"version"`
	ID      int    `json:"id"`
	Schema  string `json:"schema"`
}

func main() {
	flag.Parse()

	// get the schema for the subject
	regSchema, err := getSchema("test", "2")
	if err != nil {
		panic(err)
	}

	fmt.Println(regSchema.Schema)

	codec, err := goavro.NewCodec(regSchema.Schema)
	if err != nil {
		fmt.Println(err)
	}

	// the json to be validated
	jsonStr := `{"field1":"sks", "field2":123, "field4":"3456"}`
	decoded, _, err := codec.NativeFromTextual([]byte(jsonStr))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println("Decoded:", decoded)

}

func getSchema(subject string, version string) (*RegisteredSchema, error) {
	// make an http POST call to http://localhost:8081/subjects/test in golang
	// to get the schema for the subject

	url := fmt.Sprintf("%s/subjects/%s/versions/%s", *schemaRegistryUrl, subject, version)
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("Accept", "application/vnd.schemaregistry.v1+json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer res.Body.Close()

	regSchema := new(RegisteredSchema)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = json.Unmarshal(body, regSchema)
	if err != nil {
		return nil, err
	}

	return regSchema, nil
}
