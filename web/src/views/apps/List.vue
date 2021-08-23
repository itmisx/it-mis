<template>
  <div>
    <div class="center">
      <input class="c-header" placeholder="搜索应用~" />
      <a-divider style="margin: 0"></a-divider>
      <div class="c-body">
        <app-card
          appName="host"
          appTitle="堡垒机"
          desc="简易堡垒机系统，支持SSH、SFTP、运维授权、安全审计~"
          @enterApp="enterApp('host')"
        ></app-card>
      </div>
    </div>
  </div>
</template>

<script>
import AppCard from "@/components/apps/Card.vue";
export default {
  components: {
    AppCard,
  },
  mounted() {
    this.showApps = true;
  },
  methods: {
    enterApp: function (appName) {
      this.$router.push("/home/" + appName);
    },
  },
};
</script>

<style lang="scss" scoped>
.center {
  display: -webkit-flex;
  display: flex;
  flex-direction: column;
  flex: 1;
}
.c-header {
  border: none;
  outline: none;
  padding: 0px 50px;
  line-height: 60px;
  background: url("~@/assets/home/search.svg") left center no-repeat;
  background-size: 14px 14px;
  background-position: 30px;
}
.c-header:empty:before {
  content: attr(placeholder);
  display: block;
  color: #888;
}
.c-body {
  padding: 10px 10px;
  flex: 1;
  display: -webkit-flex;
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  align-content: flex-start;
}
</style>
