<template class="b">
  <div >
    <a-row >
      <a-col :span="1" :offset="1">
        <div class="personal-nav">
          <SubMenu />
        </div>
      </a-col>

      <a-col :span="20" :offset="1">
        <a-row justify="center"><br />
          <a-col :span="7"><br />
            <a-input-search
            
              placeholder="输入用户名或昵称"
              v-model:value="value"
               style=" border-radius: 25px;box-shadow: 3px 3px 3px #dcdfdf"
              @search="onSearch"
              
            />
          </a-col>
          <a-col :span="2" :offset="1">
            <br />
            <a-button type="primary" shape="circle" @click="onCancle"><CloseOutlined /></a-button>
          </a-col>
          <br /><br /><br /><br/>
          <a-col :span="20" >
            <a-list :grid="{ gutter: 16, column: 4 }" :data-source="this.showData">
              <template #renderItem="{ item }">
                <a-list-item>
                  <div class="follower-block">
                    <!-- <a-row >
                    <a-col>-->
                    <a-row justify="center">
                      <a-col>
                        <br />
                        <img
                         @click="gotoPerson(item.uid)"
                          slot="cover"
                          alt="example"
                          :src="item.icon"
                          style="height: 60px; border-radius: 50% ;box-shadow: 3px 3px 3px #dcdfdf"
                          
                        />
                        <br />
                        <!-- <br /> -->
                        <span class="follower-name">{{item.name}}</span>
                        <br />
                        <span class="follower-nickname">{{item.nickname}}</span>
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
      </a-col>
    </a-row>
    <br />

    <br />
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { Options, Vue } from "vue-class-component";
import { UserOutlined,CloseOutlined} from "@ant-design/icons-vue";
import SubMenu from "../../components/PersonalNavigation";
import { getRequest } from "@/http/request.js";

export default {
  components: {
    UserOutlined,
    SubMenu,CloseOutlined
  },
  data() {
    return {
      data: [],
      showData:[],
      value:''
    };
  },
  created() {
    let id = JSON.parse(sessionStorage.getItem("user")).uid;
    getRequest("/followers", this.handleCallback, {
      errorCallback: e => {
        console.log(e);
      },
      params: { uid: id }
    });
  },
  methods: {
    onCancle()
    {
      this.showData=this.data;
    },
    handleCallback(response) {
      console.log(response);
      this.data = response.result;
      this.showData = response.result;
    },
    onSearch(value) {
      this.showData=this.data.filter(item => item.name.indexOf(value)>=0 || item.nickname.indexOf(value)>=0)
    },
    gotoPerson(id) {
      console.log(id)
      this.$router.push({
        path: "/personalSetOthers",
        query: { uId: id }
      });
    }
  }
};
</script>

<style >
body {
  min-width: 700px;
  background-color: #edeeed;
}

.follower-block {
  text-align: center;
  background-color: #ffffff;
  width: 100%;
  height: 246px;
  border-radius: 20px;
  padding: 10px;
  margin: 0 auto;
  box-shadow: 3px 3px 2px #dcdfdf;
}

.follower-name {
  font-size: 18px;
  color: #485355f5;
text-shadow: 1px 1px 1px rgba(47, 56, 55, 0.384);
  font-weight: bold;
}
.follower-nickname {
  font-size: 12px;
  color: #485355f5;

  font-weight:lighter;
   overflow: hidden;

  text-overflow: ellipsis;




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