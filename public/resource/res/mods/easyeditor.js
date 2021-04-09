layui.define(['jquery', 'layer', 'form', 'element', 'upload', 'code', 'face'], function(exports) { //提示：模块也可以依赖其它模块，如：layui.define('layer', callback);

	var $ = layui.jquery,
		layer = layui.layer,
		form = layui.form,
		element = layui.element,
		upload = layui.upload,
		face = layui.face,
		device = layui.device();

	let renderer = new marked.Renderer();
	renderer.code = function(code, infostring, escaped) {
		return "<pre>" +code.replace(/\n/g, "<br>") + "</pre>";
	}
	marked.setOptions({
		renderer: renderer,
		tables: true,
		breaks: true
	});

	layui.focusInsert = function(obj, str) {
		var result, val = obj.value;
		obj.focus();
		if (document.selection) { //ie
			result = document.selection.createRange();
			document.selection.empty();
			result.text = str;
		} else {
			result = [val.substring(0, obj.selectionStart), str, val.substr(obj.selectionEnd)];
			obj.focus();
			obj.value = result.join('');
		}
	};

	let easyeditor = {
		init: function(options) {
			var html = ['<div class="layui-unselect fly-edit ' + (options.style === "fangge" ? "easyeditor-fangge" : "") +'" >',
				'<span type="face" title="插入表情"><i class="iconfont chengliangyun-md-icon-biaoqing" style="top: 1px;"></i></span>',
				'<span type="picture" title="插入图片：img[src]"><i class="iconfont chengliangyun-md-icon-tupian"></i></span>',
				'<span type="href" title="超链接格式：a(href)[text]"><i class="iconfont chengliangyun-md-icon-chaolianjie"></i></span>',
				'<span type="code" title="插入代码"><i class="iconfont chengliangyun-md-icon-daima" style="top: 1px;"></i></span>',
				'<span type="yinyong" title="引用"><i class="iconfont chengliangyun-md-icon-blockquote"></i></span>',
				'<span type="ul" title="无序列表"><i class="iconfont chengliangyun-md-icon-wuxuliebiao"></i></span>',
				'<span type="ol" title="有序列表"><i class="iconfont chengliangyun-md-icon-youxuliebiao"></i></span>',
				'<span type="table" title="表格"><i class="iconfont chengliangyun-md-icon-biaoge"></i></span>',
				'<span type="video" title="视频"><i class="iconfont chengliangyun-md-icon-shipin"></i></span>',
				'<span type="hr" title="分割线">hr</span>', '<div class="fly-right">',
				'<span type="yulan"  title="预览"><i class="iconfont chengliangyun-md-icon-yulanyulan"></i></span>',
				'<span type="fullScreen"  title="全屏"><i class="iconfont chengliangyun-md-icon-quanping"></i></span>',
				'</div>'
			].join('');

			var log = {},
				mod = {
					face: function(editor, self) { //插入表情
						var str = '',
							ul, face = easyeditor.faces;
						for (var key in face) {
							str += '<li title="' + key + '"><img src="' + face[key] + '"></li>';
						}
						str = '<ul id="LAY-editface" class="layui-clear">' + str + '</ul>';
						layer.tips(str, self, {
							tips: 3,
							time: 0,
							skin: 'layui-edit-face'
						});
						$(document).on('click', function() {
							layer.closeAll('tips');
						});
						$('#LAY-editface li').on('click', function() {
							var title = $(this).attr('title') + ' ';
							layui.focusInsert(editor[0], 'face' + title);
							editor.trigger('keyup');
						});
					},
					picture: function(editor) { //插入图片
						options = options || {}
						layer.open({
							type: 1,
							id: 'fly-jie-upload',
							title: '插入图片',
							area: 'auto',
							shade: false,
							area: '465px',
							fixed: false,
							offset: [
								editor.offset().top - $(window).scrollTop() + 'px', editor.offset().left + 'px'
							],
							skin: 'layui-layer-border', 
							content: ['<ul class="layui-form layui-form-pane" style="margin: 20px;">', '<li class="layui-form-item">',
								'<label class="layui-form-label">URL</label>', '<div class="layui-input-inline">',
								'<input required name="image" placeholder="支持直接粘贴远程图片地址" value="" class="layui-input">', '</div>',
								'</li>', '<li class="layui-form-item" style="text-align: center;">',
								'<button type="button" lay-submit lay-filter="uploadImages" class="layui-btn">确认</button>', '</li>',
								'</ul>'
							].join(''),
							success: function(layero, index) {
								var image = layero.find('input[name="image"]');

								if (options.uploadUrl == null || options.uploadUrl == '') {
									layer.msg('未配置图片上传路径,图片无法保存', {
										icon: 5
									});
								}

								//执行上传实例
								upload.render({
									elem: '#uploadImg',
									url: options.uploadUrl,
									size: options.uploadSize || 1024,
									done: function(res) {
										if (res.code == 0) {
											image.val(res.url);
										} else {
											layer.msg(res.msg, {
												icon: 5
											});
										}
									}
								});
								form.on('submit(uploadImages)', function(data) {
									var field = data.field;
									if (!field.image) return image.focus();
									layui.focusInsert(editor[0], '![图片未命名](' + field.image + ')\n');
									layer.close(index);
									editor.trigger('keyup');
								});
							}
						});
					},
					href: function(editor) { //超链接
						layer.prompt({
							title: '请输入合法链接',
							shade: false,
							fixed: false,
							id: 'LAY_flyedit_href',
							offset: [
								editor.offset().top - $(window).scrollTop() + 'px', editor.offset().left + 'px'
							]
						}, function(val, index, elem) {
							if (!/^http(s*):\/\/[\S]/.test(val)) {
								layer.tips('这根本不是个链接，不要骗我。', elem, {
									tips: 1
								})
								return;
							}
							layui.focusInsert(editor[0], ' [' + val + '](' + val + ')');
							layer.close(index);
							editor.trigger('keyup');
						});
					},
					code: function(editor) { //插入代码
						layer.prompt({
							title: '请贴入代码',
							formType: 2,
							maxlength: 10000,
							shade: false,
							id: 'LAY_flyedit_code',
							area: ['800px', '360px']
						}, function(val, index, elem) {
							layui.focusInsert(editor[0], '\n~~~\n' + val + '\n~~~\n');
							layer.close(index);
							editor.trigger('keyup');
						});
					},
					yinyong: function(editor) {
						layer.prompt({
							title: '请贴入引用内容',
							formType: 2,
							maxlength: 10000,
							shade: false,
							id: 'LAY_flyedit_code',
							area: ['800px', '360px']
						}, function(val, index, elem) {
							layui.focusInsert(editor[0], '> ' + val + '\n');
							layer.close(index);
							editor.trigger('keyup');
						});
					},
					hr: function(editor) { //插入水平分割线
						layui.focusInsert(editor[0], '-----\n');
						editor.trigger('keyup');
					},
					ul: function(editor) { //插入无序列表
						layui.focusInsert(editor[0], '\n-  \n-  \n-  \n');
						editor.trigger('keyup');
					},
					ol: function(editor) { //插入有序列表
						layui.focusInsert(editor[0], '\n1. \n2. \n3. \n');
						editor.trigger('keyup');
					},
					table: function(editor) {
						layui.focusInsert(editor[0], '\n表头|表头|表头\n:---:|:--:|:---:\n内容|内容|内容 \n');
						editor.trigger('keyup');
					},
					video :function(editor){
						layer.open({
							type: 1,
							id: 'fly-jie-upload',
							title: '插入视频',
							shade: false,
							area: '465px',
							fixed: false,
							offset: [
								editor.offset().top - $(window).scrollTop() + 'px', editor.offset().left + 'px'
							],
							skin: 'layui-layer-border',
							content: ['<ul class="layui-form layui-form-pane" style="margin: 20px;">', '<li class="layui-form-item">',
								'<label class="layui-form-label">封面图</label>', '<div class="layui-input-inline">',
								'<input required name="image" placeholder="支持远程图片地址" value="" class="layui-input">', '</div>',
								'<button type="button" class="layui-btn layui-btn-primary" id="uploadImg"><i class="iconfont chengliangyun-md-icon-shangchuan"></i>上传封面</button>',
								'</li>', '<li class="layui-form-item">',
								'<label class="layui-form-label">视频</label>', '<div class="layui-input-inline">',
								'<input required name="video" placeholder="支持远程视频地址" value="" class="layui-input">', '</div>',
								'<button type="button" class="layui-btn layui-btn-primary" id="uploadVideo"><i class="iconfont chengliangyun-md-icon-shangchuan"></i>上传视频</button>',
								'<div class="layui-progress" lay-filter="progress" id="progress" style="margin-right: 6px;margin-top: 4px;visibility:hidden;" lay-showpercent="true">',
								'<div class="layui-progress-bar" lay-percent="20%"></div>',
								'</div></li>','<li class="layui-form-item" style="text-align: center;">',
								'<button type="button" lay-submit lay-filter="uploadFile" class="layui-btn">确认</button>', '</li>',
								'</ul>'
							].join(''),
							success: function(layero, index) {
								var image = layero.find('input[name="image"]')
								,video = layero.find('input[name="video"]')
								,progress = layero.find('#progress');
								
								if (options.uploadUrl == null || options.uploadUrl == '') {
									layer.msg('未配置图片上传路径,图片无法保存', {
										icon: 5
									});
								}
								
								if (options.videoUploadUrl == null || options.videoUploadUrl == '') {
									layer.msg('未配置视频上传路径,视频无法保存', {
										icon: 5
									});
								}
						
								//执行上传图片实例
								upload.render({
									elem: '#uploadImg',
									url: options.uploadUrl,
									size: options.uploadSize || 1024,
									done: function(res) {
										if (res.code == 0) {
											image.val(res.url);
										} else {
											layer.msg(res.msg, {
												icon: 5
											});
										}
									}
								});
								
								//执行上传视频实例
								upload.render({
									elem: '#uploadVideo',
									accept :'video',
									url: options.videoUploadUrl,
									size: options.videoUploadSize || 1024*10,
									done: function(res) {
										if (res.code == 0) {
											progress.css('visibility','hidden');
											image.val(res.url);
										} else {
											layer.msg(res.msg, {
												icon: 5
											});
										}
									},
									progress: function(n){
										if(n>0){progress.css('visibility','visible')}
									    var percent = n + '%' //获取进度百分比
										element.progress('progress', percent); //可配合 layui 进度条元素使用
										if(n>=100){}
									 }
								});
								
								form.on('submit(uploadFile)', function(data) {
									var field = data.field;
									if (!field.image) return image.focus();
									if (!field.video) return video.focus();
									layui.focusInsert(editor[0], 'video('+field.image+')[' + field.video + ']\n');
									layer.close(index);
									editor.trigger('keyup');
								});
							}
						});
					},
					fullScreen: function(editor, span) { //全屏
						$(window).resize(function() { //当浏览器大小变化时
							//获取浏览器窗口高度
							var winHeight = 0;
							if (window.innerHeight)
								winHeight = window.innerHeight;
							else if ((document.body) && (document.body.clientHeight))
								winHeight = document.body.clientHeight;
							//通过深入Document内部对body进行检测，获取浏览器窗口高度
							if (document.documentElement && document.documentElement.clientHeight)
								winHeight = document.documentElement.clientHeight;
							$(options.elem).css('height', winHeight - 40 + "px");
							$(window).unbind('resize');
						});
						var othis = $(span);
						othis.attr("type", "exitScreen");
						othis.attr("title", "退出全屏");
						othis.html('<i class="iconfont chengliangyun-md-icon-tuichuquanping"></i>');
						var ele = document.documentElement,
							reqFullScreen = ele.requestFullScreen || ele.webkitRequestFullScreen ||
							ele.mozRequestFullScreen || ele.msRequestFullscreen;
						if (typeof reqFullScreen !== 'undefined' && reqFullScreen) {
							reqFullScreen.call(ele);
						};
					},
					exitScreen: function(editor, span) { //退出全屏
						var othis = $(span);
						othis.attr("type", "fullScreen");
						othis.attr("title", "全屏");
						othis.html('<i class="iconfont chengliangyun-md-icon-quanping"></i>');
						var ele = document.documentElement
						if (document.exitFullscreen) {
							document.exitFullscreen();
						} else if (document.mozCancelFullScreen) {
							document.mozCancelFullScreen();
						} else if (document.webkitCancelFullScreen) {
							document.webkitCancelFullScreen();
						} else if (document.msExitFullscreen) {
							document.msExitFullscreen();
						}
						//恢复初始高度
						$(options.elem).css("height", "270px");
					},
					yulan: function(editor, span) { //预览
						var othis = $(span),
							getContent = function() {
								var content = editor.val();
								return /^\{html\}/.test(content) ?
									content.replace(/^\{html\}/, '') :
									easyeditor.content(content)
							},
							isMobile = device.ios || device.android;

						if (mod.yulan.isOpen) return layer.close(mod.yulan.index);

						mod.yulan.index = layer.open({
							type: 1,
							title: '预览',
							shade: false,
							offset: 'r',
							id: 'LAY_flyedit_yulan',
							area: [
								isMobile ? '100%' : '50%', '100%'
							],
							scrollbar: isMobile ? false : true,
							anim: -1,
							isOutAnim: false,
							content: '<div class="detail-body layui-text easyeditor-content" style="margin:20px;">' + getContent() +
								'</div>',
							success: function(layero) {
								let layuiCode = options.codeStyle === 'layuiCode';
								layuiCode ? easyeditor.codeContent({
									elem: layero.find('pre')
								}) : "";
								editor.on('keyup', function(val) {
									layero.find('.detail-body').html(getContent());
									layuiCode ? easyeditor.codeContent({
										elem: layero.find('pre')
									}) : "";
								});
								mod.yulan.isOpen = true;
								othis.addClass('layui-this');
							},
							end: function() {
								delete mod.yulan.isOpen;
								othis.removeClass('layui-this');
							}
						});

					}
				};
			layui.use('face', function(face) {
				options = options || {};
				easyeditor.faces = face;
				$(options.elem).each(function(index) {
					var that = this,
						othis = $(that),
						parent = othis.parent();
					parent.prepend(html);
					parent.find('.fly-edit span').on('click', function(event) {
						var type = $(this).attr('type');
						mod[type].call(that, othis, this);
						if (type === 'face') {
							event.stopPropagation()
						}
					});
				});
				//按钮样式
				let buttonColor = options.buttonColor,
					hoverColor = options.hoverColor,
					hoverBgColor = options.hoverBgColor;
				$(".fly-edit span").css("color", buttonColor ? buttonColor : "");
				$(".fly-edit span").hover(function() {
					$(this).css("color", hoverColor ? hoverColor : "").css("background-color", hoverBgColor ? hoverBgColor :
						"")
				}, function() {
					$(this).css("color", buttonColor ? buttonColor : "").css("background-color", "")
				});
				//代码块
				options.codeStyle === 'layuiCode' ?renderer.code = function(code, infostring, escaped) {
					return "<pre>" + code + "</pre>";
				}:'';
			});
		},
		codeContent: function(options) {
			let params = {
				elem: options.elem,
				title: 'code',
				about: false
				//,encode: true
			}
			if (options.codeSkin === 'notepad') {
				params.skin = 'notepad';
			}
			$("pre").css("white-space","pre-wrap");
			layui.code(params);
		},
		escape: function(html) {
			return String(html || '').replace(/&(?!#?[a-zA-Z0-9]+;)/g, '&amp;')
				.replace(/</g, '&lt;').replace(/'/g, '&#39;').replace(/"/g, '&quot;');
		},
		
		content: function(content) {
			return marked(easyeditor.escape(content || '') //xss
				.replace(/  \n/g, "<br>")//强制换行
				.replace(/video\([\s\S]+?\)\[[\s\S]*?\]/g, function(str){//转义视频
					var img = (str.match(/video\(([\s\S]+?)\)\[/)||[])[1];
					var video = (str.match(/\)\[([\s\S]*?)\]/)||[])[1];
					if(!video) return str;
					return ['<video id="video" controls="" preload="none" poster="'+img+'">',
							  '<source id="mp4" src="'+video+'" type="video/mp4">',
							'</video>'].join("");
				  }))
				.replace(/<a href="\S.+">/g,function(str){
					var href = (str.match(/<a href="([\s\S]+?)">/)||[])[1];
					var rel =  /^(http(s)*:\/\/)\b(?!(\w+\.)*(chengliangyun.com|www.chengliangyun.com))\b/.test(href.replace(/\s/g, ''));
					return '<a href="'+ href +'" target="_blank"'+ (rel ? ' rel="nofollow"' : '') +'>';
				}) //转义超链接
				.replace(/<table/g, "<table class='layui-table' ") //表格样式
				.replace(/<blockquote/g, "<blockquote class='layui-elem-quote layui-text'")
				
				.replace(/face\[([^\s\[\]]+?)\]/g, function(face) { //转义表情
					let alt = face.replace(/^face/g, '');
					return '<img class="face" alt="' + alt + '" title="' + alt + '" src="' + easyeditor.faces[alt] + '">';
				});
		},
		render: function(options) {
			options = options || {};
			var photos = function() {
				if ($(window).width() > 750) {
					layer.photos({
						photos: options.elem,
						img:"img:not(.face)",
						zIndex: 9999999999,
						anim: -1
					});
				} else {
					$('body').on('click', options.elem + ' img:not(.face)', function() {
						window.open(this.src);
					});
				}
			}
			$(options.elem).each(function() {
				let othis = $(this),
					text = othis.text();
				othis.html(easyeditor.content(text));
			});
			//相册
			photos();
		}
	}

	if (!easyeditor.faces) {
		easyeditor.faces = face;
	}

	exports('easyeditor', easyeditor);
});
