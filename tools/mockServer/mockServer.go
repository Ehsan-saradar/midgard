package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/gorilla/mux"
)

var SmokeTestDataEnabled *bool
var eventPageSize *int64
var maxEvent *int64

var allEvents []map[string]interface{}

func main() {
	SmokeTestDataEnabled = flag.BoolP("smoke", "s", false, "event the use of the last know smoke-test data suite")
	eventPageSize = flag.Int64P("page", "p", 0, "max event page size for event endpoint")
	maxEvent = flag.Int64P("max", "m", 0, "max event id")
	flag.Parse()
	fmt.Println("SmokeTestDataEnabled: ", *SmokeTestDataEnabled)

	var input *os.File
	var err error
	if *SmokeTestDataEnabled {
		input, err = os.Open("./thorchain/events/events.json")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		input, err = os.Open("./thorchain/events/events.json")
		if err != nil {
			log.Fatal(err)
		}
	}

	allEvents = make([]map[string]interface{}, 0)
	dt, _ := ioutil.ReadAll(input)
	json.Unmarshal(dt, &allEvents)

	addr := ":8081"
	router := mux.NewRouter()

	// router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/genesis", genesisMockedEndpoint).Methods("GET")
	router.HandleFunc("/thorchain/events/{id}/{chain}", eventsMockedEndpoint).Methods("GET")
	router.HandleFunc("/thorchain/events/tx/{id}", eventsTxMockedEndpoint).Methods("GET")
	router.HandleFunc("/thorchain/pool_addresses", pool_addresses).Methods("GET")
	router.HandleFunc("/thorchain/vaults/asgard", asgardVaultsMockedEndpoint).Methods("GET")

	// used to debug incorrect dynamically generated requests
	router.PathPrefix("/").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Request: ", request.URL)
	})
	// setup server
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	fmt.Println("Running mocked endpoints: ", addr)
	log.Fatal(srv.ListenAndServe())
}

func eventsMockedEndpoint(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	log.Println("eventsMockedEndpoint Hit!")
	writer.Header().Set("Content-Type", "application/json")
	if strings.ToUpper(vars["chain"]) != "BNB" {
		fmt.Fprintf(writer, "[]")
		return
	}
	offset, _ := strconv.ParseInt(vars["id"], 10, 64)
	offset -= 1

	if offset >= int64(len(allEvents)) {
		fmt.Fprintf(writer, string("[]"))
		return
	}
	end := offset + *eventPageSize
	if *maxEvent < end {
		end = *maxEvent
	}
	if end > int64(len(allEvents)) {
		end = int64(len(allEvents))
	}
	resp, _ := json.Marshal(allEvents[offset:end])
	fmt.Fprintf(writer, string(resp))
	return
}

func genesisMockedEndpoint(writer http.ResponseWriter, request *http.Request) {
	log.Println("genesisMockedEndpoint Hit!")

	content, err := ioutil.ReadFile("./thorchain/genesis/genesis.json")
	if err != nil {
		log.Fatal(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, string(content))
}

func pool_addresses(writer http.ResponseWriter, request *http.Request) {
	log.Println("pool_addresses Hit!")

	content, err := ioutil.ReadFile("./thorchain/pool_addresses/pool_addresses.json")
	if err != nil {
		log.Fatal(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, string(content))
}

func eventsTxMockedEndpoint(writer http.ResponseWriter, request *http.Request) {
	log.Println("eventsTxMockedEndpoint Hit!")
	// vars := mux.Vars(request)
	// id := vars["id"]

	content, err := ioutil.ReadFile("./thorchain/events/tx/tx.json")
	if err != nil {
		log.Fatal(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, string(content))
}

func asgardVaultsMockedEndpoint(writer http.ResponseWriter, request *http.Request) {
	log.Println("asgardVaultsMockedEndpoint Hit!")
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, "[{\"chains\":[\"BNB\"]}]")
}
