/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50643
 Source Host           : localhost:3306
 Source Schema         : weeds

 Target Server Type    : MySQL
 Target Server Version : 50643
 File Encoding         : 65001

 Date: 02/06/2020 18:15:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ks_article
-- ----------------------------
DROP TABLE IF EXISTS `ks_article`;
CREATE TABLE `ks_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `content` text NOT NULL COMMENT '内容',
  `goods_id` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

-- ----------------------------
-- Records of ks_article
-- ----------------------------
BEGIN;
INSERT INTO `ks_article` VALUES (1, NULL, '<p>&nbsp;</p>\r\n<p style=\"text-align: center;\"><strong>dddddddddddddd</strong></p>\r\n<p>文章内容啊啊啊</p>\r\n<p>&nbsp;</p>', 9);
INSERT INTO `ks_article` VALUES (4, NULL, '<p>&nbsp;</p>\r\n<!--?xml version=\"1.0\" encoding=\"UTF-8\"?-->\r\n<div><strong>1, 俩表字符类型不一致列表查询</strong></div>\r\n<div>```sql</div>\r\n<div>SELECT *</div>\r\n<div>FROM `hz_slideshow`</div>\r\n<div>WHERE `live_id` IN(SELECT `live_id`</div>\r\n<div>FROM `hz_live_history`</div>\r\n<div>WHERE `start_date` BETWEEN \'2018-08-29 00:00:00\'</div>\r\n<div>AND \'2018-08-30 00:00:00\')</div>\r\n<div>```</div>\r\n<div>&gt;&gt;&gt;&gt;&gt;失败原因：Illegal mix of collations (utf8_unicode_ci,IMPLICIT) and (utf8_general_ci,IMPLICIT) for operation \'=\'</div>\r\n<div>&nbsp;</div>\r\n<div>修改后</div>\r\n<div>```sql</div>\r\n<div>SELECT *</div>\r\n<div>FROM `hz_slideshow`</div>\r\n<div>WHERE `live_id` IN(SELECT binary `live_id`</div>\r\n<div>FROM `hz_live_history`</div>\r\n<div>WHERE `start_date` BETWEEN \'2018-08-29 00:00:00\'</div>\r\n<div>AND \'2018-08-30 00:00:00\')</div>\r\n<div>&nbsp;</div>\r\n<div>&nbsp;</div>\r\n<div>&nbsp;</div>\r\n<div><strong>2,&nbsp;delete from table_name&nbsp;&nbsp;删除表的全部数据</strong></div>\r\n<div>&nbsp;&nbsp; &nbsp;对于MyISAM 会立刻释放磁盘空间 （应该是做了特别处理，也比较合理）；&nbsp; &nbsp; InnoDB 不会释放磁盘空间</div>\r\n<div>&nbsp;</div>\r\n<div><strong>3， a表字段根据b表统计修改</strong></div>\r\n<div>UPDATE hz_small_course AS b</div>\r\n<div>&nbsp;&nbsp;&nbsp;SET b.`real_buy_num`= IFNULL((</div>\r\n<div>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;SELECT c.counts FROM(</div>\r\n<div>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;SELECT small_course_id, COUNT(*) counts FROM hz_small_course_competence GROUP BY small_course_id</div>\r\n<div>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;) c</div>\r\n<div>WHERE b.`id`= c.small_course_id),0) ;</div>\r\n<div>&nbsp;</div>\r\n<div>&nbsp;</div>\r\n<div><strong>4,mysql explain</strong></div>\r\n<div>id:选择标识符</div>\r\n<div>select_type:表示查询的类型。</div>\r\n<div>table:输出结果集的表</div>\r\n<div>partitions:匹配的分区</div>\r\n<div>type:表示表的连接类型</div>\r\n<div>possible_keys:表示查询时，可能使用的索引</div>\r\n<div>key:表示实际使用的索引</div>\r\n<div>key_len:索引字段的长度</div>\r\n<div>ref:列与索引的比较</div>\r\n<div>rows:扫描出的行数(估算的行数)</div>\r\n<div>filtered:按表条件过滤的行百分比</div>\r\n<div>Extra:执行情况的描述和说明</div>\r\n<div>&nbsp;</div>\r\n<div>type</div>\r\n<div>对表访问方式，表示MySQL在表中找到所需行的方式，又称&ldquo;访问类型&rdquo;。</div>\r\n<div>常用的类型有：&nbsp;ALL、index、range、 ref、eq_ref、const、system、NULL（从左到右，性能从差到好）</div>\r\n<div>ALL：Full Table Scan， MySQL将遍历全表以找到匹配的行</div>\r\n<div>index: Full Index Scan，index与ALL区别为index类型只遍历索引树</div>\r\n<div>range:只检索给定范围的行，使用一个索引来选择行</div>\r\n<div>ref: 表示上述表的连接匹配条件，即哪些列或常量被用于查找索引列上的值</div>\r\n<div>eq_ref: 类似ref，区别就在使用的索引是唯一索引，对于每个索引键值，表中只有一条记录匹配，简单来说，就是多表连接中使用primary key或者 unique key作为关联条件</div>\r\n<div>const、system: 当MySQL对查询某部分进行优化，并转换为一个常量时，使用这些类型访问。如将主键置于where列表中，MySQL就能将该查询转换为一个常量，system是const类型的特例，当查询的表只有一行的情况下，使用system</div>\r\n<div>NULL: MySQL在优化过程中分解语句，执行时甚至不用访问表或索引，例如从一个索引列里选取最小值可以通过单独索引查找完成。</div>\r\n<div>&nbsp;</div>\r\n<div>&nbsp;</div>\r\n<div>// 在线增加索引过程中遇到的锁表导致无法更新处理方式</div>\r\n<div>// 1、执行增加索引(会除以一直等待状态)</div>\r\n<div>// 2、查询锁表的进程</div>\r\n<div>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;select * from information_schema.processlist where command not in (\'Sleep\')&nbsp;&nbsp;and user not in (\'mydba\',\'event_scheduler\',\'repl\',\'binlogbak\',\'system user\')</div>\r\n<div>// 3、判断进程是否允许被干掉（select之类的数据，主要是杀掉导致第一步处以等待状态的进程）</div>\r\n<div>kill 进程id</div>\r\n<div>&nbsp;</div>\r\n<div>&nbsp;</div>\r\n<div>&nbsp;</div>', 12);
INSERT INTO `ks_article` VALUES (8, NULL, '<p></p><p><!--?xml version=\"1.0\" encoding=\"UTF-8\"?--><b>\r\n\r\n1,安装nginx-rtmp-module模块<br></b></p><p>参考地址：<a href=\"https://github.com/arut/nginx-rtmp-module/wiki/Directives\" style=\"background-color: rgb(255, 255, 255);\">https://github.com/arut/nginx-rtmp-module/wiki/Directives</a></p><p><!--?xml version=\"1.0\" encoding=\"UTF-8\"?-->\r\n\r\n</p><div><b>2,配置文件 open /usr/local/etc/nginx 下的nginx.conf</b></div><div><pre><code>#配置nginx支持rtmp\r\nsudo vim /usr/local/nginx/conf/nginx.conf\r\nrtmp {\r\n    server{\r\n        listen 1066;\r\n        #录制配置\r\n        record all;\r\n        record_unique on;\r\n        record_path /home/weeds/weeds/rtmp/record;\r\n        record_suffix -%Y-%m-%d-%H_%M_%S.flv;\r\n        #record_max_size 300m;\r\n        #直播权限配置\r\n        publish_notify on;\r\n        #notify_method get;\r\n        on_publish http://dev.ft.com/new_api/huazhen_rtmp/pushCallBack;\r\n        #on_publish_done http://127.0.0.1/publishDoneCallBack.php;\r\n        #on_play http://127.0.0.1/playCallBack.php;\r\n        #on_play_done http://127.0.0.1/playDoneCallBack.php;\r\n        #直播流配置\r\n        application rtmplive{\r\n            live on;\r\n            #为rtmp设置最大链接数。默认为off\r\n            max_connections 1024;\r\n        }\r\n        application hls{\r\n            live on;\r\n            #hls on;\r\n            #hls_path /home/weeds/weeds/rtmp/hls;\r\n            #hls_fragment 5s;\r\n        }\r\n    \r\n        hls on;\r\n        hls_path /home/weeds/weeds/rtmp/hls;\r\n        hls_fragment 5s;#没有生效\r\n        hls_playlist_length 30s;\r\n        hls_nested on; #默认是off。打开后的作用是每条流自己有一个文件夹\r\n        hls_cleanup off;#不清理ts\r\n\r\n    }\r\n}</code></pre><p><b>3,nginx配置支持hls</b></p><pre><code>http节点的server下增加\r\n#加入hls支持\r\n        location /hls {\r\n            types {\r\n                application/vnd.apple.mpegurl m3u8; \r\n                #或 application/x-mpegURL\r\n                video/mp2t ts;\r\n            }\r\n            alias /home/weeds/weeds/rtmp/hls;  #视频流文件目录(自己创建)\r\n            expires -1;\r\n            add_header Cache-Control no-cache;\r\n        }<br></code></pre><div><span><b>4,#推流</b></span></div><div><span>ffmpeg -re -i /home/weeds/weeds/audio/4.mp4 -vcodec libx264 -acodec aac -f flv</span> <a shape=\"rect\" href=\"rtmp://localhost:1066/hls/room0000\" target=\"_blank\">rtmp://localhost:1066/hls/room</a></div><div>#拉流地址</div><div>http://localhost/hls/room/index.m3u8 <br></div><div><a shape=\"rect\" href=\"rtmp://localhost:1066/hls/room0000\" target=\"_blank\">rtmp://localhost:1066/hls/room</a><br></div><div><br></div><div>工具：</div><div>推流：ffmpeg,OBS</div><div>拉流：VLC</div><div><br></div><p><br></p><!--?xml version=\"1.0\" encoding=\"UTF-8\"?-->\r\n\r\n</div><!--?xml version=\"1.0\" encoding=\"UTF-8\"?-->', 17);
INSERT INTO `ks_article` VALUES (9, NULL, '<p></p><div><b>### 作用\r\n</b></div><div>用来查看当前线程处理情况；可以看到总共有多少链接数，哪些线程有问题(time是执行秒数，时间长的就应该多注意了)，然后可以把有问题的线程 kill 掉，这样可以临时解决一些突发性的问题\r\n</div><div><b>### 命令\r\n</b></div><div>```SQL\r\n</div><div>select id, db, user, host, command, time, state, info\r\n</div><div>from information_schema.processlist\r\n</div><div>where command != \'Sleep\'\r\n</div><div>order by time desc\r\n</div><div>```\r\n</div><div>![8d61515d49678f6e150417c63a4d761e.png](evernotecid://485807EF-3D57-4E1A-A183-F769BBEEA6DF/appyinxiangcom/18512895/ENResource/p14)\r\n</div><div><b>### 结果解析\r\n</b></div><div>| 字段 |解析 |备注 |\r\n</div><div>| --- | --- | --- |\r\n</div><div>| ID | 链接mysql 服务器线程的唯一标识，可以通过kill来终止此线程的链接 | |\r\n</div><div>| User | 当前线程链接数据库的用户 | |\r\n</div><div>| Host | 显示这个语句是从哪个ip 的哪个端口上发出的。可用来追踪出问题语句的用户 | |\r\n</div><div>| DB | 线程链接的数据库，如果没有则为null | |\r\n</div><div>| Command | 显示当前连接的执行的命令，一般就是休眠或空闲（sleep），查询（query），连接（connect） | |\r\n</div><div>| Time | 线程处在当前状态的时间，单位是秒 | |\r\n</div><div>| State | 显示使用当前连接的sql语句的状态，很重要的列，后续会有所有的状态的描述，请注意，state只是语句执行中的某一个状态，一个 sql语句，已查询为例，可能需要经过copying to tmp table，Sorting result，Sending data等状态才可以完成 | |\r\n</div><div>| Info | 线程执行的sql语句，如果没有语句执行则为null。这个语句可以使客户端发来的执行语句也可以是内部执行的语句 | |\r\n</div><div><b>### kill使用\r\n</b></div><div>-- 查询执行时间超过2分钟的线程，然后拼接成 kill 语句\r\n</div><div>```SQL\r\n</div><div>select concat(\'kill \', id, \';\')\r\n</div><div>from information_schema.processlist\r\n</div><div>where command != \'Sleep\'\r\n</div><div>and time &gt; 2*60\r\n</div><div>order by time desc\r\n</div><p><!--?xml version=\"1.0\" encoding=\"UTF-8\"?-->\r\n\r\n</p><div>```&nbsp;</div>', 43);
INSERT INTO `ks_article` VALUES (10, NULL, '<p style=\"text-align: center;\"><span style=\"font-family: \'arial black\', \'avant garde\'; font-size: 56px;\"><strong>愚人节快乐</strong></span></p>\r\n<p style=\"text-align: left;\"><span style=\"font-family: \'arial black\', \'avant garde\'; font-size: 12px;\">大家好 才是做得好</span></p>', 44);
INSERT INTO `ks_article` VALUES (11, NULL, '<p>kasdoiw</p>\r\n<p><strong>adsfkj</strong></p>', 45);
INSERT INTO `ks_article` VALUES (12, NULL, '<p>adsf</p>', 46);
COMMIT;

