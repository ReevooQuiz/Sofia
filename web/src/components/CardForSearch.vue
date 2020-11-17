<template>
  <div class="cardForSearch">
    <a-card hoverable title="金平路" size="small" style="border-radius : 3px">
     
    </a-card>
    <br/>
  </div>
</template>
       


<script >
import moment from "moment";
export default {
  data() {
    return {
      likes: 0,
      dislikes: 0,
      action: null,
    
      moment
    };
  },
  methods: {
    like() {
      this.likes = 1;
      this.dislikes = 0;
      this.action = "liked";
    },
    dislike() {
      this.likes = 0;
      this.dislikes = 1;
      this.action = "disliked";
    }
  }
};
</script>
