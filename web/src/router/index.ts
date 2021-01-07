import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import Home from "../views/Home.vue";
import { getRequest_checkSession} from "@/http/request.js";
import { useStore } from '../store/index'
const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home
  },
  {
    path: "/search",
    name: "Search",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/Search.vue")
  },
  {
    path: "/ban",
    name: "BanLift",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    meta: { requireAuth: true, roles: 0 },
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/BanLift.vue")
  },
  {
    path: "/question",
    name: "Question",

    component: () =>
      import(/* webpackChunkName: "about" */ "../views/QuestionDetail.vue")
  },
  {
    path: "/hotRank",
    name: "HotRank",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/HotRank.vue")
  },
  {
    path: "/recommend",
    name: "Recommend",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/Recommend.vue")
  },
  {
    path: "/categoryLife",
    name: "CategoryLife",
    component: () =>
        import(/* webpackChunkName: "about" */ "../views/CategoryLife.vue")
  },
  {
    path: "/categoryStudy",
    name: "/CategoryStudy",
    component: () =>
        import(/* webpackChunkName: "about" */ "../views/CategoryStudy.vue")
  },
  {
    path: "/postQuestion",
    name: "PostQuestion",
    meta: { requireAuth: true },
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/PostQuestion.vue")
  },
  {
    path: "/personal",
    name: "Personal",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    meta: { requireAuth: true },
    component: () =>
      import(/* webpackChunkName: "about" */ "../views//personalView/Personal.vue")
  },
  {
    path: "/personalFollowing",
    name: "PersonalFollowing",
    meta: { requireAuth: true },
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/personalView/PersonalFollowing.vue")
  },
  {
    path: "/personalFollower",
    name: "PersonalFollower",
    meta: { requireAuth: true },
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/personalView/PersonalFollower.vue")
  },
  {
    path: "/personalQuestion",
    name: "PersonalQuestion",
    meta: { requireAuth: true },
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/personalView/PersonalQuestion.vue")
  },
  {
    path: "/personalAnswer",
    name: "PersonalAnswer",
    meta: { requireAuth: true },
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/personalView/PersonalAnswer.vue")
  },
  {
    path: "/personalMessage",
    name: "PersonalMessage",
    meta: { requireAuth: true },
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/personalView/PersonalMessage.vue")
  },
  {
    path: "/personalSet",
    name: "PersonalSet",
    meta: { requireAuth: true },
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/personalView/PersonalSet.vue")
  },
  {
    path: "/login",
    name: "Login",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/Login.vue")
  },
  {
    path: "/register",
    name: "Register",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/Register.vue")
  },
  {
    path: "/personalSetOthers",
    name: "PersonalSetOthers",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/personalView/PersonalOthers.vue")
  },
  {
    path: "/personalCollection",
    name: "PersonalCollection",
    meta: { requireAuth: true },
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/personalView/PersonalCollection.vue")
  }

];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

// let checksession_callback =(response)=>{
//   if(response.code===1)
//   {
//     sessionStorage.removeItem("user");

//   }
// }
// 路由守卫
router.beforeEach((to, from, next) => {

  const flag = sessionStorage.getItem('user')
  let role_check = 4

  if (flag) {
    role_check = JSON.parse(flag).role;
  }
  console.log(role_check)


  if (to.meta.requireAuth == true) { // 需要登录权限进入的路由
    if (!flag) {     // 获取不到登录信息
      next({
        path: '/login'
      })
    } else {
      // 获取到登录信息，进行下一步

      // 先checksession
      getRequest_checkSession((res)=>{
        if(res.code===1)
        {
          sessionStorage.removeItem("user");
          next({
            path: '/login'
          })
        }
        else if(res.code===2){
          console.log("check session time out")
        }

      }
      )

      if (to.meta.roles===0) {
        if (role_check !== 0) {
          sessionStorage.removeItem("user");
          next({
            path: '/login'
          })
        } else {
          return next();
        }
      } else {
        return next();
      }

      return next();
    }
  } else {       // 不需要登录权限的路由直接进行下一步
    return next();
  }



}


)

export default router;
