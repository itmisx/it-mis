<template>
	<a-card hoverable class="card" :bodyStyle="{ height: '140px' }">
		<a-card-meta :title="title" :description="desc">
			<electrocardiogram
				slot="avatar"
				theme="outline"
				size="24"
				fill="#1890ff"
			/>
		</a-card-meta>
		<template slot="actions" class="ant-card-actions">
			<a-icon key="edit" type="star" />
			<a-icon key="setting" type="arrow-right" @click="enterApp" />
		</template>
	</a-card>
</template>

<script>
import { Electrocardiogram } from "@icon-park/vue";
export default {
	name: "Monitor",
	components: {
		Electrocardiogram,
	},
	data() {
		return {
			appName: "monitor",
			title: "监控",
			desc: "服务监控，支持站点、端口、Ping、进程及自定义脚本",
		};
	},
	methods: {
		enterApp: function () {
			console.log(123);
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
