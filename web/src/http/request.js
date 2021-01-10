import axios from 'axios'
import router from "@/router/index.ts";

const server1 = axios.create({
    // baseURL: "http://localhost:4000/",
    // baseURL: "https://private-74c97e-reevooapi.apiary-mock.com",
    // baseURL:"https://121.4.194.85",
    baseURL:"http://localhost:9092/",
    // timeout: 5000,
});

const server2 = axios.create({
    // baseURL: "http://localhost:4000/",
    // baseURL: "https://private-74c97e-reevooapi.apiary-mock.com",
    // baseURL:"https://121.4.194.85",
    baseURL:"http://localhost:9093/",
    // timeout: 5000,
});

const server3 = axios.create({
    // baseURL: "http://localhost:4000/",
    // baseURL: "https://private-74c97e-reevooapi.apiary-mock.com",
    // baseURL:"https://121.4.194.85",
    baseURL:"http://localhost:9094/",
    // timeout: 5000,
});
// server1.defaults.retry = 3;
// server1.defaults.retryDelay = 500;

// server2.defaults.retry = 3;
// server2.defaults.retryDelay = 500;

// server3.defaults.retry = 3;
// server3.defaults.retryDelay = 500;
// 设置拦截器
// 请求拦截器
server1.interceptors.request.use(
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
server1.interceptors.response.use(
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
    // function axiosRetryInterceptor(err) {
    //     var config = err.config;
    //     // console.log("retry")
    //     // console.log(config.retry)
    //     // If config does not exist or the retry option is not set, reject
    //     if (!config || !config.retry) return Promise.reject(err);

    //     // Set the variable for keeping track of the retry count
    //     config.__retryCount = config.__retryCount || 0;

    //     // Check if we've maxed out the total number of retries
    //     if (config.__retryCount >= config.retry) {
    //         // Reject with the error
    //         return Promise.reject(err);
    //     }

    //     // Increase the retry count
    //     config.__retryCount += 1;

    //     // Create new promise to handle exponential backoff
    //     var backoff = new Promise(function (resolve) {
    //         setTimeout(function () {
    //             resolve();
    //         }, config.retryDelay || 1);
    //     });

    //     // Return the promise in which recalls axios to retry the request
    //     return backoff.then(function () {
    //         return server1(config);
    //     });
    // }
);

server2.interceptors.request.use(
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
server2.interceptors.response.use(
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
    // function axiosRetryInterceptor(err) {
    //     var config = err.config;
    //     // console.log("retry")
    //     // console.log(config.retry)
    //     // If config does not exist or the retry option is not set, reject
    //     if (!config || !config.retry) return Promise.reject(err);

    //     // Set the variable for keeping track of the retry count
    //     config.__retryCount = config.__retryCount || 0;

    //     // Check if we've maxed out the total number of retries
    //     if (config.__retryCount >= config.retry) {
    //         // Reject with the error
    //         return Promise.reject(err);
    //     }

    //     // Increase the retry count
    //     config.__retryCount += 1;

    //     // Create new promise to handle exponential backoff
    //     var backoff = new Promise(function (resolve) {
    //         setTimeout(function () {
    //             resolve();
    //         }, config.retryDelay || 1);
    //     });

    //     // Return the promise in which recalls axios to retry the request
    //     return backoff.then(function () {
    //         return server2(config);
    //     });
    // }
);
server3.interceptors.request.use(
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
server3.interceptors.response.use(
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
    // function axiosRetryInterceptor(err) {
    //     var config = err.config;
    //     // console.log("retry")
    //     // console.log(config.retry)
    //     // If config does not exist or the retry option is not set, reject
    //     if (!config || !config.retry) return Promise.reject(err);

    //     // Set the variable for keeping track of the retry count
    //     config.__retryCount = config.__retryCount || 0;

    //     // Check if we've maxed out the total number of retries
    //     if (config.__retryCount >= config.retry) {
    //         // Reject with the error
    //         return Promise.reject(err);
    //     }

    //     // Increase the retry count
    //     config.__retryCount += 1;

    //     // Create new promise to handle exponential backoff
    //     var backoff = new Promise(function (resolve) {
    //         setTimeout(function () {
    //             resolve();
    //         }, config.retryDelay || 1);
    //     });

    //     // Return the promise in which recalls axios to retry the request
    //     return backoff.then(function () {
    //         return server3(config);
    //     });
    // }
);

let resetToken = (url,
    body,
    callback,
    post,
    put,
    session,
    { errorCallback, params }) => {

    server1
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
                else if(put)
                {
                    putRequest(url, body, callback, {
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
                if (sessionStorage.getItem("user") !== null) {
                    sessionStorage.removeItem("user");
                  }
                //   this.$store.commit("changeLogStatus",false);
                //   this.$store.commit("changeIcon","");
                    // this.$store.commit("changeAdmin",false);
                  message.success("登录超时");
                  this.$router.push({ path: "/" });
                
            }
            else {
                // 超时 先不考虑
                console.log("time out!")
                if (sessionStorage.getItem("user") !== null) {
                    sessionStorage.removeItem("user");
                  }
                //   this.$store.commit("changeLogStatus",false);
                //   this.$store.commit("changeIcon","");
                    // this.$store.commit("changeAdmin",false);
                  message.success("登录超时");
                  this.$router.push({ path: "/" });
            }

            callback(response.data);

        });

}
let postRequest = (url, body, callback, { errorCallback }) => {

    switch (url) {
        case "/approve":
        case "/ban":
        case "/banned":
        case "/checkSession":
        case "/checkToken":
        case "/collection":
        case "/favorite":
        case "/follow":
        case "/followed":
        case "/followers":
        case "/infoList":
        case "/like":
        case "/login":
        case "/notifications":
        case "/oauth/github":
        case "/passwd":
        case "/publicInfo":
        case "/refreshToken":
        case "/register":
        case "/userAnswers":
        case "/userQuestions":
        case "/verificationCode":
        case "/verify":
        case "/wordBan":
        case "/wordsBanned":
            server1
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
                        callback(response.data);
                    }
                    else if (response.data.code === 2) {
                        console.log("time out ,go to refresh token")
                        resetToken(url, body, callback, true, false,false, {
                            errorCallback: errorCallback,
                            params: {}
                        })

                    }
                    else {
                        callback(response.data);
                    }


                });
            break;


        case "/questions":
        case "/question":
        case "/answers":
        case "/answer":
        case "/comments":
        case "/criticisms":
        case "/disable_question":
        case "/delete_answer":
            server2
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
                        callback(response.data);
                    }
                    else if (response.data.code === 2) {
                        console.log("time out ,go to refresh token")
                        resetToken(url, body, callback, true,false, false, {
                            errorCallback: errorCallback,
                            params: {}
                        })

                    }
                    else {
                        callback(response.data);
                    }


                });
            break;
        case "/searchQuestions":
        case "/searchAnswers":
        case "/searchUsers":
        case "/hotlist":
        case "/search": server3
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
                    callback(response.data);
                }
                else if (response.data.code === 2) {
                    console.log("time out ,go to refresh token")
                    resetToken(url, body,  callback, true, false,false, {
                        errorCallback: errorCallback,
                        params: {}
                    })

                }
                else {
                    callback(response.data);
                }


            });
    }
    console.log("here");



};
let putRequest = (url, body, callback, { errorCallback }) => {

    switch (url) {
        case "/approve":
        case "/ban":
        case "/banned":
        case "/checkSession":
        case "/checkToken":
        case "/collection":
        case "/favorite":
        case "/follow":
        case "/followed":
        case "/followers":
        case "/infoList":
        case "/like":
        case "/login":
        case "/notifications":
        case "/oauth/github":
        case "/passwd":
        case "/publicInfo":
        case "/refreshToken":
        case "/register":
        case "/userAnswers":
        case "/userQuestions":
        case "/verificationCode":
        case "/verify":
        case "/wordBan":
        case "/wordsBanned":
            server1
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
                        callback(response.data);
                    }
                    else if (response.data.code === 2) {
                        console.log("time out ,go to refresh token")
                        resetToken(url, body,  callback,false, true, false, {
                            errorCallback: errorCallback,
                            params: {}
                        })

                    }
                    else {
                        callback(response.data);
                    }


                });
            break;


        case "/questions":
        case "/question":
        case "/answers":
        case "/answer":
        case "/comments":
        case "/criticisms":
        case "/disable_question":
        case "/delete_answer":
            server2
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
                        callback(response.data);
                    }
                    else if (response.data.code === 2) {
                        console.log("time out ,go to refresh token")
                        resetToken(url, body,  callback, false,true, false, {
                            errorCallback: errorCallback,
                            params: {}
                        })

                    }
                    else {
                        callback(response.data);
                    }


                });
            break;
        case "/searchQuestions":
        case "/searchAnswers":
        case "/searchUsers":
        case "/hotlist":
        case "/search": server3
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
                    callback(response.data);
                }
                else if (response.data.code === 2) {
                    console.log("time out ,go to refresh token")
                    resetToken(url, body, callback, false,true, false, {
                        errorCallback: errorCallback,
                        params: {}
                    })

                }
                else {
                    callback(response.data);
                }


            });
    }
    console.log("here");



};

