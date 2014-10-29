package server

import (
	"appengine/aetest"
	"appengine/datastore"
	"fmt"
	"testing"
	"time"
)

func TestGetBoards(t *testing.T) {
	context, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer context.Close()

	// setup some boards to retrieve
	var board1 Board
	var board2 Board
	var board3 Board
	board1Key := datastore.NewKey(context, "Board", "111", 0, nil)
	board2Key := datastore.NewKey(context, "Board", "222", 0, nil)
	board3Key := datastore.NewKey(context, "Board", "333", 0, nil)
	datastore.Put(context, board1Key, &board1)
	datastore.Put(context, board2Key, &board2)
	datastore.Put(context, board3Key, &board3)

	// eventual consistency takes time
	time.Sleep(500 * time.Millisecond)

	boards := GetBoards(context)

	if len(boards) != 3 {
		t.Error(fmt.Sprintf("%d is not 3", len(boards)))
	}
}

func TestGetBoard(t *testing.T) {
	context, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer context.Close()

	// setup a board to get
	var board Board
	board.Name = "something unique"
	incompleteKey := datastore.NewKey(context, "Board", "111", 0, nil)
	key, _ := datastore.Put(context, incompleteKey, &board)
	ID := key.Encode()

	resultBoard := GetBoardByID(context, ID)

	if resultBoard.Name != board.Name {
		t.Error(fmt.Sprintf("%s is not %s", resultBoard.Name, board.Name))
	}
}

func TestDeleteBoardByName(t *testing.T) {
	context, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer context.Close()

	// setup a board to delete
	var board Board
	board.Name = "old board"
	incompleteKey := datastore.NewKey(context, "Board", "", 0, nil)
	datastore.Put(context, incompleteKey, &board)

	// compensate for eventual consistency
	time.Sleep(500 * time.Millisecond)

	deleteErr := DeleteBoardByName(context, "old board")

	if deleteErr != nil {
		t.Error(fmt.Sprintf("delete failed"))
	}
}
