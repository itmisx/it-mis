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
        <a-menu-item key="apps">
          <a-icon type="appstore" />
          <span>应用中心</span>
        </a-menu-item>
        <a-menu-item key="org">
          <a-icon type="team" />
          <span>组织团队</span>
        </a-menu-item>
        <a-menu-item key="system">
          <a-icon type="setting" />
          <span>系统设置</span>
        </a-menu-item>
        <a-menu-item key="store">
          <a-icon type="shopping" theme="twoTone" two-tone-color="#eb2f96" />
          <span>应用商店</span>
        </a-menu-item>
      </a-menu>
    </div>
    <div class="avtar">
      <a-dropdown>
        <a-avatar style="background-color: #1890ff" icon="user" />
        <a-menu slot="overlay" @click="handleMenu">
          <a-menu-item key="logout">
            <a href="javascript:;">退出</a>
          </a-menu-item>
        </a-menu>
      </a-dropdown>

      <span style="margin-top: 10px">©IT-MIS</span>
    </div>
  </div>
</template>

<script>
import { logoutAPI } from "@/api/index.js";
export default {
  data: function () {
    return {
      navSelected: [],
    };
  },
  beforeMount: function () {
    // 获取上次导航菜单
    let lastNavSelected = localStorage.getItem("lastest_nav_selected");
    if (lastNavSelected) {
      this.navSelected = [lastNavSelected];
    } else {
      // 默认值
      this.navSelected = ["apps"];
    }
  },
  methods: {
    navChange(item) {
      localStorage.setItem("lastest_nav_selected", item.key);
      this.navSelected = [item.key];
      switch (item.key) {
        case "apps":
          this.$router.push("/home/apps");
          break;
        case "org":
          this.$router.push("/home/org");
          break;
        case "system":
          this.$router.push("/home/system");
          break;
        case "store":
          this.$router.push("/home/store");
          break;
      }
    },
    handleMenu(e) {
      switch (e.key) {
        case "logout":
          logoutAPI().then((res) => {
            if (res.code === 0) {
              localStorage.removeItem("user_token");
              this.$router.push("/");
            }
          });
          break;
        default:
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
    font-size: 34px;
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
