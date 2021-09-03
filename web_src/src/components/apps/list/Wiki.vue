<template>
	<a-card hoverable class="card" :bodyStyle="{ height: '140px' }">
		<a-card-meta :title="title" :description="desc">
			<bookshelf slot="avatar" theme="outline" size="24" fill="#1890ff" />
		</a-card-meta>
		<template slot="actions" class="ant-card-actions">
			<a-icon key="edit" type="star" />
			<a-icon key="setting" type="arrow-right" @click="enterApp" />
		</template>
	</a-card>
</template>

<script>
import { Bookshelf } from "@icon-park/vue";
export default {
	name: "Wiki",
	components: {
		Bookshelf,
	},
	data() {
		return {
			name: "wiki",
			title: "知识库",
			desc: "支持富文本、Markdown、脑图、流程图、excel",
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
