import Vue from "vue";
import VueRouter from "vue-router";
import NProgress from "nprogress";

Vue.use(VueRouter);

// 路由页面定义
const login = () =>
    import ("@/views/Login.vue");
const home = () =>
    import ("@/views/home/Home.vue");
const apps = () =>
    import ("@/views/apps/List.vue");
const org = () =>
    import ("@/views/org/Index.vue");
const system = () =>
    import ("@/views/system/Index.vue");
const store = () =>
    import ("@/views/store/Index.vue");
const host = () =>
    import ("@/views/apps/host/Index.vue");
const monitor = () =>
    import ("@/views/apps/monitor/Index.vue");

// 路由表
const routes = [{
        path: "/",
        component: login,
    },
    {
        path: "/home",
        component: home,
        children: [{
                path: "apps",
                component: apps,
            },
            {
                path: "org",
                component: org,
            },
            {
                path: "system",
                component: system,
            },
            {
                path: "store",
                component: store,
            },
            {
                path: "monitor",
                component: monitor,
            },
            {
                path: "host",
                component: host,
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

router.beforeEach((to, from, next) => {
    NProgress.start();
    next();
});
router.afterEach(() => {
    NProgress.done();
});

export default router;