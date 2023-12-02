package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"

	"github.com/lifei6671/chrome-douyin/chrome"
	"github.com/lifei6671/chrome-douyin/config"
	"github.com/lifei6671/chrome-douyin/web"
)

var conf = flag.String("conf", "", "yaml config file")

func main() {

	flag.Parse()
	// docker run -p 9222:9222 ghcr.io/go-rod/rod chrome --headless --no-sandbox --remote-debugging-port=9222 --remote-debugging-address=0.0.0.0
	//u := "ws://127.0.0.1:35987/devtools/browser/b5427c79-ead4-4718-8230-f7566d136044"

	//cookie = ttwid=1%7CoxsjZKbrpUa71U-8bFp2m3ilMskaVih-aVnvYEPGnRM%7C1692926812%7Cf83726279996ea09536e36c28d655a3b6bab5036a98065cb3e40de01ccc5e4cb; passport_csrf_token=c0d539b96dc09e3694edc4769396ecb1; passport_csrf_token_default=c0d539b96dc09e3694edc4769396ecb1; webcast_local_quality=null; strategyABtestKey=%221696907137.696%22; stream_recommend_feed_params=%22%7B%5C%22cookie_enabled%5C%22%3Atrue%2C%5C%22screen_width%5C%22%3A1440%2C%5C%22screen_height%5C%22%3A900%2C%5C%22browser_online%5C%22%3Atrue%2C%5C%22cpu_core_num%5C%22%3A8%2C%5C%22device_memory%5C%22%3A8%2C%5C%22downlink%5C%22%3A10%2C%5C%22effective_type%5C%22%3A%5C%224g%5C%22%2C%5C%22round_trip_time%5C%22%3A200%7D%22; volume_info=%7B%22isUserMute%22%3Afalse%2C%22isMute%22%3Afalse%2C%22volume%22%3A0.5%7D; download_guide=%222%2F20231010%2F0%22; pwa2=%220%7C0%7C2%7C1%22; FORCE_LOGIN=%7B%22videoConsumedRemainSeconds%22%3A180%2C%22isForcePopClose%22%3A1%7D; passport_assist_user=CjwQrxBFHpi423GDwg00wdkrJ_oEGJNdpgDQqg7aALWwfTyORGdyh8rk6Y3r5O_HvEUBAqYeVC9v4zJvO64aSgo8mxmrk0sd4YVjGClRhiQ4wJzM3nUuH4se-qt3kYihraK0Tez9AOvikzOSSqNPf4I8T0ADimzVPyesN7YYELiXvg0Yia_WVCABIgED0BMqRg%3D%3D; n_mh=jr10LeZpK79oyuXVKiScXlEnxw32CuXluNMUZBOzh3Y; sso_uid_tt=9fd0c3c7b33478d95107037f94f4ac58; sso_uid_tt_ss=9fd0c3c7b33478d95107037f94f4ac58; toutiao_sso_user=aef261cccd0202e306907e7b22bea48f; toutiao_sso_user_ss=aef261cccd0202e306907e7b22bea48f; sid_ucp_sso_v1=1.0.0-KDRlZjU2NmJlY2ZhOWQ0YjlmMWYzYjEyYmE4MTAxYjZjY2E2MjNkODEKHQi3nOrD5gIQwu2TqQYY7zEgDDDAtPPVBTgGQPQHGgJobCIgYWVmMjYxY2NjZDAyMDJlMzA2OTA3ZTdiMjJiZWE0OGY; ssid_ucp_sso_v1=1.0.0-KDRlZjU2NmJlY2ZhOWQ0YjlmMWYzYjEyYmE4MTAxYjZjY2E2MjNkODEKHQi3nOrD5gIQwu2TqQYY7zEgDDDAtPPVBTgGQPQHGgJobCIgYWVmMjYxY2NjZDAyMDJlMzA2OTA3ZTdiMjJiZWE0OGY; odin_tt=a0e92a8554b76c260cd49348d5691d65869ac09f4167432e42bc4736ace8afd2c1bec93a621109e5275d93f698eb316f; passport_auth_status=5a1416325c69710f7ae273ae6ffade95%2C; passport_auth_status_ss=5a1416325c69710f7ae273ae6ffade95%2C; uid_tt=726e0bf22787aeb708879885f0472379; uid_tt_ss=726e0bf22787aeb708879885f0472379; sid_tt=cb593f753b4f15aefc82140a11e4eb92; sessionid=cb593f753b4f15aefc82140a11e4eb92; sessionid_ss=cb593f753b4f15aefc82140a11e4eb92; LOGIN_STATUS=1; _bd_ticket_crypt_doamin=3; _bd_ticket_crypt_cookie=b373a8694fe2cf8c04f6cd4fcd58c227; __security_server_data_status=1; store-region=cn-bj; store-region-src=uid; sid_guard=cb593f753b4f15aefc82140a11e4eb92%7C1696921298%7C5183987%7CSat%2C+09-Dec-2023+07%3A01%3A25+GMT; sid_ucp_v1=1.0.0-KDFlNDc3ZmFlYWE1YTE0NDk0NmJhZDZiMjQ3ZTA5ZTZlOWYyNzU3NTUKGQi3nOrD5gIQ0u2TqQYY7zEgDDgGQPQHSAQaAmxxIiBjYjU5M2Y3NTNiNGYxNWFlZmM4MjE0MGExMWU0ZWI5Mg; ssid_ucp_v1=1.0.0-KDFlNDc3ZmFlYWE1YTE0NDk0NmJhZDZiMjQ3ZTA5ZTZlOWYyNzU3NTUKGQi3nOrD5gIQ0u2TqQYY7zEgDDgGQPQHSAQaAmxxIiBjYjU5M2Y3NTNiNGYxNWFlZmM4MjE0MGExMWU0ZWI5Mg; msToken=aOZnDr1A6OjTngx7zGVijih3PJob2umMYXTSORdNHL4C-11z6iiEnPeFs9fqQAGVSL3ObWk_4vvktJ890VlYy-NquvAMn0LCCiQJ9XIXIlAG6kpbgJOzFclAsIM=; publish_badge_show_info=%220%2C0%2C0%2C1696921299020%22; SEARCH_RESULT_LIST_TYPE=%22single%22; IsDouyinActive=true; home_can_add_dy_2_desktop=%220%22; FOLLOW_LIVE_POINT_INFO=%22MS4wLjABAAAAZu94yXL7DOxO3XdW0-EHCcLEjxYOHe952P03e6fsL4k%2F1696953600000%2F0%2F0%2F1696923183549%22; FOLLOW_NUMBER_YELLOW_POINT_INFO=%22MS4wLjABAAAAZu94yXL7DOxO3XdW0-EHCcLEjxYOHe952P03e6fsL4k%2F1696953600000%2F0%2F1696922583550%2F0%22"
	//s_v_web_id=verify_llpwwpci_J6h1ro48_7uXd_4T31_9lWk_gOdWj8qkrG5z; ttwid=1%7CoxsjZKbrpUa71U-8bFp2m3ilMskaVih-aVnvYEPGnRM%7C1692926812%7Cf83726279996ea09536e36c28d655a3b6bab5036a98065cb3e40de01ccc5e4cb; passport_csrf_token=c0d539b96dc09e3694edc4769396ecb1; passport_csrf_token_default=c0d539b96dc09e3694edc4769396ecb1; webcast_local_quality=null; ttcid=726c9c7fda1a44ca9212328bf38b3c8815; volume_info=%7B%22isUserMute%22%3Afalse%2C%22isMute%22%3Afalse%2C%22volume%22%3A0.5%7D; download_guide=%222%2F20231010%2F0%22; pwa2=%220%7C0%7C2%7C1%22; FORCE_LOGIN=%7B%22videoConsumedRemainSeconds%22%3A180%2C%22isForcePopClose%22%3A1%7D; passport_assist_user=CjwQrxBFHpi423GDwg00wdkrJ_oEGJNdpgDQqg7aALWwfTyORGdyh8rk6Y3r5O_HvEUBAqYeVC9v4zJvO64aSgo8mxmrk0sd4YVjGClRhiQ4wJzM3nUuH4se-qt3kYihraK0Tez9AOvikzOSSqNPf4I8T0ADimzVPyesN7YYELiXvg0Yia_WVCABIgED0BMqRg%3D%3D; n_mh=jr10LeZpK79oyuXVKiScXlEnxw32CuXluNMUZBOzh3Y; sso_uid_tt=9fd0c3c7b33478d95107037f94f4ac58; sso_uid_tt_ss=9fd0c3c7b33478d95107037f94f4ac58; toutiao_sso_user=aef261cccd0202e306907e7b22bea48f; toutiao_sso_user_ss=aef261cccd0202e306907e7b22bea48f; sid_ucp_sso_v1=1.0.0-KDRlZjU2NmJlY2ZhOWQ0YjlmMWYzYjEyYmE4MTAxYjZjY2E2MjNkODEKHQi3nOrD5gIQwu2TqQYY7zEgDDDAtPPVBTgGQPQHGgJobCIgYWVmMjYxY2NjZDAyMDJlMzA2OTA3ZTdiMjJiZWE0OGY; ssid_ucp_sso_v1=1.0.0-KDRlZjU2NmJlY2ZhOWQ0YjlmMWYzYjEyYmE4MTAxYjZjY2E2MjNkODEKHQi3nOrD5gIQwu2TqQYY7zEgDDDAtPPVBTgGQPQHGgJobCIgYWVmMjYxY2NjZDAyMDJlMzA2OTA3ZTdiMjJiZWE0OGY; passport_auth_status=5a1416325c69710f7ae273ae6ffade95%2C; passport_auth_status_ss=5a1416325c69710f7ae273ae6ffade95%2C; uid_tt=726e0bf22787aeb708879885f0472379; uid_tt_ss=726e0bf22787aeb708879885f0472379; sid_tt=cb593f753b4f15aefc82140a11e4eb92; sessionid=cb593f753b4f15aefc82140a11e4eb92; sessionid_ss=cb593f753b4f15aefc82140a11e4eb92; LOGIN_STATUS=1; _bd_ticket_crypt_doamin=3; _bd_ticket_crypt_cookie=b373a8694fe2cf8c04f6cd4fcd58c227; __security_server_data_status=1; store-region=cn-bj; store-region-src=uid; sid_guard=cb593f753b4f15aefc82140a11e4eb92%7C1696921298%7C5183987%7CSat%2C+09-Dec-2023+07%3A01%3A25+GMT; sid_ucp_v1=1.0.0-KDFlNDc3ZmFlYWE1YTE0NDk0NmJhZDZiMjQ3ZTA5ZTZlOWYyNzU3NTUKGQi3nOrD5gIQ0u2TqQYY7zEgDDgGQPQHSAQaAmxxIiBjYjU5M2Y3NTNiNGYxNWFlZmM4MjE0MGExMWU0ZWI5Mg; ssid_ucp_v1=1.0.0-KDFlNDc3ZmFlYWE1YTE0NDk0NmJhZDZiMjQ3ZTA5ZTZlOWYyNzU3NTUKGQi3nOrD5gIQ0u2TqQYY7zEgDDgGQPQHSAQaAmxxIiBjYjU5M2Y3NTNiNGYxNWFlZmM4MjE0MGExMWU0ZWI5Mg; csrf_session_id=b90d805ccaf2598ddbbea740b21fe717; SEARCH_RESULT_LIST_TYPE=%22single%22; FOLLOW_NUMBER_YELLOW_POINT_INFO=%22MS4wLjABAAAAZu94yXL7DOxO3XdW0-EHCcLEjxYOHe952P03e6fsL4k%2F1696953600000%2F0%2F1696922583550%2F0%22; strategyABtestKey=%221697011180.239%22; stream_recommend_feed_params=%22%7B%5C%22cookie_enabled%5C%22%3Atrue%2C%5C%22screen_width%5C%22%3A1440%2C%5C%22screen_height%5C%22%3A900%2C%5C%22browser_online%5C%22%3Atrue%2C%5C%22cpu_core_num%5C%22%3A8%2C%5C%22device_memory%5C%22%3A8%2C%5C%22downlink%5C%22%3A1.55%2C%5C%22effective_type%5C%22%3A%5C%223g%5C%22%2C%5C%22round_trip_time%5C%22%3A350%7D%22; FOLLOW_LIVE_POINT_INFO=%22MS4wLjABAAAAZu94yXL7DOxO3XdW0-EHCcLEjxYOHe952P03e6fsL4k%2F1697040000000%2F0%2F1697011180587%2F0%22; bd_ticket_guard_client_data=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWl0ZXJhdGlvbi12ZXJzaW9uIjoxLCJiZC10aWNrZXQtZ3VhcmQtcmVlLXB1YmxpYy1rZXkiOiJCTTZqVzFxV1l4OHlLbUVIMnZVWEJVVmZOK1RwT2NSOTFiMXZobk8xMENJTHlPc0FOK1lobTltcVRaWXRwMGl6VERyMVJmcUd0TlV3T0FlRVpkUFEvaFE9IiwiYmQtdGlja2V0LWd1YXJkLXdlYi12ZXJzaW9uIjoxfQ%3D%3D; odin_tt=3edd0291fef4f94960c21d9367c4b367ae266951a1f78dd8cf4cf7fc994b68e88edcad13b53f31c2614bb69c9a08e407; msToken=0bK746q-MoMHPfclMwa9xQn-CuNUGVAMsFAUq1Y7GBuz-FbQDIy9Tp905PXVdwxiUSsHyPMbDVw5zzW1R40TAM5L7Erx1MZdwNomiDj6gWfwJO3wKUBS8-inXJA=; msToken=71H4LpjCTDpUHP5fDQwQbGXlvAdA9e4shqHc36ojNvA0x9OAZ1x_7Ybjl2eCBaQOpAbDHASYrIvkNYR3lC1tMAwzDOr-NByi8nWag4BMeH-qc4iy13SZfsmkEdw=; home_can_add_dy_2_desktop=%221%22; publish_badge_show_info=%220%2C0%2C0%2C1697011186274%22; tt_scid=QIrp.Ch4NJDv1wv9GuXRxf8RlZt5obr1WB7rStEDMD1DoWi8X1xoPQAsTcCW5rCd1e4c; IsDouyinActive=false; passport_fe_beating_status=false; __ac_nonce=065267e7300082fcd5487; __ac_signature=_02B4Z6wo00f01lOKq0AAAIDDFLNAckJ2M1JTqq.AAPHTPyfuLrRUSyoV5N-xgFncYyNGVvx45Y9NCNEKDeBgCAxI-Xf0ooIz8cz2hfijtnynW8aOgJNELXrM3K4IXdUV0x9lmFDfdf6JfdnC4e; __ac_referer=__ac_blank
	var f string
	if conf != nil && len(*conf) > 0 {
		f = *conf
	} else {
		f = config.GetAssetLocation("config.yaml")
	}
	log.Println("config file:", f)
	b, err := os.ReadFile(f)
	if err != nil {
		log.Fatalln("read config file fail:", err)
	}
	if err := yaml.Unmarshal(b, &config.Default); err != nil {
		log.Fatalln("unmarshal config file fail:", err)
	}

	gin.SetMode(config.Default.Mode)

	err = web.Run(config.Default.Addr, func() *chrome.Chrome {
		log.Printf(" [GIN] config=[%s]", config.Default)
		log.Printf(" [GIN] service-url=[%s]", config.Default.ServiceURL)
		cookies := []chrome.Cookie{
			{
				Name:   "douyin",
				Cookie: config.Default.DouYin.Cookie,
				Domain: config.Default.DouYin.Domain,
			},
			{
				Name:   "xiaohongshu",
				Cookie: config.Default.XiaoHongShu.Cookie,
				Domain: config.Default.XiaoHongShu.Domain,
			},
		}

		return chrome.New(chrome.WithCookies(cookies...), chrome.WithServiceURL(config.Default.ServiceURL))
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
