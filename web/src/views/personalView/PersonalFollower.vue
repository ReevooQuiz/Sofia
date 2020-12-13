<template>
  <div>
    <a-row>
      <a-col :span="1" :offset="1">
        <div class="personal-nav">
        <SubMenu />
        </div>
      </a-col>

      <a-col :span="20" :offset="1">
        <br/>
        <a-list :grid="{ gutter: 16, column: 4 }" :data-source="this.data">
          <template #renderItem="{ item }">
            <a-list-item>
              <div class="follower-block">
                <!-- <a-row >
                <a-col>-->
                <a-row justify="center">
                  <a-col>
                    <br />
                    <img
                      slot="cover"
                      alt="example"
                      src="https://tse2-mm.cn.bing.net/th/id/OIP.OCLuKoXlay8WIeNZPpCfcgHaHa?pid=Api&rs=1"
                      style="height: 60px; border-radius: 50%"
                    />
                    <br />
                    <!-- <br /> -->
                    <span class="follower-name">{{item.name}}</span>
                    <br />
                    <br />
                    <span class="follower-profile">{{item.profile}}</span>
                  </a-col>
                  <!-- <a-col :span="2" :offset="1"></a-col> -->
                </a-row>

                <br />
                <!-- </a-col>
                </a-row>-->
              </div>
            </a-list-item>
          </template>
        </a-list>
      </a-col>
    </a-row>
    <br />

    <br />
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { Options, Vue } from "vue-class-component";
import { UserOutlined } from "@ant-design/icons-vue";
import SubMenu from "../../components/PersonalNavigation";
import {getRequest} from "@/http/request.js";

export default {
  components: {
    UserOutlined,
    SubMenu
  },
  data() {
    return {
      data:[]
    };
  },
  created() {
    let id=JSON.parse(sessionStorage.getItem("user")).uid;
    getRequest("/followers",this.handleCallback,{
      errorCallback:(e)=>{console.log(e)},
      params:{uid:id}
    })
  },
  methods: {
    handleCallback(response){
        console.log(response);
        this.data=response.result;
    },
   
   
  }
};
</script>

<style>
body {
  height: 100%;
  background-color: #edeeed;
}

.follower-block {
  text-align: center;
  background-color: #ffffff;
  width:100%;
  height: 246px;
  border-radius: 20px;
  padding: 10px;
  margin: 0 auto;
  box-shadow: 3px 3px 2px #dcdfdf;
}

.follower-name {
  font-size: 18px;
  color: #485355f5;

  font-weight: bold;
}

.follower-profile {
  overflow: hidden;

  text-overflow: ellipsis;

  display: -webkit-box;

  -webkit-line-clamp: 3;

  -webkit-box-orient: vertical;
}

.personal-nav .ant-menu-inline {
    width: 40%;
}
</style>