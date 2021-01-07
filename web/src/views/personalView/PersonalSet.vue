<template>
  <div>
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
      <a-row>
        <a-col :span="1" :offset="1">
          <div class="personal-nav">
            <SubMenu />
          </div>
        </a-col>

        <a-col :span="3" :offset="2">
          <br />
          <br />
          <br />

          <a-row justify="center">
            <a-col>
              <img
                v-if="this.edit===false"
                slot="cover"
                alt="example"
                :src="this.imageUrl"
                style="height: 100px; border-radius: 50%"
              />
              <br />
              <!-- <a-upload
              v-if="this.edit===true"
              v-model:fileList="fileList"
              name="file"
              :multiple="true"
              action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
              :headers="headers"
              @change="handleChange"
            >
              <CameraOutlined />
              </a-upload>-->
              <a-upload
                v-if="this.edit===true"
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
                <img style="width:100px" v-if="imageUrl" :src="imageUrl" alt="avatar" />
                <div v-else>
                  <!-- todo -->
                  <loading-outlined v-if="loading" />
                  <plus-outlined v-else />
                  <div class="ant-upload-text">上传头像</div>
                </div>
              </a-upload>
            </a-col>
          </a-row>

          <br />
          <!-- <a-tag color="#88d5d1">
            <VerifiedOutlined />学习区专家
          </a-tag> -->
          <a-tag color="#88d5d1">
            <FireOutlined />
            等级 {{info.level}}
          </a-tag>
          <a-tag color="#88d5d1" style="margin:3px">
            <UserOutlined />
            <span v-if="info.account_type===1">管理员</span>
            <span v-else>普通用户</span>
          </a-tag>

          <a-divider />

          <span>
            <TeamOutlined />
            {{info.following_count}} 关注 ·
          </span>
          <span>{{info.follower_count}} 粉丝 ·</span>
          <span>
            <LikeOutlined />
            {{info.like_count}} 赞
          </span>

          <span>
            <FormOutlined />
            {{info.question_count}} 问题 ·
          </span>
          <span>
            <CopyOutlined />
            {{info.answer_count}} 回答
          </span>
          <a-divider />

          <a-button
            v-if="this.edit===false && this.editPass===false &&this.forgetPass===false"
            block
            ghost
            @click="changeEditStatus"
          >修改个人信息</a-button>
          <div v-if="this.edit===true">
            <a-form-item>
              <a-row>
                <a-col :span="10" :offset="1">
                  <a-button ghost html-type="submit">保存</a-button>
                </a-col>
                <a-col :span="10" :offset="2">
                  <a-button ghost style="margin-left: 50px" @click="handleCancle">取消</a-button>
                </a-col>
              </a-row>
            </a-form-item>
          </div>

          <a-button
            style="margin-top: 10px"
            v-if="this.editPass===false && this.edit===false &&this.forgetPass===false"
            block
            ghost
            @click="changeEditPassStatus"
          >修改密码</a-button>

           <a-button
            style="margin-top: 10px"
            v-if="this.editPass===false && this.edit===false &&this.forgetPass===false"
            block
            ghost
            @click="changeForgetPassStatus"
          >忘记密码</a-button>
          <!-- <a-button v-if="this.editPass===true" block ghost @click="savePassWord">保存密码</a-button> -->
        </a-col>

        <a-col :span="10" :offset="2">
          <br />
          <br />
          <br />
          <div v-if="this.edit===false && this.editPass===false && this.forgetPass===false ">
            <a-row>
              <a-col :span="3">
                <span class="set-lable">姓名</span>
              </a-col>
              <a-col :span="16" :offset="2">
                <span class="set-content">{{info.name}}</span>
              </a-col>
            </a-row>

            <br />
            <a-row>
              <a-col :span="3">
                <span class="set-lable">昵称</span>
              </a-col>
              <a-col :span="17" :offset="2">
                <span class="set-content">{{info.nickname}}</span>
              </a-col>
            </a-row>
            <br />
            <a-row>
              <a-col :span="3">
                <span class="set-lable">性别</span>
              </a-col>
              <a-col :span="17" :offset="2">
                <span v-if="info.gender==0" class="set-content">男</span>
                <span v-else-if="info.gender==1" class="set-content">女</span>
              </a-col>
            </a-row>
            <br />
            <a-row>
              <a-col :span="3">
                <span class="set-lable">邮箱</span>
              </a-col>
              <a-col :span="17" :offset="2">
                <span class="set-content">{{info.email}}</span>
              </a-col>
            </a-row>
            <br />
            <a-row>
              <a-col :span="3">
                <span class="set-lable">个人简介</span>
              </a-col>
              <a-col :span="17" :offset="2">
                <span class="set-content">{{info.profile}}</span>
              </a-col>
            </a-row>
          </div>

          <div v-else-if="this.edit===true && this.editPass===false && this.forgetPass===false">
            <!-- <a-form
            style="margin-left:60px"
            name="custom-validation"
            ref="ruleForm"
            :model="ruleForm"
            :rules="rules"
            v-bind="layout"
            @finish="handleFinish"
            @finishFailed="handleFinishFailed"
            >-->
            <a-row>
              <a-col :span="3">
                <span class="set-lable">姓名</span>
              </a-col>
              <a-col :span="19" :offset="2">
                <a-form-item required has-feedback name="name">
                  <a-input v-model:value="ruleForm.name" />
                </a-form-item>
              </a-col>
            </a-row>

            <a-row>
              <a-col :span="3">
                <span class="set-lable">昵称</span>
              </a-col>
              <a-col :span="19" :offset="2">
                <a-form-item required has-feedback name="nickname">
                  <a-input v-model:value="ruleForm.nickname" />
                </a-form-item>
              </a-col>
            </a-row>

            <a-row>
              <a-col :span="3">
                <span class="set-lable">性别</span>
              </a-col>
              <a-col :span="19" :offset="2">
                <a-form-item required name="gender">
                  <a-radio-group v-model:value="ruleForm.gender">
                    <a-radio value="0">男</a-radio>
                    <a-radio value="1">女</a-radio>
                  </a-radio-group>
                </a-form-item>
              </a-col>
            </a-row>

            <a-row>
              <a-col :span="3">
                <span class="set-lable">邮箱</span>
              </a-col>
              <a-col :span="19" :offset="2">
                <a-form-item required has-feedback name="email">
                  <a-input v-model:value="ruleForm.email" />
                </a-form-item>
              </a-col>
            </a-row>

            <a-row>
              <a-col :span="3">
                <span class="set-lable">个人简介</span>
              </a-col>
              <a-col :span="19" :offset="2">
                <a-form-item required has-feedback name="profile">
                  <a-textarea :auto-size="{ minRows: 5 }" v-model:value="ruleForm.profile" />
                </a-form-item>
              </a-col>
            </a-row>
            <!-- </a-form> -->
          </div>

          <div v-else-if="this.edit===false && this.editPass===true && this.forgetPass===false">
            <a-form
              name="passwordChange"
              ref="passwordChange"
              :model="passwordChange"
              :rules="passRules"
              v-bind="layout"
              @finish="handleSavePass"
              @finishFailed="handleFinishFailed"
            >
              <a-row>
                <a-col :span="3">
                  <span class="set-lable">原密码</span>
                </a-col>
                <a-col :span="19" :offset="2">
                  <a-form-item required has-feedback name="old">
                    <a-input v-model:value="passwordChange.old" type="password" autocomplete="off" />
                  </a-form-item>
                </a-col>
              </a-row>

              <a-row>
                <a-col :span="3">
                  <span class="set-lable">新密码</span>
                </a-col>
                <a-col :span="19" :offset="2">
                  <a-form-item required has-feedback name="new">
                    <a-input v-model:value="passwordChange.new" type="password" autocomplete="off" />
                  </a-form-item>
                </a-col>
              </a-row>

              <a-row>
                <a-col :span="3">
                  <span class="set-lable">密码确认</span>
                </a-col>
                <a-col :span="19" :offset="2">
                  <a-form-item has-feedback name="newCheck">
                    <a-input
                      v-model:value="passwordChange.newCheck"
                      type="password"
                      autocomplete="off"
                    />
                  </a-form-item>
                </a-col>
              </a-row>

              <a-row>
                <a-col :span="3">
                  <a-form-item>
                    <a-button v-if="this.editPass===true" ghost html-type="submit">保存密码</a-button>
                    <!-- <a-button type="primary" html-type="submit">Submit</a-button> -->
                  </a-form-item>
                </a-col>
                <a-col :span="19" :offset="2"></a-col>
              </a-row>
            </a-form>
          </div>

           <div v-else-if="this.edit===false && this.forgetPass===true  && this.editPass===false">
            <a-form
              name="passwordChange"
              ref="passwordChange"
              :model="passwordChange"
              :rules="passRules"
              v-bind="layout"
              @finish="handleSavePass"
              @finishFailed="handleFinishFailed"
            >
              <a-row>
                <a-col :span="3">
                  <span class="set-lable">验证码</span>
                </a-col>
                <a-col :span="19" :offset="2">
                  <a-form-item required has-feedback name="old">
                    <a-input v-model:value="passwordChange.old"  autocomplete="off" />
                  </a-form-item>
                </a-col>
              </a-row>

              <a-row>
                <a-col :span="3">
                  <span class="set-lable">新密码</span>
                </a-col>
                <a-col :span="19" :offset="2">
                  <a-form-item required has-feedback name="new">
                    <a-input v-model:value="passwordChange.new" type="password" autocomplete="off" />
                  </a-form-item>
                </a-col>
              </a-row>

              <a-row>
                <a-col :span="3">
                  <span class="set-lable">密码确认</span>
                </a-col>
                <a-col :span="19" :offset="2">
                  <a-form-item has-feedback name="newCheck">
                    <a-input
                      v-model:value="passwordChange.newCheck"
                      type="password"
                      autocomplete="off"
                    />
                  </a-form-item>
                </a-col>
              </a-row>

              <a-row>
                <a-col :span="3">
                  <a-form-item>
                    <a-button v-if="this.editPass===true" ghost html-type="submit">保存密码</a-button>
                    <!-- <a-button type="primary" html-type="submit">Submit</a-button> -->
                  </a-form-item>
                </a-col>
                <a-col :span="19" :offset="2"></a-col>
              </a-row>
            </a-form>
          </div>
        </a-col>
      </a-row>

      <br />

      <br />
    </a-form>
  </div>
