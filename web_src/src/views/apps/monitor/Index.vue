<template>
	<div class="body">
		<div class="c-header">
			<electrocardiogram
				slot="avatar"
				theme="outline"
				size="20"
				fill="#1890ff"
				style="margin-top: 6px"
			/>
			<span style="margin-left: 2px">监控</span>
			<a-button icon="rollback" style="margin-left: 20px" @click="gotoApps"
				>返回应用中心
			</a-button>
		</div>
		<div style="background: #f2f0f5; flex: 1">
			<div style="margin: 10px 10px; padding: 10px 10px; background: #fff">
				<task-list></task-list>
			</div>
		</div>
	</div>
</template>
<script>
import { Electrocardiogram } from "@icon-park/vue";
import TaskList from "./TaskList.vue";
export default {
	components: {
		Electrocardiogram,
		TaskList,
	},
	data() {
		return {};
	},
	methods: {
		gotoApps() {
			this.$router.push("/home/apps");
		},
	},
};
</script>

<style lang="scss" scoped>
.body {
	height: 100%;
	width: 100%;
	display: flex;
	display: -webkit-flex;
	flex-direction: column;
}
.c-header {
	margin-left: 20px;
	height: 60px;
	font-size: 18px;
	display: flex;
	display: -webkit-flex;
	align-items: center;
}
</style>
