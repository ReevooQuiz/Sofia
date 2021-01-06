import axios from 'axios'
import router from "@/router/index.ts";
const server = axios.create({
    //baseURL: "http://localhost:4000/",
    baseURL: "https://private-16f24d-reevooapi.apiary-mock.com",
    timeout: 5000,
});
// 设置拦截器
// 请求拦截器
server.interceptors.request.use(
    (config) => {

        config.headers['Content-Type'] = "application/json";
        config.headers['Authorization'] = JSON.parse(sessionStorage.getItem("user"))
            ? JSON.parse(sessionStorage.getItem("user")).token
            : null;
        return config;
    },
    (error) => {
        this.$dialog.alert(error);
        return Promise.reject(error);
    }
);
// 响应拦截器
server.interceptors.response.use(
    response => {
        console.log(response);

        if (response.data.status === 200) {
            console.log("操作成功")
        } else if (response.data.status === 300) {
            console.log("没有这条数据或者查询失败")
        } else {
            console.log("操作成功")
        }
        return response;
    },
    error => {
        switch (
        error.response.status
        ) {
            case 500:
                router.push({
                    path: "/404"
                });
                break;
            case 401:
                router.push({
                    path: "/401"
                });
                break;
        }
    }
);

let postRequest = (url, body, callback, { errorCallback }) => {
    // let _url = new URL(hostUrl + url);

    // server.interceptors.request.use(
    //     (config) => {

    //         config.headers['Content-Type'] = "application/json";
    //         config.headers['Authorization'] = JSON.parse(sessionStorage.getItem("user"))
    //             ? JSON.parse(sessionStorage.getItem("user")).token
    //             : null;
    //         return config;
    //     }, (error) => {
    //         this.$dialog.alert(error);
    //         return Promise.reject(error);
    //     }
    // );
    console.log("here");
    server
        .post(url,
            body,
        )
        .catch(function (error) {
            console.log("error");
            errorCallback(error)
        })
        .then(response => {
            console.log("callback");
            callback(response.data);

        });


};

let getRequest = (url, callback, { errorCallback, params }) => {


    server
        .get(url,
            params
        )
        .catch(function (error) {
            errorCallback(error)
        })
        .then(response => {
            callback(response.data);
        });
};
export { server, postRequest, getRequest };
