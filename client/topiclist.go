package client

import (
	"regexp"

	"github.com/achillesss/log"
)

/*
<li class=" j_thread_list clearfix" data-field='{&quot;id&quot;:5026842582,&quot;author_name&quot;:&quot;trouble0408&quot;,&quot;first_post_id&quot;:105277795802,&quot;reply_num&quot;:74,&quot;is_bakan&quot;:null,&quot;vid&quot;:&quot;&quot;,&quot;is_good&quot;:null,&quot;is_top&quot;:null,&quot;is_protal&quot;:null,&quot;is_membertop&quot;:null,&quot;frs_tpoint&quot;:null}' >
            <div class="t_con cleafix">

                    <div class="col2_left j_threadlist_li_left">
                <span class="threadlist_rep_num center_text"
                      title="回复">74</span>
            </div>
                <div class="col2_right j_threadlist_li_right ">
            <div class="threadlist_lz clearfix">
                <div class="threadlist_title pull_left j_th_tit ">


    <a href="/p/5026842582" title="互听 互评 互花花 ٩(˃̶͈̀௰˂̶͈́)و" target="_blank" class="j_th_tit ">互听 互评 互花花 ٩(˃̶͈̀௰˂̶͈́)و</a>
</div><div class="threadlist_author pull_right">
    <span class="tb_icon_author "
          title="主题作者: trouble0408"
          data-field='{&quot;user_id&quot;:1789673488}' ><i class="icon_author"></i><span class="frs-author-name-wrap"><a data-field='{&quot;un&quot;:&quot;trouble0408&quot;}' class="frs-author-name j_user_card " href="/home/main/?un=trouble0408&ie=utf-8&fr=frs" target="_blank">trouble0408</a></span><span class="icon_wrap  icon_wrap_theme1 frs_bright_icons "></span>    </span>
    <span class="pull-right is_show_create_time" title="创建时间">3-17</span>
</div>
            </div>
                            <div class="threadlist_detail clearfix">
                    <div class="threadlist_text pull_left">
                                <div class="threadlist_abs threadlist_abs_onlyline ">
            这里Leon baby 可处qy 各位不喜勿喷呐 喜欢的可以关注 送花花 评论 合唱 也可以提意见或者私信我想听什么
        </div>

            <div class="small_wrap j_small_wrap">
                <a href="#" onclick="return false;" class="small_btn_pre j_small_pic_pre" style="display:none"></a>
                <a href="#" onclick="return false;" class="small_btn_next j_small_pic_next" style="display:none"></a>
                <div class="small_list j_small_list cleafix">
                    <div class="small_list_gallery">
                        <ul class="threadlist_media j_threadlist_media clearfix" id="fm5026842582"><li><a class="thumbnail vpic_wrap"><img src="" attr="29336" data-original="http://imgsrc.baidu.com/forum/wh%3D200%2C90%3B/sign=8da641ec7b0e0cf3a0a246f93a76de26/1b03e859252dd42a35c0ed550a3b5bb5c8eab829.jpg"  bpic="http://imgsrc.baidu.com/forum/w%3D580%3B/sign=062174a0d4c451daf6f60ce386c65066/eac4b74543a98226d208c7d98382b9014b90eb07.jpg" class="threadlist_pic j_m_pic "  /></a><div class="threadlist_pic_highlight j_m_pic_light"></div></li><li><a class="thumbnail vpic_wrap"><img src="" attr="87428" data-original="http://imgsrc.baidu.com/forum/wh%3D200%2C90%3B/sign=36c3827bc0ea15ce41bbe80b863016ca/06a9ed380cd79123a225788aa4345982b3b78035.jpg"  bpic="http://imgsrc.baidu.com/forum/w%3D580%3B/sign=e4ebaa3e8b025aafd3327ec3cbd6a964/a08b87d6277f9e2f32f0f1f61630e924b999f303.jpg" class="threadlist_pic j_m_pic "  /></a><div class="threadlist_pic_highlight j_m_pic_light"></div></li><li><a class="thumbnail vpic_wrap"><img src="" attr="90355" data-original="http://imgsrc.baidu.com/forum/wh%3D200%2C90%3B/sign=fcba0a6c05f431adbc874b3b7b068096/faf8a72eb9389b5001cf0fac8c35e5dde6116ec2.jpg"  bpic="http://imgsrc.baidu.com/forum/w%3D580%3B/sign=e77ae3dcba1bb0518f24b3200641dbb4/908fa0ec08fa513d689a8320346d55fbb2fbd918.jpg" class="threadlist_pic j_m_pic "  /></a><div class="threadlist_pic_highlight j_m_pic_light"></div></li></ul>
                        <div class="small_pic_num center_text">共&nbsp;5&nbsp;张</div>
                    </div>
                </div>
            </div>                    </div>


<div class="threadlist_author pull_right">
        <span class="tb_icon_author_rely j_replyer" title="最后回复人: 无梦旅人丶丶">
            <i class="icon_replyer"></i>
            <a data-field='{&quot;un&quot;:&quot;\u65e0\u68a6\u65c5\u4eba\u4e36\u4e36&quot;}' class="frs-author-name j_user_card " href="/home/main/?un=%E6%97%A0%E6%A2%A6%E6%97%85%E4%BA%BA%E4%B8%B6%E4%B8%B6&ie=utf-8&fr=frs" target="_blank">无梦旅人丶丶</a>        </span>
        <span class="threadlist_reply_date pull_right j_reply_data" title="最后回复时间">
            14:40        </span>
</div>
                </div>
                    </div>
    </div>
</li>

*/
// http://tieba.baidu.com/f?kw= Cookie: BDUSS=
func parseTopicListResp(resp []byte) (topic map[string]string) {
	if resp != nil {
		topic = make(map[string]string)
		// <a href="/p/5026842582" title="互听 互评 互花花 ٩(˃̶͈̀௰˂̶͈́)و" target="_blank" class="j_th_tit ">互听 互评 互花花 ٩(˃̶͈̀௰˂̶͈́)و</a>
		reg := regexp.MustCompile(`(?m:\s+)[<]a\shref="/p/(\d+)"\stitle=(?:.*\s)target=(?:.*\s)class=(?:.*")[>](.+)[<][[:graph:]][a][>]\r*\n`)
		g := reg.FindAllStringSubmatch(string(resp), -1)
		if *debug {
			for i := range g {
				log.Infofln("match message: %#v\n", g[i])
			}
		}
		for i := range g {
			topic[g[i][1]] = g[i][2]
		}
	}
	return
}

func (a *agent) getTopicList() []byte {
	return a.get(a.fDetailURL)
}
