// ajax 拦截
import { hook } from "ajax-hook";
import parseUrl from "./parse_url.js";
hook({
    open: function(arg) {
        if (URLFilter(arg[1])) {
            return true;
        }
    },
});
// fetch 拦截
var _f = fetch;
window.fetch = function() {
    if (URLFilter(arguments[0])) return {};
    return _f.apply(window, arguments);
};

// url拦截过滤
const URLFilter = function(url) {
    if (parseUrl(url).path === "/index/csrf-token") {
        return true;
    }
    return false;
};