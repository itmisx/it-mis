<template>
  <div class="left">
    <div class="title">IT-MIS</div>
    <div class="nav">
      <a-menu
        style="width: 180px"
        :selectedKeys="navSelected"
        mode="inline"
        theme="dark"
        @select="navChange"
      >
        <a-menu-item key="app-list">
          <a-icon type="appstore" />
          <span>应用中心</span>
        </a-menu-item>
        <a-menu-item key="team">
          <a-icon type="team" />
          <span>组织团队</span>
        </a-menu-item>
        <a-menu-item key="setting">
          <a-icon type="setting" />
          <span>系统设置</span>
        </a-menu-item>
        <a-menu-item key="app-store">
          <a-icon type="shopping" theme="twoTone" two-tone-color="#eb2f96" />
          <span>应用商店</span>
        </a-menu-item>
      </a-menu>
    </div>
    <div class="avtar">
      <a-avatar style="background-color: #1890ff" icon="user" />
      <span style="margin-top: 10px">©IT-MIS</span>
    </div>
  </div>
</template>

<script>
export default {
  data: function () {
    return {
      navSelected: [],
    };
  },
  beforeMount: function () {
    // 获取上次导航菜单
    let lastNavSelected = sessionStorage.getItem("last-nav-selected");
    if (lastNavSelected) {
      this.navSelected = [lastNavSelected];
    } else {
      // 默认值
      this.navSelected = ["app-list"];
    }
  },
  methods: {
    navChange: function (item) {
      sessionStorage.setItem("last-nav-selected", item.key);
      this.navSelected = [item.key];
      switch (item.key) {
        case "app-list":
          this.$router.push("/home/app-list");
          break;
        case "team":
          this.$router.push("/home/team");
          break;
        case "setting":
          this.$router.push("/home/setting");
          break;
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.left {
  width: 180px;
  height: 100%;
  background-color: rgb(0, 21, 41);
  display: flex;
  display: -webkit-flex;
  flex-direction: column;
  align-items: center;
  .title {
    margin-top: 5px;
    text-align: center;
    font-size: 28px;
    font-weight: 700;
    color: #1890ff;
  }
  .nav {
    flex: 1;
    display: flex;
    display: -webkit-flex;
    flex-direction: column;
    align-items: center;
    .nav-item {
      color: white;
      font-size: 14px;
      margin-bottom: 20px;
      background-color: teal;
    }
  }
  .avtar {
    display: flex;
    display: -webkit-flex;
    flex-direction: column;
    align-items: center;
  }
}
</style>
