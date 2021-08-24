<template>
  <div class="main" v-if="showLoginWindow">
    <spin-x loadingID="login">
      <a-card class="box-card">
        <div class="login-form-title">
          <span style="font-size: 36px">{{ `IT-MIS` }}</span>
        </div>
        <div
          style="
            text-align: center;
            margin-top: 10px;
            font-size: 16px;
            color: #999999;
          "
        >
          <span>IT信息化一站式平台</span>
        </div>
        <a-form style="margin: 10px 40px 5px">
          <a-input
            placeholder="账号"
            v-model="user"
            prefix-icon="iconfont icon-user"
            size="large"
            @keyup.enter.native="login"
            class="form-input"
          >
            <a-icon slot="prefix" type="user" />
          </a-input>
          <a-input-password
            placeholder="密码"
            v-model="password"
            show-password
            prefix-icon="iconfont icon-key"
            size="large"
            @keyup.enter.native="login"
            class="form-input"
          >
            <a-icon slot="prefix" type="lock" />
          </a-input-password>
          <a-button
            type="primary"
            style="width: 100%"
            size="large"
            @click="submit"
            class="form-input"
            >登录
          </a-button>
        </a-form>
      </a-card>
    </spin-x>
    <div class="bottom">
      <span>©IT-MIS 版权所有</span>
    </div>
  </div>
</template>
<script>
import { loginAPI } from "@/api/index.js";
export default {
  name: "Login",
  data() {
    return {
      user: "",
      password: "",
      showLoginWindow: false,
    };
  },
  methods: {
    submit() {
      loginAPI({
        user: this.user,
        passwd: this.password,
      }).then((res) => {
        if (res.code === 0) {
          this.$router.push("/home");
        }
      });
    },
  },
  mounted() {
    let userToken = localStorage.getItem("user_token");
    if (userToken) {
      this.$router.push("/home");
    } else {
      this.showLoginWindow = true;
    }
  },
};
</script>

<style scoped>
.main {
  width: 100%;
  height: 100%;
  background-image: url("~@/assets/bg.svg");
  background-repeat: no-repeat;
  background-position: center 110px;
  background-size: 100%;
  background-color: #f0f2f5;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}
.box-card {
  height: 430px;
  width: 430px;
  background-color: white;
}
.form-input {
  margin-top: 10px;
}
.login-form-title {
  margin-top: 30px;
  color: #1890ff;
  font-size: 32px;
  display: flex;
  display: -webkit-flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}
.bottom {
  width: 100%;
  height: 30px;
  position: absolute;
  bottom: 0px;
  text-align: center;
  font-size: 12px;
}
</style>
