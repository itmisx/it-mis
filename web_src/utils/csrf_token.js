import axios from "axios";
import { headers } from "@/utils/request.js";
import setting from "@/setting.js";
export const getCsrfToken = async function() {
    axios.defaults.withCredentials = true;
    await axios.get(setting.baseURL + "/index/csrf-token").then((res) => {
        let data = res.data;
        if (data.code === 0 && data.data.csrf_token) {
            headers.CSRFToken = res.data.data.csrf_token;
            console.log(123);
            // 全局ajax拦截
            import ("@/utils/interceptors.js");
        }
    });
};