<template>
	<div class="body">
		<div class="c-header">
			<config
				theme="outline"
				size="20"
				fill="#333"
				style="margin-top: 6px; margin-right: 2px"
			/>
			<span style="font-size: 16px">设置</span>
		</div>
		<div class="center">
			<a-tabs
				default-active-key="1"
				tab-position="left"
				@prevClick="callback"
				@nextClick="callback"
			>
				<a-tab-pane tab="个人设置" key="1"></a-tab-pane>
				<a-tab-pane tab="系统邮箱" key="2"></a-tab-pane>
			</a-tabs>
		</div>
	</div>
</template>
<script>
import { Config } from "@icon-park/vue";
export default {
	data() {
		return {};
	},
	components: {
		Config,
	},
	methods: {
		callback(val) {
			console.log(val);
		},
	},
};
</script>

<style lang="scss" scoped>
.body {
	width: 100%;
	height: 100%;
	display: flex;
	display: -webkit-flex;
	flex-direction: column;
}
.c-header {
	padding: 0px 20px;
	height: 60px;
	border-bottom: 1px solid #e0e0e0;
	display: flex;
	display: -webkit-flex;
	flex-direction: row;
	align-items: center;
}
.center {
	flex: 1;
	display: flex;
	display: -webkit-flex;
	flex-direction: row;
	margin: 20px 20px;
	border: 1px solid #e0e0e0;
	padding: 10px 10px;
}
</style>
