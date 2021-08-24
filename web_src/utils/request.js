import axios from "axios";
import { message } from "ant-design-vue";
import { eventBus } from "../main.js";
import setting from "@/setting.js";

const headers = {
    CSRFToken: "",
};

// request 请求封装
// options
// -- url,请求uri
// -- method，请求方法
export default function request(options) {
    var instance = axios.create({
        baseURL: setting.baseURL,
        timeout: 30000,
        withCredentials: true,
    });
    // header加上csrf-token
    instance.defaults.headers.common["CSRFToken"] = headers.CSRFToken;
    // 添加请求拦截器
    instance.interceptors.request.use(
        (config) => {
            // 触发showLoading事件
            eventBus.$emit("showLoading", options.loadingID, options.loadingTip);
            return config;
        },
        (error) => {
            return Promise.reject(error);
        }
    );
    // 添加响应拦截器
    instance.interceptors.response.use(
        (response) => {
            // 触发hideLoading事件
            eventBus.$emit("hideLoading", options.loadingID);
            if (response.data.code === 0) {
                return response.data;
            } else if (response.data.code == 200000) {
                let href = window.location.href;
                if (href.indexOf("/login") == -1) {
                    window.location.href = "/#login";
                    location.reload();
                }
            } else {
                message.error(response.data.msg);
            }
        },
        (error) => {
            eventBus.$emit("hideLoading", options.loadingID);
            message.error("请求出错，请检查网络！");
            return Promise.reject(error);
        }
    );
    return instance(options);
}

export { headers };