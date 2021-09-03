import Vue from "vue";
import App from "./App.vue";
import router from "./router";
// 引入antd
import Antd from "ant-design-vue";
import "ant-design-vue/dist/antd.css";
Vue.use(Antd);
// nprogress
import ("@/components/common/NProgress");
// 组件注册
import SpinX from "@/components/common/Spinx";
Vue.component("spin-x", SpinX);

// 获取csrf-token
import { getCsrfToken } from "@/utils/csrf_token.js";
getCsrfToken();

// 定义eventBus用于传递事件
export var eventBus = new Vue();

Vue.config.productionTip = false;

new Vue({
    router,
    render: (h) => h(App),
}).$mount("#app");