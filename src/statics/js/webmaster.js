 // 加载全局脚本
 $(function() {
   SetLayout();
   $(window).resize(function() {
     SetLayout();
   });
   SetNavActive();
   LoadAvatar();
 });

 // 按需加载脚本
 $(function() {
   var url = window.location.pathname;
   var p = url.split("/")
   for (var i = 0; i < p.length; i++) {
     if (p[i].length == 0) p.splice(i, 1);
   }
   // TODO
   
 });

 /*----------------------------------
  * Layout scripts
  *----------------------------------*/
 function SetLayout() {
   var wh = $(window).height();
   var ww = $(window).width();
   $('#leftside').height(wh);
   $('#rightside').height(wh);
   $('#rightside').width(ww - 280);
 }

 function LoadAvatar() {
   var account_email = $("#leftside .avatar").attr("title");
   $("#leftside .avatar").css("background-image", "url(http://www.gravatar.com/avatar/" + MD5(account_email) + ")");
 }

 function SetNavActive() {
   var url = window.location.pathname;
   var p = url.split("/")
   for (var i = 0; i < p.length; i++) {
     if (p[i].length == 0) p.splice(i, 1);
   }
   if (p.length == 1) {
     $('.nav li a').eq(0).addClass("active");
   } else if (p.length > 1) {
     if (p[1] == "dashboard") {
       $('.nav li a').eq(0).addClass("active");
     } else if (p[1] == "goods" || p[1] == "goodscategory") {
       $('.nav li a').eq(1).addClass("active");
     } else if (p[1] == "article" || p[1] == "articlecategory") {
       $('.nav li a').eq(2).addClass("active");
     } else if (p[1] == "link" || p[1] == "linkcategory") {
       $('.nav li a').eq(3).addClass("active");
     } else if (p[1] == "user") {
       $('.nav li a').eq(4).addClass("active");
     } else if (p[1] == "setting") {
       $('.nav li a').eq(5).addClass("active");
     } else if (p[1] == "system") {
       $('.nav li a').eq(6).addClass("active");
     }
   }
 }


 /*----------------------------------
  * Dashboard scripts
  *----------------------------------*/