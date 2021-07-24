<template>
  <div class="center">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  mounted() {
    // 获取上次路由地址
    let router = this.$router;
    let lastRoutePath = sessionStorage.getItem("last-route-path");
    if (lastRoutePath) {
      router.push(lastRoutePath);
    } else {
      // 默认
      router.push("/home/app-list");
    }
  },
  watch: {
    $route: function (newRoute) {
      sessionStorage.setItem("last-route-path", newRoute.path);
    },
  },
};
</script>

<style lang="scss" scoped>
.center {
  display: none;
  display: -webkit-flex;
  display: flex;
  flex-direction: column;
  flex: 1;
}
</style>
