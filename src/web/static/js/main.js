function fullScreen() {
    var el = document.documentElement;
    var rfs = el.requestFullScreen || el.webkitRequestFullScreen || el.mozRequestFullScreen || el.msRequestFullScreen;

    //typeof rfs != "undefined" && rfs
    if (rfs) {
        rfs.call(el);
    } else if (typeof window.ActiveXObject !== "undefined") {
        //for IE，这里其实就是模拟了按下键盘的F11，使浏览器全屏
        var wscript = new ActiveXObject("WScript.Shell");
        if (wscript != null) {
            wscript.SendKeys("{F11}");
        }
    }
}

function lastPage() {
    var pageNo = document.getElementById("pageNo").value
    pageNo = parseInt(pageNo)
    if (pageNo > 1) {
        pageNo = pageNo - 1
    }
    document.getElementById("pageNo").value = pageNo
    document.getElementById("search-form").submit()
}

function nextPage() {
    var pageNo = document.getElementById("pageNo").value
    var totalPage = document.getElementById("totalPage").value
    pageNo = parseInt(pageNo)
    totalPage = parseInt(totalPage)
    if (pageNo < totalPage) {
        pageNo = pageNo + 1
    }
    document.getElementById("pageNo").value = pageNo
    document.getElementById("search-form").submit()
}

function choosePage(pageNo) {
    document.getElementById("pageNo").value = pageNo
    document.getElementById("search-form").submit()
}

function openAjax(path) {
    console.log(path)
    $.ajax({
        type: "POST",
        url: "/play",
        data: { "id": path }
    });
}


function deleteWarring(id) {
    $("#deleteId").val(id)
    $("#deleteContext").text(id)
    $('#warningLabel').text("提示：删除不可恢复")
    $('#warning').modal('show')
}

function deleteAjax(id) {
    if (!id) {
        id = document.getElementById("deleteId").value
    }
    $.ajax({
        type: "POST",
        url: "/delete",
        data: { "id": id },
        success(data) {
            if (data.Code == 200) {
                success(data.Message)
            } else {
                fail(data.Message)
            }
        }
    });
    $('#warning').modal('hide')
}

function clickDirAjax(id) {
    $.ajax({
        type: "POST",
        url: "/removedir",
        data: { "id": id },
        success(data) {
            if (data.Code == 200) {
                var mess = "执行成功，请更新索引"
                success(mess)
                location.reload()
            } else {
                fail(data.Message)
            }
        }
    });
}


function openDirAjax(path) {
    console.log(path)
    $.ajax({
        type: "POST",
        url: "/opendir",
        data: { "id": path }
    });
}


function openCode(code) {
    var keyWord
    if (code) {
        keyWord = code
    } else {
        $('#detail').modal('hide')
        keyWord = document.getElementById("fcode").innerText
    }
    var url = "https://www.cdnbus.in/" + keyWord
    window.open(url, "_blank")
}

function searchActress(code) {
    var keyWord
    if (code) {
        keyWord = code
    } else {
        $('#detail').modal('hide')
        keyWord = document.getElementById("factress").innerText
    }
    var url = "https://www.cdnbus.in/search/" + keyWord
    window.open(url, "_blank")
}

function openActress(actress) {
    var keyWord
    if (actress) {
        keyWord = actress
    } else {
        $('#detail').modal('hide')
        keyWord = document.getElementById("factress").innerText
    }
    document.getElementById("keyWord").value = keyWord
    window.open("/views?keyWord=" + keyWord)
}

function clickActress(actress) {
    var keyWord = actress
    document.getElementById("keyWord").value = keyWord
    document.getElementById("search-form").submit()
}

function openModal(id) {
    var file;
    $.ajax({
        type: "POST",
        url: "/info",
        data: { "id": id },
        async: false,
        success(data) {
            file = data
        },
        error() {

        }
    });
    $('#fcode').text(file.Code)
    $('#factress').text(file.Actress)
    $('#fmtime').text(file.MTime)
    $('#fsize').text(file.SizeStr)
    $('#myModalLabel').text(file.Name)
    var image = file.Jpg
    if (!image) {
        image = file.Png
    }
    $('#myModalImg').attr("src", "data:image/png;base64," + image)
    $('#detail').modal('show')
}


function addDir() {
    var file = document.getElementById("addDirText").value
    $.ajax({
        type: "POST",
        url: "/adddir",
        data: { "id": file },
        success(data) {
            if (data.Code == 200) {
                $('#addDirModal').modal('hide')
                success(data.Message)
            } else {
                $('#addDirModal').modal('hide')
                fail(data.Message)
            }
        }
    });

}
function locationViews() {
    document.getElementById("keyWord").value = ""
    window.location = "/views"
}

function refreshIndex() {

    $.ajax({
        type: "GET",
        url: "/fresh",
        async: false,
        success(data) {
            if (data.Code == 200) {
                var href = window.location.href
                window.location = href
            } else {
                fail(data.Message)
            }
        }
    });

}

function syncAv(id) {
    $.ajax({
        type: "Post",
        url: "/sync",
        data: { "id": id },
        success(data) {
            console.log(data)
            if (data.Code == 200) {
                success(data.Message)
            } else {
                fail(data.Message)
            }
        }
    });
}

function success(msg) {
    selfAlert(msg, "alert-success")
}

function fail(msg) {
    selfAlert(msg, "alert-danger")
}

function selfAlert(msg, clazz) {
    var nodeId = new Date().getTime()
    var div = document.createElement("div")
    div.setAttribute("id", nodeId)
    var html = "<div  class=\"msg alert " + clazz + " alert-dismissable\">\n" +
        "    <button type=\"button\" class=\"close\" data-dismiss=\"alert\"\n" +
        "            aria-hidden=\"true\">\n" +
        "        &times;\n" +
        "    </button>\n" +
        msg + "   \n" +
        "</div>";
    document.getElementById("msg").append(div)
    document.getElementById(nodeId).innerHTML = html
    setTimeout(function () {
        document.getElementById(nodeId).remove()
    }, "4000");
}

