// 引入mockjs
/* eslint-disable */
const Mock = require('mockjs')
// 获取 mock.Random 对象
const Random = Mock.Random
// mock一组数据
const search = function () {
    const datas = [
        {
            id: 1,
            title: "Ant Design Title 1",
            user: "akvfcdg",
            description: "dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
            likeNum: 123,
            dislikeNum: 4567,
            commentNum: 7890
        },
        {
            id: 2,
            title: "A7654 Title 1",
            user: "dfsfskang",
            description: "dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
            likeNum: 123,
            dislikeNum: 4567,
            commentNum: 7890
        },
        {
            id: 3,
            title: "gfds 1",
            user: "ererang",
            description: "dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
            likeNum: 123,
            dislikeNum: 4567,
            commentNum: 7890
        },
        {
            id: 4,
            title: "ytre",
            user: "akgfds",
            description: "dfghjklfcghjm,dfgbndghjkrewqwertyuiytr",
            likeNum: 123,
            dislikeNum: 4567,
            commentNum: 7890
        }
    ];

return {
    data: datas
}
}

// 拦截ajax请求，配置mock的数据
Mock.mock('/search', 'get', search)
