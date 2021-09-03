<template>
	<a-card hoverable class="card" :bodyStyle="{ height: '140px' }">
		<a-card-meta :title="title" :description="desc">
			<server slot="avatar" theme="outline" size="24" fill="#1890ff" />
		</a-card-meta>
		<template slot="actions" class="ant-card-actions">
			<a-icon key="edit" type="star" />
			<a-icon key="setting" type="arrow-right" @click="enterApp" />
		</template>
	</a-card>
</template>

<script>
import { Server } from "@icon-park/vue";
export default {
	name: "Host",
	components: {
		Server,
	},
	data() {
		return {
			appName: "host",
			title: "Teleport",
			desc: "堡垒机系统，支持SSH、SFTP、运维授权、安全审计",
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
