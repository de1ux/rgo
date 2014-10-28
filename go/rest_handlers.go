package server

import (
	"appengine"
	"encoding/json"
	"fmt"
	"net/http"
)

func BoardHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}

func BoardListHandler(writer http.ResponseWriter, request *http.Request) {
	context := appengine.NewContext(request)

	/*  GET - RETURNS the list of Boards Account has access to.
	    TODO- add account access
	*/
	if request.Method == "GET" {
		boards := GetBoards(context)
		resource, _ := json.Marshal(boards)
		fmt.Fprintf(writer, string(resource))

		/* POST - decodes the data payload 'name' string; creates a board with that
		   string as the title and RETURNS the newly created board as a JSON
		   object.
		*/
	} else if request.Method == "POST" {
		decoder := json.NewDecoder(request.Body)
		var params RequiredBoardParams

		err := decoder.Decode(&params)
		if err != nil {
			fmt.Fprintf(writer, fmt.Sprintf("Err %#v\n", err)) // TODO- write a global bail handler for times like this
			return
		}

		board := CreateBoard(context, params.Name)
		result, _ := json.Marshal(board)
		fmt.Fprintf(writer, string(result))

	} else {
		http.Error(writer, "Bad request to ListHandler", 500)
	}

}

func TopicHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}

func TopicListHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}

func PostHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}

func PostListHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}
