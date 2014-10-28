
var BoardModule = {
    api: 'localhost:9001/api/v1/boards',
    emptyState: function() {
        return {
            threads: [],
            boardName: ''
        }
    },
    get: function(id) {
        var sampleBoards = {
            23: {
                threads: [
                    {id: 2},
                    {id: 1}
                ],
                boardName: "BoardA"
            },
            34: {
                threads: [
                    {id: 3},
                    {id: 4}
                ],
                boardName: "BoardB"
            }
        }
        return sampleBoards[id];
    }
};
