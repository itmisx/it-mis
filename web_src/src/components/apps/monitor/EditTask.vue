<template>
	<spin-x loadingID="add_monitor_task">
		<a-modal
			:title="record ? `编辑监听` : `新增监听`"
			visible
			@ok="onSubmit"
			ok-text="确定"
			cancel-text="取消"
			@cancel="handleCancel"
		>
			<a-form-model
				ref="ruleForm"
				:model="form"
				:label-col="{ span: 4 }"
				:wrapper-col="{ span: 18 }"
				:rules="rules"
			>
				<a-form-model-item label="名称" prop="name">
					<a-input v-model="form.name" placeholder="请输入监控名称" />
				</a-form-model-item>
				<a-form-model-item label="类型" prop="monitor_type">
					<a-select v-model="form.monitor_type" placeholder="请输入监控类型">
						<a-select-option :value="1"> 站点监控 </a-select-option>
						<a-select-option :value="2"> 端口监控 </a-select-option>
					</a-select>
				</a-form-model-item>
				<a-form-model-item label="地址" prop="host">
					<a-input v-model="form.host" placeholder="请输入监控地址(域名/IP)" />
				</a-form-model-item>
				<a-form-model-item label="端口" prop="port">
					<a-input v-model="form.port" placeholder="请输入端口" />
				</a-form-model-item>
			</a-form-model>
		</a-modal>
	</spin-x>
</template>
<script>
import { addMonitorTask } from "@/api/apps/monitor.js";
export default {
	props: {
		record: {
			type: Object,
			default: null,
		},
	},
	data() {
		return {
			form: {
				name: "",
				monitor_type: 1,
				host: "",
				port: "",
			},
			rules: {
				name: {
					required: true,
					message: "请输入监控名称",
					trigger: "blur",
				},
				monitor_type: {
					required: true,
					message: "请选择监控类型",
					trigger: "blur",
				},
				host: {
					required: true,
					message: "请输入监控地址",
					trigger: "blur",
				},
				port: {
					required: true,
					message: "请输入监控端口",
					trigger: "blur",
				},
			},
		};
	},
	methods: {
		// 提交信息
		onSubmit() {
			this.$refs.ruleForm.validate((valid) => {
				if (valid) {
					addMonitorTask(this.form).then((res) => {
						if (res.code === 0) {
							this.$message.success("操作成功");
							this.$emit("hideEditTask");
						} else {
							this.$message.error(res.msg);
						}
					});
				} else {
					this.$message.error("输入有误，请检查！");
				}
			});
		},
		// 取消，关闭窗口
		handleCancel() {
			this.$emit("hideEditTask");
		},
		// 重置表单
		resetForm() {
			this.$refs.ruleForm.resetFields();
		},
	},
};
</script>

<style lang="scss" scoped></style>
