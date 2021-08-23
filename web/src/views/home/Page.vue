<template>
  <div class="center">
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  mounted() {
    let lastestRoutePath = sessionStorage.getItem("lastest_nav_selected");
    if (lastestRoutePath) {
      this.$router.push(lastestRoutePath);
    } else {
      // 默认
      this.$router.push("/home/apps");
    }
  },
  watch: {},
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
