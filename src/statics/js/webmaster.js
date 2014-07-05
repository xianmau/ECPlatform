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
   if (p.length == 2) {
     if (p[1] == "zktree") {
       ZonesSetDragDiv()
       ZonesLoadZkTree();
     }
   } else if (p.length == 3) {
     if (p[1] == "reports" && p[2] == "brokerdetail") {
       ReportsLoadBrokerChart()
     } else if (p[1] == "reports" && p[2] == "loggerdetail") {
       ReportsLoadLoggerChart()
     } else if (p[1] == "reports" && (p[2] == "brokers" || p[2] == "loggers")) {
       ReportsDealFloatData()
     } else if (p[1] == "reports" && p[2] == "topics") {
       ReportsDealSegment()
     }
   }

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
     } else if (p[1] == "zones" || p[1] == "zktree") {
       $('.nav li a').eq(2).addClass("active");
     } else if (p[1] == "accounts") {
       $('.nav li a').eq(3).addClass("active");
     } else if (p[1] == "settings") {
       $('.nav li a').eq(4).addClass("active");
     }
   }
 }

 function LoadAvatar() {
   var account_email = $("#leftside .avatar").attr("title");
   $("#leftside .avatar").css("background-image", "url(http://www.gravatar.com/avatar/" + MD5(account_email) + ")");
 }


 /*----------------------------------
  * Dashboard scripts
  *----------------------------------*/
 function DashboardExtractorRun() {
   $.ajax({
     type: 'POST',
     url: '/webmaster/dashboard/SetExtractor',
     data: {
       act: 'run'
     },
     success: function(data) {
       location.reload();
     },
     dataType: 'text'
   });
 }

 function DashboardExtractorStop() {
   if (!confirm("Are your sure?")) {
     return;
   }
   $.ajax({
     type: 'POST',
     url: '/webmaster/dashboard/SetExtractor',
     data: {
       act: 'stop'
     },
     success: function(data) {
       location.reload();
     },
     dataType: 'text'
   });
 }

 function DashboardDbTruncate(range) {
   if (!confirm("Are your sure?")) {
     return;
   }

   var sqlstr = "";
   if (range == "broker") {
     sqlstr = "truncate `tb_broker`"
     DashboardExecuteSql(sqlstr);
     sqlstr = "truncate `tb_broker_stat`"
     DashboardExecuteSql(sqlstr);
     sqlstr = "truncate `tb_latest_broker_stat`"
     DashboardExecuteSql(sqlstr);
     return
   } else if (range == "logger") {
     sqlstr = "truncate `tb_logger`"
     DashboardExecuteSql(sqlstr);
     sqlstr = "truncate `tb_logger_stat`"
     DashboardExecuteSql(sqlstr);
     sqlstr = "truncate `tb_latest_logger_stat`"
     DashboardExecuteSql(sqlstr);
     return
   } else if (range == "app") {
     sqlstr = "truncate `tb_app`"
   } else if (range == "topic") {
     sqlstr = "truncate `tb_topic`"
   } else if (range == "zone") {
     sqlstr = "truncate `tb_zone`"
   } else {
     alert("error")
     return
   }

   DashboardExecuteSql(sqlstr);
 }

 function DashboardDbDeleteHistory(range) {
   if (!confirm("Are your sure?")) {
     return;
   }

   var sqlstr = "";
   if (range == "all") {
     sqlstr = "truncate `tb_broker_stat`"
     DashboardExecuteSql(sqlstr);
     sqlstr = "truncate `tb_logger_stat`"
     DashboardExecuteSql(sqlstr);
     return
   } else if (range == "1") {
     sqlstr = "delete from `tb_broker_stat` where date(`Timestamp`) <= date_sub(curdate(), INTERVAL 1 DAY)"
     DashboardExecuteSql(sqlstr);
     sqlstr = "delete from `tb_logger_stat` where date(`Timestamp`) <= date_sub(curdate(), INTERVAL 1 DAY)"
     DashboardExecuteSql(sqlstr);
     return
   } else if (range == "7") {
     sqlstr = "delete from `tb_broker_stat` where date(`Timestamp`) <= date_sub(curdate(), INTERVAL 7 DAY)"
     DashboardExecuteSql(sqlstr);
     sqlstr = "delete from `tb_logger_stat` where date(`Timestamp`) <= date_sub(curdate(), INTERVAL 7 DAY)"
     DashboardExecuteSql(sqlstr);
     return
   } else if (range == "15") {
     sqlstr = "delete from `tb_broker_stat` where date(`Timestamp`) <= date_sub(curdate(), INTERVAL 15 DAY)"
     DashboardExecuteSql(sqlstr);
     sqlstr = "delete from `tb_logger_stat` where date(`Timestamp`) <= date_sub(curdate(), INTERVAL 15 DAY)"
     DashboardExecuteSql(sqlstr);
     return
   } else {
     alert("error")
     return
   }

   DashboardExecuteSql(sqlstr);
 }

 function DashboardExecSql() {
   if (!confirm("Are your sure?")) {
     return;
   }

   var sqlstr = $('#sqlstr').val();
   DashboardExecuteSql(sqlstr);
 }

 function DashboardExecuteSql(sqlstr) {
   $.ajax({
     type: 'POST',
     url: '/webmaster/dashboard/ExecuteSql',
     data: {
       sqlstr: sqlstr
     },
     success: function(data) {
       location.reload();
     },
     dataType: 'text'
   });
 }

 function DashboardResetSql() {
   $('#sqlstr').val("");
 }



 /*----------------------------------
  * Settings scripts
  *----------------------------------*/
 function SettingsShowCreateForm() {
   $('.create-form').toggle('slow');
 }

 function SettingsCancel() {
   $('#createzoneid').val("");
   $('#createhosts').val("");
   $('.create-form').toggle('slow');
 }

 function SettingsCreateZone() {
   var createzoneid = $('#createzoneid').val();
   var createhosts = $('#createhosts').val();

   $.ajax({
     type: 'POST',
     url: '/webmaster/settings/CreateZone',
     data: {
       createzoneid: createzoneid,
       createhosts: createhosts
     },
     success: function(data) {
       location.reload();
     },
     dataType: 'text'
   });
 }

 function SettingsEditZone(zoneid) {
   var edithosts = $('#edithosts_' + zoneid).val();

   $.ajax({
     type: 'POST',
     url: '/webmaster/settings/EditZone',
     data: {
       editzoneid: zoneid,
       edithosts: edithosts
     },
     success: function(data) {
       location.reload();
     },
     dataType: 'text'
   });
 }

 function SettingsDeleteZone(zoneid) {
   if (!confirm("Are your sure?")) {
     return;
   }
   $.ajax({
     type: 'POST',
     url: '/webmaster/settings/DeleteZone',
     data: {
       deletezoneid: zoneid
     },
     success: function(data) {
       location.reload();
     },
     dataType: 'text'
   });
 }

 /*----------------------------------
  * Accounts scripts
  *----------------------------------*/
 function AccountsShowCreateForm() {
   $('.create-form').toggle('slow');
 }

 function AccountsCancel() {
   $('#createname').val("");
   $('#createpassword').val("");
   $('#createremark').val("");
   $('.create-form').toggle('slow');
 }

 function AccountsCreate() {
   var createname = $('#createname').val();
   var createpassword = $('#createpassword').val();
   var createremark = $('#createremark').val();

   $.ajax({
     type: 'POST',
     url: '/webmaster/accounts/Create',
     data: {
       createname: createname,
       createpassword: createpassword,
       createremark: createremark
     },
     success: function(data) {
       location.reload();
     },
     dataType: 'text'
   });
 }

 function AccountsEdit(name) {
   var editpassword = $('#editpassword_' + name).val();
   var editremark = $('#editremark_' + name).val();

   $.ajax({
     type: 'POST',
     url: '/webmaster/accounts/Edit',
     data: {
       editname: name,
       editpassword: editpassword,
       editremark: editremark
     },
     success: function(data) {
       location.reload();
     },
     dataType: 'text'
   });
 }

 function AccountsDelete(name) {
   if (!confirm("Are your sure?")) {
     return;
   }
   $.ajax({
     type: 'POST',
     url: '/webmaster/accounts/Delete',
     data: {
       deletename: name
     },
     success: function(data) {
       location.reload();
     },
     dataType: 'text'
   });
 }


 /*----------------------------------
  * Zones scripts
  *----------------------------------*/

 function ZonesLoadCreateForm() {
   var curNode = $('#cur_node').val();
   //alert(curNode);
   $('.zktree-forms').slideUp('fast');
   $('#zktree-forms-update').hide();
   $('#zktree-forms-create').show();
   $('.zktree-forms').slideDown('slow');

   $('#zktree_forms_create_path').val(curNode)

 }

 function ZonesCreateNode() {
   var curZone = $('#cur_zone').val();
   var path = $('#zktree_forms_create_path').val();
   var name = $('#zktree_forms_create_name').val();
   var data = $('#zktree_forms_create_data').val();
   //alert(path + name + data);
   $.ajax({
     type: 'POST',
     url: '/webmaster/zones/CreateNode',
     data: {
       zone: curZone,
       path: path + "/" + name,
       data: data
     },
     success: function(data) {
       location.reload();
     },
     dataType: 'text'
   });
 }

 function ZonesLoadUpdateForm() {
   var curNode = $('#cur_node').val();
   //alert(curNode);
   $('.zktree-forms').slideUp('fast');
   $('#zktree-forms-create').hide();
   $('#zktree-forms-update').show();
   $('.zktree-forms').slideDown('slow');

   $('#zktree_forms_update_path').val(curNode)
 }

 function ZonesUpdateNode() {
   var curZone = $('#cur_zone').val();
   var path = $('#zktree_forms_update_path').val();
   var data = $('#zktree_forms_update_data').val();
   //alert(path + data);
   $.ajax({
     type: 'POST',
     url: '/webmaster/zones/UpdateNode',
     data: {
       zone: curZone,
       path: path,
       data: data
     },
     success: function(data) {
       location.reload();
     },
     dataType: 'text'
   });
 }


 function ZonesCancel() {
   $('.zktree-forms').slideUp('fast');
 }

 function ZonesDeleteNode() {
   var curZone = $('#cur_zone').val();
   var curNode = $('#cur_node').val();
   if (!confirm("Are you sure?")) {
     return;
   }
   //alert(curNode);
   $.ajax({
     type: 'POST',
     url: '/webmaster/zones/DeleteNode',
     data: {
       zone: curZone,
       path: curNode
     },
     success: function(data) {
       location.reload();
     },
     dataType: 'text'
   });
 }



 function ZonesSetDragDiv() {
   // 设置结点信息层可拖曳
   $(".zktree-right").drag({
     handler: $('.dragdiv-title'),
     opacity: 0.7
   });
 }

 function ZonesLoadZkTree() {
   $('#zktree_div')
     .on('init.jstree', function() {})
     .on('select_node.jstree', function(e, data) {
       var current_node = "/" + data.instance.get_path(data.selected[0], '/', false);
       $('.node_path').html(current_node);
       $.ajax({
         url: '/webmaster/zktree/getdata',
         data: {
           'zoneid': $('#cur_zone').val(),
           'znode': current_node
         },
         complete: function(XHR, status) {
           if (status == "parsererror") {
             var code = "";
             code += '<h4>PATH: </h4>';
             code += '<p class="znodeinfo-path">' + current_node + '</p>';
             code += '<h4>DATA: </h4>';
             code += '<p class="znodeinfo-rawdata">';
             code += XHR.responseText;
             code += '</p>';
             code += '';
             code += '';
             $('#cur_node').val(current_node)

             // form field
             $('#zktree_forms_create_path').val(current_node)
             $('#zktree_forms_update_path').val(current_node)

             $('.node_show').html(code);
             $(".nano").nanoScroller();
           }
           //alert(XHR.responseText + status);
         },
         success: function(data) {
           var code = "";
           var o = data;
           if (o == null) {
             o = [];
           }
           code += '<h4>PATH: </h4>';
           code += '<p class="znodeinfo-path">' + current_node + '</p>';
           code += '<h4>DATA: </h4>';
           code += '<p class="znodeinfo-data">';
           code += JSON.stringify(o, null, 4);
           code += '</p>';
           code += '';
           code += '';
           $('#cur_node').val(current_node)

           // form field
           $('#zktree_forms_create_path').val(current_node)
           $('#zktree_forms_update_path').val(current_node)

           $('.node_show').html(code);
           $(".nano").nanoScroller();
         },
         dataType: 'json'
       });
     })
     .jstree({
       'core': {
         'data': {
           'dataType': 'json',
           'url': function(node) {
             return node.id === '#' ?
               '/static/js/jstree.data.json' :
               '/webmaster/zktree/getchildren';
           },
           'data': function(node) {
             if (node.id == '#') {
               return;
             }
             var nd = $.jstree.reference('#zktree_div');
             return {
               'zoneid': $('#cur_zone').val(),
               'znode': nd.get_path(node, '/', false)
             };
           }
         }
       },
       // 其它一些参数设置
       'plugins': ["sort"]
     });
 }


 /*----------------------------------
  * reports scripts
  *----------------------------------*/
 function ReportsDealFloatData() {
   $('.float-data').each(function(index) {
     var data = $(this).text()
     if (data == 0) {
       $(this).html("--")
     } else {

       $(this).html(Math.round((data * 100)) + "%")
     }
   })
 }

 function ReportsDealSegment() {
   // 处理段数据
   $('.segments-data').each(function() {
     var data = $(this).text();
     var o = jQuery.parseJSON(data);
     data = "";
     for (var i = 0; i < o.length; i++) {
       var tmp = "<strong>" + o[i].last_confirm_entry + "</strong><br />";
       tmp += "<ul>";
       for (var j = 0; j < o[i].loggers.length; j++) {
         tmp += "<li>" + o[i].loggers[j] + "</li>";
       }
       tmp += "</ul>";
       data += tmp;
     }
     var s = '<span class="label label-info" data-toggle="popover" title data-html="true" data-content="' + data + '" data-original-title="Segments Details">Segments Details</span>';
     $(this).html(s);
   });
   $('.segments-data .label').popover({
     placement: 'left'
   })
   $('.segments-data .label').popover('hide');
   $('.segments-data .label').click(function() {
     $('.segments-data .label').not(this).popover('hide');
   });
 }

 function ReportsLoadBrokerChart() {
   Highcharts.setOptions({
     global: {
       useUTC: true
     }
   });

   // Create the chart
   $('#container').highcharts('StockChart', {
     chart: {
       events: {
         load: function() {
           // set up the updating of the chart each second
           var series0 = this.series[0];
           var series1 = this.series[1];
           var series2 = this.series[2];
           setInterval(function() {
             $.get('/webmaster/reports/brokerdetail/getlatestdata?zoneid=' + $('#zoneid').val() + '&brokerid=' + $('#brokerid').val(), function(data) {
               series0.addPoint(jQuery.parseJSON(data)[0], true, true);
               series1.addPoint(jQuery.parseJSON(data)[1], true, true);
               series2.addPoint(jQuery.parseJSON(data)[2], true, true);
             });
           }, 60000);
         }
       }
     },

     rangeSelector: {
       buttons: [{
         count: 30,
         type: 'minute',
         text: '30M'
       }, {
         count: 1,
         type: 'hours',
         text: '1H'
       }, {
         type: 'all',
         text: 'All'
       }],
       inputEnabled: false,
       selected: 0
     },

     title: {
       text: 'Performance Evaluation of Broker: ' + $('#brokerid').val()
     },

     subtitle: {
       text: 'from zone: ' + $('#zoneid').val()
     },

     yAxis: {
       title: {
         text: 'percentage (%)'
       }
     },

     exporting: {
       enabled: false
     },

     tooltip: {
       shared: true
     },

     legend: {
       enabled: true,
       align: 'right',
       backgroundColor: '#FCFFC5',
       borderColor: 'black',
       borderWidth: 0,
       layout: 'vertical',
       verticalAlign: 'top',
       y: 0,
       shadow: true,
       floating: true
     },

     series: [{
       name: 'Cpu Rate(%)',
       data: jQuery.parseJSON($('#cpuData').val())
     }, {
       name: 'Net Rate(%)',
       data: jQuery.parseJSON($('#netData').val())
     }, {
       name: 'Disk Rate(%)',
       data: jQuery.parseJSON($('#diskData').val())
     }]
   });
 }

 function ReportsLoadLoggerChart() {
   Highcharts.setOptions({
     global: {
       useUTC: true
     }
   });

   // Create the chart
   $('#container').highcharts('StockChart', {
     chart: {
       events: {
         load: function() {
           // set up the updating of the chart each second
           var series0 = this.series[0];
           var series1 = this.series[1];
           var series2 = this.series[2];
           setInterval(function() {
             $.get('/webmaster/reports/loggerdetail/getlatestdata?zoneid=' + $('#zoneid').val() + '&loggerid=' + $('#loggerid').val(), function(data) {
               series0.addPoint(jQuery.parseJSON(data)[0], true, true);
               series1.addPoint(jQuery.parseJSON(data)[1], true, true);
               series2.addPoint(jQuery.parseJSON(data)[2], true, true);
             });
           }, 60000);
         }
       }
     },

     rangeSelector: {
       buttons: [{
         count: 30,
         type: 'minute',
         text: '30M'
       }, {
         count: 1,
         type: 'hours',
         text: '1H'
       }, {
         type: 'all',
         text: 'All'
       }],
       inputEnabled: false,
       selected: 0
     },

     title: {
       text: 'Performance Evaluation of Logger: ' + $('#loggerid').val()
     },

     subtitle: {
       text: 'from zone: ' + $('#zoneid').val()
     },

     yAxis: {
       title: {
         text: 'percentage (%)'
       }
     },

     exporting: {
       enabled: false
     },

     tooltip: {
       shared: true
     },

     legend: {
       enabled: true,
       align: 'right',
       backgroundColor: '#FCFFC5',
       borderColor: 'black',
       borderWidth: 0,
       layout: 'vertical',
       verticalAlign: 'top',
       y: 0,
       shadow: true,
       floating: true
     },

     series: [{
       name: 'Cpu Rate(%)',
       data: jQuery.parseJSON($('#cpuData').val())
     }, {
       name: 'Net Rate(%)',
       data: jQuery.parseJSON($('#netData').val())
     }, {
       name: 'Disk Rate(%)',
       data: jQuery.parseJSON($('#diskData').val())
     }]
   });
 }


 var MD5 = function(string) {
   function RotateLeft(lValue, iShiftBits) {
     return (lValue << iShiftBits) | (lValue >>> (32 - iShiftBits));
   }

   function AddUnsigned(lX, lY) {
     var lX4, lY4, lX8, lY8, lResult;
     lX8 = (lX & 0x80000000);
     lY8 = (lY & 0x80000000);
     lX4 = (lX & 0x40000000);
     lY4 = (lY & 0x40000000);
     lResult = (lX & 0x3FFFFFFF) + (lY & 0x3FFFFFFF);
     if (lX4 & lY4) {
       return (lResult ^ 0x80000000 ^ lX8 ^ lY8);
     }
     if (lX4 | lY4) {
       if (lResult & 0x40000000) {
         return (lResult ^ 0xC0000000 ^ lX8 ^ lY8);
       } else {
         return (lResult ^ 0x40000000 ^ lX8 ^ lY8);
       }
     } else {
       return (lResult ^ lX8 ^ lY8);
     }
   }

   function F(x, y, z) {
     return (x & y) | ((~x) & z);
   }

   function G(x, y, z) {
     return (x & z) | (y & (~z));
   }

   function H(x, y, z) {
     return (x ^ y ^ z);
   }

   function I(x, y, z) {
     return (y ^ (x | (~z)));
   }

   function FF(a, b, c, d, x, s, ac) {
     a = AddUnsigned(a, AddUnsigned(AddUnsigned(F(b, c, d), x), ac));
     return AddUnsigned(RotateLeft(a, s), b);
   };

   function GG(a, b, c, d, x, s, ac) {
     a = AddUnsigned(a, AddUnsigned(AddUnsigned(G(b, c, d), x), ac));
     return AddUnsigned(RotateLeft(a, s), b);
   };

   function HH(a, b, c, d, x, s, ac) {
     a = AddUnsigned(a, AddUnsigned(AddUnsigned(H(b, c, d), x), ac));
     return AddUnsigned(RotateLeft(a, s), b);
   };

   function II(a, b, c, d, x, s, ac) {
     a = AddUnsigned(a, AddUnsigned(AddUnsigned(I(b, c, d), x), ac));
     return AddUnsigned(RotateLeft(a, s), b);
   };

   function ConvertToWordArray(string) {
     var lWordCount;
     var lMessageLength = string.length;
     var lNumberOfWords_temp1 = lMessageLength + 8;
     var lNumberOfWords_temp2 = (lNumberOfWords_temp1 - (lNumberOfWords_temp1 % 64)) / 64;
     var lNumberOfWords = (lNumberOfWords_temp2 + 1) * 16;
     var lWordArray = Array(lNumberOfWords - 1);
     var lBytePosition = 0;
     var lByteCount = 0;
     while (lByteCount < lMessageLength) {
       lWordCount = (lByteCount - (lByteCount % 4)) / 4;
       lBytePosition = (lByteCount % 4) * 8;
       lWordArray[lWordCount] = (lWordArray[lWordCount] | (string.charCodeAt(lByteCount) << lBytePosition));
       lByteCount++;
     }
     lWordCount = (lByteCount - (lByteCount % 4)) / 4;
     lBytePosition = (lByteCount % 4) * 8;
     lWordArray[lWordCount] = lWordArray[lWordCount] | (0x80 << lBytePosition);
     lWordArray[lNumberOfWords - 2] = lMessageLength << 3;
     lWordArray[lNumberOfWords - 1] = lMessageLength >>> 29;
     return lWordArray;
   };

   function WordToHex(lValue) {
     var WordToHexValue = "",
       WordToHexValue_temp = "",
       lByte, lCount;
     for (lCount = 0; lCount <= 3; lCount++) {
       lByte = (lValue >>> (lCount * 8)) & 255;
       WordToHexValue_temp = "0" + lByte.toString(16);
       WordToHexValue = WordToHexValue + WordToHexValue_temp.substr(WordToHexValue_temp.length - 2, 2);
     }
     return WordToHexValue;
   };

   function Utf8Encode(string) {
     string = string.replace(/\r\n/g, "\n");
     var utftext = "";
     for (var n = 0; n < string.length; n++) {
       var c = string.charCodeAt(n);
       if (c < 128) {
         utftext += String.fromCharCode(c);
       } else if ((c > 127) && (c < 2048)) {
         utftext += String.fromCharCode((c >> 6) | 192);
         utftext += String.fromCharCode((c & 63) | 128);
       } else {
         utftext += String.fromCharCode((c >> 12) | 224);
         utftext += String.fromCharCode(((c >> 6) & 63) | 128);
         utftext += String.fromCharCode((c & 63) | 128);
       }
     }
     return utftext;
   };
   var x = Array();
   var k, AA, BB, CC, DD, a, b, c, d;
   var S11 = 7,
     S12 = 12,
     S13 = 17,
     S14 = 22;
   var S21 = 5,
     S22 = 9,
     S23 = 14,
     S24 = 20;
   var S31 = 4,
     S32 = 11,
     S33 = 16,
     S34 = 23;
   var S41 = 6,
     S42 = 10,
     S43 = 15,
     S44 = 21;
   string = Utf8Encode(string);
   x = ConvertToWordArray(string);
   a = 0x67452301;
   b = 0xEFCDAB89;
   c = 0x98BADCFE;
   d = 0x10325476;
   for (k = 0; k < x.length; k += 16) {
     AA = a;
     BB = b;
     CC = c;
     DD = d;
     a = FF(a, b, c, d, x[k + 0], S11, 0xD76AA478);
     d = FF(d, a, b, c, x[k + 1], S12, 0xE8C7B756);
     c = FF(c, d, a, b, x[k + 2], S13, 0x242070DB);
     b = FF(b, c, d, a, x[k + 3], S14, 0xC1BDCEEE);
     a = FF(a, b, c, d, x[k + 4], S11, 0xF57C0FAF);
     d = FF(d, a, b, c, x[k + 5], S12, 0x4787C62A);
     c = FF(c, d, a, b, x[k + 6], S13, 0xA8304613);
     b = FF(b, c, d, a, x[k + 7], S14, 0xFD469501);
     a = FF(a, b, c, d, x[k + 8], S11, 0x698098D8);
     d = FF(d, a, b, c, x[k + 9], S12, 0x8B44F7AF);
     c = FF(c, d, a, b, x[k + 10], S13, 0xFFFF5BB1);
     b = FF(b, c, d, a, x[k + 11], S14, 0x895CD7BE);
     a = FF(a, b, c, d, x[k + 12], S11, 0x6B901122);
     d = FF(d, a, b, c, x[k + 13], S12, 0xFD987193);
     c = FF(c, d, a, b, x[k + 14], S13, 0xA679438E);
     b = FF(b, c, d, a, x[k + 15], S14, 0x49B40821);
     a = GG(a, b, c, d, x[k + 1], S21, 0xF61E2562);
     d = GG(d, a, b, c, x[k + 6], S22, 0xC040B340);
     c = GG(c, d, a, b, x[k + 11], S23, 0x265E5A51);
     b = GG(b, c, d, a, x[k + 0], S24, 0xE9B6C7AA);
     a = GG(a, b, c, d, x[k + 5], S21, 0xD62F105D);
     d = GG(d, a, b, c, x[k + 10], S22, 0x2441453);
     c = GG(c, d, a, b, x[k + 15], S23, 0xD8A1E681);
     b = GG(b, c, d, a, x[k + 4], S24, 0xE7D3FBC8);
     a = GG(a, b, c, d, x[k + 9], S21, 0x21E1CDE6);
     d = GG(d, a, b, c, x[k + 14], S22, 0xC33707D6);
     c = GG(c, d, a, b, x[k + 3], S23, 0xF4D50D87);
     b = GG(b, c, d, a, x[k + 8], S24, 0x455A14ED);
     a = GG(a, b, c, d, x[k + 13], S21, 0xA9E3E905);
     d = GG(d, a, b, c, x[k + 2], S22, 0xFCEFA3F8);
     c = GG(c, d, a, b, x[k + 7], S23, 0x676F02D9);
     b = GG(b, c, d, a, x[k + 12], S24, 0x8D2A4C8A);
     a = HH(a, b, c, d, x[k + 5], S31, 0xFFFA3942);
     d = HH(d, a, b, c, x[k + 8], S32, 0x8771F681);
     c = HH(c, d, a, b, x[k + 11], S33, 0x6D9D6122);
     b = HH(b, c, d, a, x[k + 14], S34, 0xFDE5380C);
     a = HH(a, b, c, d, x[k + 1], S31, 0xA4BEEA44);
     d = HH(d, a, b, c, x[k + 4], S32, 0x4BDECFA9);
     c = HH(c, d, a, b, x[k + 7], S33, 0xF6BB4B60);
     b = HH(b, c, d, a, x[k + 10], S34, 0xBEBFBC70);
     a = HH(a, b, c, d, x[k + 13], S31, 0x289B7EC6);
     d = HH(d, a, b, c, x[k + 0], S32, 0xEAA127FA);
     c = HH(c, d, a, b, x[k + 3], S33, 0xD4EF3085);
     b = HH(b, c, d, a, x[k + 6], S34, 0x4881D05);
     a = HH(a, b, c, d, x[k + 9], S31, 0xD9D4D039);
     d = HH(d, a, b, c, x[k + 12], S32, 0xE6DB99E5);
     c = HH(c, d, a, b, x[k + 15], S33, 0x1FA27CF8);
     b = HH(b, c, d, a, x[k + 2], S34, 0xC4AC5665);
     a = II(a, b, c, d, x[k + 0], S41, 0xF4292244);
     d = II(d, a, b, c, x[k + 7], S42, 0x432AFF97);
     c = II(c, d, a, b, x[k + 14], S43, 0xAB9423A7);
     b = II(b, c, d, a, x[k + 5], S44, 0xFC93A039);
     a = II(a, b, c, d, x[k + 12], S41, 0x655B59C3);
     d = II(d, a, b, c, x[k + 3], S42, 0x8F0CCC92);
     c = II(c, d, a, b, x[k + 10], S43, 0xFFEFF47D);
     b = II(b, c, d, a, x[k + 1], S44, 0x85845DD1);
     a = II(a, b, c, d, x[k + 8], S41, 0x6FA87E4F);
     d = II(d, a, b, c, x[k + 15], S42, 0xFE2CE6E0);
     c = II(c, d, a, b, x[k + 6], S43, 0xA3014314);
     b = II(b, c, d, a, x[k + 13], S44, 0x4E0811A1);
     a = II(a, b, c, d, x[k + 4], S41, 0xF7537E82);
     d = II(d, a, b, c, x[k + 11], S42, 0xBD3AF235);
     c = II(c, d, a, b, x[k + 2], S43, 0x2AD7D2BB);
     b = II(b, c, d, a, x[k + 9], S44, 0xEB86D391);
     a = AddUnsigned(a, AA);
     b = AddUnsigned(b, BB);
     c = AddUnsigned(c, CC);
     d = AddUnsigned(d, DD);
   }
   var temp = WordToHex(a) + WordToHex(b) + WordToHex(c) + WordToHex(d);
   return temp.toLowerCase();
 }