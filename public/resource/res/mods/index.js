/**

 @Name: Fly社区主入口

 */
 

layui.define(['layer', 'laytpl', 'form', 'element', 'upload', 'util', 'laydate'], function(exports){
  
  var $ = layui.jquery
  ,layer = layui.layer
  ,laytpl = layui.laytpl
  ,form = layui.form
  ,element = layui.element
  ,upload = layui.upload
  ,util = layui.util
  ,device = layui.device()
  ,TOPIC_CATGORY_CACHE=null

  ,DISABLED = 'layui-btn-disabled';
  
  //阻止IE7以下访问
  if(device.ie && device.ie < 8){
    layer.alert('如果您非得使用 IE 浏览器访问Fly社区，那么请使用 IE8+');
  }
  
  layui.focusInsert = function(obj, str){
    var result, val = obj.value;
    obj.focus();
    if(document.selection){ //ie
      result = document.selection.createRange(); 
      document.selection.empty(); 
      result.text = str; 
    } else {
      result = [val.substring(0, obj.selectionStart), str, val.substr(obj.selectionEnd)];
      obj.focus();
      obj.value = result.join('');
    }
  };

  refreshVerify = function(){ 
      fly.json("/captcha/get/img", {}, function(res){ 
        if(res.code == 0){ 
            //console.log("res.data:"+res['data'].idKeyC);//res.base64stringC
            $("#idKeyC").val(res['data'].idKeyC);
            $("#verCode").removeClass("layui-hide");
            $("#verCodeImg").attr("src",res['data'].base64stringC);
        };
      });  
  }

  //重新渲染表单
  //重新渲染表单
  renderForm = function(){
    layui.use('form', function(){
     var form = layui.form;//高版本建议把括号去掉，有的低版本，需要加()
     form.render();
    });
   } 
  updateDefaultCatForm =  function(res){
    
    if (!res['data']){
      return;
    }
    ///div>/pub6/topic/type 
    var header =  '<ul class="fly-panel-main fly-list-static">';
    var end = '</ul>';
    var content = header;
    for(let i = 0; i< res['data'].length; i++){ 

      var html= [
            '<li>',
              '<a href="/jie/index.html?type=TYPE">CAT_NAME</a>',
            '</li> '  
      ].join("");

      if ($("#topicType")){
        $("#topicType").append("<option value="+res['data'][i]['id']+">"+res['data'][i]['name']+"</option>"); 
      }
    
      html = html.replace(/TYPE/g, res['data'][i]['id']); 
      html = html.replace(/CAT_NAME/g, res['data'][i]['name']); 
      content+= html;
    }
    content += end;
    if ($("#catory_list")){
        $("#catory_list").html(content);
    }
    renderForm();//表单重新渲染，要不然添加完显示不出来新的option 
  }

  saveToCache = function(name, value){ 
      var saveData = {key:name, value:value};
      layui.sessionData('pub6_com_user_cache_cat', saveData);//把AJSON对象存储为字符串
  }
  loadFromCache = function (name) { 

      var cacheUser = layui.sessionData('pub6_com_user_cache_cat');
      return cacheUser[name] || "{}";
  }
  saveTopicToCache = function(name, value){ 
      var saveData = {key:name, value:value};
      layui.sessionData('pub6_com_user_cache_topic', saveData);//把AJSON对象存储为字符串
  }
  loadTopicFromCache = function (name) { 

      var cacheUser = layui.sessionData('pub6_com_user_cache_topic');
      return cacheUser[name] || {};
  }

  var CATEGORY_CACHE_TAG = "category"; 
  refreshCatgory = function(){  
      var cacheUser = loadFromCache(CATEGORY_CACHE_TAG);
      if(cacheUser && cacheUser.localeCompare("{}") != 0){ 
        //var category = cacheUser[CATEGORY_CACHE_TAG] ;
        res =  JSON.parse(cacheUser);//JSON.stringify(cacheUser);
        updateDefaultCatForm(res);
        return;
      }
      //layui.data('user_cache', saveData);//把AJSON对象存储为字符串
      fly.json("/pub6/topic/category", {}, function(res){ 
        if(res.code == 0){ 
            ////console.log("refreshCatgory res.data:"+JSON.stringify(res));//res.base64stringC 
                // 
            saveToCache(CATEGORY_CACHE_TAG, JSON.stringify(res));
            
            updateDefaultCatForm(res);
        };
      });  
  } 
  addTopicTopList = function(res, contentId, pageId, currPage){ 
    if (res['data']['total'] == 0 || res['data']['list']==null || res['data']['list'].length == 0){
      //
      $(contentId).html('<div class="fly-none">暂时没有相关数据</div>');
      return;
    }

    var content1 = "";
    for(let i = 0; i< res['data']['list'].length; i++){

      var html = ['<ul class="fly-list">'
          ,'<li>'
              ,'<a alt= "TOPIC_OWNER_NAME" href="/user/home.html?uid=USER_ID" class="fly-avatar">'
                          ,'<img src="USER_OWNER_SRC" alt="TOPIC_OWNER_NAME">'
              ,'</a>'
              ,'<h2>'
                ,'<a href="/jie/index.html?type=CAT_TYPE_ID" class="layui-badge">TOPIC_CAT</a>'
                ,'<a href="/jie/detail.html?tid=TOPIC_ID&uid=USER_OWNER_ID">TOPIC_TITLE</a>'
              ,'</h2>'
              ,'<div class="fly-list-info">'
                ,'<a href="/user/home.html?uid=USER_ID" link>'
                  ,'<cite>TOPIC_OWNER_NAME</cite>'
                  ,'<!--'
                  ,'<i class="iconfont icon-renzheng" title="认证信息：XXX"></i>'
                  ,'<i class="layui-badge fly-badge-vip">VIP3</i>'
                  ,'-->'
                ,'</a>'
                ,'<span>TOPIC_CREATE_TIME</span>' 
                ,'<span class="fly-list-kiss layui-hide-xs" title="成员"><i class="iconfont icon-renshu" title="回答"></i> MEMEBER_NUMBERS</span>'
                ,'<!--<span class="layui-badge fly-badge-accept layui-hide-xs">已结</span>-->'
                ,'<span class="fly-list-nums"> '
                ,'<i class="iconfont"  title="人气">&#xe60b;</i> SEE_NUMBER'
                ,'</span>'
              ,'</div> '
            ,'</li>'
        ,'</ul>'].join('');
 
      var name = res['data']['list'][i]['name'];
      var uid = res['data']['list'][i]['create_owner_id'];
                var avatar = res['data']['list'][i]['ext'];
      var tid = res['data']['list'][i]['id'];
      var title = res['data']['list'][i]['topic_name'];
      //转义
      title = fly.htmlEscape(title).trim();

      var catgory_type = res['data']['list'][i]['topic_type'];
      var topic_content = res['data']['list'][i]['topic_content'];
      var create_time = res['data']['list'][i]['create_time'];
      var see_num = res['data']['list'][i]['see_num'];
      var mem_number = res['data']['list'][i]['members_number'];

      var category_type_id = catgory_type;
      var category = loadFromCache(CATEGORY_CACHE_TAG);
      if(category && category.localeCompare("{}") != 0){ 
        var categoryres =  JSON.parse(category);//JSON.stringify(cacheUser); 
        for(let p = 0; p< categoryres['data'].length; p++){   
          if(categoryres['data'][p]['id'] == catgory_type){
            catgory_type = categoryres['data'][p]['name'];
            break;
          }
        } 
      }

      html = html.replace(/SEE_NUMBER/g, see_num);
      html = html.replace(/MEMEBER_NUMBERS/g, mem_number);
      html = html.replace(/TOPIC_ID/g, tid);
      html = html.replace(/USER_ID/g, uid);
      html = html.replace(/USER_OWNER_ID/g, uid);
      html = html.replace(/TOPIC_OWNER_NAME/g, name);
      if (avatar && avatar.length > 0){
        html = html.replace(/USER_OWNER_SRC/g, avatar);

      }else{

        html = html.replace(/USER_OWNER_SRC/g, "/res/images/avatar/default.png");
      }
      html = html.replace(/TOPIC_CAT/g, catgory_type);
      html = html.replace(/CAT_TYPE_ID/g, category_type_id);
      html = html.replace(/TOPIC_TITLE/g, title);
      html = html.replace(/TOPIC_CONTENT/g, topic_content);
      html = html.replace(/TOPIC_CREATE_TIME/g, create_time); 
      content1+= html;
    }
    $(contentId).html(content1); 

    //for page css

    if (layui.cache.innerName == "topic_index"){ 
      var pageSize = 20;
      //topic_page
      var pageTotal = ((parseInt(res['data']['total']/pageSize)) > 0)?( parseInt(res['data']['total']/pageSize )+ ((res['data']['total']%pageSize)>0?1:0) ):1;
      var page_html=[
          '<div class="laypage-main">',
          '<a href="/jie/index.html?page=1" class="laypage-next">首页</a>',
          '<span class="laypage-curr">1</span>', 
          '<span>…</span>',
          '<a href="/jie/index.html?page='+pageTotal+'" class="laypage-last" title="尾页">尾页</a>',
          '<a href="/jie/index.html?page='+currPage+1+'" class="laypage-next">下一页</a>',
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
          pageRealHtml += '<a href="/jie/index.html?page='+j+'" >'+j+'</a>'; 
        }
      }//end for
      pageRealHtml += page_html[6]; 
      $(pageId).html(pageRealHtml);
    }//end if 

  }

  saveTopicAllToCache = function(name, value){ 
      var saveData = {key:name, value:value};
      layui.sessionData('pub6_com_user_cache_topic_all', saveData);//把AJSON对象存储为字符串
  } 
  getTopicTopList = function(){ 

      var topic_top = loadTopicFromCache('topic_top');
      if(JSON.stringify(topic_top) != '{}'){ 
        res =  JSON.parse(topic_top);

        addTopicTopList(res, "#topic-list-top", "#topic_page", 1);
        return;
      }

      fly.json("/pub6/topic/all", {"order_num":1}, function(res){ 
        if(res.code == 0){ 
            //console.log("getTopicTopList res.total:"+res['data']['total'] );//+ ", res:"+ JSON.stringify(res)
            saveTopicToCache("topic_top", JSON.stringify(res));

            addTopicTopList(res, "#topic-list-top", "#topic_page", 1);
        };
      });  
  }

  getTopicAllList = function(currPage){ 
      var pageSize = 20;

      if (layui.cache.innerName == 'topic_index' && layui.cache.topictype != 0){ 
          
        fly.json("/pub6/topic/type", {"topic_type":layui.cache.topictype, "pageSize":pageSize, "order_num":0, "pageNum":currPage}, function(res){ 
          if(res.code == 0){ 
              //console.log("getTopicTopList res.total:"+res['data']['total'] );//+ ", res:"+ JSON.stringify(res)
              //saveTopicToCache("topic_top", JSON.stringify(res));

            addTopicTopList(res, "#topic-list", "#topic_page", currPage);
          };
        });  
        return;
      }

      fly.json("/pub6/topic/all", {"pageSize":pageSize, "order_num":0, "pageNum":currPage}, function(res){ 
        if(res.code == 0){ 
            ////console.log("getTopicAllList res.total:" +res['data']['total']  +", res:"+ JSON.stringify(res));//res.base64stringC   +", res:"+ JSON.stringify(res)
            saveTopicAllToCache("alllist_1", JSON.stringify(res));

            addTopicTopList(res, "#topic-list", "#topic_page", currPage);
        //     for(let i = 0; i< res['data']['list'].length; i++){
        //         var html = ['<ul class="fly-list">'
        //             ,'<li>'
        //                 ,'<a alt= "TOPIC_OWNER_NAME" href="/user/home.html?uid=USER_ID" class="fly-avatar">'
        //                   ,'<img src="USER_OWNER_SRC" alt="TOPIC_OWNER_NAME">'
        //                 ,'</a>'
        //                 ,'<h2>'
        //                   ,'<a class="layui-badge">TOPIC_CAT</a>'
        //                   ,'<a href="/jie/detail.html?tid=TOPIC_ID&uid=USER_ID">TOPIC_TITLE</a>'
        //                 ,'</h2>'
        //                 ,'<div class="fly-list-info">'
        //                   ,'<a href="/user/home.html?uid=USER_ID" link>'
        //                     ,'<cite>TOPIC_OWNER_NAME</cite>' 
        //                   ,'</a>'
        //                   ,'<span>TOPIC_CREATE_TIME</span>' 
        //                   ,'<span class="fly-list-kiss layui-hide-xs" title="加入"><i class="iconfont icon-renshu" title="回答"></i> MEMEBER_NUMBERS</span>'
        //                   ,'<!--<span class="layui-badge fly-badge-accept layui-hide-xs">已结</span>-->'
        //                   ,'<span class="fly-list-nums"> '
        //                     ,'<i class="iconfont"  title="人气">&#xe60b;</i> SEE_NUMBER'
        //                   ,'</span>'
        //                 ,'</div> '
        //               ,'</li>' 
        //           ,'</ul>'].join('');
        //         if (res['data']['list'][i] && res['data']['list'][i]['name']){
        //           var name = res['data']['list'][i]['name'];
        //           var uid = res['data']['list'][i]['create_owner_id'];
        //           var avatar = res['data']['list'][i]['ext']; 
        //         }
        //         var tid = res['data']['list'][i]['id'];
        //         var title = res['data']['list'][i]['topic_name'];
        //         //转义
        //         title = fly.htmlEscape(title).trim();

        //         var catgory_type = res['data']['list'][i]['topic_type'];
        //         var topic_content = res['data']['list'][i]['topic_content'];
        //         var create_time = res['data']['list'][i]['create_time'];
        //         var see_num = res['data']['list'][i]['see_num'];
        //         var mem_number = res['data']['list'][i]['members_number'];

        //         var category = loadFromCache(CATEGORY_CACHE_TAG);
        //         if(category && category.localeCompare("{}") != 0){ 
        //           var categoryres =  JSON.parse(category);//JSON.stringify(cacheUser);
        //           for(let p = 0; p< categoryres['data'].length; p++){   
        //             if(categoryres['data'][p]['id'] == catgory_type){
        //               catgory_type = categoryres['data'][p]['name'];
        //               break;
        //             }
        //           }
        //         }
                
        //         html = html.replace(/TOPIC_ID/g, tid);
        //         html = html.replace(/USER_ID/g, uid);
        //         html = html.replace(/SEE_NUMBER/g, see_num);
        //         html = html.replace(/MEMEBER_NUMBERS/g, mem_number);
        //         if (avatar && avatar.length > 0){
        //           html = html.replace(/USER_OWNER_SRC/g, avatar);

        //         }else{

        //         html = html.replace(/USER_OWNER_SRC/g, "/res/images/avatar/default.png");
        //         }
        //         html = html.replace(/TOPIC_OWNER_NAME/g, name);
        //         html = html.replace(/TOPIC_CAT/g, catgory_type);
        //         html = html.replace(/TOPIC_TITLE/g, title);
        //         html = html.replace(/TOPIC_CONTENT/g, topic_content);
        //         html = html.replace(/TOPIC_CREATE_TIME/g, create_time);
        //         $("#topic-list").append(html);
        //       }
        }

      });  
  }


  //数字前置补零
  layui.laytpl.digit = function(num, length, end){
    var str = '';
    num = String(num);
    length = length || 2;
    for(var i = num.length; i < length; i++){
      str += '0';
    }
    return num < Math.pow(10, length) ? str + (num|0) : num;
  };

  var layerIndex = 0;
  function showLoad() {
 
    layerIndex =  layer.msg('拼命执行中...', {icon: 16,shade: [0.5, '#f5f5f5'],scrollbar: false,offset: 'auto', time:100000});

  }

  function closeLoad() {
    if (layerIndex == 0){
      return;
    }
    layer.close(layerIndex);
    layerIndex = 0;
  } 

  var fly = {

    //Ajax
    json: function(url, data, success, options){
      var that = this, type = typeof data === 'function';
      
      if(type){
        options = success
        success = data;
        data = {};
      }

      options = options || {}; 
      var authorilize = "";
      if (layui.cache.user['token'] && layui.cache.user['token'].length > 0){
        authorilize = "Bearer "+layui.cache.user['token'] ;
        //console.log("url:"+url +",authorilize:"+authorilize);
      }else{
         //console.log("url:"+url);
      }
      //增加等待框
      showLoad();


      return $.ajax({
        type: options.type || 'post',
        dataType: options.dataType || 'json',
        data: data,
        headers: {
            Accept: "application/json; charset=utf-8",
            Authorization:authorilize  
        },
        url: url,
        success: function(res){
          //console.log("code:"+res.code);

          if(res.code == 0) {
            //增加等待框,关闭
            closeLoad();
            ////console.log("  token:"+ res.data.token);
            success && success(res);
          } else { 
            //增加等待框,关闭
            closeLoad();

            if (layui.cache.page == 'user'  || layui.cache.innerName == "topic_add"){
              refreshVerify();
            }
            if (res.code == -401){
              //登录失败
              layui.cache.user = {
                username: ''
                ,uid: -1
                ,token:'' 
              }; 
              if (layui.cache.innerName != "login"){
                //判断异常原因
                layer.msg('请求异常，请先登录:'+ (res.msg || res.code), {shift: 6});
                setTimeout(function () {
                                    location.href = "/user/login.html"; 
                                  }, 2000);   
              }
            } else{ 
              layer.msg('请求异常:'+res.msg || res.code, {shift: 6}); 
            }
            options.error && options.error();
          }
        }, error: function(e){ 


          setTimeout(function () {
            //增加等待框,关闭
            closeLoad();
          }, 1000);   
          layer.msg('请求异常:'+e.status, {shift: 6});  
        }
      });
    }

    //计算字符长度
    ,charLen: function(val){
      var arr = val.split(''), len = 0;
      for(var i = 0; i <  val.length ; i++){
        arr[i].charCodeAt(0) < 299 ? len++ : len += 2;
      }
      return len;
    }
    
    ,form: {}

    //简易编辑器
    ,layEditor: function(options){
      var html = ['<div class="layui-unselect fly-edit">'
        ,'<span type="face" title="插入表情"><i class="iconfont icon-yxj-expression" style="top: 1px;"></i></span>'
        ,'<span type="picture" title="插入图片：img[src]"><i class="iconfont icon-tupian"></i></span>'
        ,'<span type="href" title="超链接格式：a(href)[text]"><i class="iconfont icon-lianjie"></i></span>'
        ,'<span type="code" title="插入代码或引用"><i class="iconfont icon-emwdaima" style="top: 1px;"></i></span>'
        ,'<span type="hr" title="插入水平线">hr</span>'
        ,'<span type="yulan" title="预览"><i class="iconfont icon-yulan1"></i></span>'
      ,'</div>'].join('');

      var log = {}, mod = {
        face: function(editor, self){ //插入表情
          var str = '', ul, face = fly.faces;
          for(var key in face){
            str += '<li title="'+ key +'"><img src="'+ face[key] +'"></li>';
          }
          str = '<ul id="LAY-editface" class="layui-clear">'+ str +'</ul>';
          layer.tips(str, self, {
            tips: 3
            ,time: 0
            ,skin: 'layui-edit-face'
          });
          $(document).on('click', function(){
            layer.closeAll('tips');
          });
          $('#LAY-editface li').on('click', function(){
            var title = $(this).attr('title') + ' ';
            layui.focusInsert(editor[0], 'face' + title);
          });
        }
        ,picture: function(editor){ //插入图片
          layer.open({
            type: 1
            ,id: 'fly-jie-upload'
            ,title: '插入图片'
            ,area: 'auto'
            ,shade: false
            ,area: '465px'
            ,fixed: false
            ,offset: [
              editor.offset().top - $(window).scrollTop() + 'px'
              ,editor.offset().left + 'px'
            ]
            ,skin: 'layui-layer-border'
            ,content: ['<ul class="layui-form layui-form-pane" style="margin: 20px;">'
              ,'<li class="layui-form-item">'
                ,'<label class="layui-form-label">URL</label>'
                ,'<div class="layui-input-block">'
                    ,'<input required name="image" placeholder="支持直接粘贴远程图片地址" value="" class="layui-input" >'
                  ,'</div>' 
              ,'</li>'
              ,'<li class="layui-form-item" style="text-align: center;">'
                ,'<button type="button" lay-submit lay-filter="uploadImages" class="layui-btn">确认</button>'
              ,'</li>'
            ,'</ul>'].join('')
            ,success: function(layero, index){
              var image =  layero.find('input[name="image"]');

              //执行上传实例
              upload.render({
                elem: '#uploadImg'
                ,url: '/api/upload/'
                ,size: 200
                ,done: function(res){
                  if(res.code == 0){
                    image.val(res.url);
                  } else {
                    layer.msg(res.msg, {icon: 5});
                  }
                }
              });
              
              form.on('submit(uploadImages)', function(data){
                var field = data.field;
                if(!field.image) return image.focus();
                layui.focusInsert(editor[0], 'img['+ field.image + '] ');
                layer.close(index);
              });
            }
          });
        }
        ,href: function(editor){ //超链接
          layer.prompt({
            title: '请输入合法链接'
            ,shade: false
            ,fixed: false
            ,id: 'LAY_flyedit_href'
            ,offset: [
              editor.offset().top - $(window).scrollTop() + 'px'
              ,editor.offset().left + 'px'
            ]
          }, function(val, index, elem){
            if(!/^http(s*):\/\/[\S]/.test(val)){
              layer.tips('这根本不是个链接，不要骗我。', elem, {tips:1})
              return;
            }
            layui.focusInsert(editor[0], ' a('+ val +')['+ val + '] ');
            layer.close(index);
          });
        }
        ,code: function(editor){ //插入代码
          layer.prompt({
            title: '请贴入代码或任意文本'
            ,formType: 2
            ,maxlength: 10000
            ,shade: false
            ,id: 'LAY_flyedit_code'
            ,area: ['800px', '360px']
          }, function(val, index, elem){
            layui.focusInsert(editor[0], '[pre]\n'+ val + '\n[/pre]');
            layer.close(index);
          });
        }
        ,hr: function(editor){ //插入水平分割线
          layui.focusInsert(editor[0], '[hr]');
        }
        ,yulan: function(editor){ //预览
          var content = editor.val();
          
          content = /^\{html\}/.test(content) 
            ? content.replace(/^\{html\}/, '')
          : fly.content(content);

          layer.open({
            type: 1
            ,title: '预览'
            ,shade: false
            ,area: ['100%', '100%']
            ,scrollbar: false
            ,content: '<div class="detail-body" style="margin:20px;">'+ content +'</div>'
          });
        }
      };
      
      layui.use('face', function(face){
        options = options || {};
        fly.faces = face;
        $(options.elem).each(function(index){
          var that = this, othis = $(that), parent = othis.parent();
          parent.prepend(html);
          parent.find('.fly-edit span').on('click', function(event){
            var type = $(this).attr('type');
            mod[type].call(that, othis, this);
            if(type === 'face'){
              event.stopPropagation()
            }
          });
        });
      });
      
    }

    ,escape: function(html){
      return String(html||'').replace(/&(?!#?[a-zA-Z0-9]+;)/g, '&amp;')
      .replace(/</g, '&lt;').replace(/>/g, '&gt;').replace(/'/g, '&#39;').replace(/"/g, '&quot;');
    }

    //内容转义
    ,content: function(content){
      //支持的html标签
      var html = function(end){
        return new RegExp('\\n*\\['+ (end||'') +'(pre|hr|div|span|p|table|thead|th|tbody|tr|td|ul|li|ol|li|dl|dt|dd|h2|h3|h4|h5)([\\s\\S]*?)\\]\\n*', 'g');
      };
      content = fly.escape(content||'') //XSS
      .replace(/img\[([^\s]+?)\]/g, function(img){  //转义图片
        return '<img src="' + img.replace(/(^img\[)|(\]$)/g, '') + '">';
      }).replace(/@(\S+)(\s+?|$)/g, '@<a href="javascript:;" class="fly-aite">$1</a>$2') //转义@
      .replace(/face\[([^\s\[\]]+?)\]/g, function(face){  //转义表情
        var alt = face.replace(/^face/g, '');
        return '<img alt="'+ alt +'" title="'+ alt +'" src="' + fly.faces[alt] + '">';
      }).replace(/a\([\s\S]+?\)\[[\s\S]*?\]/g, function(str){ //转义链接
        var href = (str.match(/a\(([\s\S]+?)\)\[/)||[])[1];
        var text = (str.match(/\)\[([\s\S]*?)\]/)||[])[1];
        if(!href) return str;
        var rel =  /^(http(s)*:\/\/)\b(?!(\w+\.)*(sentsin.com|layui.com))\b/.test(href.replace(/\s/g, ''));
        return '<a href="'+ href +'" target="_blank"'+ (rel ? ' rel="nofollow"' : '') +'>'+ (text||href) +'</a>';
      })
      .replace(html(), '\<$1 $2\>').replace(html('/'), '\</$1\>') //转移HTML代码
      .replace(/\n/g, '<br>') //转义换行   
      return content;
    }
    
    //新消息通知
    ,newmsg: function(){
      var elemUser = $('.fly-nav-user');
      if(layui.cache.user && layui.cache.user.uid !== -1 && elemUser[0]){
        fly.json('/message/nums/', {
          _: new Date().getTime()
        }, function(res){
          if(res.code === 0 && res.count > 0){
            var msg = $('<a class="fly-nav-msg" href="javascript:;">'+ res.count +'</a>');
            elemUser.append(msg);
            msg.on('click', function(){
              fly.json('/message/read', {}, function(res){
                if(res.code === 0){
                  location.href = '/user/message/';
                }
              });
            });
            layer.tips('你有 '+ res.count +' 条未读消息', msg, {
              tips: 3
              ,tipsMore: true
              ,fixed: true
            });
            msg.on('mouseenter', function(){
              layer.closeAll('tips');
            })
          }
        });
      }
      return arguments.callee;
    }
    , saveCurrUrlToCache :function(value){ 
        var saveData = {key:"pub6_curr_url", value:value};
        layui.sessionData('pub6_com_user_cache__curr_url', saveData);//把AJSON对象存储为字符串
    }
    , loadFromCurrUrlToCache :function () { 

        var cacheUser = layui.sessionData('pub6_com_user_cache__curr_url');
        return cacheUser["pub6_curr_url"] || "../";
    }

    
  };

  //签到
  var tplSignin = ['{{# if(d.signed){ }}'
    ,'<button class="layui-btn layui-btn-disabled">今日已签到</button>'
    ,'<span>获得了<cite>{{ d.experience }}</cite>飞吻</span>'
  ,'{{# } else { }}'
    ,'<button class="layui-btn layui-btn-danger" id="LAY_signin">今日签到</button>'
    ,'<span>可获得<cite>{{ d.experience }}</cite>飞吻</span>'
  ,'{{# } }}'].join('')
  ,tplSigninDay = '已连续签到<cite>{{ d.days }}</cite>天'

  ,signRender = function(data){
    laytpl(tplSignin).render(data, function(html){
      elemSigninMain.html(html);
    });
    laytpl(tplSigninDay).render(data, function(html){
      elemSigninDays.html(html);
    });
  }

  ,elemSigninHelp = $('#LAY_signinHelp')
  ,elemSigninTop = $('#LAY_signinTop')
  ,elemSigninMain = $('.fly-signin-main')
  ,elemSigninDays = $('.fly-signin-days');
  
  if(elemSigninMain[0]){
    /*
    fly.json('/sign/code', function(res){
      if(!res.data) return;
      signRender.token = res.data.token;
      signRender(res.data);
    });
    */
  }
  // $('body').on('click', '#LAY_signin', function(){
  //   var othis = $(this);
  //   if(othis.hasClass(DISABLED)) return;

  //   fly.json('/sign/in', {
  //     token: signRender.token || 1
  //   }, function(res){
  //     signRender(res.data);
  //   }, {
  //     error: function(){
  //       othis.removeClass(DISABLED);
  //     }
  //   });

  //   othis.addClass(DISABLED);
  // });
   //刷新二维码
   $('body').on('click', "#verCodeImg", function(){
        //refreshVerify
        refreshVerify();
   });

  //登出
  $('body').on('click', '#loginOut', function(){
    var othis = $(this);
    if(othis.hasClass(DISABLED)) return;

    //console.log("token:"+layui.cache.user['token']);

    fly.json('/pub6Login/logout ', {
      token: layui.cache.user['token'] || 1
    }, function(res){
            //
      var saveData = {
                key:"userinfo", remove: true};
      layui.data('user_cache', saveData);//把AJSON对象存储为字符串
      layui.sessionData('user_cache', saveData);//

      location.href = "/";

      othis.removeClass(DISABLED);
      signRender(res.data);
      return;
    }, {
      error: function(){
        othis.removeClass(DISABLED);
      }
    });

    othis.addClass(DISABLED);
  });


  //签到说明
  elemSigninHelp.on('click', function(){
    layer.open({
      type: 1
      ,title: '签到说明'
      ,area: '300px'
      ,shade: 0.8
      ,shadeClose: true
      ,content: ['<div class="layui-text" style="padding: 20px;">'
        ,'<blockquote class="layui-elem-quote">“签到”可获得社区飞吻，规则如下</blockquote>'
        ,'<table class="layui-table">'
          ,'<thead>'
            ,'<tr><th>连续签到天数</th><th>每天可获飞吻</th></tr>'
          ,'</thead>'
          ,'<tbody>'
            ,'<tr><td>＜5</td><td>5</td></tr>'
            ,'<tr><td>≥5</td><td>10</td></tr>'
            ,'<tr><td>≥15</td><td>15</td></tr>'
            ,'<tr><td>≥30</td><td>20</td></tr>'
          ,'</tbody>'
        ,'</table>'
        ,'<ul>'
          ,'<li>中间若有间隔，则连续天数重新计算</li>'
          ,'<li style="color: #FF5722;">不可利用程序自动签到，否则飞吻清零</li>'
        ,'</ul>'
      ,'</div>'].join('')
    });
  });

  //签到活跃榜
  var tplSigninTop = ['{{# layui.each(d.data, function(index, item){ }}'
    ,'<li>'
      ,'<a href="/u/{{item.uid}}" target="_blank">'
        ,'<img src="{{item.user.avatar}}">'
        ,'<cite class="fly-link">{{item.user.username}}</cite>'
      ,'</a>'
      ,'{{# var date = new Date(item.time); if(d.index < 2){ }}'
        ,'<span class="fly-grey">签到于 {{ layui.laytpl.digit(date.getHours()) + ":" + layui.laytpl.digit(date.getMinutes()) + ":" + layui.laytpl.digit(date.getSeconds()) }}</span>'
      ,'{{# } else { }}'
        ,'<span class="fly-grey">已连续签到 <i>{{ item.days }}</i> 天</span>'
      ,'{{# } }}'
    ,'</li>'
  ,'{{# }); }}'
  ,'{{# if(d.data.length === 0) { }}'
    ,'{{# if(d.index < 2) { }}'
      ,'<li class="fly-none fly-grey">今天还没有人签到</li>'
    ,'{{# } else { }}'
      ,'<li class="fly-none fly-grey">还没有签到记录</li>'
    ,'{{# } }}'
  ,'{{# } }}'].join('');

  elemSigninTop.on('click', function(){
    var loadIndex = layer.load(1, {shade: 0.8});
    fly.json('/json/signin.js', function(res){ //实际使用，请将 url 改为真实接口
      var tpl = $(['<div class="layui-tab layui-tab-brief" style="margin: 5px 0 0;">'
        ,'<ul class="layui-tab-title">'
          ,'<li class="layui-this">最新签到</li>'
          ,'<li>今日最快</li>'
          ,'<li>总签到榜</li>'
        ,'</ul>'
        ,'<div class="layui-tab-content fly-signin-list" id="LAY_signin_list">'
          ,'<ul class="layui-tab-item layui-show"></ul>'
          ,'<ul class="layui-tab-item">2</ul>'
          ,'<ul class="layui-tab-item">3</ul>'
        ,'</div>'
      ,'</div>'].join(''))
      ,signinItems = tpl.find('.layui-tab-item');

      layer.close(loadIndex);

      layui.each(signinItems, function(index, item){
        var html = laytpl(tplSigninTop).render({
          data: res.data[index]
          ,index: index
        });
        $(item).html(html);
      });

      layer.open({
        type: 1
        ,title: '签到活跃榜 - TOP 20'
        ,area: '300px'
        ,shade: 0.8
        ,shadeClose: true
        ,id: 'layer-pop-signintop'
        ,content: tpl.prop('outerHTML')
      });

    }, {type: 'get'});
  });


  //回帖榜
  var tplReply = ['{{# layui.each(d.data, function(index, item){ }}'
    ,'<dd>'
      ,'<a href="/u/{{item.uid}}">'
        ,'<img src="{{item.user.avatar}}">'
        ,'<cite>{{item.user.username}}</cite>'
        ,'<i>{{item["count(*)"]}}次回答</i>'
      ,'</a>'
    ,'</dd>'
  ,'{{# }); }}'].join('')
  ,elemReply = $('#LAY_replyRank');

  if(elemReply[0]){
    /*
    fly.json('/top/reply/', {
      limit: 20
    }, function(res){
      var html = laytpl(tplReply).render(res);
      elemReply.find('dl').html(html);
    });
    */
  };

  //相册
  if($(window).width() > 750){
    layer.photos({
      photos: '.photos'
      ,zIndex: 9999999999
      ,anim: -1
    });
  } else {
    $('body').on('click', '.photos img', function(){
      window.open(this.src);
    });
  }


  //搜索
  $('.fly-search').on('click', function(){

  });

  // //新消息通知
  // fly.newmsg();

  // //发送激活邮件
  // fly.activate = function(email){
  //   fly.json('/api/activate/', {}, function(res){
  //     if(res.code === 0){
  //       layer.alert('已成功将激活链接发送到了您的邮箱，接受可能会稍有延迟，请注意查收。', {
  //         icon: 1
  //       });
  //     };
  //   });
  // };
  // $('#LAY-activate').on('click', function(){
  //   fly.activate($(this).attr('email'));
  // });

  /*传入html字符串源码即可*/
   
  fly.htmlEscape = function (text){ 
    return text.replace(/[<>"&]/g, function(match, pos, originalText){
      switch(match){
      case "<": return "&lt;"; 
      case ">":return "&gt;";
      case "&":return "&amp;"; 
      case "\"":return "&quot;"; 
    } 
    }); 
  }


  fly.saveInputFromCache = function(name, value){ 
      var saveData = {key:name, value:value};
      layui.data('pub6_com_user_cache_input_topic', saveData);//把AJSON对象存储为字符串
  }
  fly.loadInputFromCache = function (name) {  
      var cacheUser = layui.data('pub6_com_user_cache_input_topic');
      return cacheUser[name] || {};
  }  

  //点击@
  var default_sort = 0;
  var sort_default_selected = 1;
  var sort_indo_selected = 2;
  var sort_end_selected = 3;
  $('body').on('click', '#sort_default', function(){
    var othis = $(this), text = othis.text();
    // if(othis.attr('href') !== 'javascript:;'){
    //   return;
    // } 

    //判断当前选择的是不是sort_default
    if (default_sort == sort_default_selected){
      return;
    }

    default_sort = sort_default_selected;
    // if (layui.cache.innerName == 'topic_index' && layui.cache.topictype != 0){ 
        
        pageSize = 20;
        currPage = 1;
      if (layui.cache.topictype == 0){ 
        fly.json("/pub6/topic/all", {"pageSize":pageSize, "order_num":0, "pageNum":currPage}, function(res){ 
          if(res.code == 0){ 
              //console.log("sort_default res.total:"+res['data']['total'] );//+ ", res:"+ JSON.stringify(res)
              //saveTopicToCache("topic_top", JSON.stringify(res));

            addTopicTopList(res, "#topic-list", "#topic_page", currPage);
          };
        });   
      }else{ 
        fly.json("/pub6/topic/type", { "topic_type":layui.cache.topictype, "pageSize":pageSize, "order_num":0, "pageNum":currPage}, function(res){ 
          if(res.code == 0){ 
              //console.log("sort_default res.total:"+res['data']['total'] );//+ ", res:"+ JSON.stringify(res)
              //saveTopicToCache("topic_top", JSON.stringify(res));

            addTopicTopList(res, "#topic-list", "#topic_page", currPage);
          };
        });   
      }
      $("#sort_indo").removeClass("layui-this");
      $("#sort_end").removeClass("layui-this");
      othis.addClass("layui-this");
      return;
    // }
  });
  //点击@ 
  $('body').on('click', '#sort_indo', function(){
    var othis = $(this), text = othis.text();
    // if(othis.attr('href') !== 'javascript:;'){
    //   return;
    // } 

    //判断当前选择的是不是 sort_indo
    if (default_sort == sort_indo_selected){
      return;
    }
    default_sort = sort_indo_selected;
    // if (layui.cache.innerName == 'topic_index' && layui.cache.topictype != 0){ 
      
      pageSize = 20;
      currPage = 1;
      if (layui.cache.topictype == 0){ 

        fly.json("/pub6/topic/all", {"status":1, "pageSize":pageSize, "order_num":0, "pageNum":currPage}, function(res){ 
          if(res.code == 0){ 
              //console.log("sort_indo getTopicTopList res.total:"+res['data']['total'] );//+ ", res:"+ JSON.stringify(res)
              //saveTopicToCache("topic_top", JSON.stringify(res));

            addTopicTopList(res, "#topic-list", "#topic_page", currPage);
          };
        });  
      }else{  
        fly.json("/pub6/topic/type", {"status":1, "topic_type":layui.cache.topictype,  "pageSize":pageSize, "order_num":0, "pageNum":currPage}, function(res){ 
          if(res.code == 0){ 
            //console.log("sort_end getTopicTopList res.total:"+res['data']['total'] );//+ ", res:"+ JSON.stringify(res)
            //saveTopicToCache("topic_top", JSON.stringify(res));

            addTopicTopList(res, "#topic-list", "#topic_page", currPage);
          };
        });   
      }
      $("#sort_default").removeClass("layui-this"); 
      $("#sort_end").removeClass("layui-this");
      $("#sort_indo").addClass("layui-this");

      return;
    // }

  });
  //点击@ 
  $('body').on('click', '#sort_end', function(){
    var othis = $(this), text = othis.text();
    // if(othis.attr('href') !== 'javascript:;'){
    //   return;
    // } 

    //判断当前选择的是不是sort_end_selected
    if (default_sort == sort_end_selected){
      return;
    }
    default_sort = sort_end_selected;
 
    // if (layui.cache.innerName == 'topic_index' && layui.cache.topictype != 0){ 
    pageSize = 20;
    currPage = 1;
    if (layui.cache.topictype == 0){ 
      fly.json("/pub6/topic/all", {"status":3, "pageSize":pageSize, "order_num":0, "pageNum":currPage}, function(res){ 
        if(res.code == 0){ 
            //console.log("getTopicTopList res.total:"+res['data']['total'] );//+ ", res:"+ JSON.stringify(res)
            //saveTopicToCache("topic_top", JSON.stringify(res));

          addTopicTopList(res, "#topic-list", "#topic_page", currPage);
        };
      });  
    }else{

      fly.json("/pub6/topic/type", {"status":3, "topic_type":layui.cache.topictype,  "pageSize":pageSize, "order_num":0, "pageNum":currPage}, function(res){ 
        if(res.code == 0){ 
          //console.log("sort_end getTopicTopList res.total:"+res['data']['total'] );//+ ", res:"+ JSON.stringify(res)
          //saveTopicToCache("topic_top", JSON.stringify(res));

          addTopicTopList(res, "#topic-list", "#topic_page", currPage);
        };
      });  
    }
    $("#sort_default").removeClass("layui-this");
    $("#sort_indo").removeClass("layui-this"); 
    othis.addClass("layui-this");
    return;
    // }
  });

  //点击@
  // $('body').on('click', '.fly-aite', function(){
  //   var othis = $(this), text = othis.text();
  //   if(othis.attr('href') !== 'javascript:;'){
  //     return;
  //   }
  //   text = text.replace(/^@|（[\s\S]+?）/g, '');
  //   othis.attr({
  //     href: '/jump?username='+ text
  //     ,target: '_blank'
  //   });
  // });

  var onFinishCallback = function(){
    currPage = layui.cache.currPage;
      //暂时放在这里，更新详情页面的评论
      //console.log("onFinishCallback!");
 
    //取这个主题的评论列表
    fly.json('/api.v2/topic/commentList', {topic_id:layui.cache.tid, "pageNum":currPage}, function(res){
      //console.log("[1]commentList res:"+JSON.stringify(res));  
           
      if (!res['data']['list']){
        return;
      }
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
              ,'<span class="jieda-zan" type="zan">'
                ,'<i class="iconfont icon-zan"></i>'
                ,'<em>LIKE_NUMBER</em>'
              ,'</span>'
              ,'<span type="reply">'
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
      var pageTotal = (res['data']['total']/pageSize)?(res['data']['total']/pageSize+ ((res['data']['total']%pageSize)>0?1:0)):1;
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
 
  } ;

  //表单提交
  form.on('submit(*)', function(data){
    //console.log(data.field);
    for (var i=0;i<data.field.length; i++){
      //对提交的表单内容进行转义处理
      var old = $(data.form)[i].val();
      //防止cscr
      $(data.form)[i].val(fly.htmlEscape(old));
    } 
    var action = $(data.form).attr('action'), button = $(data.elem);

    //記錄本地緩存
    var cacheName = $(data.form).attr('cacheName');
    var cacheNum = 0;
    if (cacheName && cacheName.length > 0){
      var cacheNumStr = $(data.form).attr('cacheNum');
      cacheNum = new Number(cacheNumStr); 
    }
    //end.

    var onFinishCallbackFlag = $(data.form).attr('onFinishCallback');
    fly.json(action, data.field, function(res){
      var end = function(){
        if(res.action){
          location.href = res.action;
        } else {
          //fly.form[action||button.attr('key')](data.field, data.form);
          var url = button.attr('after_sumbit_url');
          layer.msg("提交成功", {shift: 6});

          //add start
          if (cacheNum > 0){
            //有本地緩存，清除緩存
            for (var i = 1; i <= cacheNum; i++) { 
              fly.saveInputFromCache(cacheName+i, ""); 
            };
          }
          //end.
 
          if (url ){
            if (url != "#"){
              setTimeout(function () {
                                    location.href = url; 
                                  }, 1000);  
            } else{
              //不需要跳转
              $(data.form)[0].reset();
              layui.form.render();
              if (onFinishCallbackFlag){
                onFinishCallback();
              }
            }
          }else{ 
              setTimeout(function () {
                                    location.href = fly.loadFromCurrUrlToCache(); 
                                  }, 1000);  
          }           
        }
      } 
      if(res.code == 0){
        button.attr('alert') ? layer.alert(res.msg, {
          icon: 1,
          time: 10*1000,
          end: end
        }) : end();
      };
    });
    return false;
  });
  //表单提交
  form.on('submit(reply)', function(data){
    var action = $(data.form).attr('action'), button = $(data.elem);
    fly.json(action, data.field, function(res){
      var end = function(){
        if(res.action){
          location.href = res.action;
        } else {
          //fly.form[action||button.attr('key')](data.field, data.form);

          layer.msg("提交成功", {shift: 6});
            
        }
      };
      if(res.code == 0){
        button.attr('alert') ? layer.alert(res.msg, {
          icon: 1,
          time: 10*1000,
          end: end
        }) : end();
      };
    });
    return false;
  });

  var updateHeaderInfo = function (){
    if (layui.cache.user.uid == -1 || !layui.cache.userdetail){
      return;
    }

    //登录成功，根据缓存信息更新顶端数据
    //console.log("updateHeaderInfo email:"+layui.cache.userdetail.email +",nickname: "+ layui.cache.userdetail.nickname 
      //  + " ,cityname:"+ layui.cache.userdetail.cityname +
      // ",blogSite: "+ layui.cache.userdetail.blogSite +
      // ",signNotes: "+layui.cache.userdetail.signNotes );
    if (layui.cache.userdetail.nickname  && layui.cache.userdetail.nickname.length >0){
      $("#user_nickname").html(layui.cache.userdetail.nickname); 
    }else{
      $("#user_nickname").html(layui.cache.user.username); 
    }
    avatar = layui.cache.user.avatar;
    if (avatar && avatar.length > 0){
      $("#user_avater").attr("src", avatar); 
    }else{ 
      $("#user_avater").attr("src", "/res/images/avatar/default.png"); 
    }
  };

  //登录表单
  form.on('submit(login)', function(data){
    var action = $(data.form).attr('action'), button = $(data.elem);
      fly.json(action, data.field, function(res){
        var end = function(){
            //登录成功
            layui.cache.user = {
              username: data.field['username']
              ,uid: res.data.uid
              ,token:res.data.token
              ,avatar: res.data.ext
              ,joinTime:res.data.joinTime
              ,experience: 83
              ,sex: '男'
            };  
            if (!res.data.ext || res.data.ext.length == 0){
              layui.cache.user.avatar = "/res/images/avatar/default.png";
            }

            //
            var saveData = {key:"userinfo", value:JSON.stringify(layui.cache.user)};
            //layui.data('user_cache', saveData);//把AJSON对象存储为字符串
            layui.sessionData('user_cache', saveData);//
            //layui.data('user_cache', layui.cache.user);
            layer.msg("登录成功", {shift: 6});
          
            setTimeout(function () {
                                  if (layui.cache.lasturl){
                                    location.href = layui.cache.lasturl; 
                                  }else{
                                    location.href = "../"; 
                                  }
                                }, 1000);
            //加载用户详情

            fly.json('/api.v2/user/getUserProfileByUserId  ', {userid:layui.cache.user.uid}, function(res){
                //console.log("getUserProfileByUserId res:"+JSON.stringify(res));  
                if (!res['data']){
                    //登录成功，加载用户详情
                    layui.cache.userdetail = {
                      email: ''
                      ,nickname: layui.cache.user.username
                      ,cityname:  ''
                      ,blogSite: ''
                      ,signNotes:  '生活不只是眼前，还有梦想'
                    };  
                    var saveData = {key:"userinfo", value:JSON.stringify(layui.cache.userdetail)}; 
                    layui.sessionData('user_cache_more', saveData);//
                    return;
                }
                //登录成功，加载用户详情
                layui.cache.userdetail = {
                  email: res['data']['email']
                  ,nickname: res['data']['nickname']
                  ,cityname:res['data']['cityname']
                  ,blogSite: res['data']['blogSite'] 
                  ,signNotes: res['data']['signNotes'] 
                };  
                var saveData = {key:"userinfo", value:JSON.stringify(layui.cache.userdetail)}; 
                layui.sessionData('user_cache_more', saveData);//
            }); 
          };
     
          if(res.code == 0){
            button.attr('alert') ? layer.alert(res.msg, {
              icon: 1,
              time: 10*1000,
              end: end
            }) : end();
          };
      });
      return false;
  });

  //注册表单
  form.on('submit(register)', function(data){
    var action = $(data.form).attr('action'), button = $(data.elem);
    fly.json(action, data.field, function(res){
      var end = function(){
        //注册成功 
        layer.msg("注册成功，请使用用户名密码登录", {shift: 6});
        
        setTimeout(function () {
              location.href = "/user/login.html";
                            }, 1000);  
      };
      if(res.code == 0){
        button.attr('alert') ? layer.alert(res.msg, {
          icon: 1,
          time: 10*1000,
          end: end
        }) : end();
      };
    });
    return false;
  });

  //加载特定模块
  if(layui.cache.page && layui.cache.page !== 'index'){
    var extend = {};
    extend[layui.cache.page] = layui.cache.page;
    layui.extend(extend);
    layui.use(layui.cache.page);
  }
  
  //加载IM
  if(!device.android && !device.ios){
    //layui.use('im');
  }

  //加载编辑器
  fly.layEditor({
    elem: '.fly-editor'
  });

  //手机设备的简单适配
  var treeMobile = $('.site-tree-mobile')
  ,shadeMobile = $('.site-mobile-shade')

  treeMobile.on('click', function(){
    $('body').addClass('site-mobile');
  });

  shadeMobile.on('click', function(){
    $('body').removeClass('site-mobile');
  });

  //获取统计数据
  $('.fly-handles').each(function(){
    var othis = $(this);
    $.get('/api/handle?alias='+ othis.data('alias'), function(res){
      othis.html('（下载量：'+ res.number +'）');
    })
  });
  
  //固定Bar
  util.fixbar({
    bar1: '&#xe642;'
    ,bgcolor: '#009688'
    ,click: function(type){
      if(type === 'bar1'){
        layer.msg('打开 index.js，开启发表新帖的路径');
        //location.href = 'jie/add.html';
      }
    }
  }); 


  //读取缓存，然后设置layui.cache
  if (!layui.cache.user){ 
    var cacheUser = layui.sessionData('user_cache');
    var localConfig = cacheUser['userinfo'] || {};
    if(JSON.stringify(localConfig) == '{}'){ 
      layui.cache.user = {  
        username: '游客'
        ,uid: -1
        ,avatar: '/res/images/avatar/00.jpg'
        ,experience: 83
        ,sex: '男'
      }; 
    }else{
        layui.cache.user =  JSON.parse(localConfig);//JSON.stringify(cacheUser);
      
        var cacheUserMore = layui.sessionData('user_cache_more');
        var userMoreStr = cacheUserMore['userinfo'] || {};
        if(JSON.stringify(userMoreStr) == '{}'){ 
          //return;
        }else{
          layui.cache.userdetail = JSON.parse(userMoreStr);
          updateHeaderInfo();
        }    
      
      }
  }

  if (layui.cache.page == 'user'  || layui.cache.innerName == 'topic_add'){
    refreshVerify();
  }
  if (layui.cache.innerName == 'topic_add' || layui.cache.innerName == "topic_index"){ 

    var category = loadFromCache(CATEGORY_CACHE_TAG);
    if(!category || category.localeCompare("{}") == 0){ 
      refreshCatgory(); 
    }else{ 
      if(category && category.localeCompare("{}") != 0){  
        res =  JSON.parse(category);//JSON.stringify(cacheUser);
        updateDefaultCatForm(res);
        //return;
      }
    }
  } 
  if (layui.cache.page == "index"){  
    refreshCatgory();  
        
    getTopicTopList();
    getTopicAllList(1);  
  }
  if (layui.cache.innerName == "topic_index"){
    getTopicAllList(layui.cache.currPage);  
  }
  if (!layui.cache.user || layui.cache.user.uid == -1){ 
    $('#nologin').removeClass("layui-hide");
  }else{ 
    $('#loginafter').removeClass("layui-hide");
  }

  exports('fly', fly);

});

