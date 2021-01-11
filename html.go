package main

const ActivityDetailHtml = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title></title>
    <meta name="viewport" content="width=500, user-scalable=0, target-densitydpi=device-dpi">
    <meta name="format-detection" content="telephone=no, email=no">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="author" content="">
</head>
<style>
    body {
        width: 500px;
        height: 100%;
        margin: 0 auto;
        overflow: hidden;
        font-family: "PingFang SC" !important;
    }

    .main_box {
        width: 100%;
        height: 400px;
        overflow: hidden;
        margin: 0 auto;
    }

    .act_item_list_box {
        padding: 26px;
        width: 448px;
    }

    .item_list {
        overflow: hidden;
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }

   .item_list .img_pic {
    width: 138px;
    height: 183px;
    overflow: hidden;
    border-radius: 3.5px;
    margin-right: 15px;
  }
 .item_list img {
    width: 138px;
    height: 183px;
    display: block;
    object-fit: cover;
    border-radius: 3.5px;
  }
   .item_list .goods_info{
    width: 304px;
    height: 180px;
    margin-top:3px;
    overflow: hidden;
    position: relative;
  }
    .item_list .goods_info .type {
        overflow: hidden;
        margin-top: 4px;
    }

    .item_list .goods_info h3 {
    margin-bottom:8px;
    width: 100%;
    font-size: 26px;
    font-weight: 500;
    color: #000000;
    max-height: 56px;
    margin-top: -5px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
  }

    .item_list .run {
        display: inline-block;
        color: #ffffff;
        font-size: 20px;
        border-radius: 20px;
        padding: 4px 14px;
        margin-right: 4px;
        overflow: hidden;
        max-height: 35px;
    }

    .item_list .run1 {
        background-image: linear-gradient(to right, #fd683d, #f99d0d);
    }

    .item_list .run2 {
        max-width: 80px;
        background-image: linear-gradient(to right, #00be72, #00ba9b);
    }

    .item_list .run3 {
        color: #00bc71;
        border: 1px solid #00bc71;
        padding: 3px 8px;
    }

    .item_list .list_info_box {
        width: 100%;
        position: absolute;
        bottom: 0;
    }

    .item_list .time_box, .item_list .addr, .item_list .team_box, .item_list .team_box_create {
        font-size: 20px;
        color: #555555;
        margin-bottom: 2px;
        width: 100%;
        display: flex;
        justify-content: flex-start;
        align-items: center;
    }

    .time_box span, .addr span, .team_box span, .team_box_create span {
        flex-shrink: 0;
        margin-right: 11px;
        display: inline-block;
        width: 23px;
        height: 23px;
        background: url('https://webapp.codoon.com/mini_sports_app/upload/ic-time-position.png') no-repeat left center;
        background-size: cover;
        vertical-align: middle;
    }

    .time_box p, .addr p, .team_box p, .team_box_create p {
        flex: 1;
        flex-shrink: 0;
        display: inline-block;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
        margin: 0;
    }

    .item_list .team_box span {
        background: url('https://webapp.codoon.com/mini_sports_app/upload/ic-share-time-position.png') no-repeat center center;
        background-size: cover;
    }

    .item_list .team_box_create span {
        background: url('https://webapp.codoon.com/mini_sports_app/upload/ic-share-team-people.png') no-repeat center center;
        background-size: cover;
    }

    .item_list .addr span {
        width: 30px;
        height: 30px;
        background: url('https://webapp.codoon.com/mini_sports_app/upload/ic_location.png') no-repeat center center;
        background-size: cover;
    }

    .item_list .time_box {
        margin-left: 2px;
    }

    .team_member {
        margin-top: 30px;
        height: 45px;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .team_member .member_list {
        float: left;
        width: 60%;
        display: flex;
        flex-direction: row-reverse;
        justify-content: flex-end;
        align-items: center;
        padding-left: 16px;
    }

    .team_member .member_list .item {
        width: 45px;
        height: 45px;
        background-color: #e6e6e6;
        background-color: var(--white);
        border: solid 3px #ffffff;
        border-radius: 100%;
        overflow: hidden;
        margin-left: -20px;
    }

    .team_member .member_list img {
        border-radius: 100%;
        display: inline-block;
        width: 45px;
        height: 45px;
    }

    .team_member .num_join {
        font-size: 26px;
        color: #00bc71;
    }

    .btn {
        margin-top: 30px;
        text-align: center;
        color: #ffffff;
        font-size: 28px;
        font-weight: 500;
        line-height: 66px;
        width: 100%;
        height: 66px;
        border-radius: 16px;
        background-image: linear-gradient(to right, #00be72, #00ba9b);
    }
</style>
<body>
<div class="main_box">
    <div class="act_item_list_box">
        <!-- 活动详情 -->
        <div class="item_list">
			<div class="img_pic"><img src="{{.Placard}}"/></div>
            <div class="goods_info">
                <h3>{{.Name}}</h3>
                <div class="type">
					{{if .IsVip}} <span class="run run1">{{.VipCertField}}</span> {{end}}
                    <span class="run run2">{{.SpType}}</span>
                    <span class="run run3">{{.PayType}}</span>
                </div>
                <div class="list_info_box">
                    <div class="time_box"><span></span>
                        <p>{{.DateArea}}</p></div>
                    <div class="addr"><span></span>
                        <p>{{.Address}}</p></div>
                </div>
            </div>
        </div>

        <!-- 团成员 -->
        <div class="team_member">
            <div class="member_list">
				{{range .ImgList}}
					<div class="item">
                    	<img src="{{.}}!256m256"/>
                	</div>
				{{end}}
            </div>
            <div class="num_join">{{.JoinNum}}</div>
        </div>

        <div class="btn">点击报名</div>
    </div>
</div>
</body>
</html>

`