-- ----------------------------
-- Table structure for ks_comments
-- ----------------------------
DROP TABLE IF EXISTS `ks_comments`;
CREATE TABLE `ks_comments` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) unsigned NOT NULL COMMENT '商品类型:1-文章',
  `goods_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '父级id',
  `name` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '商品名称',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `description` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '商品简介',
  `image` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '商品大图',
  `order` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品排序',
  `is_verify` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否审核上线 0未审核 1已审核',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='评论表';

-- ----------------------------
-- Records of ks_comments
-- ----------------------------
BEGIN;
INSERT INTO `ks_comments` VALUES (1, 1, 17, 0, '野草ks', 'weedsyl@163.com', '评论内容', '', 0, 1, NULL, '2019-08-10 16:54:13', '2019-08-10 16:54:13');
INSERT INTO `ks_comments` VALUES (2, 1, 17, 0, '野草ks', 'weedsyl@163.com', '评论内容', '', 0, 1, NULL, '2019-08-10 16:55:26', '2019-08-10 16:55:26');
INSERT INTO `ks_comments` VALUES (3, 1, 17, 0, 'weeds', 'weedsyl@163.com', 's爱迪生手动阀', '', 0, 1, NULL, '2019-08-11 15:43:23', '2019-08-11 15:43:23');
INSERT INTO `ks_comments` VALUES (4, 1, 17, 0, '周三', 'weeds7785@gmail.com', '微测试', '', 0, 1, NULL, '2019-08-11 15:49:19', '2019-08-11 15:49:19');
INSERT INTO `ks_comments` VALUES (5, 1, 12, 0, '野草ks', 'weedsyl@163.com', '我可以老瞧瞧吗', '', 0, 1, NULL, '2019-08-14 06:22:39', '2019-08-14 06:22:39');
INSERT INTO `ks_comments` VALUES (6, 0, 0, 0, '张三', 'zhangsan@163.com', '我就是来看看的 不要介意', '', 0, 1, NULL, '2019-08-14 06:22:39', '2019-08-14 06:22:39');
INSERT INTO `ks_comments` VALUES (7, 1, 17, 0, '张三', 'zs@174.com', '斯柯达附近斯柯达附近斯柯达附近斯柯达附近斯柯达附近斯柯达附近斯柯达附近斯柯达附近斯柯达附近斯柯达附近斯柯达附近斯柯达附近斯柯达附近斯柯达附近', '', 0, 0, NULL, '0000-00-00 00:00:00', '0000-00-00 00:00:00');
COMMIT;

-- ----------------------------
-- Table structure for ks_goods
-- ----------------------------
DROP TABLE IF EXISTS `ks_goods`;
CREATE TABLE `ks_goods` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) unsigned NOT NULL COMMENT '商品类型:1-文章',
  `name` varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '商品名称',
  `description` text CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '商品简介',
  `image` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '商品大图',
  `small_image` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '商品小图',
  `tag` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '1:免费;2:会员;3:付费',
  `price` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '价格[原价]',
  `cover_tag` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '封面标签 ks_type_common id',
  `view_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '点击人数',
  `favorite_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '收藏数',
  `praise_num` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '点赞数',
  `is_recommend` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否推荐；0:不推荐;1:推荐',
  `order` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '商品排序',
  `is_verify` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否审核上线 0未审核 1已审核',
  `recommend_stars` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '推荐指数',
  `is_comment` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否开启评论',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_type_verify_recommend_order` (`type`,`is_verify`,`is_recommend`,`order`) USING BTREE,
  KEY `idx_tag` (`tag`) USING BTREE,
  KEY `idx_recommend` (`is_recommend`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COMMENT='商品表';

-- ----------------------------
-- Records of ks_goods
-- ----------------------------
BEGIN;
INSERT INTO `ks_goods` VALUES (9, 1, '欢迎光临我的新文章s', '简要s', 'article/2020-03-08-22-41-10-1847.jpeg', 'article/2020-03-08-18-41-37-7887.jpg', 1, 100, 1, 3, 2, 3, 0, 5, 0, 4, 1, '2019-07-04 09:52:14', '2019-07-26 03:45:12', NULL, 0);
INSERT INTO `ks_goods` VALUES (12, 1, 'mysql 积累', '记录一些日常sql知识和查询语句优化', 'WechatIMG103.jpeg', '2020-03-07-20-55-37-4059.jpeg', 1, 0, 1, 49, 0, 0, 1, 0, 1, 3, 1, '2019-07-04 09:52:14', '2019-08-14 16:46:18', NULL, 0);
INSERT INTO `ks_goods` VALUES (17, 1, 'nginx rtmp 应用层协议', '本文使用的流媒体服务器的搭建是基于rtmp（Real Time Message Protocol）协议的，rtmp协议是应用层的协议，要依靠底层的传输层协议，比如tcp协议来保证信息传输的可靠性。来搭建自己的直播间测试服务,用的的技术 nginx HLS rtmp', 'images/5ca43a92533706a045b8dfb0fbdcdd4b.jpg', 'images/403122bbf8e5611db5bbdd5bb57a5fd1.jpg', 1, 0, 0, 116, 0, 5, 0, 1, 1, 5, 1, '2019-08-10 11:56:20', '2019-08-10 12:08:30', NULL, 0);
INSERT INTO `ks_goods` VALUES (18, 2, '2019年恐怖《魔童》BD中英双字幕', '◎译 名 魔童/明亮的燃烧/灵异乍现(台)/灼烧 ◎片 名 Brightburn ◎年 代 2019 ◎产 地 美国 ◎类 别 恐怖 ◎语 言 英语 ◎上映日期 2019-05-24(美国) ◎IMDb评分 6.2/10 from 29037 users ◎豆瓣评分 5.5/10 from 4,961 users ◎片 长 91分钟 ◎导 演 大卫雅洛维斯基', 'https://extraimage.net/images/2019/08/12/c642dfc25e135a6d468d11faf3372acb.jpg', 'https://extraimage.net/images/2019/08/12/c642dfc25e135a6d468d11faf3372acb.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:20:39', '2019-08-14 16:20:39', NULL, 0);
INSERT INTO `ks_goods` VALUES (19, 2, '2019年奇幻动作《雷霆沙赞！》BD国英双语双字', '雷霆沙赞！ BD国英双语双字 2019年奇幻动作 ◎译 名 雷霆沙赞！/奇迹队长/沙赞/沙赞！(台)/沙赞！神力集结(港)/神奇上尉/雷霆沙赞 ◎片 名 Shazam!/Billy Batson and the Legend of Shazam!/Franklin ◎年 代 2019 ◎产 地 美国 ◎类 别 动作/奇幻/冒险 ◎语 言 英语', 'https://extraimage.net/images/2019/06/27/de2f70224c435a441d5bede56effed44.jpg', 'https://extraimage.net/images/2019/06/27/de2f70224c435a441d5bede56effed44.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:24:36', '2019-08-14 16:25:07', NULL, 0);
INSERT INTO `ks_goods` VALUES (20, 2, '2019年喜剧《全民追女王》BD中英双字幕', '全民追女王 BD中英双字幕 2019年喜剧 ◎译 名 全民追女王/全民追女神/希望渺茫/弗拉斯基/弗莱斯基/漫长的过程/索爆高低恋(港)/选情尬翻天(台)/高风险赌注 ◎片 名 Long Shot/Flarsky ◎年 代 2019 ◎产 地 美国 ◎类 别 喜剧 ◎语 言 英语 ◎字 幕 中英双字幕 ◎上映日', 'https://extraimage.net/images/2019/07/25/f7d0eac37404497b869c49697b253319.jpg', 'https://extraimage.net/images/2019/07/25/f7d0eac37404497b869c49697b253319.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:26:18', '2019-08-14 16:26:46', NULL, 0);
INSERT INTO `ks_goods` VALUES (21, 2, '2019年剧情音乐《火箭人》BD中英双字幕', '◎译 名 火箭人/摇滚太空人(港)/火箭客 ◎片 名 Rocketman ◎年 代 2019 ◎产 地 英国/美国 ◎类 别 剧情/同性/音乐/传记 ◎语 言 英语 ◎字 幕 中英双字幕 ◎文件格式 x264 + aac ◎视频尺寸 1280 x 720 ◎文件大 1CD ◎上映日期 2019-05-16(戛纳电影节)/2019-05-22', 'https://extraimage.net/images/2019/08/11/556929f5b65b1ae8693bafd0f0ae73bb.jpg', 'https://extraimage.net/images/2019/08/11/556929f5b65b1ae8693bafd0f0ae73bb.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:00', '2019-08-14 16:31:00', NULL, 0);
INSERT INTO `ks_goods` VALUES (22, 2, '2019年剧情传记《JT・莱罗伊》BD中英双字幕', '◎译 名 JT莱罗伊/JT勒罗伊/绝世大作 ◎片 名 JT Leroy/Jeremiah Terminator LeRoy ◎年 代 2018 ◎产 地 美国 ◎类 别 剧情/传记 ◎语 言 英语 ◎字 幕 中英双字幕 ◎文件格式 x264 + aac ◎视频尺寸 1280 x 720 ◎文件大 1CD ◎上映日期 2018-09-10(多伦多电影节)', 'https://extraimage.net/images/2019/08/11/79ab54a8f343dea1e71960c16804cac9.jpg', 'https://extraimage.net/images/2019/08/11/79ab54a8f343dea1e71960c16804cac9.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:01', '2019-08-14 16:31:01', NULL, 0);
INSERT INTO `ks_goods` VALUES (23, 2, '2018年惊悚恐怖《八四年夏天》BD英语中字', '◎译 名 八四年夏天 ◎片 名 Summer of \'84 ◎年 代 2017 ◎产 地 美国 ◎类 别 剧情/惊悚/恐怖 ◎语 言 英语 ◎字 幕 中文 ◎上映日期 2018-01-22(圣丹斯电影节) / 2018-08-10(美国) ◎IMDb评分 6.7/10 from 23923 users ◎豆瓣评分 6.1/10 from 1,585 users ◎文件格', 'https://extraimage.net/images/2018/09/04/355fd4163593ead728c7116fd1c9d23c.jpg', 'https://extraimage.net/images/2018/09/04/355fd4163593ead728c7116fd1c9d23c.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:01', '2019-08-14 16:31:01', NULL, 0);
INSERT INTO `ks_goods` VALUES (24, 2, '2019年奇幻冒险《阿拉丁》BD中英双字幕', '◎译 名 阿拉丁/阿拉丁真人版 ◎片 名 Aladdin ◎年 代 2019 ◎产 地 美国 ◎类 别 爱情 / 歌舞 / 奇幻 / 冒险 ◎语 言 英语 ◎字 幕 中英双字幕 ◎文件格式 x264 + aac ◎视频尺寸 1280 x 720 ◎文件大 1CD ◎上映日期 2019-05-24(美国/中国) ◎IMDb评分 7.3/10 fro', 'https://extraimage.net/images/2019/08/10/4dd5a9e7babfec548c4f103475e1eb82.jpg', 'https://extraimage.net/images/2019/08/10/4dd5a9e7babfec548c4f103475e1eb82.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:02', '2019-08-14 16:31:02', NULL, 0);
INSERT INTO `ks_goods` VALUES (25, 2, '2019年喜剧《一条狗的使命2》BD中英双字幕', '◎译 名 一条狗的使命2/一条狗的旅程/为了与你相遇2/再见亦是狗朋友2(港)/狗狗的旅程(台) ◎片 名 A Dog\'s Journey/A Dog\'s Purpose 2 ◎年 代 2019 ◎产 地 美国 ◎类 别 剧情/喜剧/家庭 ◎语 言 英语 ◎字 幕 中英双字幕 ◎文件格式 x264 + aac ◎视频尺寸 1280 x 72', 'https://extraimage.net/images/2019/08/09/30e6b763e291dd23359bdf1ca7f08c91.jpg', 'https://extraimage.net/images/2019/08/09/30e6b763e291dd23359bdf1ca7f08c91.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:03', '2019-08-14 16:31:03', NULL, 0);
INSERT INTO `ks_goods` VALUES (26, 2, '2019年惊悚悬疑《复仇》BD中字', '◎译 名 复仇 ◎片 名 Badla ◎年 代 2019 ◎产 地 印度 ◎类 别 惊悚/悬疑 ◎语 言 印地语 ◎字 幕 中文 ◎文件格式 x264 + aac ◎视频尺寸 1280 x 720 ◎文件大 1CD ◎上映日期 2019-03-08(印度) ◎IMDb评分 8.0/10 from 11733 users ◎豆瓣评分 6.6/10 from 159 u', 'https://extraimage.net/images/2019/08/09/ccdd49b26839bf26f708e9ae6fbf83f6.jpg', 'https://extraimage.net/images/2019/08/09/ccdd49b26839bf26f708e9ae6fbf83f6.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:04', '2019-08-14 16:31:04', NULL, 0);
INSERT INTO `ks_goods` VALUES (27, 2, '2019年惊悚剧情《女特工》BD中英双字幕', '女特工 BD中英双字幕 2019年惊悚剧情 ◎译 名 女特工 ◎片 名 The Operative ◎年 代 2019 ◎产 地 美国/以色列 ◎类 别 惊悚 ◎语 言 英语 ◎字 幕 中英双字幕 ◎文件格式 x264 + aac ◎视频尺寸 1280 x 720 ◎文件大 1CD ◎上映日期 2019-02-10(柏林电影节) ◎IMDb', 'https://extraimage.net/images/2019/08/08/91431644a9af61d8c1f1da025baa27e3.jpg', 'https://extraimage.net/images/2019/08/08/91431644a9af61d8c1f1da025baa27e3.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:07', '2019-08-14 16:31:07', NULL, 0);
INSERT INTO `ks_goods` VALUES (28, 2, '2019年高分获奖剧情《寄生虫》BD韩语中字', '寄生虫 BD韩语中字 2019年高分获奖剧情 ◎译 名 寄生虫/上流寄生族(港)/寄生上流(台) ◎片 名 Gisaengchung/Parasite ◎年 代 2019 ◎产 地 韩国 ◎类 别 剧情 ◎语 言 韩语 ◎字 幕 中文 ◎文件格式 x264 + aac ◎视频尺寸 1280 x 720 ◎文件大 1CD ◎上映日期 2019', 'https://extraimage.net/images/2019/08/08/5f23a28434de4516a7694deda137ecb4.jpg', 'https://extraimage.net/images/2019/08/08/5f23a28434de4516a7694deda137ecb4.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:08', '2019-08-14 16:31:08', NULL, 0);
INSERT INTO `ks_goods` VALUES (29, 2, '2019年悬疑《十二个想死的孩子》BD日语中字', '◎译 名 十二个想死的孩子/12个想死的孩子/12个想死的孩子们 ◎片 名 十二人の死にたい子どもたち ◎年 代 2019 ◎产 地 日本 ◎类 别 悬疑 ◎语 言 日语 ◎字 幕 中文 ◎文件格式 x264 + aac ◎视频尺寸 1280 x 720 ◎文件大 1CD ◎上映日期 2019-01-25(日本) ◎IMD', 'https://extraimage.net/images/2019/08/02/9bf80660d593caca3cd8f912be334480.jpg', 'https://extraimage.net/images/2019/08/02/9bf80660d593caca3cd8f912be334480.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:09', '2019-08-14 16:31:09', NULL, 0);
INSERT INTO `ks_goods` VALUES (30, 2, '2019年惊悚剧情《伦敦往事》BD中英双字幕', '伦敦往事 BD中英双字幕 2019年惊悚剧情 ◎译 名 伦敦往事 ◎片 名 Once Upon a Time in London ◎年 代 2019 ◎产 地 英国 ◎类 别 犯罪/剧情/惊悚 ◎语 言 英语 ◎字 幕 中英双字幕 ◎上映日期 2019-04-19(英国) ◎IMDb评分 7.2/10 from 2959 users ◎豆瓣评分 5.6/10', 'https://extraimage.net/images/2019/08/06/59bf94e7c946de91ae841c45986466c6.jpg', 'https://extraimage.net/images/2019/08/06/59bf94e7c946de91ae841c45986466c6.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:10', '2019-08-14 16:31:10', NULL, 0);
INSERT INTO `ks_goods` VALUES (31, 2, '019年惊悚动作《报仇雪恨/血债血偿》BD中英双字幕', '◎译 名 报仇雪恨/血债血偿 ◎片 名 A Score To Settle ◎年 代 2018 ◎产 地 美国 ◎类 别 动作/惊悚/犯罪 ◎语 言 英语 ◎字 幕 中英双字幕 ◎上映日期 2019-08-02(美国) ◎IMDb评分 4.5/10 from 381 users ◎文件格式 x264 + aac ◎视频尺寸 1280 x 720 ◎文件大', 'https://extraimage.net/images/2019/08/05/446f6655019f190968cee62f363af8eb.jpg', 'https://extraimage.net/images/2019/08/05/446f6655019f190968cee62f363af8eb.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:11', '2019-08-14 16:31:11', NULL, 0);
INSERT INTO `ks_goods` VALUES (32, 2, '2019年获奖剧情《沦落人》BD粤语中英双字', '沦落人 BD粤语中英双字 2019年剧情 ◎译 名 沦落人/Still Human ◎片 名 S落人 ◎年 代 2018 ◎产 地 香港 ◎类 别 剧情 ◎语 言 粤语 ◎字 幕 中英双字幕 ◎上映日期 2018-11-08(台湾金马奖) / 2019-04-11(香港) ◎IMDb评分 7.4/10 from 351 users ◎豆瓣评分 8.2/10', 'https://extraimage.net/images/2019/08/04/733e0191fe8ed58b0dc0df4f63c87590.jpg', 'https://extraimage.net/images/2019/08/04/733e0191fe8ed58b0dc0df4f63c87590.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:12', '2019-08-14 16:31:12', NULL, 0);
INSERT INTO `ks_goods` VALUES (33, 2, '2019年剧情《妈阁是座城》HD国语中英双字', '妈阁是座城 HD国语中英双字 2019年剧情 ◎译 名 A City Called Macau/Macau Is The Seat of The City ◎片 名 妈阁是座城 ◎年 代 2018 ◎产 地 中国 ◎类 别 剧情 ◎语 言 普通话 ◎字 幕 中英双字幕 ◎上映日期 2018-12-27(香港点映) / 2019-06-14(中国) ◎IMDb评分', 'https://extraimage.net/images/2019/08/02/4011c97f79991f6128b1362e7a96844f.jpg', 'https://extraimage.net/images/2019/08/02/4011c97f79991f6128b1362e7a96844f.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:13', '2019-08-14 16:31:13', NULL, 0);
INSERT INTO `ks_goods` VALUES (34, 2, '2018年剧情《他们(国际版)》BD中英双字幕', '◎译 名 他们（国际版）/上流世界(台) ◎片 名 Loro ◎年 代 2018 ◎产 地 意大利 ◎类 别 剧情 ◎语 言 意大利语 ◎字 幕 中英双字幕 ◎上映日期 2018-09-13(意大利) ◎IMDb评分 6.7/10 from 830 users ◎豆瓣评分 7.5/10 from 94 users ◎文件格式 x264 + aac ◎视频', 'https://extraimage.net/images/2019/08/02/8984e1a85fab753d40a9bfb24b6ecf58.jpg', 'https://extraimage.net/images/2019/08/02/8984e1a85fab753d40a9bfb24b6ecf58.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:14', '2019-08-14 16:31:14', NULL, 0);
INSERT INTO `ks_goods` VALUES (35, 2, '2018年惊悚喜剧《银湖之底》BD英语中字', '◎译 名 银湖之底/银湖之下 ◎片 名 Under the Silver Lake ◎年 代 2018 ◎产 地 美国 ◎类 别 剧情 / 喜剧 / 惊悚 / 犯罪 ◎语 言 英语 ◎字 幕 中文 ◎上映日期 2018-05-15(戛纳电影节) / 2019-04-19(美国) ◎IMDb评分 6.5/10 from 19799 users ◎豆瓣评分 6.5/10 f', 'https://extraimage.net/images/2019/08/02/d4aefaf25821aab6f3b037b5a65a90d2.jpg', 'https://extraimage.net/images/2019/08/02/d4aefaf25821aab6f3b037b5a65a90d2.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:15', '2019-08-14 16:31:15', NULL, 0);
INSERT INTO `ks_goods` VALUES (36, 2, '2019年惊悚剧情《红海潜水俱乐部》BD中英双字幕', '红海潜水俱乐部 BD中英双字幕 2019年惊悚剧情 ◎译 名 红海潜水俱乐部/红海潜水度假村/红海救援/红海深潜(台) ◎片 名 The Red Sea Diving Resort ◎年 代 2019 ◎产 地 美国 ◎类 别 剧情/惊悚/历史 ◎语 言 英语 ◎字 幕 中英双字幕 ◎上映日期 2019-06-28(旧金山犹', 'https://extraimage.net/images/2019/08/01/ca7c7e6f39f113e0ff2b2ddce0bd926c.jpg', 'https://extraimage.net/images/2019/08/01/ca7c7e6f39f113e0ff2b2ddce0bd926c.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:16', '2019-08-14 16:31:16', NULL, 0);
INSERT INTO `ks_goods` VALUES (37, 2, '2018年剧情《顿巴斯/疯狂的边境》BD中英双字幕', '顿巴斯/疯狂的边境 BD中英双字幕 2018年剧情 ◎译 名 顿巴斯/疯狂的边境(港)/着魔的国境(台) ◎片 名 Донбас/Donbass ◎年 代 2018 ◎产 地 德国/乌克兰/法国/荷兰/罗马尼亚 ◎类 别 剧情 ◎语 言 乌克兰语/俄语 ◎字 幕 中英双字幕 ◎上映日期 2018-05-09(戛纳', 'https://extraimage.net/images/2019/07/31/5bb15d810dfd930919483bb26cf9fb83.jpg', 'https://extraimage.net/images/2019/07/31/5bb15d810dfd930919483bb26cf9fb83.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:17', '2019-08-14 16:31:17', NULL, 0);
INSERT INTO `ks_goods` VALUES (38, 2, '2019年科幻动作《黑衣人：全球追缉》HD韩版中字', '黑衣人：全球追缉 HD韩版中字 2019年科幻动作 ◎译 名 黑衣人：全球追缉/MIB星际战警：跨国行动(台)/黑衣人：国际守护者/黑衣人23/黑衣人4/黑衣人外传/黑超特警组：反转世界(港)/黑超特警队4 ◎片 名 Men in Black International/Men in Black 4/Men in Black IV/MIB 2', 'https://extraimage.net/images/2019/07/29/79b913efe06f6bc2cce8544003b40862.jpg', 'https://extraimage.net/images/2019/07/29/79b913efe06f6bc2cce8544003b40862.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:20', '2019-08-14 16:31:20', NULL, 0);
INSERT INTO `ks_goods` VALUES (39, 2, '2019年科幻动作《复仇者联盟4：终局之战》HD中英双字幕', '复仇者联盟4：终局之战 HD中英双字幕 2019年科幻动作 ◎译 名 复仇者联盟4：终局之战/复仇者联盟3：无尽之战(下)/复联4 ◎片 名 Avengers: Endgame/Avengers: Infinity War - Part II/AVG4/The Avengers 3: Part 2/The Avengers 4: Endgame ◎年 代 2019 ◎产 地 美国', 'https://extraimage.net/images/2019/07/29/3cc4813f9364426c8beb4da39eb132d1.jpg', 'https://extraimage.net/images/2019/07/29/3cc4813f9364426c8beb4da39eb132d1.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:21', '2019-08-14 16:31:21', NULL, 0);
INSERT INTO `ks_goods` VALUES (40, 2, '2019年动作喜剧《城市猎人》BD法语中字', '城市猎人 BD法语中字 2019年动作喜剧 ◎译 名 城市猎人 ◎片 名 Nicky Larson/Nicky Larson and Cupidons Perfume/Nicky Larson et le Parfum de Cupidon ◎年 代 2019 ◎产 地 法国 ◎类 别 喜剧/动作/犯罪 ◎语 言 法语 ◎字 幕 中文 ◎上映日期 2019-02-06(法国) ◎', 'https://extraimage.net/images/2019/07/29/0e698985a424dab293826d167a36baa0.jpg', 'https://extraimage.net/images/2019/07/29/0e698985a424dab293826d167a36baa0.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:23', '2019-08-14 16:31:23', NULL, 0);
INSERT INTO `ks_goods` VALUES (41, 2, '2019年剧情传记《托尔金》BD中英双字幕', '托尔金 BD中英双字幕 2019年剧情传记 ◎译 名 托尔金 ◎片 名 Tolkien / A Light in the Darkness / J.R.R. Tolkien ◎年 代 2019 ◎产 地 美国 ◎类 别 剧情 / 传记 ◎语 言 英语 ◎字 幕 中英双字幕 ◎上映日期 2019-05-03(英国) / 2019-05-10(美国) ◎IMDb评分 7.0/', 'https://extraimage.net/images/2019/07/28/fea6c02b803aed78d3fd7f1b32d80c41.jpg', 'https://extraimage.net/images/2019/07/28/fea6c02b803aed78d3fd7f1b32d80c41.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:25', '2019-08-14 16:31:25', NULL, 0);
INSERT INTO `ks_goods` VALUES (42, 2, '2019年历史战争《凯萨里》BD中英双字幕', '◎译 名 凯萨里/橙黄色的头巾 ◎片 名 Kesari ◎年 代 2019 ◎产 地 印度 ◎类 别 历史 / 战争 ◎语 言 印地语 ◎字 幕 中英双字幕 ◎文件格式 x264 + aac ◎视频尺寸 1280 x 720 ◎文件大 1CD ◎上映日期 2019-03-21(印度) ◎IMDb评分 7.5/10 from 9152 users ◎豆', 'https://extraimage.net/images/2019/08/12/4f4c0a3b94cbcdddb34eb6e74061fa20.jpg', 'https://extraimage.net/images/2019/08/12/4f4c0a3b94cbcdddb34eb6e74061fa20.jpg', 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, '2019-08-14 16:31:53', '2019-08-14 16:31:53', NULL, 0);
INSERT INTO `ks_goods` VALUES (43, 1, 'Mysql show processlist 排查问题', 'Mysql show processlist 排查问题', 'images/531565678883_.pic.jpg', 'images/57e590811619b8bebd4343c161a2d965.jpg', 1, 0, 0, 2, 3, 4, 0, 1, 1, 2, 0, '2019-08-14 16:50:03', '2019-08-14 16:50:03', NULL, 0);
INSERT INTO `ks_goods` VALUES (46, 1, '测试', '大事发生', '', '', 0, 0, 0, 0, 0, 0, 0, 1, 1, 4, 1, NULL, NULL, NULL, 0);
INSERT INTO `ks_goods` VALUES (47, 3, 'aaa', 'dddd', '', 'article/2020-05-04-07-47-30-8081.jpg', 0, 0, 0, 34, 0, 0, 0, 1, 1, 3, 1, NULL, NULL, NULL, 0);
COMMIT;

-- ----------------------------
-- Table structure for ks_goods_type_relation
-- ----------------------------
DROP TABLE IF EXISTS `ks_goods_type_relation`;
CREATE TABLE `ks_goods_type_relation` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `type_common_id` int(11) unsigned NOT NULL DEFAULT '0',
  `goods_id` int(11) unsigned NOT NULL DEFAULT '0',
  `created_at` varchar(20) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COMMENT='商品标签关系表';

-- ----------------------------
-- Records of ks_goods_type_relation
-- ----------------------------
BEGIN;
INSERT INTO `ks_goods_type_relation` VALUES (14, 2, 12, '');
INSERT INTO `ks_goods_type_relation` VALUES (15, 2, 44, '');
INSERT INTO `ks_goods_type_relation` VALUES (16, 3, 44, '');
INSERT INTO `ks_goods_type_relation` VALUES (17, 2, 9, '');
INSERT INTO `ks_goods_type_relation` VALUES (18, 3, 9, '');
INSERT INTO `ks_goods_type_relation` VALUES (19, 6, 45, '');
INSERT INTO `ks_goods_type_relation` VALUES (20, 8, 45, '');
INSERT INTO `ks_goods_type_relation` VALUES (21, 6, 46, '');
INSERT INTO `ks_goods_type_relation` VALUES (22, 7, 46, '');
INSERT INTO `ks_goods_type_relation` VALUES (27, 7, 47, '');
INSERT INTO `ks_goods_type_relation` VALUES (28, 8, 47, '');
COMMIT;

-- ----------------------------
-- Table structure for ks_mac
-- ----------------------------
DROP TABLE IF EXISTS `ks_mac`;
CREATE TABLE `ks_mac` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `download_url` varchar(255) NOT NULL COMMENT '下载链接',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `goods_id` int(11) unsigned NOT NULL DEFAULT '0',
  `link` varchar(255) DEFAULT '' COMMENT '详情链接',
  `content` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ks_mac
-- ----------------------------
BEGIN;
INSERT INTO `ks_mac` VALUES (1, 'ddd', NULL, 47, '', '<p>addfdadfdddd</p>');
COMMIT;

-- ----------------------------
-- Table structure for ks_menus
-- ----------------------------
DROP TABLE IF EXISTS `ks_menus`;
CREATE TABLE `ks_menus` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `order` int(11) NOT NULL DEFAULT '0',
  `title` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `icon` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `url` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '路由',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of ks_menus
-- ----------------------------
BEGIN;
INSERT INTO `ks_menus` VALUES (1, 0, 3, '网站管理', 'layui-icon-set', '/', '2020-02-16 23:47:06', '2020-02-22 11:49:44', 0);
INSERT INTO `ks_menus` VALUES (2, 1, 4, '用户列表', 'layui-icon-username', '/admin/userIndex', '2020-02-16 23:48:01', '2020-04-09 23:25:55', 0);
INSERT INTO `ks_menus` VALUES (3, 1, 2, '菜单列表', 'layui-icon-menu-fill', '/admin/menuIndex', '2020-02-16 23:48:34', '2020-04-09 23:26:44', 0);
INSERT INTO `ks_menus` VALUES (4, 0, 1, '商品管理', 'layui-icon-read', '/', '2020-02-16 23:49:05', '2020-05-02 18:05:55', 0);
INSERT INTO `ks_menus` VALUES (5, 4, 101, '文章列表', 'layui-icon-survey', '/admin/articleIndex', '2020-02-16 23:49:32', '2020-04-09 23:28:10', 0);
INSERT INTO `ks_menus` VALUES (6, 0, 2, '会员管理', 'layui-icon-user', '/', '2020-02-20 00:19:26', '2020-02-22 11:49:54', 0);
INSERT INTO `ks_menus` VALUES (7, 6, 201, '会员列表', 'layui-icon-right', '/admin/customerIndex', '2020-02-20 10:13:57', '2020-04-09 09:16:52', 0);
INSERT INTO `ks_menus` VALUES (11, 0, 0, 'weeds', 'layui-icon-password', '', '2020-02-23 01:00:12', '2020-02-23 01:02:46', 1);
INSERT INTO `ks_menus` VALUES (12, 11, 2, '下级', 'layui-icon-password', '/', '2020-02-23 01:02:27', '2020-02-23 01:02:27', 0);
INSERT INTO `ks_menus` VALUES (13, 1, 4, '商品标签', 'layui-icon-note', '/admin/goodsTypeIndex', '2020-04-04 20:56:38', '2020-04-09 09:16:30', 0);
INSERT INTO `ks_menus` VALUES (14, 4, 3, 'MAC列表', 'layui-icon-ios', '/admin/macIndex', '2020-05-02 18:07:19', '2020-05-03 23:14:05', 0);
COMMIT;

-- ----------------------------
-- Table structure for ks_movie
-- ----------------------------
DROP TABLE IF EXISTS `ks_movie`;
CREATE TABLE `ks_movie` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `download_url` varchar(255) NOT NULL COMMENT '下载链接',
  `deleted_at` timestamp NULL DEFAULT NULL,
  `goods_id` int(11) unsigned NOT NULL DEFAULT '0',
  `link` varchar(255) DEFAULT '' COMMENT '详情链接',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ks_movie
-- ----------------------------
BEGIN;
INSERT INTO `ks_movie` VALUES (2, 'ftp://ygdy8:ygdy8@yg90.dydytt.net:8596/阳光电影www.ygdy8.com.魔童.BD.720p.中英双字幕.mkv', NULL, 18, '/html/gndy/dyzz/20190813/58984.html');
INSERT INTO `ks_movie` VALUES (3, 'ftp://ygdy8:ygdy8@yg90.dydytt.net:8582/阳光电影www.ygdy8.com.雷霆沙赞！.BD.720p.国英双语双字.mkv', NULL, 19, '/html/gndy/dyzz/20190727/58905.html');
INSERT INTO `ks_movie` VALUES (4, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:4206/阳光电影www.ygdy8.com.全民追女王.BD.720p.中英双字幕.mkv', NULL, 20, '/html/gndy/dyzz/20190727/58906.html');
INSERT INTO `ks_movie` VALUES (5, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:8492/阳光电影www.ygdy8.com.火箭人.BD.720p.中英双字幕.mkv', NULL, 21, '/html/gndy/dyzz/20190812/58983.html');
INSERT INTO `ks_movie` VALUES (6, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:7492/阳光电影www.ygdy8.com.JT・莱罗伊.BD.720p.中英双字幕.mkv', NULL, 22, '/html/gndy/dyzz/20190812/58982.html');
INSERT INTO `ks_movie` VALUES (7, 'ftp://ygdy8:ygdy8@yg90.dydytt.net:8595/阳光电影www.ygdy8.com.八四年夏天.BD.720p.英语中字.mkv', NULL, 23, '/html/gndy/dyzz/20190812/58981.html');
INSERT INTO `ks_movie` VALUES (8, 'ftp://ygdy8:ygdy8@yg90.dydytt.net:8594/阳光电影www.ygdy8.com.阿拉丁.BD.720p.中英双字幕.mkv', NULL, 24, '/html/gndy/dyzz/20190811/58978.html');
INSERT INTO `ks_movie` VALUES (9, 'ftp://ygdy8:ygdy8@yg90.dydytt.net:8593/阳光电影www.ygdy8.com.一条狗的使命2.BD.720p.中英双字幕.mkv', NULL, 25, '/html/gndy/dyzz/20190810/58975.html');
INSERT INTO `ks_movie` VALUES (10, 'ftp://ygdy8:ygdy8@yg72.dydytt.net:8381/阳光电影www.ygdy8.com.复仇.BD.720p.中字.mkv', NULL, 26, '/html/gndy/dyzz/20190810/58974.html');
INSERT INTO `ks_movie` VALUES (11, 'ftp://ygdy8:ygdy8@yg90.dydytt.net:8592/阳光电影www.ygdy8.com.女特工.BD.720p.中英双字幕.mkv', NULL, 27, '/html/gndy/dyzz/20190809/58971.html');
INSERT INTO `ks_movie` VALUES (12, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:8490/阳光电影www.ygdy8.com.寄生虫.BD.720p.韩语中字.mkv', NULL, 28, '/html/gndy/dyzz/20190809/58970.html');
INSERT INTO `ks_movie` VALUES (13, 'ftp://ygdy8:ygdy8@yg90.dydytt.net:8591/阳光电影www.ygdy8.com.十二个想死的孩子.BD.720p.日语中字.mkv', NULL, 29, '/html/gndy/dyzz/20190808/58967.html');
INSERT INTO `ks_movie` VALUES (14, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:4208/阳光电影www.ygdy8.com.伦敦往事.BD.720p.中英双字幕.mkv', NULL, 30, '/html/gndy/dyzz/20190807/58961.html');
INSERT INTO `ks_movie` VALUES (15, 'ftp://ygdy8:ygdy8@yg90.dydytt.net:8590/阳光电影www.ygdy8.com.报仇雪恨.BD.720p.中英双字幕.mkv', NULL, 31, '/html/gndy/dyzz/20190806/58956.html');
INSERT INTO `ks_movie` VALUES (16, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:8487/阳光电影www.ygdy8.com.沦落人.BD.720p.粤语中字.mkv', NULL, 32, '/html/gndy/dyzz/20190805/58947.html');
INSERT INTO `ks_movie` VALUES (17, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:7487/阳光电影www.ygdy8.com.妈阁是座城.HD.720p.国语中字.mkv', NULL, 33, '/html/gndy/dyzz/20190805/58946.html');
INSERT INTO `ks_movie` VALUES (18, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:8486/阳光电影www.ygdy8.com.他们(国际版).BD.720p.中英双字幕.mkv', NULL, 34, '/html/gndy/dyzz/20190803/58928.html');
INSERT INTO `ks_movie` VALUES (19, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:7486/阳光电影www.ygdy8.com.银湖之底.BD.720p.中字.mkv', NULL, 35, '/html/gndy/dyzz/20190803/58927.html');
INSERT INTO `ks_movie` VALUES (20, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:8485/阳光电影www.ygdy8.com.红海潜水俱乐部.BD.720p.中英双字幕.mkv', NULL, 36, '/html/gndy/dyzz/20190802/58921.html');
INSERT INTO `ks_movie` VALUES (21, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:6178/阳光电影www.ygdy8.com.顿巴斯.BD.720p.中英双字幕.mkv', NULL, 37, '/html/gndy/dyzz/20190801/58918.html');
INSERT INTO `ks_movie` VALUES (22, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:8483/阳光电影www.ygdy8.com.黑衣人：全球追缉.HD.720p.韩版中字.mkv', NULL, 38, '/html/gndy/dyzz/20190731/58915.html');
INSERT INTO `ks_movie` VALUES (23, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:4207/阳光电影www.ygdy8.com.复仇者联盟4：终局之战.HD.720p.中英双字幕.mkv', NULL, 39, '/html/gndy/dyzz/20190730/58912.html');
INSERT INTO `ks_movie` VALUES (24, 'ftp://ygdy8:ygdy8@yg90.dydytt.net:8584/阳光电影www.ygdy8.com.城市猎人.BD.720p.法语中字.mkv', NULL, 40, '/html/gndy/dyzz/20190730/58911.html');
INSERT INTO `ks_movie` VALUES (25, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:8482/阳光电影www.ygdy8.com.托尔金.BD.720p.中英双字幕.mkv', NULL, 41, '/html/gndy/dyzz/20190728/58910.html');
INSERT INTO `ks_movie` VALUES (26, 'ftp://ygdy8:ygdy8@yg45.dydytt.net:4209/阳光电影www.ygdy8.com.凯萨里.BD.720p.中英双字幕.mkv', NULL, 42, '/html/gndy/dyzz/20190813/58985.html');
COMMIT;

-- ----------------------------
-- Table structure for ks_roles
-- ----------------------------
DROP TABLE IF EXISTS `ks_roles`;
CREATE TABLE `ks_roles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL COMMENT '名称',
  `fun_ids` varchar(255) NOT NULL COMMENT '拥有的权限',
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  `updated_at` datetime NOT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='权限表';

-- ----------------------------
-- Records of ks_roles
-- ----------------------------
BEGIN;
INSERT INTO `ks_roles` VALUES (1, '超管', ',1,', 0, '2020-02-07 14:43:54', '2020-02-07 14:43:57');
INSERT INTO `ks_roles` VALUES (2, '运营', ',1,', 0, '2020-02-07 14:43:54', '2020-02-07 14:43:57');
INSERT INTO `ks_roles` VALUES (3, '财务', ',1,', 0, '2020-02-07 14:43:54', '2020-02-07 14:43:57');
COMMIT;

-- ----------------------------
-- Table structure for ks_tutors
-- ----------------------------
DROP TABLE IF EXISTS `ks_tutors`;
CREATE TABLE `ks_tutors` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `sex` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '性别  0:未知 1:男  2:女',
  `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '讲师名',
  `image` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '讲师图片',
  `small_image` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '讲师小图片',
  `description` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT '简介',
  `content` text COLLATE utf8_unicode_ci NOT NULL COMMENT '详细内容',
  `is_verify` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否审核上线： 0:否 1:是  ',
  `created_at` int(11) NOT NULL DEFAULT '0',
  `updated_at` int(11) NOT NULL DEFAULT '0',
  `deleted_at` int(11) DEFAULT '0',
  `recommend_stars` tinyint(1) NOT NULL DEFAULT '1' COMMENT '推荐指数',
  PRIMARY KEY (`id`),
  KEY `tutors_name_unique` (`name`),
  KEY `idx_verify` (`is_verify`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of ks_tutors
-- ----------------------------
BEGIN;
INSERT INTO `ks_tutors` VALUES (1, 1, '野草ks', 'www.baidu.com', 'www.google.com', '野草ks简介', '这个是内容', 1, 0, 0, 0, 4);
INSERT INTO `ks_tutors` VALUES (2, 2, 'weeds', 'www.baidu.com', 'www.google.com', '野草ks简介', '这个是内容', 1, 0, 0, 0, 4);
COMMIT;

-- ----------------------------
-- Table structure for ks_type_common
-- ----------------------------
DROP TABLE IF EXISTS `ks_type_common`;
CREATE TABLE `ks_type_common` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '标签名',
  `is_verify` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '审核上线：0 否； 1 是',
  `order` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序 越大越前',
  `type` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '1:文章',
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '图标',
  `created_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_type` (`type`) USING BTREE,
  KEY `is_verify` (`is_verify`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='公共类型标签表';

-- ----------------------------
-- Records of ks_type_common
-- ----------------------------
BEGIN;
INSERT INTO `ks_type_common` VALUES (1, '标签1', 1, 12, 1, '', '0000-00-00 00:00:00', '2020-04-09 22:44:29');
INSERT INTO `ks_type_common` VALUES (2, '标签2', 1, 2, 1, '', '0000-00-00 00:00:00', '2020-04-07 00:06:06');
INSERT INTO `ks_type_common` VALUES (3, '标签3', 1, 2, 1, '', '0000-00-00 00:00:00', '2020-04-09 23:16:28');
INSERT INTO `ks_type_common` VALUES (4, '电影标签1', 1, 1, 2, '', '2020-04-06 17:22:13', '0000-00-00 00:00:00');
INSERT INTO `ks_type_common` VALUES (5, '电影标签2', 0, 2, 2, '', '2020-04-06 17:22:37', '0000-00-00 00:00:00');
INSERT INTO `ks_type_common` VALUES (6, '开发软件', 1, 1, 3, '', '2020-04-06 23:15:04', '0000-00-00 00:00:00');
INSERT INTO `ks_type_common` VALUES (7, '系统软件', 1, 2, 3, '', '2020-04-06 23:15:20', '0000-00-00 00:00:00');
INSERT INTO `ks_type_common` VALUES (8, '图形图像', 1, 3, 3, '', '2020-04-06 23:16:54', '0000-00-00 00:00:00');
INSERT INTO `ks_type_common` VALUES (9, '媒体工具', 1, 4, 3, '', '2020-04-06 23:17:10', '0000-00-00 00:00:00');
INSERT INTO `ks_type_common` VALUES (10, '设计素材', 1, 5, 3, '', '2020-04-06 23:17:27', '0000-00-00 00:00:00');
INSERT INTO `ks_type_common` VALUES (11, '网络工具', 1, 6, 3, '', '2020-04-06 23:18:01', '0000-00-00 00:00:00');
COMMIT;

-- ----------------------------
-- Table structure for ks_users
-- ----------------------------
DROP TABLE IF EXISTS `ks_users`;
CREATE TABLE `ks_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '登录名',
  `mobile` varchar(11) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '手机号',
  `password` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '邮箱',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `salt` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `is_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色',
  PRIMARY KEY (`id`),
  KEY `idx_roleId` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of ks_users
-- ----------------------------
BEGIN;
INSERT INTO `ks_users` VALUES (1, 'weeds', '17620940805', '3a3d95dca53caacaf2ba6acbc9f9d4d6', 'weedsyl@163.com', '2020-02-14 17:09:16', '2020-06-02 18:14:02', '', 0, 1);
INSERT INTO `ks_users` VALUES (2, 'weedsks', '17620940905', '8ac51085d5b4550d3b40f54d67fdf490', 'weeds@163.com', '2020-02-15 22:02:45', '2020-04-06 00:13:31', '', 1, 2);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
