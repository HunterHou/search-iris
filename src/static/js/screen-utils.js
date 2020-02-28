//进行切换
var fullScreenClickCount = 0;
//调用各个浏览器提供的全屏方法
var handleFullScreen = function () {
    var de = document.documentElement;

    if (de.requestFullscreen) {
        de.requestFullscreen();
    } else if (de.mozRequestFullScreen) {
        de.mozRequestFullScreen();
    } else if (de.webkitRequestFullScreen) {
        de.webkitRequestFullScreen();
    } else if (de.msRequestFullscreen) {
        de.msRequestFullscreen();
    }
    else {
        wtx.info("当前浏览器不支持全屏！");
    }

};
//调用各个浏览器提供的退出全屏方法
var exitFullscreen = function () {
    if (document.exitFullscreen) {
        document.exitFullscreen();
    } else if (document.mozCancelFullScreen) {
        document.mozCancelFullScreen();
    } else if (document.webkitExitFullscreen) {
        document.webkitExitFullscreen();
    }
}
//获取.fullscreen样式（选择器）的点击事件，再该事件中调用已封装好的方法

$(".fullscreen").click(function () {
    if (fullScreenClickCount % 2 == 0) {
        handleFullScreen();
    } else {
        exitFullscreen();
    }
    fullScreenClickCount++;
});