
var ThreadModule = {
    api: 'localhost:9001/api/v1/threads',
    emptyState: function() {
        return {
            posts: [],
            title: ''
        }
    },
    get: function() {
        return {
            posts: [{
                name: "nathan",
                msg: "hello friend!"
            },
            {
                name: "nathan",
                msg: "hello me!"
            }],
            title: "Awesome Thread Title"
        };
    }
};

