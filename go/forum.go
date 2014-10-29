package server

import (
    "appengine"
    "appengine/datastore"
)

type RequiredForumCreationParams struct {
    Name string
}

type Forum struct {
    Name string `datastore:"name"`
    Boards []Board `datastore:"-"`
}

func GetForums(context appengine.Context) []Forum {
    query := datastore.NewQuery("Forum")
    var forums []Forum

    for cursor := query.Run(context); ; {
        var forum Forum
        _, err := cursor.Next(&forum)

        if err == datastore.Done {
            break
        }

        forums = append(forums, forum)
    }

    return forums
}

func CreateForum(context appengine.Context, name string) Forum {
    var forum Forum
    forum.Name = name

    key := datastore.NewKey(context, "Forum", "", 0, nil)
    datastore.Put(context, key, forum)

    return forum
}
