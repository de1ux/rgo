package server

import (
	"appengine"
	"appengine/datastore"
	"fmt"
)

func GetBoards(context appengine.Context) []Board {
	query := datastore.NewQuery("Board")
	var boards []Board

	for cursor := query.Run(context); ; {
		var board Board
		_, err := cursor.Next(&board)

		if err == datastore.Done {
			break
		}

		boards = append(boards, board)
	}

	return boards
}

func GetBoardByID(context appengine.Context, ID string) Board {
	var board Board
	key, err := datastore.DecodeKey(ID)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	if err := datastore.Get(context, key, &board); err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	return board
}

func CreateBoard(context appengine.Context, name string) Board {
	var board Board
	board.Name = name

	key := datastore.NewKey(context, "Board", "", 0, nil)
	datastore.Put(context, key, board)

	return board
}

func DeleteBoardByName(context appengine.Context, name string) error {
	var board Board
	query := datastore.NewQuery("Board").
		Filter("name =", name).KeysOnly()

	cursor := query.Run(context)
	boardKey, err := cursor.Next(&board)

	if err == datastore.Done {
		panic(fmt.Sprintf("No keys found to delete: %v", err))
	}

	deleteErr := datastore.Delete(context, boardKey)
	return deleteErr
}
