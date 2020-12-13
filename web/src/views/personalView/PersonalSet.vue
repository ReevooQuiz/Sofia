<template>
  <div>
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
              slot="cover"
              alt="example"
              src="https://tse2-mm.cn.bing.net/th/id/OIP.OCLuKoXlay8WIeNZPpCfcgHaHa?pid=Api&rs=1"
              style="height: 100px; border-radius: 50%"
            />
            <br />
            <a-upload
              v-if="this.edit===true"
              v-model:fileList="fileList"
              name="file"
              :multiple="true"
              action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
              :headers="headers"
              @change="handleChange"
            >
              <CameraOutlined />
            </a-upload>
          </a-col>
        </a-row>

        <br />
        <a-tag color="#88d5d1">
          <VerifiedOutlined />学习区专家
        </a-tag>
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
          {{info.follower_count}} 关注 ·
        </span>
        <span>{{info.followed_count}} 粉丝 ·</span>
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

        <a-button v-if="this.edit===false" block ghost @click="changeEditStatus">修改个人信息</a-button>
        <div v-if="this.edit===true">
          <a-button ghost style="padding-left:5%;width:45%" @click="handleFinish">保存</a-button>
          <a-button ghost style="padding-left:5%;width:45%" @click="handleCancle">取消</a-button>
        </div>
      </a-col>

      <a-col :span="10" :offset="2">
        <br />
        <br />
        <br />
        <div v-if="this.edit===false">
          <a-row>
            <a-col :span="3">
              <span class="set-lable">姓名</span>
            </a-col>
            <a-col :span="16" :offset="2">
              <span class="set-content">{{info.username}}</span>
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
              <span v-if="info.gender===0" class="set-content">男</span>
              <span v-else class="set-content">女</span>
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

        <div v-else>
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
              <a-col :span="3">
                <span class="set-lable">姓名</span>
              </a-col>
              <a-col :span="19" :offset="2">
                <a-form-item required has-feedback name="username">
                  <a-input v-model:value="ruleForm.username" />
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
          </a-form>
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
import SubMenu from "../../components/PersonalNavigation";
import { message } from "ant-design-vue";
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
import { postRequest } from "@/http/request.js";
const data = {
  username: "akangakang",
  nickname: "aaaaaaa",
  gender: 0,
  email: "11111111@sjtu.edu.cn",
  profile:
    "个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介个人简介",
  icon: "icon data",
  level: 2,

  account_type: 1,
  label: "math",
  like_count: 10,
  question_count: 10,
  answer_count: 10,
  follower_count: 10,
  followed_count: 10
};
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
    CopyOutlined
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

    return {
      fileList: [],
      headers: {
        authorization: "authorization-text"
      },
      info: {},
      edit: false,
      ruleForm: {
        username: "",
        nickname: "",
        email: "",
        gender: 0,
        icon: "",
        profile: ""
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
    this.info = data;
    this.ruleForm.username = this.info.username;
    this.ruleForm.nickname = this.info.nickname;
    this.ruleForm.email = this.info.email;
    this.ruleForm.gender = this.info.gender;
    this.ruleForm.profile = this.info.profile;
    let id = JSON.parse(sessionStorage.getItem("user")).uid;
    // getRequest("/followers",this.handleCallback,{
    //   errorCallback:(e)=>{console.log(e)},
    //   params:{uid:id}
    // })
  },
  methods: {
    handleCallback(response) {
      console.log(response);
      this.info = response.result;
    },
    handleChange(info) {
      if (info.file.status !== "uploading") {
        console.log(info.file, info.fileList);
      }
      if (info.file.status === "done") {
        message.success(`${info.file.name} file uploaded successfully`);
      } else if (info.file.status === "error") {
        message.error(`${info.file.name} file upload failed.`);
      }
    },
    handleFinish() {
      console.log(this.ruleForm);
      console.log("finished");
      postRequest("/publicInfo", this.ruleForm, this.mycallback, {
        errorCallback: error => {
          console.log(JSON.stringify(error));
        }
      });
    },
    mycallback(response) {
      if (response.code === 0) {
        this.info.nickname = this.ruleForm.username;
        this.info.nickname = this.ruleForm.nickname;
        this.info.email = this.ruleForm.email;
        this.info.gender = this.ruleForm.gender;
        this.info.profile = this.ruleForm.profile;
        this.edit = !this.edit;
      }
    },
    handleFinishFailed(errors) {
      console.log(JSON.stringify(errors));
    },
    changeEditStatus() {
      this.edit = !this.edit;
    },
    handleCancle() {
      this.ruleForm.username = this.info.username;
      this.ruleForm.nickname = this.info.nickname;
      this.ruleForm.email = this.info.email;
      this.ruleForm.gender = this.info.gender;
      this.ruleForm.profile = this.info.profile;
      this.edit = !this.edit;
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