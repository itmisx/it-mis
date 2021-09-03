<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  mounted() {
    // 判断登录状态，决定进入登录页面或主页
    this.$router.push("/");
  },
};
</script>

<style lang="scss">
* {
  font-family: "PingFang SC", "Helvetica Neue", Helvetica, "Hiragino Sans GB",
    "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
}
html,
body,
#app {
  margin: 0;
  padding: 0;
  height: 100%;
  width: 100%;
  color: #2c3e50;
  background-color: #fafafa;
}
</style>
