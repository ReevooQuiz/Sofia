<template>
  <div id="register">
    <a-row justify="center">
      <a-col>
        <div id="register-block">
          <a-row justify="center">
            <a-col>
              <span id="login-font">注 册</span>
              <br />
              <br />
            </a-col>
          </a-row>
          <a-form
            style="margin-left:60px"
            name="custom-validation"
            ref="ruleForm"
            :model="ruleForm"
            :rules="rules"
            v-bind="layout"
            @finish="handleFinish"
            @finishFailed="handleFinishFailed"
          >
            <a-row justify="start">
              <a-col :span="8">
                <a-form-item label="头像" name="icon">
                  <a-upload
                    v-model:value="ruleForm.icon"
                    v-model:fileList="fileList"
                    name="icon"
                    list-type="picture-card"
                    class="avatar-uploader"
                    :show-upload-list="false"
                    action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
                    :before-upload="beforeUpload"
                    @change="handleChange"
                  >
                    <img v-if="imageUrl" :src="imageUrl" alt="avatar" />
                    <div v-else>
                      <!-- todo -->
                      <loading-outlined v-if="loading" />
                      <plus-outlined v-else />
                      <div class="ant-upload-text">上传头像</div>
                    </div>
                  </a-upload>
                </a-form-item>
              </a-col>

              <a-col :span="14">
                <a-form-item required has-feedback label="用户名" name="name">
                  <a-input v-model:value="ruleForm.name" />
                </a-form-item>
                <a-form-item required has-feedback label="昵称" name="nickname">
                  <a-input v-model:value="ruleForm.nickname" />
                </a-form-item>
                <a-form-item required has-feedback label="邮箱" name="email">
                  <a-input v-model:value="ruleForm.email" />
                </a-form-item>
                <a-form-item required label="性别" name="gender">
                  <a-radio-group v-model:value="ruleForm.gender">
                    <a-radio value="0">男</a-radio>
                    <a-radio value="1">女</a-radio>
                  </a-radio-group>
                </a-form-item>
                <a-form-item required has-feedback label="密码" name="pass">
                  <a-input v-model:value="ruleForm.pass" type="password" autocomplete="off" />
                </a-form-item>
                <a-form-item class="register-form-item" has-feedback label="确认密码" name="checkPass">
                  <a-input v-model:value="ruleForm.checkPass" type="password" autocomplete="off" />
                </a-form-item>
              </a-col>
            </a-row>
            <a-form-item :wrapper-col="{ span: 20 }">
              <a-button type="primary" html-type="submit">注册</a-button>
              <a-button style="margin-left: 50px" @click="resetForm">重置</a-button>
            </a-form-item>
          </a-form>
          <!-- </a-col>
          </a-row>-->
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
import { PlusOutlined, LoadingOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import {postRequest} from "@/http/request.js";
function getBase64(img, callback) {
  const reader = new FileReader();
  reader.addEventListener("load", () => callback(reader.result));
  reader.readAsDataURL(img);
}
export default {
  components: {
    "a-button": Button,
    UserOutlined,
    LockOutlined,
    LoadingOutlined,
    PlusOutlined
  },
  data() {
    let checkName = async (rule, value, callback) => {
      if (value === "") {
        return Promise.reject("请输入用户名");
      } else {
        return Promise.resolve();
      }
    };
    // let checkIcon = async (rule, value, callback) => {
    //   if (this.ruleForm.icon === "") {
    //     return Promise.reject("请上传头像");
    //   } else {
    //     return Promise.resolve();
    //   }
    // };
    let checkNickName = async (rule, value, callback) => {
      if (value === "") {
        return Promise.reject("请输入昵称");
      } else {
        return Promise.resolve();
      }
    };
    let checkGender = async (rule, value, callback) => {
      if (value === 3) {
        return Promise.reject("请输入性别");
      } else {
        return Promise.resolve();
      }
    };
    let checkEmail = async (rule, value, callback) => {
      if (value === "") {
        return Promise.reject("请输入邮箱");
      } else {
        if (value !== "") {
          var reg = /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/;
          if (!reg.test(value)) {
            return Promise.reject("请输入正确邮箱格式");
          }
        }
        return Promise.resolve();
      }
    };
    let validatePass = async (rule, value) => {
      if (value === "") {
        return Promise.reject("请输入密码");
      } else {
        if (this.ruleForm.checkPass !== "") {
          this.$refs.ruleForm.validateField("checkPass");
        }
        return Promise.resolve();
      }
    };
    let validatePass2 = async (rule, value, callback) => {
      if (value === "") {
        return Promise.reject("请再次输入密码");
      } else if (value !== this.ruleForm.pass) {
        return Promise.reject("确认密码与密码不相同");
      } else {
        return Promise.resolve();
      }
    };
    return {
      ruleForm: {
        pass: "",
        checkPass: "",
        name: "",
        nickname: "",
        email: "",
        gender: 3,
        icon: ""
      },
      rules: {
        pass: [{ validator: validatePass, trigger: "change" }],
        checkPass: [{ validator: validatePass2, trigger: "change" }],
        email: [{ validator: checkEmail, trigger: "change" }],
        name: [{ validator: checkName, trigger: "change" }],
        nickname: [{ validator: checkNickName, trigger: "change" }],
        gender: [{ validator: checkGender, trigger: "change" }]
        // icon: [{ validator: checkIcon }]
      },
      layout: {
        labelCol: { span: 8 },
        wrapperCol: { span: 16 }
        // labelAlign:'left'
      },
      fileList: [],
      loading: false,
      imageUrl: ""
    };
  },

  methods: {
    handleFinish(values) {
      console.log(values);
      console.log("finished");
      postRequest('/register',values,this.mycallback,{
            errorCallback:  (error) => {console.log(error)},
        })
    },
    mycallback(response)
    {
      console.log(response);
      if(response.code==0)
      {
        message.error("注册失败：用户名已被使用");
      }
      else if(response.code==1)
      {
        message.error("注册失败：邮箱已被使用");
      }
      else if(response.code==2)
      {
        message.success("注册成功");
        this.$router.push({ path:'/login'  });
      }
    },
    handleFinishFailed(errors) {
      console.log(JSON.stringify(errors));
    },
    resetForm() {
      this.$refs.ruleForm.resetFields();
    },
    handleChange(info) {
      if (info.file.status === "uploading") {
        this.loading = true;
        return;
      }
      if (info.file.status === "done") {
        // Get this url from response in real world.
        getBase64(info.file.originFileObj, imageUrl => {
          this.imageUrl = imageUrl;
          this.loading = false;
          let arr = imageUrl.split(",");
          this.ruleForm.icon = arr[1];
          console.log(arr);
        });
      }
      if (info.file.status === "error") {
        this.loading = false;
      }
    },
    beforeUpload(file) {
      const isJpgOrPng =
        file.type === "image/jpeg" || file.type === "image/png";
      if (!isJpgOrPng) {
        message.error("You can only upload JPG file!");
      }
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isLt2M) {
        message.error("Image must smaller than 2MB!");
      }
      return isJpgOrPng && isLt2M;
    }
  }
};
</script>

<style>
#register {
  min-height: 683px;
  background-color: #edeeed;
}

#register-block {
  top: 18%;
  /* position: fixed; */
  left: 27%;
  text-align: center;
  background-color: #ffffff;
  width: 700px;
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
/* 
.register-form-item label {
  position: relative;
  margin-right: 100px;
} */

.avatar-uploader > .ant-upload {
  width: 128px;
  height: 128px;
}
.ant-upload-select-picture-card i {
  font-size: 32px;
  color: #999;
}

.ant-upload-select-picture-card .ant-upload-text {
  margin-top: 8px;
  color: #666;
}
</style>
