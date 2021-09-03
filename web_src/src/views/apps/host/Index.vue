<template>
  <div>
    <a-card title="JumpServer" style="height: 100%; width: 100%">
      <a-tabs default-active-key="1" @change="callback">
        <a-tab-pane key="1">
          <span slot="tab"> <a-icon type="schedule" />服务主机</span>
          <div style="height: 100%; width: 100%">
            <host-list></host-list>
          </div>
        </a-tab-pane>
        <a-tab-pane key="2">
          <span slot="tab"> <a-icon type="project" />分组授权</span>
          Content of Tab Pane 1
        </a-tab-pane>
        <a-tab-pane key="3">
          <span slot="tab"> <a-icon type="bug" />运维审计</span>
          Content of Tab Pane 1
        </a-tab-pane>
      </a-tabs>
    </a-card>
  </div>
</template>
<script>
import HostList from "./HostList.vue";
export default {
  components: {
    HostList,
  },
  data: function () {
    return {};
  },
  methods: {
    callback: function () {},
  },
};
</script>

<style lang="scss" scoped></style>
