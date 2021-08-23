import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
// 引入antd
import Antd from "ant-design-vue";
import "ant-design-vue/dist/antd.css";
// 引入nprogress
import NProgress from "nprogress";
import "nprogress/nprogress.css";
NProgress.configure({
    easing: "ease", // 动画方式
    speed: 500, // 递增进度条的速度
    showSpinner: false, // 是否显示加载ico
    trickleSpeed: 200, // 自动递增间隔
    minimum: 0.3, // 初始化时的最小百分比
});
Vue.use(Antd);
// 组件注册
import SpinX from "@/components/common/spinx";
Vue.component("spin-x", SpinX);

// 获取csrf-token
import { getCsrfToken } from "@/utils/csrf_token.js";
getCsrfToken();

// 定义eventBus用于传递事件
export var eventBus = new Vue();

Vue.config.productionTip = false;

new Vue({
    router,
    store,
    render: (h) => h(App),
}).$mount("#app");