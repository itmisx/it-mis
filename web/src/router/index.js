import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: () => import("../views/SignIn.vue"),
  },
  {
    path: "/home",
    component: () => import("../views/home/Index.vue"),
    children: [
      {
        path: "app-list",
        component: () => import("../views/app/List.vue"),
      },
      {
        path: "deployment",
        component: () => import("../views/deployment/Index.vue"),
      },
      {
        path: "host",
        component: () => import("../views/host/Index.vue"),
      },
      {
        path: "team",
        component: () => import("../views/team/Index.vue"),
      },
      {
        path: "setting",
        component: () => import("../views/setting/Index.vue"),
      },
    ],
  },
];

//以下是为了屏蔽路由重复push的错误打印
//获取原型对象上的push函数
//修改原型对象中的push方法
const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch((err) => err);
};

const router = new VueRouter({
  routes,
});

export default router;
