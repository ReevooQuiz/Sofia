<template>
  <div id="login">
    <a-row justify="center">
      <a-col>
        <div id="login-block">
          <a-row justify="center">
            <a-col>
              <span id="login-font">登 录</span>
              <br />
              <br />
              <a-form :model="formInline" @submit="handleSubmit" @submit.native.prevent>
                <a-form-item class="login-input">
                  <a-input v-model:value="formInline.name" placeholder="请输入用户名">
                    <template #prefix>
                      <UserOutlined style="color:rgba(0,0,0,.25)" />
                    </template>
                  </a-input>
                </a-form-item>
                <a-form-item class="login-input">
                  <a-input v-model:value="formInline.password" type="password" placeholder="请输入密码">
                    <template #prefix>
                      <LockOutlined style="color:rgba(0,0,0,.25)" />
                    </template>
                  </a-input>
                </a-form-item>
                <a-tag color="#272c31" @click="goToLoginG">
                  <GithubOutlined />使用Github登录
                </a-tag>
                <br />
                <br />
                <a-form-item>
                  <a-button @click="goToRegister">注册</a-button>

                  <a-button
                    style="margin-left:90px;"
                    type="primary"
                    html-type="submit"
                    :disabled="formInline.user === '' || formInline.password === ''"
                  >登录</a-button>
                </a-form-item>
              </a-form>
            </a-col>
          </a-row>
        </div>
      </a-col>
    </a-row>
    <br />

    <br />
  </div>
</template>

<script>
import axios from 'axios';
import { defineComponent } from "vue";
import { Options, Vue } from "vue-class-component";
import {
  UserOutlined,
  LockOutlined,
  GithubOutlined
} from "@ant-design/icons-vue";
import { Button } from "ant-design-vue";
import { postRequest, getRequest } from "@/http/request.js";
import { message } from "ant-design-vue";
import Axios from 'axios';
export default {
  components: {
    "a-button": Button,
    UserOutlined,
    LockOutlined,
    GithubOutlined
  },
  data() {
    return {
      formInline: {
        name: "",
        password: ""
      }
    };
  },

  methods: {
    handleSubmit(e) {
      console.log(this.formInline);
      postRequest("/login", this.formInline, this.handleLogin, {
        errorCallback: error => {
          console.log(error);
        }
      });
    },
    handleLogin(response) {
      // 问题： 应该返回头像信息
      console.log(response);
      if (response.code == 0) {
        this.$store.commit("changeLogStatus",true);
        message.success("登录成功");
        if (response.result.role==0)
            this.$store.commit("changeAdmin",true);
        else  this.$store.commit("changeAdmin",false);
        console.log(response.result);
        this.$store.commit("changeIcon",response.result.icon);
        sessionStorage.setItem("user", JSON.stringify(response.result));
        // this.$store.commit('modify',"改变值！")
        //  this.$store.state.loginStatus=true;
        this.$router.back();
      }
      // else{
      //     this.$dialog.alert('您的账户已被禁用，请联系管理员解禁。').then(
      //         (dialog)=>{
      //             dialog.close();
      //         })
      // }
    },
    goToRegister() {
      this.$router.push({ path: "/register" });
    },
    goToLoginG() {

      axios
        .get("http://github.com/login/oauth/authorize?client_id=51f0dde36e2f4fcee97c&redirect_uri=http://rv-s.cn:9092/oauth/github",
            {}
        )
        .catch(function (error) {
            console.log(error);
        })
        .then(response => {
            console.log(response);


            if (response.data.code === 1) {
                console.log("code failed")
            }
            else if (response.data.code === 2) {
                console.log("time out ,go to refresh token")


            }
            else {
                this.handleLogin(response.data);
            }

        });
      // getRequest("/oauth/github", this.handleLogin, {
      //   errorCallback: error => {
      //     console.log(error);
      //   }
      // });
    }
  }
};
</script>

<style >
#login {
  min-height: 683px;
  background-color: #edeeed;
}

#login-block {
  top: 26%;
  position: fixed;
  left: 33%;
  text-align: center;
  background-color: #ffffff;
  width: 480px;
  border-radius: 40px;
  padding: 20px;
  margin: 0 auto;
  box-shadow: 3px 3px 2px #dcdfdf;
}
#login-font {
  font-size: 35px;
  color: #485355f5;

  font-weight: bold;
}
.login-input .ant-input-affix-wrapper > input.ant-input {
  padding: 0;
  border: none;
  outline: none;
  background-color: #bae7e500;
}
.login-input .ant-input-affix-wrapper {
  border: 1px solid #ffffff00;
  border-radius: 12px;
  padding: 4px 14px;
  width: 100%;

  background-color: #bae7e5d3;
}
</style>
