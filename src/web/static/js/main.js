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
        data: {"id": path}
    });
}

function deleteAjax(id) {
    $.ajax({
        type: "POST",
        url: "/delete",
        data: {"id": id},
        success(data) {
            if (data.Code == 200) {
                success(data.Message)
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
        data: {"id": path}
    });
}

function clickCode(code) {
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

function clickAcress(actress) {
    var keyWord
    if (actress) {
        keyWord = actress
    } else {
        $('#detail').modal('hide')
        keyWord = document.getElementById("factress").innerText
    }

    document.getElementById("keyWord").value = keyWord
    document.getElementById("search-form").submit()
}

function openModal(id) {
    var file;
    $.ajax({
        type: "POST",
        url: "/info",
        data: {"id": id},
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
    $('#myModalImg').attr("src", "data:image/png;base64," + file.Jpg)
    $('#detail').modal('show')
}

function removeDirAjax(path) {
    console.log(path)
    $.ajax({
        type: "POST",
        url: "/play",
        data: {"id": path}
    });
}

function refresh() {
    $.ajax({
        type: "GET",
        url: "/fresh",
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

function syncAv(id) {
    $.ajax({
        type: "Post",
        url: "/sync",
        data: {"id": id},
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

    var html = "<div class=\"alert alert-success alert-dismissable\">\n" +
        "    <button type=\"button\" class=\"close\" data-dismiss=\"alert\"\n" +
        "            aria-hidden=\"true\">\n" +
        "        &times;\n" +
        "    </button>\n" +
        msg + "   \n" +
        "</div>";
    document.getElementById("msg").innerHTML = html

    setTimeout(function () {
        document.getElementById("msg").innerHTML = "";
    }, "2000");
}

function fail(msg) {
    var html = "<div class=\"alert alert-danger alert-dismissable\">\n" +
        "    <button type=\"button\" class=\"close\" data-dismiss=\"alert\"\n" +
        "            aria-hidden=\"true\">\n" +
        "        &times;\n" +
        "    </button>\n" +
        msg + "   \n" +
        "</div>";
    document.getElementById("msg").innerHTML = html

    setTimeout(function () {
        document.getElementById("msg").innerHTML = "";
    }, "2000");
}