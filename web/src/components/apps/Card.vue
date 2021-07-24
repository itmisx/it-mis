<template>
  <a-card hoverable class="card" @native:click="enterApp">
    <a-card-meta :title="appTitle" :description="desc">
      <a-icon
        slot="avatar"
        type="project"
        style="font-size: 24px; color: rgb(33, 171, 243)"
      />
    </a-card-meta>
    <template slot="actions" class="ant-card-actions">
      <a-icon key="edit" type="star" />
      <a-icon key="setting" type="arrow-right" @click="enterApp" />
    </template>
  </a-card>
</template>

<script>
export default {
  name: "Card",
  props: {
    appName: {
      type: String,
    },
    appTitle: {
      type: String,
    },
    desc: {
      type: String,
      default: "版本管理、软件提测、软件发布~",
    },
  },
  data: function () {
    return {};
  },
  methods: {
    enterApp: function () {
      this.$emit("enterApp", this.appName);
    },
  },
};
</script>

<style lang="scss" scoped>
.card {
  margin: 10px 10px;
  width: 250px;
}
</style>
