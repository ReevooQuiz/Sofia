<template>
  <div id="login">
    <a-row justify="center">
      <a-col>
        <div id="login-block">
          <a-row justify="center">
            <a-col>
              <span id="login-font">登 录</span>
              <br/>
              <a-form :model="formInline" @submit="handleSubmit" @submit.native.prevent>
                <a-form-item class="login-input">
                  <a-input v-model:value="formInline.user" placeholder="请输入用户名">
                    <template #prefix>
                      <UserOutlined style="color:rgba(0,0,0,.25)" />
                    </template>
                  </a-input>
                </a-form-item>
                <a-form-item class="login-input">
                  <a-input
                    v-model:value="formInline.password"
                    type="password"
                    placeholder="请输入密码"
                  >
                    <template #prefix>
                      <LockOutlined style="color:rgba(0,0,0,.25)" />
                    </template>
                  </a-input>
                </a-form-item>

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
import { defineComponent } from "vue";
import { Options, Vue } from "vue-class-component";
import { UserOutlined, LockOutlined } from "@ant-design/icons-vue";
import { Button } from "ant-design-vue";
export default {
  components: {
    "a-button": Button,
    UserOutlined,
    LockOutlined
  },
  data() {
    return {
      formInline: {
        user: "",
        password: ""
      }
    };
  },

  methods: {
    handleSubmit(e) {
      console.log(this.formInline);
    },
    goToRegister() {
      console.log("?");
      this.$router.push({ path: "/register" });
    }
  },
 
};
</script>

<style >
#login {
  min-height: 683px;
  background-color: #edeeed;
}

#login-block {
    top: 30%;position: fixed;
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

/* .login-block .ant-input-affix-wrapper > input.ant-input {
    padding: 0;
    border: none;
    outline: none;
     background-color: #bae7e5d3;
} */
</style>
