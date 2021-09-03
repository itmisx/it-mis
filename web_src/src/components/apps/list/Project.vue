<template>
	<a-card hoverable class="card" :bodyStyle="{ height: '140px' }">
		<a-card-meta :title="title" :description="desc">
			<bookmark slot="avatar" theme="outline" size="24" fill="#1890ff" />
		</a-card-meta>
		<template slot="actions" class="ant-card-actions">
			<a-icon key="edit" type="star" />
			<a-icon key="setting" type="arrow-right" @click="enterApp" />
		</template>
	</a-card>
</template>

<script>
import { Bookmark } from "@icon-park/vue";
export default {
	name: "Project",
	components: {
		Bookmark,
	},
	data() {
		return {
			name: "project",
			title: "项目管理",
			desc: "产品线、项目、需求、任务、测试",
		};
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
