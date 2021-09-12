<template>
	<div class="left">
		<div class="title">IT-MIS</div>
		<div class="nav">
			<a-menu
				style="width: 130px"
				:selectedKeys="navSelected"
				mode="inline"
				theme="dark"
				@select="navChange"
			>
				<a-menu-item key="apps">
					<a-icon type="appstore" />
					<span>应用</span>
				</a-menu-item>
				<a-menu-item key="org">
					<a-icon type="team" />
					<span>组织</span>
				</a-menu-item>
				<a-menu-item key="system">
					<a-icon type="setting" />
					<span>设置</span>
				</a-menu-item>
				<a-menu-item key="message">
					<a-icon type="message" two-tone-color="#eb2f96" />
					<a-badge dot><span>消息</span></a-badge>
				</a-menu-item>
				<a-menu-item key="store">
					<a-icon type="shopping" two-tone-color="#eb2f96" />
					<span>商店</span>
				</a-menu-item>
			</a-menu>
		</div>
		<div class="left-bottom">
			<logout
				theme="outline"
				size="24"
				fill="#1890ff"
				style="cursor: pointer"
				@click.native="logout"
			/>
			<span style="margin-top: 10px">©IT-MIS</span>
		</div>
	</div>
</template>

<script>
import { Logout } from "@icon-park/vue";
import { logoutAPI } from "@/api/index.js";
export default {
	data: function () {
		return {
			navSelected: [],
		};
	},
	components: {
		Logout,
	},
	beforeMount: function () {
		// 获取上次导航菜单
		let lastNavSelected = localStorage.getItem("lastest_nav_selected");
		if (lastNavSelected) {
			this.navSelected = [lastNavSelected];
		} else {
			// 默认值
			this.navSelected = ["apps"];
		}
	},
	methods: {
		navChange(item) {
			localStorage.setItem("lastest_nav_selected", item.key);
			this.navSelected = [item.key];
			switch (item.key) {
				case "apps":
					this.$router.push("/home/apps");
					break;
				case "org":
					this.$router.push("/home/org");
					break;
				case "system":
					this.$router.push("/home/system");
					break;
				case "store":
					this.$router.push("/home/store");
					break;
			}
		},
		logout() {
			logoutAPI().then((res) => {
				if (res.code === 0) {
					this.$router.push("/");
				}
			});
		},
	},
};
</script>

<style lang="scss" scoped>
.left {
	width: 130px;
	height: 100%;
	background-color: rgb(0, 21, 41);
	display: flex;
	display: -webkit-flex;
	flex-direction: column;
	align-items: center;
	.title {
		height: 60px;
		width: 100%;
		line-height: 60px;
		text-align: center;
		font-size: 28px;
		font-weight: 700;
		color: #1890ff;
	}
	.nav {
		flex: 1;
		display: flex;
		display: -webkit-flex;
		flex-direction: column;
		align-items: center;
		.nav-item {
			color: white;
			font-size: 32px;
			margin-bottom: 20px;
			background-color: teal;
		}
	}
	.left-bottom {
		display: flex;
		display: -webkit-flex;
		flex-direction: column;
		align-items: center;
	}
}
</style>
