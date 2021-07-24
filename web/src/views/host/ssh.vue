<template>
  <div id="xterm-container"></div>
</template>

<script>
import "xterm/css/xterm.css";
import { Terminal } from "xterm";

export default {
  data: function () {
    return {
      cmd: "",
      term: {},
    };
  },
  mounted() {
    this.init();
    this.term = new Terminal({
      rendererType: "canvas", //渲染类型
      rows: 40, //行数
      fontSize: 16,
      convertEol: true, //启用时，光标将设置为下一行的开头
      scrollback: 10, //终端中的回滚量
      disableStdin: false, //是否应禁用输入。
      cursorStyle: "bar", //光标样式
      cursorBlink: true, //光标闪烁
      // theme: {
      //   foreground: "yellow", //字体
      //   background: "#060101", //背景色
      //   cursor: "help", //设置光标
      // },
    });
    this.term.open(document.getElementById("xterm-container"));
    this.term.write("\r\n~$ ");
    this.term.focus();
    this.term.onKey((keyEvent) => {
      let key = keyEvent.key;
      this.send(key);
    });
  },
  methods: {
    init: function () {
      if (typeof WebSocket === "undefined") {
        alert("您的浏览器不支持socket");
      } else {
        // 实例化socket
        this.socket = new WebSocket("ws://localhost:8080/host/console");
        // 监听socket连接
        this.socket.onopen = this.open;
        // 监听socket错误信息
        this.socket.onerror = this.error;
        // 监听socket消息
        this.socket.onmessage = this.getMessage;
      }
    },
    open: function () {
      console.log("socket连接成功");
    },
    error: function () {
      console.log("连接错误");
    },
    getMessage: function (msg) {
      console.log(msg);
      let msgData = msg.data;
      msgData = msgData.replaceAll("\n", "\r\n");
      this.term.write(msgData);
    },
    send: function (params) {
      this.socket.send(params);
    },
    close: function () {
      console.log("socket已经关闭");
    },
  },
};
</script>

<style lang="scss" scoped></style>
