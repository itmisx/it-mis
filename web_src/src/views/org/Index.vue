<template>
	<div class="body">
		<div class="c-header">
			<city
				theme="outline"
				size="20"
				fill="#333"
				style="margin-top: 6px; margin-right: 2px"
			/>
			<span style="font-size: 16px">组织</span>
		</div>
		<div class="center">
			<div class="department">
				<div class="department-search">
					<a-input-search placeholder="搜索部门" />
				</div>
				<a-tree
					:tree-data="treeData"
					show-icon
					default-expand-all
					:default-selected-keys="['0-0-0']"
				>
					<a-icon slot="switcherIcon" type="down" />
					<a-icon slot="smile" type="smile-o" />
					<a-icon slot="meh" type="smile-o" />
					<template slot="custom" slot-scope="{ selected }">
						<a-icon :type="selected ? 'frown' : 'frown-o'" />
					</template>
				</a-tree>
			</div>
			<div style="border-left: 1px solid #e0e0e0; width: 30px"></div>
			<div class="member-list">
				<div class="department-search">
					<a-input-search placeholder="搜索人员" />
				</div>
				<a-table :columns="columns" :data-source="data">
					<a slot="name" slot-scope="text">{{ text }}</a>
					<span slot="customTitle"><a-icon type="smile-o" /> Name</span>
					<span slot="tags" slot-scope="tags">
						<a-tag
							v-for="tag in tags"
							:key="tag"
							:color="
								tag === 'loser'
									? 'volcano'
									: tag.length > 5
									? 'geekblue'
									: 'green'
							"
						>
							{{ tag.toUpperCase() }}
						</a-tag>
					</span>
					<span slot="action" slot-scope="">
						<a>编辑</a>
						<a-divider type="vertical" />
						<a>启用</a>
						<a-divider type="vertical" />
						<a>删除</a>
						<a-divider type="vertical" />
					</span>
				</a-table>
			</div>
		</div>
	</div>
</template>
<script>
import { City } from "@icon-park/vue";
const columns = [
	{
		dataIndex: "name",
		title: "姓名",
		key: "name",
		slots: { title: "customTitle" },
		scopedSlots: { customRender: "name" },
	},
	{
		title: "部门",
		dataIndex: "age",
		key: "age",
	},
	{
		title: "职位",
		dataIndex: "address",
		key: "address",
	},
	{
		title: "操作",
		key: "action",
		scopedSlots: { customRender: "action" },
	},
];

const data = [
	{
		key: "1",
		name: "John Brown",
		age: 32,
		address: "New York No. 1 Lake Park",
		tags: ["nice", "developer"],
	},
	{
		key: "2",
		name: "Jim Green",
		age: 42,
		address: "London No. 1 Lake Park",
		tags: ["loser"],
	},
	{
		key: "3",
		name: "Joe Black",
		age: 32,
		address: "Sidney No. 1 Lake Park",
		tags: ["cool", "teacher"],
	},
];
const treeData = [
	{
		title: "我的公司",
		key: "0-0",
		slots: {
			icon: "smile",
		},
		children: [
			{ title: "leaf", key: "0-0-0", slots: { icon: "meh" } },
			{ title: "leaf", key: "0-0-1", scopedSlots: { icon: "custom" } },
		],
	},
];

export default {
	components: {
		City,
	},
	data() {
		return {
			treeData,
			data,
			columns,
		};
	},
	methods: {
		onSelect(selectedKeys, info) {
			console.log("selected", selectedKeys, info);
		},
		onCheck(checkedKeys, info) {
			console.log("onCheck", checkedKeys, info);
		},
	},
};
</script>
<style scoped>
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
.department {
	min-width: 250px;
}
.department-search {
	width: 200px;
	margin: 5px 5px;
}
.member-list {
	flex: 3;
}
.member-search {
	width: 200px;
	margin: 5px 5px;
}
</style>
