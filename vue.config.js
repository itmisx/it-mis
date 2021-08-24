const path = require("path");
module.exports = {
    publicPath: "./",
    pages: {
        index: {
            // page 的入口
            entry: "web_src/main.js",
            // 模板来源
            template: "web_src/public/index.html",
            // 在 dist/index.html 的输出
            filename: "index.html",
        },
    },
    chainWebpack: (config) => {
        config.resolve.alias
            .set("@", path.resolve(__dirname, "web_src"))
            .set("views", path.resolve(__dirname, "web_src/views"));
    },
};