layui.use("layer", function () {
});

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
        type: "POST",
        dataType: "json",
        data: data,
        success: function (json) {
            if (json.status == 0) {
                layui.layer.msg(json.message);
            } else {
                layui.layer.msg(json.message);
                setTimeout(function () {
                    location.href = "/";
                }, 500);
            }
        }
    });
}

// function doing(id){
//     $.ajax({
//         type: "POST",
//         dataType: "json",
//         data: data,
//         success: function (json) {
//             if (json.status == 0) {
//                 layui.layer.msg(json.message);
//             } else {
//                 layui.layer.msg(json.message);
//                 setTimeout(function () {
//                     location.href = "/";
//                 }, 500);
//             }
//         }
//     });
// }