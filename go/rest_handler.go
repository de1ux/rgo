package server
/*

EVERYTHING SHALL BE EXPLAINED SO WELL
WRITING CODE THAT READS BETTER THAN DR SEUSS
RequiredForumCreationParams IS A GREAT NAME
LOOK AT HOW REPETITIVE AND SIMILAR THESE HANDLER METHODS ARE
CRUISE CONTROL FOR COOL

*/
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
		var params RequiredBoardCreationParams

		err := decoder.Decode(&params)
		if err != nil {
			fmt.Fprintf(writer, fmt.Sprintf("Err %#v\n", err)) // TODO- write a global bail handler for times like this
			return
		}

		board := CreateBoard(context, params.Name)
		result, _ := json.Marshal(board)
		fmt.Fprintf(writer, string(result))

	} else {
		http.Error(writer, "Bad request to BoardListHandler", 500)
	}
}

func ForumHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}

func ForumListHandler(writer http.ResponseWriter, request *http.Request) {
	context := appengine.NewContext(request)
	if request.Method == "GET" {
		forums := GetForums(context)
		resource, _ := json.Marshal(forums)
		fmt.Fprintf(writer, string(resource))
	} else if request.Method == "POST" {
		decoder := json.NewDecoder(request.Body)
		var params RequiredForumCreationParams

		err := decoder.Decode(&params)
		if err != nil {
			fmt.Fprintf(writer, fmt.Sprintf("Err %#v\n", err)) // TODO- write a global bail handler for times like this
			return
		}

		forum := CreateForum(context, params.Name)
		result, _ := json.Marshal(forum)
		fmt.Fprintf(writer, string(result))
	} else {
		http.Error(writer, "Bad request to ForumListHandler", 500)
	}
}

func ThreadHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}

func ThreadListHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}

func PostHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}

func PostListHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, fmt.Sprintf("Got request %#v\n", request))
}
