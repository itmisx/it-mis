export default function parseUrl(url) {
    let urlObj = {
        protocol: new RegExp("^(.+)\\:\\/\\/"),
        host: new RegExp("\\:\\/\\/(.+?)[\\?\\#\\s\\/]"),
        path: new RegExp("\\w(\\/.*?)[\\?\\#\\s]"),
        query: new RegExp("\\?(.+?)[\\#\\/\\s]"),
        hash: new RegExp("\\#(\\w+)\\s$"),
    };
    url += " ";

    function formatQuery(str) {
        return str.split("&").reduce((a, b) => {
            let arr = b.split("=");
            a[arr[0]] = arr[1];
            return a;
        }, {});
    }
    for (let key in urlObj) {
        let pattern = urlObj[key];
        urlObj[key] =
            key === "query" ?
            pattern.exec(url) && formatQuery(pattern.exec(url)[1]) :
            pattern.exec(url) && pattern.exec(url)[1];
    }
    return urlObj;
}