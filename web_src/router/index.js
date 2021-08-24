import Vue from "vue";
import VueRouter from "vue-router";
import NProgress from "nprogress";

Vue.use(VueRouter);

const routes = [{
        path: "/",
        component: () =>
            import ("@/views/Login.vue"),
    },
    {
        path: "/home",
        component: () =>
            import ("@/views/home/Home.vue"),
        children: [{
                path: "apps",
                component: () =>
                    import ("@/views/apps/List.vue"),
            },
            {
                path: "host",
                component: () =>
                    import ("@/views/apps/host/Index.vue"),
            },
            {
                path: "org",
                component: () =>
                    import ("@/views/org/Index.vue"),
            },
            {
                path: "system",
                component: () =>
                    import ("@/views/system/Index.vue"),
            },
            {
                path: "store",
                component: () =>
                    import ("@/views/store/Index.vue"),
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