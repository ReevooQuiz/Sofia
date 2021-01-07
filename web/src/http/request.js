import axios from 'axios'
import router from "@/router/index.ts";
const server = axios.create({
    // baseURL: "http://localhost:4000/",
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

export { server, postRequest, getRequest ,getRequest_checkSession};
