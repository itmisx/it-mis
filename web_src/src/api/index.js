import request from "@/utils/request.js";

// 登录
export function loginAPI(data) {
    return request({
        url: "/index/login",
        method: "post",
        data,
        loadingID: "login",
        loadingTip: "登录中",
    });
}
// 登出
export function logoutAPI() {
    return request({
        url: "/index/logout",
        method: "post",
    });
}

// 登录状态
export function loginStatusAPI() {
    return request({
        url: "/index/login-status",
        method: "get",
    });
}