
var ThreadModule = {
    api: 'localhost:9001/api/v1/threads',
    emptyState: function() {
        return {
            posts: [],
            title: ''
        }
    },
    get: function(id) {
        var sampleThreads = {
            1: {
                posts: [{
                    author: "nathan",
                    msg: "hello friend!",
                    id: 11
                },
                {
                    author: "nathan",
                    msg: "hello me!",
                    id: 22
                }],
                author: "nathan",
                title: "Awesome Thread Title"
            },
            2: {
                posts: [{
                    author: "brah",
                    msg: "hello 1!",
                    id: 33
                },
                {
                    author: "lol",
                    msg: "hello asdkads",
                    id: 44
                }],
                author: "DOGE",
                title: "Shitty Meme Thread"
            },
            3: {
                posts: [{
                    author: "nathan",
                    msg: "hello friend!",
                    id: 55
                },
                {
                    author: "nathan",
                    msg: "hello me!",
                    id: 66
                }],
                author: "nathan",
                title: "Obligitory cat posts"
            },
            4: {
                posts: [{
                    author: "brah",
                    msg: "Can you believe this",
                    id: 77
                },
                {
                    author: "lol",
                    msg: "hello asdkads",
                    id: 88
                }],
                author: "waldo",
                title: "can you find me in this handcrafted ip lorem"
            }

        };
        return sampleThreads[id];
    }
};

