 
-------------

drop table if exists guaniu_study_topic;
create table guaniu_study_topic (
  id           bigint(20)    auto_increment    comment '编号',
  parent_id         bigint(20)      default 0                  comment '父主题id',
  topic_type        int(11)      default '0'        comment 'topic 分类类型, 关联到type表',
  topic_owner_type        int(11)      default '0'        comment '主题类型1.自主学习/2.组织引导学习/3.混合类型',
  see_num int(11) default '0'  comment '围观人数',
  topic_name        varchar(255)     default ''        comment '主题名称',
  topic_pinyin       varchar(1024)     default ''        comment '主题名称拼音',
  topic_tag       varchar(255)     default ''        comment '主题tag',
  topic_img        varchar(1024)     default ''        comment '主题图片路径',
  topic_content    text          default null      comment '主题描述',
  order_num         int(4)          default 0                  comment '显示顺序',
  status            int(4)         default 0               comment '是否开放状态（1免费开放 2自己可见 3已经结束 4 已经删除 5 收费开放 ）',
  pay_money         FLOAT         default 0           comment '付费金额',
  create_owner_id            bigint(20)                                     comment '创建者',
  create_time               datetime                                   comment '时间',
  modify_time               datetime                                  comment '时间', 
  members_number         int(11) default '1'                              comment '加入成员数',

  topic_table_index       int(8)         default '1'       comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1  CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci comment = '课程主题表' ;
  
drop table if exists guaniu_study_topic_type;
create table guaniu_study_topic_type (
  id           bigint(20)    auto_increment    comment '编号', 
  name    varchar(1024)  default null        comment '主题类型',  
  primary key (id)
) engine=innodb auto_increment=1  CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci comment = '主题类型';
insert into guaniu_study_topic_type(name) values ('编程语言'), ('操作系统'), ('数据库'), ('前端移动'), ('后端架构'), ('运维测试'), ('软件工程'), ('考试认证'), ('AI大数据'), ('游戏开发'), ('软技能'), ('产品运营'), ('工具软件');
 alter table guaniu_study_topic_type modify name varchar(1024) character set gbk;

drop table if exists guaniu_study_topic_notices;
create table guaniu_study_topic_notices (
  id           bigint(20)    auto_increment    comment '编号',
  topic_id        bigint(20)             comment '主题id',    
  user_id            int(11)                                     comment '发起用户',
  notice_desc    varchar(1024)  default null        comment '公告内容',
  create_time               datetime                                   comment '时间',  
  topic_table_index       int(8)         default '1'       comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1 DEFAULT  CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci comment = '主题公告';

drop table if exists guaniu_study_topic_comments;
create table guaniu_study_topic_comments (
  id           bigint(20)    auto_increment    comment '编号',
  topic_id        bigint(20)             comment '主题id',    
  user_id            int(11)                                     comment '发起用户',
  content    varchar(1024)  default null        comment '评论内容',
  status            int(4)         default 0               comment '是否开放状态（0没有审核  1 已经审核）',
  orderSeq            int(4)         default 0               comment '排序',
  likeNum            int(4)         default 0               comment '点赞',
  create_time               datetime                                   comment '时间',  
  topic_table_index       int(8)         default '1'       comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1 comment = '主题评论';

drop table if exists guaniu_study_topic_pay_members;
create table guaniu_study_topic_pay_members (
  id           bigint(20)    auto_increment    comment '编号',
  topic_id        bigint(20)             comment '主题id',   
  user_id            int(11)                                     comment '加入用户id',
  create_time               datetime                                   comment '时间',  
  topic_table_index       int(8)         default '1'       comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1 DEFAULT CHARSET=utf8mb4 comment = '课程付费表';


drop table if exists guaniu_study_topic_members;
create table guaniu_study_topic_members (
  id           bigint(20)    auto_increment    comment '编号',
  topic_id        bigint(20)             comment '主题id',  
  user_id        bigint(11)       comment '成员id',  
  status            int(4)         default 0          comment '是否开放状态（0正常 1 禁言 2被删除）',
  admin_flag          int(4)         default 0          comment '是否管理员（0不是 1 管理员）',

  create_time               datetime                                   comment '时间', 
  topic_table_index       int(8)         default '1'       comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1 DEFAULT CHARSET=utf8mb4 comment = '主题学习成员表';

drop table if exists guaniu_study_topic_members_history;
create table guaniu_study_topic_members_history (
  id           bigint(20)    auto_increment    comment '编号',
  topic_id        bigint(20)             comment '主题id',  
  user_id        bigint(11)       comment '成员id',  
  module     varchar(256)  default null        comment '模块',    
  operatorion            int(4)         default 0          comment '0 进入，1退出', 

  create_time               datetime                                   comment '时间', 
  topic_table_index       int(8)         default '1'       comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1 comment = '主题操作历史';

drop table if exists guaniu_study_topic_config;
create table guaniu_study_topic_config (
  id           bigint(20)    auto_increment    comment '编号',
  topic_id        bigint(20)        comment '主题id', 
  create_time               datetime                                   comment '时间',
  modify_time               datetime                               comment '时间', 
  topic_table_index       int(8)         default '1'       comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1 DEFAULT CHARSET=utf8mb4 comment = '主题配置表';


drop table if exists guaniu_study_topic_resource;
create table guaniu_study_topic_resource (
  id           bigint(20)    auto_increment    comment '编号',
  topic_id        bigint(11)         comment '主题id',
  user_id        bigint(11)         comment '用户id',

  resource_type_id        int(11)  default 0       comment '关联资源类型，资源类型1.公开网课/2.付费网课/3.专业书籍/4.文档材料/5.其他类型',
  book_id        int(11)     default 0     comment '书的id',
  resource_link varchar(1024)  default null        comment '链接',
  resource_content    text default null       comment '资源内容，可以是资源链接',
  status            int(4)         default 0          comment '是否开放状态（0正常 1 禁言）',
 
  create_time               datetime                                   comment '时间', 
  topic_table_index       int(8)         comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1 DEFAULT CHARSET=utf8mb4 comment = '主题资源表';

drop table if exists guaniu_study_book_resource;
create table guaniu_study_book_resource (
  id           bigint(20)    auto_increment    comment '编号',  
  book_name    varchar(1024)           comment '书的名称',
  book_author  varchar(256)  default null        comment '书的作者', 
  book_desc    varchar(1024)  default null        comment '书的描述',
  book_tags varchar(256) default null          comment '书的标签', 
  book_cbs     varchar(256) default null          comment '出版社', 
  print_time   datetime     default null                       comment '发行时间', 
  book_isbn    varchar(256) default null          comment 'isbn', 
  status            int(4)         default 0          comment '是否开放状态（0正常 1 删除）',
 
  create_time               datetime                                   comment '时间', 
  topic_table_index       int(8)         comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1 DEFAULT CHARSET=utf8mb4  comment = '专业书资源表';

drop table if exists guaniu_study_book_category;
create table guaniu_study_book_category (
  id           bigint(20)    auto_increment    comment '编号',  
  book_id        bigint(20)        comment '主题id', 
  catgory_name    varchar(1024)           comment '目录：第1章 核心套路篇',
  catgory_page  int(11)  default 0        comment '页码',   
 
  create_time               datetime                                   comment '时间', 
  topic_table_index       int(8)         comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1 DEFAULT CHARSET=utf8mb4 comment = '专业书目录资源表';

drop table if exists guaniu_study_topic_task;
create table guaniu_study_topic_task (
  id           bigint(20)    auto_increment    comment '编号',
  topic_id        bigint(20)          comment '主题id',

  task_title        varchar(1024)  default ''       comment '任务标题',
  task_content    text          comment '任务描述',
  time_line        varchar(248)  default ''       comment '任务截点',
  resource_type_id        int(11)  default 0       comment '关联资源类型，关联字典表',
  book_id        int(11)     default 0     comment '书的id',
  book_catgory_id        int(11)  default 0        comment '书的目录id',
  resource_content    text default null       comment '资源内容',
  order_num         int(4)          default 0                  comment '显示顺序',

   create_time               datetime                                   comment '时间',
  modify_time               datetime             comment '时间', 
  topic_table_index       int(8)         default '1'       comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1 DEFAULT CHARSET=utf8mb4 comment = '主题任务表';

drop table if exists guaniu_study_topic_taskprogress;
create table guaniu_study_topic_taskprogress (
  id           bigint(20)    auto_increment    comment '编号',
  topic_id        bigint(20)          comment '主题id',
  task_id        bigint(20)       comment '任务id',
  user_id        bigint(20)          comment '用户id',
 
  content    text         comment '进度补充描述',
  status            int(3)         default '0'                comment '状态（0开始 1 完成）',
  create_owner_id            int(11)                                     comment '创建者', 
 
  create_time               datetime                                   comment '时间',
  modify_time               datetime           comment '时间', 
  topic_table_index       int(8)         default '1'       comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1 DEFAULT CHARSET=utf8mb4 comment = '任务进度表';


drop table if exists guaniu_study_topic_notes;
create table guaniu_study_topic_notes (
  id           bigint(20)    auto_increment    comment '编号',
  topic_id        bigint(20)            comment '主题id', 
  user_id        bigint(20)          comment '用户id',

  title        varchar(1024)    default null          comment '笔记标题',
  content    text           comment '笔记描述',
  like_num             int(11)    default 0  comment  '喜欢数',
  order_num         int(4)          default 0                  comment '显示顺序',
  status            int(4)         default 0               comment '是否开放状态（0免费 1 收费  2 只对自己可见 3 只有关注者可见 ）',
  create_owner_id            int(11)                                     comment '创建者',  

   create_time               datetime                                   comment '时间',
  modify_time               datetime              comment '时间', 
  topic_table_index       int(8)         default '1'       comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1 DEFAULT CHARSET=utf8mb4 comment = '主题笔记表';
 
drop table if exists guaniu_study_questions;
create table guaniu_study_questions (
  id           bigint(20)    auto_increment    comment '编号',
  topic_id        bigint(20)    default '0'          comment '主题id',
  task_id        bigint(20)     default '0'        comment '任务id',
  user_id        bigint(20)    default '0'        comment '用户id',
  question_type    tinyint(3)   default '0'        comment '选择题0, 1:对错题  2:填空题',
  question_title                     text               comment '问题', 
  option_a varchar(2048)  default null         comment '描述A', 
  option_b varchar(2048)  default null         comment '描述B', 
  option_c varchar(2048)  default null         comment '描述C', 
  option_d varchar(2048) default null         comment '描述D', 
  option_e varchar(1024)  default null         comment '描述E',
  answer1 varchar(2048)  default null         comment '选择题答案', 
  answer2 varchar(2048)  default null         comment '描述题答案', 
  answer3 varchar(2048)  default null         comment '对错题答案', 
  question_answer_desc      text          comment '问题描述',  
  question_diff    tinyint(3)   default '1'        comment '题目难度 1，越大越难',

  create_owner_id            int(11)                                     comment '创建者',  
  create_time               datetime                                   comment '时间',
  modify_time               datetime             comment '时间', 
  topic_table_index       int(8)         default '1'       comment  '表id 与表topic 相关',
  primary key (id)
) engine=innodb auto_increment=1 DEFAULT CHARSET=utf8mb4 comment = '题库表';

drop table if exists guaniu_study_questions_answer_history;
create table guaniu_study_questions_answer_history (
  id           bigint(20)    auto_increment    comment '编号', 
  topic_id        bigint(20)     default 0        comment '主题id',
  user_id        bigint(20)          comment '用户id',
  status            int(4)         default 0               comment '是否开放状态（0开放  1  只对自己可见 3 只有关注者可见 ）',
  create_owner_id            bigint(20)                   comment '创建者',  
  create_time               datetime                          comment '时间',
  primary key (id)
) engine=innodb auto_increment=1 DEFAULT CHARSET=utf8mb4 comment = '做题记录表';

drop table if exists guaniu_study_questions_answer_history_detail;
create table guaniu_study_questions_answer_history_detail (
  id           bigint(20)    auto_increment    comment '编号',
  history_id        bigint(20)     default 0        comment '答题记录id',
  question_id        bigint(20)     default 0        comment '问题id',
  question_answer     varchar(2048)  default ''         comment '描述题答案', 
  topic_id        bigint(20)     default null        comment '主题id',
  answer_right    tinyint(3)   default '0'        comment '1:准确，2错误  3未知',

  primary key (id)
) engine=innodb auto_increment=1 DEFAULT CHARSET=utf8mb4 comment = '做题详细记录表';
 






