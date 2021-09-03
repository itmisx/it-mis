<template>
  <a-spin :tip="tip" :spinning="spinning">
    <slot></slot>
  </a-spin>
</template>

<script setup>
import { eventBus } from "@/main.js";
export default {
  name: "SpinX",
  props: {
    loadingID: {
      type: String,
      default: "",
      require: true,
    },
  },
  data() {
    return {
      spinning: false,
      tip: "loading",
    };
  },
  mounted() {
    // 显示加载
    // 匹配loadingID
    eventBus.$on("showLoading", (loadingID, tip) => {
      if (this.loadingID == loadingID) {
        this.spinning = true;
        if (tip) {
          this.tip = tip;
        }
      }
    });
    // 隐藏加载
    // 匹配loadingID
    eventBus.$on("hideLoading", (loadingID) => {
      if (this.loadingID == loadingID) {
        this.spinning = false;
      }
    });
  },
};
</script>

<style lang="scss" scoped></style>