let getRequest = (url, callback, { errorCallback, params }) => {

    switch (url) {
        case "/approve":
        case "/ban":
        case "/banned":
        // case "/checkSession":
        case "/checkToken":
        case "/collection":
        case "/favorite":
        case "/follow":
        case "/followed":
        case "/followers":
        case "/infoList":
        case "/like":
        case "/login":
        case "/notifications":
        case "/oauth/github":
        case "/passwd":
        case "/publicInfo":
        case "/refreshToken":
        case "/register":
        case "/userAnswers":
        case "/userQuestions":
        case "/verificationCode":
        case "/verify":
        case "/wordBan":
        case "/wordsBanned":
            console.log(params)
            server1
                .get(url,
                    {params:params}
                )
                .catch(function (error) {
                    errorCallback(error)
                })
                .then(response => {
                    console.log(response);


                    if (response.data.code === 1) {
                        console.log("code failed")
                        callback(response.data);
                    }
                    else if (response.data.code === 2) {
                        console.log("time out ,go to refresh token")
                        resetToken(url, {},  callback, false,false, false, {
                            errorCallback: errorCallback,
                            params: params
                        })

                    }
                    else {
                        callback(response.data);
                    }

                });
            break;


        case "/questions":
        case "/question":
        case "/answers":
        case "/answer":
        case "/comments":
        case "/criticisms":
        case "/disable_question":
        case "/delete_answer":
            server2
                .get(url,
                    {params:params}
                )
                .catch(function (error) {
                    errorCallback(error)
                })
                .then(response => {
                    console.log(response);


                    if (response.data.code === 1) {
                        console.log("code failed")
                        callback(response.data);
                    }
                    else if (response.data.code === 2) {
                        console.log("time out ,go to refresh token")
                        resetToken(url, {},  callback, false,false, false, {
                            errorCallback: errorCallback,
                            params: params
                        })

                    }
                    else {
                        callback(response.data);
                    }

                });
            break;
        case "/searchQuestions":
        case "/searchAnswers":
        case "/searchUsers":
        case "/hotlist":
        case "/search": server3
            .get(url,
                {params:params}
            )
            .catch(function (error) {
                errorCallback(error)
            })
            .then(response => {
                console.log(response);


                if (response.data.code === 1) {
                    console.log("code failed")
                    callback(response.data);
                }
                else if (response.data.code === 2) {
                    console.log("time out ,go to refresh token")
                    resetToken(url, {},callback, false,false, false, {
                        errorCallback: errorCallback,
                        params: params
                    })

                }
                else {
                    callback(response.data);
                }

            });
    }
    console.log("here");

};


let getRequest_checkSession = (callback) => {
    const url = `/checkSession`;

    // getRequest(url,(rep)=>{if(rep.code )}, {
    //     errorCallback: errorCallback,
    //     params: {},
    // });

    server1
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

export { server1, postRequest, getRequest, getRequest_checkSession, putRequest };