</template>



<script>
import { defineComponent } from "vue";
import { Options, Vue } from "vue-class-component";
import SubMenu from "../../components/PersonalNavigation";
import { message } from "ant-design-vue";
import { PlusOutlined, LoadingOutlined } from "@ant-design/icons-vue";
import {
  CameraOutlined,
  FireOutlined,
  VerifiedOutlined,
  UserOutlined,
  TeamOutlined,
  LikeOutlined,
  FormOutlined,
  CopyOutlined
} from "@ant-design/icons-vue";
import { postRequest, getRequest ,putRequest} from "@/http/request.js";

function getBase64(img, callback) {
  const reader = new FileReader();
  reader.addEventListener("load", () => callback(reader.result));
  reader.readAsDataURL(img);
}

export default {
  components: {
    SubMenu,
    CameraOutlined,
    FireOutlined,
    VerifiedOutlined,
    UserOutlined,
    TeamOutlined,
    LikeOutlined,
    FormOutlined,
    CopyOutlined,
    PlusOutlined,
    LoadingOutlined
  },
  data() {
    let checkName = async (rule, value, callback) => {
      if (value === "") {
        return Promise.reject("请输入用户名");
      } else {
        return Promise.resolve();
      }
    };
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
        if (value !== "") {
          var reg = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,16}$/;
          if (!reg.test(value))
            return Promise.reject("需包含大小写字母和数字，至少8位");
        }
        if (this.ruleForm.checkPass !== "") {
          this.$refs.ruleForm.validateField("checkPass");
        }
        return Promise.resolve();
      }
    };
    let validatePass2 = async (rule, value, callback) => {
      if (value === "") {
        return Promise.reject("请再次输入密码");
      } else if (value !== this.passwordChange.new) {
        return Promise.reject("确认密码与密码不相同");
      } else {
        if (value !== "") {
          var reg = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,16}$/;
          if (!reg.test(value))
            return Promise.reject("需包含大小写字母和数字，至少8位");
        }
        return Promise.resolve();
      }
    };

    return {
      fileList: [],
      loading: false,
      imageUrl: "",
      done: false,
      headers: {
        authorization: "authorization-text"
      },
      info: {},
      edit: false,
      editPass: false,
      forgetPass:false,
      ruleForm: {
        name: "",
        nickname: "",
        email: "",
        gender: 3,
        icon: "",
        profile: ""
      },
      passwordChange: {
        old: "",
        new: "",
        newCheck: ""
      },
      passwordForget: {
        verify:"",
        new: "",
        newCheck: ""
      },
      
      passRules: {
        old: [{ validator: validatePass, trigger: "change" }],
        new: [{ validator: validatePass, trigger: "change" }],
        newCheck: [{ validator: validatePass2, trigger: "change" }]
      },
      rules: {
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
      }
    };
  },
  created() {
    // this.info = data;
    let uid = JSON.parse(sessionStorage.getItem("user")).uid;

    getRequest("/publicInfo", this.handleCallback, {
      errorCallback: e => {
        console.log(JSON.stringify(e));
      },
      params: { id: uid }
    });
    this.imageUrl = JSON.parse(sessionStorage.getItem("user")).icon;
    // this.ruleForm.username = this.info.username;

    // this.ruleForm.nickname = this.info.nickname;
    // this.ruleForm.email = this.info.email;
    // this.ruleForm.gender = this.info.gender.toString();

    // this.ruleForm.icon = this.info.icon;
    // this.ruleForm.profile = this.info.profile;
  },
  methods: {
    handleCallback(response) {
      console.log(response);
      this.info = response.result;
      this.imageUrl = this.info.icon;
    },

    handleFinish() {
      console.log(this.ruleForm);
      console.log("finished");
      putRequest("/publicInfo", this.ruleForm, this.mycallback, {
        errorCallback: error => {
          console.log(JSON.stringify(error));
        }
      });
    },
    mycallback(response) {
      if (response.code === 0) {
        this.info.nickname = this.ruleForm.name;
        this.info.nickname = this.ruleForm.nickname;
        this.info.email = this.ruleForm.email;
        this.info.gender = this.ruleForm.gender.toString();
        this.info.profile = this.ruleForm.profile;
        this.edit = !this.edit;
      }
    },
    handleFinishFailed(errors) {
      console.log(JSON.stringify(errors));
    },
    changeEditStatus() {
      this.ruleForm.name = this.info.name;

      this.ruleForm.nickname = this.info.nickname;
      this.ruleForm.email = this.info.email;
      // this.ruleForm.gender = this.info.gender;
      this.ruleForm.gender = this.info.gender.toString();
      this.ruleForm.icon = this.info.icon;
      this.ruleForm.profile = this.info.profile;
      this.edit = !this.edit;
    },
    changeEditPassStatus() {
      this.editPass = !this.editPass;
    },
    changeForgetPassStatus()
    {
      this.forgetPass =!this.forgetPass;
    },
    handleCancle() {
      this.ruleForm.name = this.info.name;
      this.ruleForm.nickname = this.info.nickname;
      this.ruleForm.email = this.info.email;
      this.ruleForm.gender = this.info.gender;
      this.ruleForm.profile = this.info.profile;
      this.edit = !this.edit;
    },
    handleChange(info) {
      if (info.file.status === "uploading") {
        this.loading = true;
        return;
      }
      if (info.file.status === "done") {
        this.done = true;
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
    },
    passCallback(res)
    {
      if(res.code===1)
      {
        if(res.type===0)
        {
          message.error("更新密码失败：原密码输入错误")
        }
        else{
           message.error("更新密码失败")
        }
      }
      else if(res.code===0)
      {
        message.success("更新密码成功")
        this.editPass=!this.editPass
      }
    },
    handleSavePass()
    {
      console.log("change pass")
      postRequest("/passwd",this.passwordChange,this.passCallback,{
        errorCallback: error => {
          console.log(JSON.stringify(error));
        }
      })
    }
  }
};
</script>

<style>
.set-lable {
  font-size: 18px;
  color: #485355f5;

  font-weight: bold;
}

.set-content {
  font-size: 16px;
  color: #33393af5;

  /* font-weight: bold; */
}
.avatar-uploader > .ant-upload {
  width: 128px;
  height: 128px;
}
body {
  height: 100%;
  background-color: #edeeed;
}

.personal-nav .ant-menu-inline {
  width: 40%;
}

.ant-divider-horizontal {
  display: block;
  clear: both;
  width: 100%;
  min-width: 100%;
  height: 1px;
  margin: 16px 0;
  background-color: #acbab9;
}

.ant-btn-background-ghost {
  color: #98a7a6;
  background: transparent !important;
  border-color: #98a7a6;
}
</style>