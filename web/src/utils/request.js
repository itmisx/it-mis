import axios from "axios";
import { Message } from "element-ui";

export default function fetch(options) {
  var instance = axios.create({
    baseURL: domain,
    timeout: 5000,
    withCredentials: true,
  });
  // header加上token
  if (localStorage.getItem("token")) {
    axios.defaults.headers.common["token"] = localStorage.getItem("token");
  }
  // 添加请求拦截器
  instance.interceptors.request.use(
    (config) => {
      // alert('request')
      // 在发送请求之前做些什么
      return config;
    },
    (error) => {
      return Promise.reject(error);
    }
  );

  // 添加响应拦截器
  instance.interceptors.response.use(
    function (response) {
      if (response.data.code == 200000) {
        let href = window.location.href;
        if (href.indexOf("/login") == -1) {
          window.location.href = "/#login";
          location.reload();
        }
        return response;
      } else {
        return response;
      }
    },
    function (error) {
      let err = error.data ? error.data : error;
      Message({
        message: err,
      });
      return Promise.reject(error);
    }
  );
  return instance(options);
}
