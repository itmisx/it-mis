<template>
  <div>
    <a-card
      title="项目管理"
      style="height: 100%; width: 100%"
      :bodyStyle="{
        fontSize: '24px',
        height: '100px',
        width: '200px',
        backgroudColor: 'red',
      }"
      ><p>card content</p>
      <!-- <a-tabs default-active-key="1" @change="callback">
        <a-tab-pane key="1">
          <span slot="tab"> <a-icon type="schedule" />产品</span>
          Content of Tab Pane 1
        </a-tab-pane>
        <a-tab-pane key="2">
          <span slot="tab"> <a-icon type="project" />项目</span>
          Content of Tab Pane 1
        </a-tab-pane>
        <a-tab-pane key="3">
          <span slot="tab"> <a-icon type="bug" />测试</span>
          Content of Tab Pane 1
        </a-tab-pane>
        <a-tab-pane key="6">
          <span slot="tab"> <a-icon type="thunderbolt" />发布</span>
          Content of Tab Pane 3
        </a-tab-pane>
        <a-tab-pane key="7">
          <span slot="tab"> <a-icon type="setting" />设置</span>
          Content of Tab Pane 3
        </a-tab-pane>
      </a-tabs> -->
    </a-card>
  </div>
</template>

<script>
export default {
  data: function () {
    return {};
  },
  methods: {
    callback: function () {},
  },
};
</script>

<style lang="scss" scoped></style>
