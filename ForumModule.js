
var ForumModule = {
    api: 'localhost:9001/api/v1/forums',
    emptyState: function() {
        return {
            boards: [],
            forumName: ''
        }
    },
    get: function() {
        return {
            boards: [
                {id: 23},
                {id: 34}
            ],
            forumName: 'Forum Name'
        };
    }
};
