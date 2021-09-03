<template>
	<div class="body">
		<div class="tbar">
			<a-form-item label="名称" labelAlign="left">
				<a-input style="width: 300px; margin-right: 5px" placeholder="名称" />
			</a-form-item>
			<a-select
				default-value="站点检测"
				style="width: 150px; margin-right: 5px"
				@change="handleChange"
			>
				<a-select-option value="站点检测">站点检测</a-select-option>
				<a-select-option value="端口检测">端口检测</a-select-option>
			</a-select>

			<a-select
				default-value="正常"
				style="width: 150px; margin-right: 5px"
				@change="handleChange"
			>
				<a-select-option value="正常">正常</a-select-option>
				<a-select-option value="异常">异常</a-select-option>
			</a-select>
			<a-button-group>
				<a-button type="primary"><a-icon type="search" />查询</a-button>
				<a-button type="primary"><a-icon type="plus" />新建监听</a-button>
			</a-button-group>
		</div>
		<a-table :columns="columns" :data-source="data">
			<span slot="action" slot-scope="record, index">
				<a @click="enable">禁用</a>
				<a-divider type="vertical" />
				<a @click="edit">编辑</a>
				<a-divider type="vertical" />
				<a @click="del(record, index)">删除</a>
			</span>
		</a-table>
	</div>
</template>
<script>
const columns = [
	{
		title: "名称",
		dataIndex: "name",
		key: "name",
	},
	{
		title: "类型",
		dataIndex: "type_text",
		key: "type_text",
	},
	{
		title: "地址",
		dataIndex: "address",
		key: "address",
	},
	{
		title: "监测间隔",
		dataIndex: "interval",
		key: "interval",
	},
	{
		title: "状态",
		key: "tags",
		dataIndex: "status_text",
	},
	{
		title: "操作",
		key: "action",
		width: 200,
		scopedSlots: { customRender: "action" },
	},
];

const data = [
	{
		key: "1",
		name: "端口",
		type: 1,
		type_text: "端口检测",
		address: "127.0.0.1",
		interval: "5分钟",
		status: 1,
		status_text: "正常",
	},
	{
		key: "2",
		name: "站点检测",
		type: 1,
		type_text: "端口检测",
		address: "127.0.0.1",
		interval: "5分钟",
		status: 1,
		status_text: "正常",
	},
];
export default {
	data() {
		return {
			data,
			columns,
		};
	},
	methods: {
		gotoApps() {
			this.$router.push("/home/apps");
		},
		enable() {
			alert("enable");
		},
		edit() {
			alert("edit");
		},
		del(record, index) {
			console.log(record, index);
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
.tbar {
	display: flex;
	display: -webkit-flex;
	flex-direction: row;
	margin-bottom: 10px;
}
</style>
