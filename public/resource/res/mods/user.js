/**

 @Name: 用户模块

 */
 
layui.define(['laypage', 'fly', 'element', 'flow', 'qiniuyun'], function(exports){

  var $ = layui.jquery;
  var layer = layui.layer;
  var util = layui.util;
  var laytpl = layui.laytpl;
  var form = layui.form;
  var laypage = layui.laypage;
  var fly = layui.fly;
  var flow = layui.flow;
  var element = layui.element;
  var upload = layui.upload;
  var qiniuyun = layui.qiniuyun;

  var gather = {}, dom = {
    mine: $('#LAY_mine')
    ,mineview: $('.mine-view')
    ,minemsg: $('#LAY_minemsg')
    ,infobtn: $('#LAY_btninfo')
  };

 
  if ( (layui.cache.page != 'user') && (!layui.cache.user || layui.cache.user.uid == -1 || !layui.cache.user.token || layui.cache.user.token.length == 0) ){ 
    //检查登录状态
    //location.href = "/";
    return;
  } 

  parseToHTML = function(res, me, uid, currPage, pageSize, contentId, contentpageId, joinOrNot){
    var header = '<ul class="mine-view jie-row">';
    var end = '</ul>';
    var otherheader = ' <ul class="jie-row">';
    var othernocontent = '<div style="min-height: 50px;  height:auto;"><i style="font-size:14px;">没有与之相关的任何主题</i></div> ';
    var content ;
    if (me){
      content = header;
    }else{
      content = otherheader;
    }
    if (!res || !res['data'] || !res['data']['list'] ||res['data']['total'] == 0 || res['data']['list']==null|| res['data']['list'].length == 0){
      content+= othernocontent;
      content += end;
      $(contentId).html(content);
      return;
    }

    for(let i = 0; i< res['data']['list'].length; i++){
      var html = ['<li>',
        '<a class="jie-title" href="/jie/detail.html?tid=TOPIC_ID&uid=USER_ID">TOPIC_TITLE</a>',
        '<i>TOPIC_CREATE_TIME</i>',
        '<a class="mine-edit" href="/jie/add.html?tid=TOPIC_ID">编辑</a>',
        '<em>SEE_NUMBER阅/MEMEBER_NUMBERS入</em>',
        '</li> '].join("");
      var htmlother= ['      <li>',
            '<a href="/jie/detail.html?tid=TOPIC_ID&uid=USER_ID" class="jie-title">TOPIC_TITLE</a>',
            '<i>TOPIC_CREATE_TIME</i>',
            '<em class="layui-hide-xs">SEE_NUMBER阅/MEMEBER_NUMBERS入</em>',
          '</li> '].join("");
 
      if (res['data']['list'][i] && res['data']['list'][i]['name']){
        var name = res['data']['list'][i]['name'];
        var create_owner_id = res['data']['list'][i]['create_owner_id'];
        var avatar = res['data']['list'][i]['ext']; 
      }
      var tid = res['data']['list'][i]['id'];
      var title = res['data']['list'][i]['topic_name']; 
      var create_time = res['data']['list'][i]['create_time'];
      var see_num = res['data']['list'][i]['see_num'];
      var mem_number = res['data']['list'][i]['members_number']; 
      
      var max_length = 20;
      var len = title.length >max_length?max_length:title.length;
      title = fly.htmlEscape(title).trim();
      title = title.substring(0, len) + (len ==max_length?"...":"");


      htmlother = htmlother.replace(/TOPIC_ID/g, tid);
      htmlother = htmlother.replace(/USER_ID/g, uid);
      htmlother = htmlother.replace(/SEE_NUMBER/g, see_num);
      htmlother = htmlother.replace(/MEMEBER_NUMBERS/g, mem_number);  
      htmlother = htmlother.replace(/TOPIC_CREATE_TIME/g, create_time);
      htmlother = htmlother.replace(/TOPIC_TITLE/g, title); 

      html = html.replace(/TOPIC_ID/g, tid);
      html = html.replace(/USER_ID/g, uid);
      html = html.replace(/SEE_NUMBER/g, see_num);
      html = html.replace(/MEMEBER_NUMBERS/g, mem_number);  
      html = html.replace(/TOPIC_TITLE/g, title); 
      html = html.replace(/TOPIC_CREATE_TIME/g, create_time);
      if (me){
        if (joinOrNot){
          //join
          content += htmlother;
        }else{
          content += html; 
        }
      }else{
        content += htmlother;
      }
    }
    content += end;
    $(contentId).html(content);

    //topic_page
    
    var pageTotal = ((parseInt(res['data']['total']/pageSize)) > 0)?( parseInt(res['data']['total']/pageSize )+ ((res['data']['total']%pageSize)>0?1:0) ):1;
    var page_html=[
        '<div class="laypage-main">',
        '<a href="/user/index.html?page=1" class="laypage-next">首页</a>',
        '<span class="laypage-curr">1</span>', 
        '<span>…</span>',
        '<a href="/user/index.html?page='+pageTotal+'"&joinpage'+layui.cache.currJoinPage+'" class="laypage-last" title="尾页">尾页</a>',
        '<a href="/user/index.html?page='+currPage+1+'"&joinpage'+layui.cache.currJoinPage+'" class="laypage-next">下一页</a>',
        '</div>'
    ];
    var page_html_home=[
        '<div class="laypage-main">',
        '<a href="/user/home.html?page=1&uid="'+ uid+' class="laypage-next">首页</a>',
        '<span class="laypage-curr">1</span>', 
        '<span>…</span>',
        '<a href="/user/home.html?joinpage='+pageTotal+'&page='+layui.cache.currPage+'&uid='+ uid+'" class="laypage-last" title="尾页">尾页</a>',
        '<a href="/user/home.html?joinpage='+currPage+1+'&page='+layui.cache.currPage+'&uid='+ uid+'" class="laypage-next">下一页</a>',
        '</div>'
    ];
    ////console.log("pageTotal:"+pageTotal+",currPage:"+currPage);
    var pageRealHtml = "";
    pageRealHtml += page_html[0];
    if (currPage >= 5){ 
      if (me){
        pageRealHtml += page_html[1];
      }else{
        pageRealHtml += page_html_home[1]; 
      }
    }
    var start = 1;
    var end = 0;
    if (currPage < 5){
      start = 1;
      end = pageTotal;
    }else{
      start = currPage - 3;
      end = pageTotal;
    }
    if (currPage == pageTotal){
      start = (pageTotal > 4)?(pageTotal - 4):1 ;
      end = pageTotal;
    }
    var index = 0;
    for (var j=start; j<=end; j++){
      if (index++ >= 5){ 
        //由于我的主页和其他用户的主页公用一个界面，导致逻辑较复杂，后续优化！
        if (!me){ 
          if (joinOrNot){
            pageRealHtml += '<a href="/user/home.html?joinpage='+pageTotal+'&page='+layui.cache.currPage+'&uid='+ uid+'" class="laypage-last" title="尾页">尾页</a>';  
            pageRealHtml +=  '<a href="/user/home.html?joinpage='+currPage+1+'&page='+layui.cache.currPage+'&uid='+ uid+'" class="laypage-next">下一页</a>';
          }else{
            pageRealHtml += '<a href="/user/home.html?page='+pageTotal+'&joinpage='+layui.cache.currJoinPage+'&uid='+ uid+'" class="laypage-last" title="尾页">尾页</a>';  
            pageRealHtml +=  '<a href="/user/home.html?page='+currPage+1+'&joinpage='+layui.cache.currJoinPage+'&uid='+ uid+'" class="laypage-next">下一页</a>';
          }
        }else{ 
          if (joinOrNot){
            pageRealHtml += '<a href="/user/index.html?joinpage='+pageTotal+'"&page'+layui.cache.currPage+'&uid='+ uid+'" class="laypage-last" title="尾页">尾页</a>';
            pageRealHtml += '<a href="/user/index.html?joinpage='+currPage+1+'"&page'+layui.cache.currPage+'&uid='+ uid+'" class="laypage-next">下一页</a>';
          }else{
            pageRealHtml += '<a href="/user/index.html?page='+pageTotal+'"&joinpage'+layui.cache.currJoinPage+'&uid='+ uid+'" class="laypage-last" title="尾页">尾页</a>';
            pageRealHtml += '<a href="/user/index.html?page='+currPage+1+'"&joinpage'+layui.cache.currJoinPage+'&uid='+ uid+'" class="laypage-next">下一页</a>';
          }
        }
        break;
      }
      if (j == currPage){ 
        pageRealHtml += '<span class="laypage-curr">'+j+'</span>';
      }else{
        if (!me){ 
          if (joinOrNot){
            pageRealHtml += '<a href="/user/home.html?joinpage='+j+'&page='+layui.cache.currPage+'&uid='+ uid+'" >'+j+'</a>';  

          }else{
            pageRealHtml += '<a href="/user/home.html?page='+j+'&joinpage='+layui.cache.currJoinPage+'&uid='+ uid+'" >'+j+'</a>';  

          }
        }else{
          if (joinOrNot){
            pageRealHtml += '<a href="/user/index.html?joinpage='+j+'&page='+layui.cache.currPage+'&uid='+ uid+'" >'+j+'</a>';  
          }else{
            pageRealHtml += '<a href="/user/index.html?page='+j+'&joinpage='+layui.cache.currJoinPage+'&uid='+ uid+'" >'+j+'</a>';   
          }
        }
      }
    }
    pageRealHtml += page_html[6]; 
    $(contentpageId).html(pageRealHtml); 
  }

  getUserTopicAllList = function(me, uid, currPage, joinOrNot) {
    // body...
    ///api.v2/topic/userTopicList

    var pageSize = 15;
    fly.json("/api.v2/topic/userTopicList", {"userId":uid, "pageSize":pageSize, "order_num":0, "pageNum":currPage}, function(res){ 
      if(res.code == 0){ 
          //console.log("getUserTopicAllList res.total:" +res['data']['total']  +", res:"+ JSON.stringify(res));//res.base64stringC   +", res:"+ JSON.stringify(res)
          //saveTopicAllToCache("alllist_1", JSON.stringify(res)); 
          parseToHTML(res, me, uid, currPage, pageSize, "#topic-list-join", "#topic_page-join", joinOrNot);
      };
    });  
  }

  getTopicAllList = function(me, uid, currPage, joinOrNot){ 
      var pageSize = 15;
      fly.json("/pub6/topic/all", {"createOwnerId":uid, "pageSize":pageSize, "order_num":0, "pageNum":currPage}, function(res){ 
        if(res.code == 0){ 
            //console.log("getTopicAllList res.total:" +res['data']['total']  +", res:"+ JSON.stringify(res));//res.base64stringC   +", res:"+ JSON.stringify(res)
            //saveTopicAllToCache("alllist_1", JSON.stringify(res)); 
            parseToHTML(res, me, uid, currPage, pageSize, "#topic-list", "#topic_page", joinOrNot);
        };
      });  
  }


  //点击@
  $('body').on('click', '#joinTopicContent', function(){
    var othis = $(this), text = othis.text(); 

    if (layui.cache.innerName == 'me'){
      //加载
      getUserTopicAllList(true, layui.cache.user.uid, layui.cache.currPage , true);
      return; 
    }
    if(layui.cache.innerName == 'other'){
      //加载
      getUserTopicAllList(false, layui.cache.curruid, layui.cache.currPage, true );
    }

  });

  //异步渲染
  var asyncRender = function(){
    //加载这个用户的信息

    if (!layui.cache.user || layui.cache.user.uid == -1 || !layui.cache.user.token || layui.cache.user.token.length == 0) {
        return;
    } 

    //修改我的主页url
    $("#my_home_page").attr("href", "/user/home.html?uid="+layui.cache.user.uid);

    if (layui.cache.innerName == 'me'){
      //加载 我的页面的数据
      getTopicAllList(true, layui.cache.user.uid, layui.cache.currPage, false );
      getUserTopicAllList(true, layui.cache.user.uid, layui.cache.currJoinPage , true);
      return; 
    }
    if(layui.cache.innerName == 'other'){
      //加载 其他人主页的数据
      getTopicAllList(false, layui.cache.curruid, layui.cache.currPage , false);
      //加载
      getUserTopicAllList(false, layui.cache.curruid, layui.cache.currJoinPage , true);
      //加载这个人的个人信息  
    
      //$("#user_nickname_other").html(layui.cache.user.username);
      $("#user_avater_header").attr("src", "/res/images/avatar/default.png"); 
      $("#user_avater_header").attr("alt", layui.cache.user.username);

      fly.json('/api.v2/user/getUserProfileByUserId  ', {userid:layui.cache.curruid}, function(res){
          //console.log("getUserProfileByUserId res:"+JSON.stringify(res));  
          if (!res['data']){
              return;
          }
          avatar = res['data']['avaterPath'];
          if (avatar && avatar.length > 0){
            $("#user_avater_header").attr("src", avatar); 
          }else{ 
            $("#user_avater_header").attr("src", "/res/images/avatar/default.png"); 
          }
          if(res['data']['nickname'] && res['data']['nickname'].length >0){

            $("#user_nickname_other").html(res['data']['nickname']);
          } 
          $("#user_joinTime").html(layui.cache.user.joinTime);
          $("#user_city").html(res['data']['cityname']);
          if(res['data']['user_signnote'] && res['data']['user_signnote'].length >0){
            $("#user_signnote").html(res['data']['user_signnote']);
          }
          if(res['data']['blogSite'] && res['data']['blogSite'].length >0){
            $("#user_homesite").html(res['data']['blogSite']);
          }

      }); 
      return;
    }

    if (layui.cache.innerName != 'set'){
      return;
    }

    fly.json('/api.v2/user/getUserProfileByUserId  ', {userid:layui.cache.user.uid}, function(res){
        //console.log("getUserProfileByUserId res:"+JSON.stringify(res));  
        if (!res['data']){
            return;
        }
        $("#email").val(res['data']['email']);
        $("#nickname").val(res['data']['nickname']);
        $("#cityname").val(res['data']['cityname']);
        $("#signNotes").val(res['data']['signNotes']);
        $("#doubanId").val(res['data']['doubanId']);
        $("#weiboId").val(res['data']['csdnId']);
        $("#csdnId").val(res['data']['csdnId']);
        $("#githubId").val(res['data']['githubId']);
        $("#blogSite").val(res['data']['blogSite']);
    }); 

  }();

  //我的相关数据
  var elemUC = $('#LAY_uc'), elemUCM = $('#LAY_ucm');
  gather.minelog = {};
  gather.mine = function(index, type, url){
    var tpl = [
      //求解
      '{{# for(var i = 0; i < d.rows.length; i++){ }}\
      <li>\
        {{# if(d.rows[i].collection_time){ }}\
          <a class="jie-title" href="/jie/{{d.rows[i].id}}/" target="_blank">{{= d.rows[i].title}}</a>\
          <i>{{ d.rows[i].collection_time }} 收藏</i>\
        {{# } else { }}\
          {{# if(d.rows[i].status == 1){ }}\
          <span class="fly-jing layui-hide-xs">精</span>\
          {{# } }}\
          {{# if(d.rows[i].accept >= 0){ }}\
            <span class="jie-status jie-status-ok">已结</span>\
          {{# } else { }}\
            <span class="jie-status">未结</span>\
          {{# } }}\
          {{# if(d.rows[i].status == -1){ }}\
            <span class="jie-status">审核中</span>\
          {{# } }}\
          <a class="jie-title" href="/jie/{{d.rows[i].id}}/" target="_blank">{{= d.rows[i].title}}</a>\
          <i class="layui-hide-xs">{{ layui.util.timeAgo(d.rows[i].time, 1) }}</i>\
          {{# if(d.rows[i].accept == -1){ }}\
          <a class="mine-edit layui-hide-xs" href="/jie/edit/{{d.rows[i].id}}" target="_blank">编辑</a>\
          {{# } }}\
          <em class="layui-hide-xs">{{d.rows[i].hits}}阅/{{d.rows[i].comment}}答</em>\
        {{# } }}\
      </li>\
      {{# } }}'
    ];

    var view = function(res){
      var html = laytpl(tpl[0]).render(res);
      dom.mine.children().eq(index).find('span').html(res.count);
      elemUCM.children().eq(index).find('ul').html(res.rows.length === 0 ? '<div class="fly-msg">没有相关数据</div>' : html);
    };

    var page = function(now){
      var curr = now || 1;
      if(gather.minelog[type + '-page-' + curr]){
        view(gather.minelog[type + '-page-' + curr]);
      } else {
        //我收藏的帖
        if(type === 'collection'){
          var nums = 10; //每页出现的数据量
          fly.json(url, {}, function(res){
            res.count = res.rows.length;

            var rows = layui.sort(res.rows, 'collection_timestamp', 'desc')
            ,render = function(curr){
              var data = []
              ,start = curr*nums - nums
              ,last = start + nums - 1;

              if(last >= rows.length){
                last = curr > 1 ? start + (rows.length - start - 1) : rows.length - 1;
              }

              for(var i = start; i <= last; i++){
                data.push(rows[i]);
              }

              res.rows = data;
              
              view(res);
            };

            render(curr)
            gather.minelog['collect-page-' + curr] = res;

            now || laypage.render({
              elem: 'LAY_page1'
              ,count: rows.length
              ,curr: curr
              ,jump: function(e, first){
                if(!first){
                  render(e.curr);
                }
              }
            });
          });
        } else {
          fly.json('/api/'+ type +'/', {
            page: curr
          }, function(res){
            view(res);
            gather.minelog['mine-jie-page-' + curr] = res;
            now || laypage.render({
              elem: 'LAY_page'
              ,count: res.count
              ,curr: curr
              ,jump: function(e, first){
                if(!first){
                  page(e.curr);
                }
              }
            });
          });
        }
      }
    };

    if(!gather.minelog[type]){
      page();
    }
  };

  if(elemUC[0]){
    layui.each(dom.mine.children(), function(index, item){
      var othis = $(item)
      gather.mine(index, othis.data('type'), othis.data('url'));
    });
  }

  //显示当前tab
  if(location.hash){
    element.tabChange('user', location.hash.replace(/^#/, ''));
  }

  element.on('tab(user)', function(){
    var othis = $(this), layid = othis.attr('lay-id');
    if(layid){
      location.hash = layid;
    }
  });

  //根据ip获取城市
  if($('#L_city').val() === ''){
    // $.getScript('http://int.dpool.sina.com.cn/iplookup/iplookup.php?format=js', function(){
    //   $('#L_city').val(remote_ip_info.city||'');
    // });
  }

  var avatarAdd = $('.avatar-add');

  fly.json('/api.v2/user/getUploadToken', '', function(res){
    //console.log("getUploadToken res:"+JSON.stringify(res));  
    var server= res['data']['server'];
    qiniuyun.loader({
      domain: 'https://guaniu.it3q.com'              // 后台设置的域名项
      , elem: "#upload-img"          // 绑定的element
      , token: res['data']['token']
      , retryCount: 6                  // 重连次数，默认6(可选) 
      , next: function(response){
        //element.progress('video-progress', response.total.percent + '%');       // 进度条
      }
      , complete: function(res){
        // layer.closeAll('loading'); // 关闭loading关闭
        layer.msg("上传成功！");
        //console.log(res) 
        var imgUrl = server+res['key']+"?imageView2/2/480/";
        $("#upload_img_display").attr("src", imgUrl); 
        fly.json('/api.v2/user/editHeaderImage', {"avaterPath":imgUrl}, function(res){
          //console.log("EditHeaderImage res:"+JSON.stringify(res));  
          layui.cache.user.avatar = imgUrl;
        });
      }
    });
  });


  //上传图片
  // if($('.upload-img')[0]){ 
  //     fly.json('/api.v2/user/getUploadToken', '', function(res){
  //       //console.log("--getUploadToken res:"+JSON.stringify(res));  
       
  //       layui.use('upload', function(upload){
  //         var avatarAdd = $('.avatar-add');
  //         upload.render({
  //           elem: '.upload-img'
  //           ,url: 'https://up-z0.qiniup.com/'
  //           ,size: 4500
  //           ,before: function(){
  //             avatarAdd.find('.loading').show();
  //           } 
  //           , method: 'post'
  //           , data: {
  //                 //key: 'aaa.png',  //自定义文件名
  //                 token: res['data']['token']
  //             } 
  //           ,done: function(res){
  //             if(res.status == 0){
  //               $.post('/user/set/', {
  //                 avatar: res.url
  //               }, function(res){
  //                 location.reload();
  //               });
  //             } else {
  //               layer.msg(res.msg, {icon: 5});
  //             }
  //             avatarAdd.find('.loading').hide();
  //           }
  //           ,error: function(){
  //             avatarAdd.find('.loading').hide();
  //           }
  //         });
  //       });
  //     });

  //}

  //合作平台
  if($('#LAY_coop')[0]){

     qiniuyun.loader({
        domain: '{$domain}'              // 后台设置的域名项
        , elem: "#upload-img"          // 绑定的element
        , token: "{$token}"              // 授权token
        , retryCount: 6                  // 重连次数，默认6(可选)
        , region: qiniu.region.z0        // 选择上传域名区域，默认自动分析(可选)
        , next: function(response){
          element.progress('video-progress', response.total.percent + '%');       // 进度条
        }
        , complete: function(res){
          // layer.closeAll('loading'); // 关闭loading关闭
          layer.msg("上传成功！");
          //console.log(res) 
        }
      });
    //资源上传
    $('#LAY_coop .uploadRes').each(function(index, item){
      var othis = $(this);



      upload.render({
        elem: item
        ,url: '/api/upload/cooperation/?filename='+ othis.data('filename')
        ,accept: 'file'
        ,exts: 'zip'
        ,size: 30*1024
        ,before: function(){
          layer.msg('正在上传', {
            icon: 16
            ,time: -1
            ,shade: 0.7
          });
        }
        ,done: function(res){
          if(res.code == 0){
            layer.msg(res.msg, {icon: 6})
          } else {
            layer.msg(res.msg)
          }
        }
      });
    });

    //成效展示
    var effectTpl = ['{{# layui.each(d.data, function(index, item){ }}'
    ,'<tr>'
      ,'<td><a href="/u/{{ item.uid }}" target="_blank" style="color: #01AAED;">{{ item.uid }}</a></td>'
      ,'<td>{{ item.authProduct }}</td>'
      ,'<td>￥{{ item.rmb }}</td>'
      ,'<td>{{ item.create_time }}</td>'
      ,'</tr>'
    ,'{{# }); }}'].join('');

    var effectView = function(res){
      var html = laytpl(effectTpl).render(res);
      $('#LAY_coop_effect').html(html);
      $('#LAY_effect_count').html('你共有 <strong style="color: #FF5722;">'+ (res.count||0) +'</strong> 笔合作授权订单');
    };

    var effectShow = function(page){
      fly.json('/cooperation/effect', {
        page: page||1
      }, function(res){
        effectView(res);
        laypage.render({
          elem: 'LAY_effect_page'
          ,count: res.count
          ,curr: page
          ,jump: function(e, first){
            if(!first){
              effectShow(e.curr);
            }
          }
        });
      });
    };

    effectShow();

  }

  //提交成功后刷新
  fly.form['set-mine'] = function(data, required){
    layer.msg('修改成功', {
      icon: 1
      ,time: 1000
      ,shade: 0.1
    }, function(){
      location.reload();
    });
  }

  //帐号绑定
  $('.acc-unbind').on('click', function(){
    var othis = $(this), type = othis.attr('type');
    layer.confirm('整的要解绑'+ ({
      qq_id: 'QQ'
      ,weibo_id: '微博'
    })[type] + '吗？', {icon: 5}, function(){
      fly.json('/api/unbind', {
        type: type
      }, function(res){
        if(res.status === 0){
          layer.alert('已成功解绑。', {
            icon: 1
            ,end: function(){
              location.reload();
            }
          });
        } else {
          layer.msg(res.msg);
        }
      });
    });
  });


  //我的消息
  gather.minemsg = function(){
    var delAll = $('#LAY_delallmsg')
    ,tpl = '{{# var len = d.rows.length;\
    if(len === 0){ }}\
      <div class="fly-none">您暂时没有最新消息</div>\
    {{# } else { }}\
      <ul class="mine-msg">\
      {{# for(var i = 0; i < len; i++){ }}\
        <li data-id="{{d.rows[i].id}}">\
          <blockquote class="layui-elem-quote">{{ d.rows[i].content}}</blockquote>\
          <p><span>{{d.rows[i].time}}</span><a href="javascript:;" class="layui-btn layui-btn-sm layui-btn-danger fly-delete">删除</a></p>\
        </li>\
      {{# } }}\
      </ul>\
    {{# } }}'
    ,delEnd = function(clear){
      if(clear || dom.minemsg.find('.mine-msg li').length === 0){
        dom.minemsg.html('<div class="fly-none">您暂时没有最新消息</div>');
      }
    }
    
    
    /*
    fly.json('/message/find/', {}, function(res){
      var html = laytpl(tpl).render(res);
      dom.minemsg.html(html);
      if(res.rows.length > 0){
        delAll.removeClass('layui-hide');
      }
    });
    */
    
    //阅读后删除
    dom.minemsg.on('click', '.mine-msg li .fly-delete', function(){
      var othis = $(this).parents('li'), id = othis.data('id');
      fly.json('/message/remove/', {
        id: id
      }, function(res){
        if(res.status === 0){
          othis.remove();
          delEnd();
        }
      });
    });

    //删除全部
    $('#LAY_delallmsg').on('click', function(){
      var othis = $(this);
      layer.confirm('确定清空吗？', function(index){
        fly.json('/message/remove/', {
          all: true
        }, function(res){
          if(res.status === 0){
            layer.close(index);
            othis.addClass('layui-hide');
            delEnd(true);
          }
        });
      });
    });

  };

  dom.minemsg[0] && gather.minemsg();

  exports('user', null);
  
});