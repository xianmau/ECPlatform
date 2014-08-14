//
$(function () {
    SetNavActive();
})


// 前台导航条选中效果
function SetNavActive() {
    var url = window.location.pathname;
    var p = url.split("/")
    for (var i = 0; i < p.length; i++) {
        if (p[i].length == 0) p.splice(i, 1);
    }
    if (p.length == 1) {
        $('#nav .item').eq(0).addClass("active");
    } else if (p.length > 1) {
        if (p[1] == "home") {
            $('#nav .item').eq(0).addClass("active");
        } else if (p[1] == "gcat") {
            var c = GetQueryString("c");
            if(c == "优质粮油" || decodeURI(c) == "优质粮油"){
                $('#nav .item').eq(1).addClass("active");
            }else if(c == "农家干货" || decodeURI(c) == "农家干货"){
                $('#nav .item').eq(2).addClass("active");
            }else if(c == "新鲜果蔬" || decodeURI(c) == "新鲜果蔬"){
                $('#nav .item').eq(3).addClass("active");
            }else if(c == "华农出品" || decodeURI(c) == "华农出品"){
                $('#nav .item').eq(4).addClass("active");
            }
        } else if (p[1] == "acat") {
            var c = GetQueryString("c");
            if(c == "健康食谱" || decodeURI(c) == "健康食谱"){
                $('#nav .item').eq(5).addClass("active");
            }
        } else if (p[1] == "message") {
            $('#nav .item').eq(6).addClass("active");
        }
    }
}



///////////////////////////////////////////

function GetQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)","i");
    var r = window.location.search.substr(1).match(reg);
    if (r!=null) return (r[2]); return null;
}