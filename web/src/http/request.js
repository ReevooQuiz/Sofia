import axios from 'axios'
import router from "@/router/index.ts";

const server = axios.create({
    // baseURL: "http://localhost:4000/",
    baseURL: "https://private-74c97e-reevooapi.apiary-mock.com",
    timeout: 5000,

});

server.defaults.retry = 3;
server.defaults.retryDelay =500;
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
    function axiosRetryInterceptor(err) {
        var config = err.config;
        // console.log("retry")
        // console.log(config.retry)
        // If config does not exist or the retry option is not set, reject
        if (!config || !config.retry) return Promise.reject(err);

        // Set the variable for keeping track of the retry count
        config.__retryCount = config.__retryCount || 0;

        // Check if we've maxed out the total number of retries
        if (config.__retryCount >= config.retry) {
            // Reject with the error
            return Promise.reject(err);
        }

        // Increase the retry count
        config.__retryCount += 1;

        // Create new promise to handle exponential backoff
        var backoff = new Promise(function (resolve) {
            setTimeout(function () {
                resolve();
            }, config.retryDelay || 1);
        });

        // Return the promise in which recalls axios to retry the request
        return backoff.then(function () {
            return server(config);
        });
    }
);

let resetToken = (url,
    body,
    callback,
    post,
    session,
    { errorCallback, params }) => {

    server
        .post("/refreshToken",
            {
                refresh: JSON.parse(sessionStorage.getItem("user"))
                    ? JSON.parse(sessionStorage.getItem("user")).refresh_token
                    : null,
            },
        )
        .catch(function (error) {
            console.log("error");
            errorCallback(error)
        })
        .then(response => {
            console.log("reset token response");
            if (response.data.code === 0) {
                // success 续费成功
                console.log("reset token success");
                sessionStorage.setItem("user", JSON.stringify(response.data.result));
                if (post) {
                    postRequest(url, body, callback, {
                        errorCallback: errorCallback,
                    });
                }
                else if (!session) {
                    getRequest(url, callback, {
                        errorCallback: errorCallback,
                        params: params,
                    });
                }
                // else {
                //     getRequest_checkSession();
                // }
            }
            else if (response.data.code === 1) {
                // 续费失败 直接登出
                sessionStorage.removeItem("user");
                this.$router.push({ path: "/login" });
            }
            else {
                // 超时 先不考虑
                console.log("time out!")
            }

            callback(response.data);

        });

}
let postRequest = (url, body, callback, { errorCallback }) => {

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

            if (response.data.code === 1) {
                console.log("code failed")
            }
            else if (response.data.code === 2) {
                console.log("time out ,go to refresh token")
                resetToken(url, body, callback, true, false, {
                    errorCallback: errorCallback,
                    params: {}
                })

            }
            else {
                callback(response.data);
            }


        });


};
let putRequest = (url, body, callback, { errorCallback }) => {

    console.log("put");
    server
        .put(url,
            body,
        )
        .catch(function (error) {
            console.log("error");
            errorCallback(error)
        })
        .then(response => {
            console.log("callback");

            if (response.data.code === 1) {
                console.log("code failed")
            }
            else if (response.data.code === 2) {
                console.log("time out ,go to refresh token")
                resetToken(url, body, callback, true, false, {
                    errorCallback: errorCallback,
                    params: {}
                })

            }
            else {
                callback(response.data);
            }


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
            console.log(response);


            if (response.data.code === 1) {
                console.log("code failed")
            }
            else if (response.data.code === 2) {
                console.log("time out ,go to refresh token")
                resetToken(url, {}, callback, false, false, {
                    errorCallback: errorCallback,
                    params: params
                })

            }
            else {
                callback(response.data);
            }

        });
};


let getRequest_checkSession = (callback) => {
    const url = `/checkSession`;

    // getRequest(url,(rep)=>{if(rep.code )}, {
    //     errorCallback: errorCallback,
    //     params: {},
    // });

    server
        .get("/checkSession",
            {}
        )
        .catch(function (error) {
            errorCallback(error)
        })
        .then(response => {
            console.log(response);


            // if (response.data.code === 1) {
            //     console.log("code failed")
            // }
            // else if (response.data.code === 2) {
            //     console.log("time out ,go to refresh token")
            //     resetToken(url, {}, callback, false, false, {
            //         errorCallback: (e)=>{console.log(e)},
            //         params:{}
            //     })

            // }
            // else {
            callback(response.data);
            // }

        });
};

export { server, postRequest, getRequest, getRequest_checkSession, putRequest };
