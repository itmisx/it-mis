const path = require("path");
module.exports = {
    publicPath: "./",
    productionSourceMap: false,
    chainWebpack: (config) => {
        config.resolve.alias
            .set("@", path.resolve(__dirname, "src"))
            .set("@views", path.resolve(__dirname, "src/views"))
            .set("@static", "public/static");
    },
};