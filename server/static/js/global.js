layui.use("layer", function () {});

function addTask() {

    var name = $(".add-content").val();
    if (name == "") {
        alert("请输入要添加的事务名");
        return;
    }
    var data = {
        "name": name
    };
    $.ajax({
        url: "add",
        type: "POST",
        dataType: "json",
        data: data,
        success: function (json) {
            if (json.status == 0) {
                layui.layer.msg(json.message);
            } else {
                location.href = "/";
            }
        }
    });
}

function changing(id, method) {
    var status = 99;
    if (method == "Finish") {
        status = 1;
    } else if (method == "Being") {
        status = 0;
    }
    if (status == 99) {
        layer.msg("无法更改");
        return;
    }
    var data = {
        "id": id,
        "status": status
    };
    $.ajax({
        url: "update",
        type: "POST",
        dataType: "json",
        data: data,
        success: function (json) {
            if (json.status == 0) {
                layui.layer.msg(json.message);
            } else {
                location.href = "/";
            }
        }
    });
}

function deleteTask(id) {
    var data = {
        "id": id
    }
    $.ajax({
        url: "delete",
        type: "POST",
        dataType: "json",
        data: data,
        success: function (json) {
            if (json.status == 0) {
                layui.layer.msg(json.message);
            } else {
                location.href = "/";
            }
        }
    });
}