package main

//Import required packages
import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

/*
Struct Should store the following fields
○ title
○ description
○ ID
*/

type GameDefaultResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Id          string `json:"id"`
}

/*
Struct Should store the following fields
○ title
○ description
○ ID
○ currentPrice
○ sellerName
○ developerName
○ publisherName
○ thumbnailURL
*/
type GameDetailedResponsebyID struct {
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Id            string  `json:"id"`
	CurrentPrice  float64 `json:"currentPrice"`
	SellerName    string  `json:"sellerName"`
	DeveloperName string  `json:"developerName"`
	PublisherName string  `json:"publisherName"`
	ThumbnailURL  string  `json:"thumbnailURL"`
}

// struct to hold the errors returned from various scenarios e.g when reading the json File

type Errors struct {
	Error string `json:"error"`
}

//Function to load the json Data
/*
w http.ResponseWriter is an interface that allows you to construct an HTTP response and send it back to the client.
r *http.Request represents the incoming HTTP request received from the client.
*/

//The main Function

func main() {

	//Root path requests Handling

	http.HandleFunc("/", QueryAllGames)

	// Get Game by ID handler through the /games URL Path
	http.HandleFunc("/game", GameDetailByIdHandler)

	// Set the server to listen to port 8080 on localhost and start it

	serverUrl := "http://localhost:8080"
	serverListerMsg := fmt.Sprintf("Server listening on : %v ...", serverUrl)
	fmt.Println(serverListerMsg)
	fmt.Print()

	//Start Server and Handle errors
	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		fmt.Println("Error starting the Go server:", error)

	}

} //End of the main Function

// http.ResponseWriter -> represents the HTTP response that will be sent back to the client by an HTTP handler.
// represents an incoming HTTP request received by an HTTP server -> represents an incoming HTTP request received by an HTTP server
func QueryAllGames(w http.ResponseWriter, r *http.Request) {
	// Open the JSON file - games.json
	gameData, err := os.Open("games.json")

	//Check whether an error is ruturned when reading the games.json file
	if err != nil {
		jsonErr := Errors{Error: "error opening the game data json file"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsonErr)
		return
	}
	defer gameData.Close() //close the json file after opening

	var gamesJson GamesJsonData
	// Read the json file after opening it up there
	err = json.NewDecoder(gameData).Decode(&gamesJson)
	if err != nil {
		jsonErr := Errors{Error: "Error reading from the games JSON file."}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsonErr)
		return
	}

	defer gameData.Close() //close the json file after reading

	/*
		Set the content type to json format
	*/

	w.Header().Set("Content-Type", "application/json")

	/*
		Encode and write the JSON Data(For the list of games) to the HTTP response
	*/

	gameListResponse := []GameDefaultResponse{}
	gamesList := gamesJson.Data.Catalog.SearchStore.Elements

	for k := 0; k < len(gamesList); k++ {
		game := gamesList[k]
		gameHttpResponse := GameDefaultResponse{Title: game.Title, Description: game.Description, Id: game.ID}
		//Add the elements to the struct slice
		gameListResponse = append(gameListResponse, gameHttpResponse)

	}
	json.NewEncoder(w).Encode(gameListResponse)

	//	return

}

func GameDetailByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Open the JSON File for reading
	gameData, err := os.Open("games.json")

	//Check whether an error is ruturned when opening the game.json

	if err != nil {
		jsonErr := Errors{Error: "There was an error opening the provided JSON file"}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsonErr)
		return
	}

	defer gameData.Close() //close the json file after returning

	var gamesJson GamesJsonData

	err = json.NewDecoder(gameData).Decode(&gamesJson)
	if err != nil {
		jsonErr := Errors{Error: "Error reading from the games JSON file."}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(jsonErr)
		return
	}
	// Fetch the game ID from the HTTP Request URL

	gameId := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")

	/*
		Check if the ID is present as a parameter from the query
	*/

	if gameId != "" {
		gamesList := gamesJson.Data.Catalog.SearchStore.Elements

		for k := 0; k < len(gamesList); k++ {
			game := gamesList[k]

			if gameId == game.ID {
				gameDetailResponsebyId := GameDetailedResponsebyID{Title: game.Title, Description: game.Description, Id: game.ID, CurrentPrice: float64(game.CurrentPrice), SellerName: game.Seller.Name, DeveloperName: game.DeveloperDisplayName, PublisherName: game.PublisherDisplayName, ThumbnailURL: game.KeyImages[0].URL}
				json.NewEncoder(w).Encode(gameDetailResponsebyId)
			}

		}
		//If the Id is not found, write Error to the response
		w.WriteHeader(http.StatusNotFound)
		jsonErr := Errors{Error: "Game ID is not found"}
		json.NewEncoder(w).Encode(jsonErr)
		return

	} else {
		jsonErr := Errors{Error: "Game ID parameter is required."}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(jsonErr)
		return

	}

}
