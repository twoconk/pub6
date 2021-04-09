/**

 @Name: 求解板块

 */
 
layui.define('fly', function(exports){

  var $ = layui.jquery;
  var layer = layui.layer;
  var util = layui.util;
  var laytpl = layui.laytpl;
  var form = layui.form;
  var fly = layui.fly;
  
  var gather = {}, dom = {
    jieda: $('#jieda')
    ,content: $('#L_content')
    ,jiedaCount: $('#jiedaCount')
  };

  //监听专栏选择
  form.on('select(column)', function(obj){
    var value = obj.value
    ,elemQuiz = $('#LAY_quiz')
    ,tips = {
      tips: 1
      ,maxWidth: 250
      ,time: 10000
    };
    elemQuiz.addClass('layui-hide');
    if(value === '0'){
      layer.tips('下面的信息将便于您获得更好的答案', obj.othis, tips);
      elemQuiz.removeClass('layui-hide');
    } else if(value === '99'){
      layer.tips('系统会对【分享】类型的帖子予以飞吻奖励，但我们需要审核，通过后方可展示', obj.othis, tips);
    }
  });
 
  //求解管理
  gather.jieAdmin = {
    //删求解
    del: function(div){
      layer.confirm('确认删除该求解么？', function(index){
        layer.close(index);
        fly.json('/api/jie-delete/', {
          id: div.data('id')
        }, function(res){
          if(res.status === 0){
            location.href = '/jie/';
          } else {
            layer.msg(res.msg);
          }
        });
      });
    }
    
    //设置置顶、状态
    ,set: function(div){
      var othis = $(this);
      fly.json('/api/jie-set/', {
        id: div.data('id')
        ,rank: othis.attr('rank')
        ,field: othis.attr('field')
      }, function(res){
        if(res.status === 0){
          location.reload();
        }
      });
    }

    //收藏
    ,collect: function(div){
      var othis = $(this), type = othis.data('type');
      fly.json('/collection/'+ type +'/', {
        cid: div.data('id')
      }, function(res){
        if(type === 'add'){
          othis.data('type', 'remove').html('取消收藏').addClass('layui-btn-danger');
        } else if(type === 'remove'){
          othis.data('type', 'add').html('收藏').removeClass('layui-btn-danger');
        }
      });
    }
  };

  $('body').on('click', '#LAY_join_topic', function(){
    var othis = $(this);
    if(othis.hasClass("layui-btn-disabled")) return;

    fly.json('/api.v2/topicmembers/add', {topic_id:layui.cache.tid}, function(res){
      othis.removeClass("layui-btn-disabled");
      $("#LAY_join_topic").removeClass("layui-btn-danger");
      $("#LAY_join_topic").addClass("layui-btn-disabled");
      $("#LAY_quit_topic").removeClass("layui-btn-disabled");
      $("#LAY_quit_topic").addClass("layui-btn-danger");
    }, {
      error: function(){
        othis.removeClass("layui-btn-disabled");
      }
    });
 
  });

  //for 任务 
  $('body').on('click', '#onClickTaskItem', function(){
    var othis = $(this);
    if(othis.hasClass("layui-btn-disabled")) return; 
    if (!taskCacheList){
      return;
    }  
    index = othis.attr("taskId");
    if(!taskCacheList[index] || !taskCacheList[index]['id']){
      return;
    }
    taskId = taskCacheList[index]['id'];  

    //从网络加载详情
    fly.json('/api.v2/topictask/edit', {
      id: taskId
    }, function(res){

      //console.log("onClickTaskItem res:"+ ", res:"+ JSON.stringify(res) );//
      var content = res['data']['task_content'];
      
      content = /^\{html\}/.test(content) 
        ? content.replace(/^\{html\}/, '')
      : fly.content(content);
      header = ['<fieldset class="layui-elem-field layui-field-title" style="margin-top: 30px;">',
          '<legend>时间计划：'+ "[" + taskCacheList[index]['time_line'] +"]" +'</legend>',
          '</fieldset>'].join("");

      layer.open({
        type: 1
        ,title: '预览:' + taskCacheList[index]['task_title'] 
        ,shade: false
        ,area: ['80%', '60%']
        ,scrollbar: false
        ,content: '<div class="detail-body" style="margin:20px;">'+ header + content +'</div>'
      });
    }, {type:"Get"});

  });
  //for 笔记 
  var resourceCacheList; 
  $('body').on('click', '#onClickResource', function(){
    var othis = $(this);
    if(othis.hasClass("layui-btn-disabled")) return; 
    if (!resourceCacheList){
      return;
    }
    resourceId = othis.attr("resourceId");

    //从网络加载详情
    fly.json('/api.v2/topicnotes/edit', {
      id: resourceCacheList[resourceId]['id']
    }, function(res){

      //console.log("onClickResource res:"+ ", res:"+ JSON.stringify(res) );//
      var content = res['data']['content'];
      
      content = /^\{html\}/.test(content) 
        ? content.replace(/^\{html\}/, '')
      : fly.content(content);

      layer.open({
        type: 1
        ,title: '预览:' + resourceCacheList[resourceId]['title']
        ,shade: false
        ,area: ['80%', '60%']
        ,scrollbar: false
        ,content: '<div class="detail-body" style="margin:20px;">'+ content +'</div>'
      });
    }, {type:"Get"});

  });
  //for 任务 编辑
  saveTopicNoteToCache = function(name, value){ 
      var saveData = {key:name, value:value};
      layui.sessionData('pub6_com_user_cache_topic_task', saveData);//把AJSON对象存储为字符串
  }
  loadTopicNoteToCache = function (name) { 

      var cacheUser = layui.sessionData('pub6_com_user_cache_topic_task');
      return cacheUser[name] || {};
  } 
  var taskCacheList; 
  $('body').on('click', '#onClickEditTask', function(){
    var othis = $(this);
    if(othis.hasClass("layui-btn-disabled")) return; 
    if (!taskCacheList){
      return;
    }
    index = othis.attr("taskId");
    taskId = taskCacheList[index]['id'];  
    saveTopicNoteToCache(taskId, taskCacheList[index]);
    location.href = "/jie/task.html?tid="+layui.cache.tid+"&taskId="+taskId;
  });

  $('body').on('click', '#LAY_add_task', function(){
    var othis = $(this);
    if(othis.hasClass("layui-btn-disabled")) return;

    if (layui.cache.user.uid == -1 || layui.cache.tid == 0){
      return;
    }
    //layui.cache.lasturl = "/jie/detail.html?tid="+layui.cache.tid;
    location.href = "/jie/task.html?tid="+layui.cache.tid;
  });
  $('body').on('click', '#LAY_add_resource', function(){
    var othis = $(this);
    if(othis.hasClass("layui-btn-disabled")) return;
    
    if (layui.cache.user.uid == -1 || layui.cache.tid == 0){
      return;
    }
    //layui.cache.lasturl = "/jie/detail.html?tid="+layui.cache.tid;
    location.href = "/jie/resource.html?tid="+layui.cache.tid;
  });
  $('body').on('click', '#LAY_add_notes', function(){
    var othis = $(this);
    if(othis.hasClass("layui-btn-disabled")) return;
    
    if (layui.cache.user.uid == -1 || layui.cache.tid == 0){
      return;
    }
    //layui.cache.lasturl = "/jie/detail.html?tid="+layui.cache.tid;
    location.href = "/jie/note.html?tid="+layui.cache.tid;
  });

  $('body').on('click', '#LAY_quit_topic', function(){
    var othis = $(this);
    if (layui.cache.user.uid == -1 || layui.cache.tid == 0){
      return;
    }
    if(othis.hasClass("layui-btn-disabled")) return;

      layer.confirm('确认退出该主题学习组么？', function(index){
        layer.close(index);


        fly.json('/api.v2/topicmembers/delete', {topic_id:layui.cache.tid}, function(res){
            othis.removeClass("layui-btn-disabled");
          $("#LAY_quit_topic").removeClass("layui-btn-danger");
          $("#LAY_quit_topic").addClass("layui-btn-disabled");
          $("#LAY_join_topic").removeClass("layui-btn-disabled");
          $("#LAY_join_topic").addClass("layui-btn-danger");
        }, {
          error: function(){
            othis.removeClass("layui-btn-disabled");
          }
        });

      }); 
  });
  //

  $('body').on('click', '#LAY_more_resources', function(){
    var othis = $(this);
    if (layui.cache.user.uid == -1 || layui.cache.tid == 0){
      return;
    }
    if(othis.hasClass("layui-btn-disabled")) return;
   
    //layui.cache.lasturl = "/jie/detail.html?tid="+layui.cache.tid;
    location.href = "/jie/list.html?type=1&tid="+layui.cache.tid;
  });

  var commentList;
  $('body').on('click', '#reply_user', function(){
    var othis = $(this);
    if (layui.cache.user.uid == -1 || layui.cache.tid == 0){
      return;
    }
    if(othis.hasClass("layui-btn-disabled")) return;
    
    var index = othis.attr("key");
    var user = commentList[index]['name'];
    var old = $("#content").val();
    $("#content").val(old + " @"+user+" ");
    layui.form.render();
  });
 
  $('body').on('click', '#LAY_more_notes', function(){
    var othis = $(this);
    if (layui.cache.user.uid == -1 || layui.cache.tid == 0){
      return;
    }
    if(othis.hasClass("layui-btn-disabled")) return;
  
    if (layui.cache.user.uid == -1 || layui.cache.tid == 0){
      return;
    }
    //layui.cache.lasturl = "/jie/detail.html?tid="+layui.cache.tid;
    location.href = "/jie/list.html?type=2&tid="+layui.cache.tid;
  });


  // var postCommentEnd = function(){
  //     //暂时放在这里，更新详情页面的评论
  //     //console.log("postCommentEnd!");
  //     loadComment(layui.cache.currPage);
  // } ;
  
  var loadComment = function(currPage){

    commentList = null;
    //取这个主题的评论列表
    fly.json('/api.v2/topic/commentList', {topic_id:layui.cache.tid, "pageNum":currPage}, function(res){
      //console.log("[1]commentList res:"+JSON.stringify(res));  
           
      if (!res['data']['list']){
        return;
      }
      commentList = res['data']['list'];

      var task_list = "";

      for(let i = 0; i< res['data']['list'].length; i++){

          var otherhtml = [
            '<li data-id="111">'
            ,'<a name="item-1111111111"></a>'
            ,'<div class="detail-about detail-about-reply">'
              ,'<a class="fly-avatar" href="">'
                ,'<img src="USER_IMAGE_RESOURCE" alt="USERNAME">'
              ,'</a>'
              ,'<div class="fly-detail-user">'
               ,' <a href="" class="fly-link">'
                ,'  <cite>USERNAME</cite>       '
                ,'</a>'
              ,'</div>'
              ,'<div class="detail-hits">'
               ,' <span>TIME</span>'
              ,'</div>'
            ,'</div>'
            ,'<div class="detail-body jieda-body photos">'
             ,' <p>CONTENT</p>'
            ,'</div>'
            ,'<div class="jieda-reply">' 
              ,'<span id="reply_user" key="REPLAY_COMMENT_ID" type="reply">'
               ,' <i class="iconfont icon-svgmoban53"></i>'
                ,'回复'
              ,'</span>' 
            ,'</div>'
          ,'</li>'
          ].join(""); 
        var userId = res['data']['list'][i]['user_id'];
        var avatar = res['data']['list'][i]['ext'];
        var name = res['data']['list'][i]['name'];
        var content = res['data']['list'][i]['content'];
                //转义
                content = fly.htmlEscape(content).trim();


        //增加转义处理
        content = fly.content(content);
        if (userId == layui.cache.topic_owner_id){
            //是owner  
        }else{
        }

        otherhtml = otherhtml.replace(/REPLAY_COMMENT_ID/g, i);
        otherhtml = otherhtml.replace(/TIME/g, res['data']['list'][i]['create_time']);
        otherhtml = otherhtml.replace(/USERNAME/g, res['data']['list'][i]['name']); 
        otherhtml = otherhtml.replace(/CONTENT/g, content); 
        otherhtml = otherhtml.replace(/LIKE_NUMBER/g, res['data']['list'][i]['likeNum']); 
        if (avatar && avatar.length > 0){
          otherhtml = otherhtml.replace(/USER_IMAGE_RESOURCE/g, avatar);

        }else{

        otherhtml = otherhtml.replace(/USER_IMAGE_RESOURCE/g, "/res/images/avatar/default.png");
        }
        task_list+= otherhtml;

      } 

      $("#jieda").html(task_list);



      //comment_page
      var pageSize = 10;
      
      var pageTotal = ((parseInt(res['data']['total']/pageSize)) > 0)?( parseInt(res['data']['total']/pageSize )+ ((res['data']['total']%pageSize)>0?1:0) ):1;
      var page_html=[
          '<div class="laypage-main">',
          '<a href="/jie/detail.html?tid='+layui.cache.tid+'" class="laypage-next">首页</a>',
          '<span class="laypage-curr">1</span>', 
          '<span>…</span>',
          '<a href="/jie/detail.html?tid='+layui.cache.tid+'&page='+pageTotal+'" class="laypage-last" title="尾页">尾页</a>',
          '<a href="/jie/detail.html?tid='+layui.cache.tid+'&page='+currPage+1+'" class="laypage-next">下一页</a>',
          '</div>'
      ];
      //console.log("pageTotal:"+pageTotal+",currPage:"+currPage);
      var pageRealHtml = "";
      pageRealHtml += page_html[0];
      if (currPage >= 5){ 
        pageRealHtml += page_html[1];
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
          pageRealHtml += page_html[3];
          pageRealHtml += page_html[4]; 
          break;
        }
        if (j == currPage){ 
          pageRealHtml += '<span class="laypage-curr">'+j+'</span>';
        }else{
          pageRealHtml += '<a href="/jie/detail.html?tid='+layui.cache.tid+'&page='+j+'" >'+j+'</a>'; 
        }
      }
      pageRealHtml += page_html[6]; 
      $("#comment_page").html(pageRealHtml);
    });
  }

  var loadTopicNotes = function(currPage, needPage) {
    fly.json('/api.v2/topicnotes/list', {topic_id:layui.cache.tid, "pageNum":currPage}, function(res){
      //console.log("topicnotes res:"+JSON.stringify(res));  
         
      var task_list = [' <dt >' , 
        ,'<div class="fly-panel-title">'
          ,'主题笔记 '
          ,'<span class="fly-signin-days"><a href="#"  id="LAY_more_notes">更多</a> </span>'
        ,'</div> </dt>'
      ].join("");

      if (needPage){ 
          task_list = ""; 
      }
      var nodata = '<div class="fly-none">没有相关数据</div>';
      if (!res['data']['list']){
        task_list += nodata; 
        if (needPage){
          $("#content_list").append(task_list);

        }else{
          $("#topic_notes").append(task_list);

        }
        return;
      }
        
      resourceCacheList = res['data']['list'];
      for(let i = 0; i< res['data']['list'].length; i++){ 
        var html = [
          '<dd>'
            ,'<a href="#" id="onClickUser" userId="USERNAME_ID">[USERNAME]</a><a href="#" id="onClickResource" resourceId="RESOURCE_ID">TITLE</a>' 
          ,'</dd>' 
        ].join('');
 
        //HTML2
        var html2 = ['<ul class="fly-list">'
            ,'<li>'
                ,'<a alt= "USERNAME" href="/user/home.html?uid=USERNAME_ID" class="fly-avatar">'
                            ,'<img src="USER_OWNER_SRC" alt="USERNAME">'
                ,'</a>'
                ,'<h2>'
                  ,'<a  href="#"  id="onClickResource" resourceId="RESOURCE_ID">TITLE</a>'
                ,'</h2>'
                ,'<div class="fly-list-info">'
                  ,'<a href="/user/home.html?uid=USERNAME_ID" link>'
                    ,'<cite>USERNAME</cite>'
                    ,'<!--'
                    ,'<i class="iconfont icon-renzheng" title="认证信息：XXX"></i>'
                    ,'<i class="layui-badge fly-badge-vip">VIP3</i>'
                    ,'-->'
                  ,'</a>'
                  ,'<span>TOPIC_CREATE_TIME</span>' 

                ,'</div> '
              ,'</li>'
          ,'</ul>'].join('');
   
        var name = res['data']['list'][i]['name'];
        var uid = res['data']['list'][i]['user_id'];
        var avatar = res['data']['list'][i]['ext'];

        var tid = res['data']['list'][i]['id'];
        var title = res['data']['list'][i]['title']; 
        //转义
        title = fly.htmlEscape(title).trim();

        var create_time = res['data']['list'][i]['create_time'];
    
        html2 = html2.replace(/USERNAME_ID/g, uid); 
        if (avatar && avatar.length > 0){
          html2 = html2.replace(/USER_OWNER_SRC/g, avatar);

        }else{

          html2 = html2.replace(/USER_OWNER_SRC/g, "/res/images/avatar/default.png");
        }
        html2 = html2.replace(/USERNAME/g, name);
        html2 = html2.replace(/TITLE/g, title); 
        html2 = html2.replace(/TOPIC_CREATE_TIME/g, create_time);
        html2 = html2.replace(/RESOURCE_ID/g, i); 

        html = html.replace(/RESOURCE_ID/g, i); 
        html = html.replace(/USERNAME_ID/g, res['data']['list'][i]['user_id']);  
        html = html.replace(/USERNAME/g, res['data']['list'][i]['name']);  
        html = html.replace(/TITLE/g, res['data']['list'][i]['title']);  
        html = html.replace(/TYPE/g, res['data']['list'][i]['like_num']); 

        if (needPage){

          task_list+= html2;
        }else{

          task_list+= html;
        }
      } 

      if (needPage){
        $("#content_list").append(task_list);

      }else{
        $("#topic_notes").append(task_list);

      }


      if (needPage){


        //comment_page
        var pageSize = 10;
        var pageTotal = (res['data']['total']/pageSize)?(res['data']['total']/pageSize+ ((res['data']['total']%pageSize)>0?1:0)):1;
        var page_html=[
            '<div class="laypage-main">',
            '<a href="/jie/list.html?tid='+layui.cache.tid+'&type='+layui.cache.type+'" class="laypage-next">首页</a>',
            '<span class="laypage-curr">1</span>', 
            '<span>…</span>',
            '<a href="/jie/list.html?tid='+layui.cache.tid+'&page='+currPage+'&type='+layui.cache.type+'" class="laypage-last" title="尾页">尾页</a>',
            '<a href="/jie/list.html?tid='+layui.cache.tid+'&type='+layui.cache.type+'&page='+currPage+1+'" class="laypage-next">下一页</a>',
            '</div>'
        ];
        //console.log("pageTotal:"+pageTotal+",currPage:"+currPage);
        var pageRealHtml = "";
        pageRealHtml += page_html[0];
        if (currPage >= 5){ 
          pageRealHtml += page_html[1]; 
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
            pageRealHtml += page_html[3];
            pageRealHtml += page_html[4]; 
            break;
          }
          if (j == currPage){ 
            pageRealHtml += '<span class="laypage-curr">'+j+'</span>';
          }else{
            pageRealHtml += '<a href="/jie/list.html?tid='+layui.cache.tid+'&type='+layui.cache.type+'&page='+j+'" >'+j+'</a>'; 
          }
        }
        pageRealHtml += page_html[6]; 
        $("#content_page").append(pageRealHtml);
      }
    });
  };


  var loadTopicResources = function(currPage, needPage) {
    // body...

    fly.json('/api.v2/topicresource/list', {topic_id:layui.cache.tid, "pageNum":currPage}, function(res){
      //console.log("topicresource res:"+JSON.stringify(res));  
            
      var task_list = [' <dt >' , 
        ,'<div class="fly-panel-title">'
          ,'主题资源 '
          ,'<span class="fly-signin-days"><a href="#"  id="LAY_more_resources">更多</a> </span>'
        ,'</div> </dt>'
      ].join("");
      var nodata = '<div class="fly-none">没有相关数据</div>';

      if (needPage){ 
          task_list = ""; 
      }

      if (!res['data']['list']){
        task_list += nodata;
        if (needPage){
          $("#content_list").append(task_list);

        }else{
          $("#topic_resource").append(task_list);

        }
        return;
      }

      for(let i = 0; i< res['data']['list'].length; i++){ 
        var html = [
          '<dd>'
            ,'<span><i class="iconfont"></i> [TYPE]</span><a href="HREF_TARGET"  title="CONTENT" target="_blank">TITLE</a>' 
          ,'</dd>' 
        ].join('');  

        //HTML2
        var html2 = ['<ul class="fly-list">'
            ,'<li>'
                ,'<a alt= "USERNAME" href="/user/home.html?uid=USERNAME_ID" class="fly-avatar">'
                            ,'<img src="USER_OWNER_SRC" alt="USERNAME">'
                ,'</a>'
                ,'<h2>'
                  ,'<a href="HREF_TARGET"  target="_blank" title="CONTENT"> [TYPE] TITLE</a>'
                ,'</h2>'
                ,'<div class="fly-list-info">'
                  ,'<a href="/user/home.html?uid=USERNAME_ID" link>'
                    ,'<cite>USERNAME</cite>'
                    ,'<!--'
                    ,'<i class="iconfont icon-renzheng" title="认证信息：XXX"></i>'
                    ,'<i class="layui-badge fly-badge-vip">VIP3</i>'
                    ,'-->'
                  ,'</a>'
                  ,'<span>TOPIC_CREATE_TIME</span>' 
                  ,'<span class="fly-list-kiss layui-hide-xs" title="悬赏飞吻"><i class="iconfont icon-kiss"></i> 60</span>'
                  ,'<!--<span class="layui-badge fly-badge-accept layui-hide-xs">已结</span>-->'
                  ,'<span class="fly-list-nums"> '
                    ,'<i class="iconfont icon-renshu" title=""></i> 66'
                  ,'</span>'
                ,'</div> '
              ,'</li>'
          ,'</ul>'].join('');
   
        var name = res['data']['list'][i]['name'];
        var uid = res['data']['list'][i]['user_id'];
        var avatar = res['data']['list'][i]['ext'];

        var tid = res['data']['list'][i]['id'];
        var title = res['data']['list'][i]['resource_link']; 
        var content = res['data']['list'][i]['resource_content']; 
        var create_time = res['data']['list'][i]['create_time'];
    
        html2 = html2.replace(/USERNAME_ID/g, uid); 
        if (avatar && avatar.length > 0){
          html2 = html2.replace(/USER_OWNER_SRC/g, avatar);

        }else{

          html2 = html2.replace(/USER_OWNER_SRC/g, "/res/images/avatar/default.png");
        }
        html2 = html2.replace(/USERNAME/g, name);
        html2 = html2.replace(/TITLE/g, title); 
        html2 = html2.replace(/TOPIC_CREATE_TIME/g, create_time);
        html2 = html2.replace(/CONTENT/g, content);
 

        html = html.replace(/TITLE/g, res['data']['list'][i]['resource_link']);
        html = html.replace(/CONTENT/g, content);

        html = html.replace(/HREF_TARGET/g, res['data']['list'][i]['resource_link']); 
        html2 = html2.replace(/HREF_TARGET/g, res['data']['list'][i]['resource_link']);  

        switch (res['data']['list'][i]['resource_type_id']){
          case 1:
            html = html.replace(/TYPE/g,  '网课链接');
            html2 = html2.replace(/TYPE/g,  '网课链接');
          break;
          case 2:
            html = html.replace(/TYPE/g, "专业书籍");
            html2 = html2.replace(/TYPE/g,  '专业书籍');
          break;
          case 3:
            html = html.replace(/TYPE/g, "文档材料");
            html2 = html2.replace(/TYPE/g,  '文档材料');
          break;
          case 4:
            html = html.replace(/TYPE/g, "其他类型");
            html2 = html2.replace(/TYPE/g,  '其他类型');
          break;
          default:
            html = html.replace(/TYPE/g, "其他类型");
            html2 = html2.replace(/TYPE/g,  '其他类型');
          break;
        }
        if (needPage){

          task_list+= html2;
        }else{

          task_list+= html;
        }
      } 
      if (needPage){
        $("#content_list").append(task_list);

      }else{
        $("#topic_resource").append(task_list);

      }

      if (needPage){
          

        //comment_page
        var pageSize = 10;
        var pageTotal = (res['data']['total']/pageSize)?(res['data']['total']/pageSize+ ((res['data']['total']%pageSize)>0?1:0)):1;
        var page_html=[
            '<div class="laypage-main">',
            '<a href="/jie/list.html?tid='+layui.cache.tid+'&page='+currPage+'&type='+layui.cache.type+'" class="laypage-next">首页</a>',
            '<span class="laypage-curr">1</span>', 
            '<span>…</span>',
            '<a href="/jie/list.html?tid='+layui.cache.tid+'&page='+currPage+'&type='+layui.cache.type+'" class="laypage-last" title="尾页">尾页</a>',
            '<a href="/jie/list.html?tid='+layui.cache.tid+'&type='+layui.cache.type+'&page='+currPage+1+'" class="laypage-next">下一页</a>',
            '</div>'
        ];
        //console.log("pageTotal:"+pageTotal+",currPage:"+currPage);
        var pageRealHtml = "";
        pageRealHtml += page_html[0];
        if (currPage >= 5){ 
          pageRealHtml += page_html[1];
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
            pageRealHtml += page_html[3];
            pageRealHtml += page_html[4]; 
            break;
          }
          if (j == currPage){ 
            pageRealHtml += '<span class="laypage-curr">'+j+'</span>';
          }else{
            pageRealHtml += '<a href="/jie/list.html?tid='+layui.cache.tid+'&type='+layui.cache.type+'&page='+j+'" >'+j+'</a>'; 
          }
        }
        pageRealHtml += page_html[6]; 
        $("#content_page").append(pageRealHtml);
      }
    });
  }

  //异步渲染
  var asyncRender = function(){
    var div = $('.fly-admin-box'), jieAdmin = $('#LAY_jieAdmin');
    //查询帖子是否收藏
    //if(jieAdmin[0] && layui.cache.user.uid != -1){
      /*
      fly.json('/collection/find/', {
        cid: div.data('id')
      }, function(res){
        jieAdmin.append('<span class="layui-btn layui-btn-xs jie-admin '+ (res.data.collection ? 'layui-btn-danger' : '') +'" type="collect" data-type="'+ (res.data.collection ? 'remove' : 'add') +'">'+ (res.data.collection ? '取消收藏' : '收藏') +'</span>');
      });
      */
    //}

    if (layui.cache.innerName == 'topic_add' && layui.cache.tid == 0){
      //no need for async load
      return;
    } 
    if (layui.cache.innerName == 'topic_index'){
      //no need for async load
      return;
    }
    if (layui.cache.innerName == 'more_list'){
      //列表
      if (layui.cache.type == 1){
        //resource
        //取这个主题的资源列表
        loadTopicResources(layui.cache.currPage,true);

      }else if (layui.cache.type == 2){
        //notes
        //取这个主题的笔记列表
        loadTopicNotes(layui.cache.currPage, true);
      }
      return;
    }
    if (layui.cache.innerName == 'topic_task'){
      //no need for async load
      if (layui.cache.tid && layui.cache.taskId){
        //loadTopicNoteToCache
        var note = loadTopicNoteToCache(layui.cache.taskId);
        if(note && (JSON.stringify(note).localeCompare("{}") != 0)){        
          var tidInner = note['id'];
          var title = note['task_title']; 
          var content = note['task_content'];
          $("#taskTitle").val(title);
          $("#taskContent").val(content);
          //修改URL为编辑的url
          $("#EDIT_FORM_ACTION").attr("action",  "/api.v2/topictask/edit?id="+tidInner);
        }
      }
      return;
    }

    if (layui.cache.innerName == 'note_task'){
      //no need for async load
      if (layui.cache.tid && layui.cache.taskId){
        //loadTopicNoteToCache 
      }
      return;
    }

    if (layui.cache.tid == 0){

      layer.msg("页面错误", {shift: 6});
    
      setTimeout(function () {
                            location.href = "../"; 
                          }, 1000);  
      return;
    }

    //取这个主题的任务列表

    fly.json('/api.v2/topictask/list', {topic_id:layui.cache.tid}, function(res){
      //console.log("topictask res:"+JSON.stringify(res));  
           
      if (!res['data']['list']){
        return;
      }
      var task_list = "";
      taskCacheList = res['data']['list'];

      for(let i = 0; i< res['data']['list'].length; i++){
 
           
        var html = [
          '<li class="layui-timeline-item">'
          ,'<i class="layui-icon layui-timeline-axis"></i>'
          ].join("");
        var edit= [
         ,'<ul class="layui-layout-right"> '
         ,' <button type="button" class="layui-btn layui-btn-primary" id="onClickEditTask"  taskId="TASK_ID">编辑</button>'
        ,'</ul>' ].join("");
        var end=[
          ,'<div class="layui-timeline-content layui-text">'
           ,'<h3 class="layui-timeline-title">TIME</h3>'
          ,'<p><a href="#" id="onClickTaskItem" taskId="TASK_ID">TITLE</a></p>'
          ,'</div>'
        ,'</li>'
        ].join(""); 
          //转义
          content = fly.htmlEscape(res['data']['list'][i]['task_title']).trim();
        edit = edit.replace(/TASK_ID/g, i);
        end = end.replace(/TIME/g, res['data']['list'][i]['create_time']);
        end = end.replace(/TASK_ID/g, i); 
        end = end.replace(/TITLE/g, content); 
        // html = html.replace(/CONTENT/g, res['data']['list'][i]['resource_content']); 
        task_list+= html;

        //比較當前topic的note id和當前的
        var id = new Number(layui.cache.uid)
        if (id == layui.cache.user.uid){ 
          task_list+= edit; 
        }
        task_list+= end;
      } 

      $("#task_list").append(task_list);
    });

    //评论
    loadComment(layui.cache.currPage);

    //取这个主题的资源列表
    loadTopicResources(1, false);

    //取这个主题的笔记列表
    loadTopicNotes(1, false);
 
    //如果已经登录，判断该用户是否已经加入该主题
    if (layui.cache.user.uid != -1){

      if (layui.cache.innerName  != "detail"){
        // $("#members_count").html("1");
        return;
      }


      fly.json('/api.v2/topicmembers/list ', {topic_id:layui.cache.tid}, function(res){
        //console.log("checkJoin res:"+JSON.stringify(res)); 
        if (!res['data'] || !res['data']['list']){
          $("#members_join_topic_no_member").removeClass("layui-hide");
          // $("#members_count").html("1");
          return;
        }
            
        //更新加入人数 
        // $("#members_count").html(res['data']['list'].length + 1);
        for(let i = 0; i< res['data']['list'].length; i++){
          var html = [' <dd>'
                ,'<a href="/user/home.html?uid=USER_ID">'
                ,'<img src="USER_IMAGE_RESOURCE" alt="TOPIC_OWNER_NAME"><cite>TOPIC_OWNER_NAME</cite> '
                ,'</a>'
                ,'</dd>'
          ].join('');
          var name = res['data']['list'][i]['name'];
          var uid = res['data']['list'][i]['id'];
          var avatar = res['data']['list'][i]['ext']; 
          html = html.replace(/USER_ID/g, uid);
          html = html.replace(/TOPIC_OWNER_NAME/g, name); 

          if (avatar && avatar.length > 0){
            html = html.replace(/USER_IMAGE_RESOURCE/g, avatar);

          }else{

          html = html.replace(/USER_IMAGE_RESOURCE/g, "/res/images/avatar/default.png");
          }

          $("#members_join_topic").append(html);
        }

        if (res['data']['list'].length == 0){
          $("#members_join_topic_no_member").removeClass("layui-hide");
        }
      });

      if (layui.cache.topic_owner){
        //是owner就不需要判断是否加入该组
        return;
      }

      fly.json('/api.v2/topicmembers/checkJoin ', {topic_id:layui.cache.tid}, function(res){
        //console.log("checkJoin res:"+JSON.stringify(res));
        if (layui.cache.topic_owner){
          //是owner就不需要判断是否加入该组
          return;
        }
        if (res['data']['result'] == 1){
          //标识已经加入这个topic 
          $("#LAY_join_topic").removeClass("layui-btn-danger");
          $("#LAY_join_topic").addClass("layui-btn-disabled");
          $("#LAY_quit_topic").removeClass("layui-btn-disabled");
          $("#LAY_quit_topic").addClass("layui-btn-danger"); 
          $("#LAY_add_notes").addClass("layui-btn-danger"); 
          $("#LAY_add_notes").removeClass("layui-btn-disabled");

        }else{
          //标识没有加入
          $("#LAY_quit_topic").removeClass("layui-btn-danger");
          $("#LAY_quit_topic").addClass("layui-btn-disabled");
          $("#LAY_join_topic").removeClass("layui-btn-disabled");
          $("#LAY_join_topic").addClass("layui-btn-danger");  
          
          $("#LAY_add_notes").addClass("layui-btn-disabled"); 
          $("#LAY_add_notes").removeClass("layui-btn-danger");
        }
      });
    }
  }();

  //解答操作
  gather.jiedaActive = {
    zan: function(li){ //赞
      var othis = $(this), ok = othis.hasClass('zanok');
      fly.json('/api/jieda-zan/', {
        ok: ok
        ,id: li.data('id')
      }, function(res){
        if(res.status === 0){
          var zans = othis.find('em').html()|0;
          othis[ok ? 'removeClass' : 'addClass']('zanok');
          othis.find('em').html(ok ? (--zans) : (++zans));
        } else {
          layer.msg(res.msg);
        }
      });
    }
    ,reply: function(li){ //回复
      var val = dom.content.val();
      var aite = '@'+ li.find('.fly-detail-user cite').text().replace(/\s/g, '');
      dom.content.focus()
      if(val.indexOf(aite) !== -1) return;
      dom.content.val(aite +' ' + val);
    }
    ,accept: function(li){ //采纳
      var othis = $(this);
      layer.confirm('是否采纳该回答为最佳答案？', function(index){
        layer.close(index);
        fly.json('/api/jieda-accept/', {
          id: li.data('id')
        }, function(res){
          if(res.status === 0){
            $('.jieda-accept').remove();
            li.addClass('jieda-daan');
            li.find('.detail-about').append('<i class="iconfont icon-caina" title="最佳答案"></i>');
          } else {
            layer.msg(res.msg);
          }
        });
      });
    }
    ,edit: function(li){ //编辑
      fly.json('/jie/getDa/', {
        id: li.data('id')
      }, function(res){
        var data = res.rows;
        layer.prompt({
          formType: 2
          ,value: data.content
          ,maxlength: 100000
          ,title: '编辑回帖'
          ,area: ['728px', '300px']
          ,success: function(layero){
            fly.layEditor({
              elem: layero.find('textarea')
            });
          }
        }, function(value, index){
          fly.json('/jie/updateDa/', {
            id: li.data('id')
            ,content: value
          }, function(res){
            layer.close(index);
            li.find('.detail-body').html(fly.content(value));
          });
        });
      });
    }
    ,del: function(li){ //删除
      layer.confirm('确认删除该回答么？', function(index){
        layer.close(index);
        fly.json('/api/jieda-delete/', {
          id: li.data('id')
        }, function(res){
          if(res.status === 0){
            var count = dom.jiedaCount.text()|0;
            dom.jiedaCount.html(--count);
            li.remove();
            //如果删除了最佳答案
            if(li.hasClass('jieda-daan')){
              $('.jie-status').removeClass('jie-status-ok').text('求解中');
            }
          } else {
            layer.msg(res.msg);
          }
        });
      });    
    }
  };

  $('.jieda-reply span').on('click', function(){
    var othis = $(this), type = othis.attr('type');
    gather.jiedaActive[type].call(this, othis.parents('li'));
  });


  //定位分页
  if(/\/page\//.test(location.href) && !location.hash){
    var replyTop = $('#flyReply').offset().top - 80;
    $('html,body').scrollTop(replyTop);
  }

  exports('jie', null);
});