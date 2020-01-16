function openAjax(path) {
    console.log(path)
    $.ajax({
        type: "POST",
        url: "/play",
        data: { "id": path }
    });
}

function removeDirAjax(path) {
    console.log(path)
    $.ajax({
        type: "POST",
        url: "/play",
        data: { "id": path }
    });
}