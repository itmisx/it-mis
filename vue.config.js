const path = require("path");
module.exports = {
    publicPath: "./",
    pages: {
        index: {
            entry: "web_src/main.js",
            template: "web_public/index.html",
            filename: "index.html",
            title: "itmis",
        },
    },
    chainWebpack: (config) => {
        config.resolve.alias
            .set("@", path.resolve(__dirname, "web_src"))
            .set("views", path.resolve(__dirname, "web_src/views"));
    },
};