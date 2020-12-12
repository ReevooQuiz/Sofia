<template>
  <div>
    <a-row justify="center">
      <a-col>
        <img
          slot="cover"
          alt="example"
          src="https://img.zcool.cn/community/0147a05664f33432f8754573f3e555.jpg"
          style="height: 150px; border-radius: 50%"
        />
        <!-- <img src="../assets/cactus2.png" height="345" width="180"    style="marginTop: 340px"/> -->
      </a-col>
    </a-row>
    <br />
    <a-row justify="center">
      <a-tag color="#88d5d1">学习区专家</a-tag>
    </a-row>
    <br />

    <a-row justify="center">
      <a-list :grid="{ gutter: 16, column: 4 }"  :data-source="dataa">
        <a-list-item slot="renderItem" slot-scope="item,index">
        <!-- {{item.title}} -->
        {{index}}
        </a-list-item>
      </a-list>
    </a-row>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { Options, Vue } from "vue-class-component";
import { UserOutlined } from "@ant-design/icons-vue";
const data1 = [
  {
    title: "Title 1"
  },
  {
    title: "Title 2"
  },
  {
    title: "Title 3"
  },
  {
    title: "Title 4"
  }
];
export default {
  components: {
    UserOutlined
  },
  data() {
    return {
      dataa : data1,
    };
  },
  methods: {
    onSearch(value) {
      if (value != "") {
        this.$router.push({ path: "/search", query: { content: value } });
      }
    }
  }
};
</script>

<style>
#search {
  background-color: #edeeed;
}
</style>