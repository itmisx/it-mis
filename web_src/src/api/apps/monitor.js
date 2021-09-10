import request from "@/utils/request.js";

// 监听任务列表
export function MonitorTaskList(data) {
    return request({
        url: "/apps/monitor/task-list",
        method: "get",
        data,
        loadingID: "monitor_task_list",
        loadingTip: "保存中",
    });
}

// 新增监听
export function addMonitorTask(data) {
    return request({
        url: "/apps/monitor/add-task",
        method: "post",
        data,
        loadingID: "add_monitor_task",
        loadingTip: "保存中",
    });
}