package chrome

import "github.com/lifei6671/chrome-douyin/utils"

type DouYinDetail struct {
	AwemeDetail AwemeDetail `json:"aweme_detail"`
	StatusCode  int         `json:"status_code"`
	ShareURL    string      `json:"share_url"`
}

type DouYinImages struct {
	AwemeList  []AwemeDetail `json:"aweme_list"`
	StatusCode int           `json:"status_code"`
}

type AwemeDetail struct {
	ActivityVideoType   int    `json:"activity_video_type"`
	Anchors             any    `json:"anchors"`
	AuthenticationToken string `json:"authentication_token"`
	Author              struct {
		AvatarSchemaList any `json:"avatar_schema_list"`
		AvatarThumb      struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"avatar_thumb"`
		AwemehtsGreetInfo string `json:"awemehts_greet_info"`
		CfList            any    `json:"cf_list"`
		CloseFriendType   int    `json:"close_friend_type"`
		ContactsStatus    int    `json:"contacts_status"`
		ContrailList      any    `json:"contrail_list"`
		CoverUrl          []struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"cover_url"`
		CreateTime                             int    `json:"create_time"`
		CustomVerify                           string `json:"custom_verify"`
		DataLabelList                          any    `json:"data_label_list"`
		EndorsementInfoList                    any    `json:"endorsement_info_list"`
		EnterpriseVerifyReason                 string `json:"enterprise_verify_reason"`
		FamiliarVisitorUser                    any    `json:"familiar_visitor_user"`
		FavoritingCount                        int    `json:"favoriting_count"`
		FollowStatus                           int    `json:"follow_status"`
		FollowerCount                          int    `json:"follower_count"`
		FollowerListSecondaryInformationStruct any    `json:"follower_list_secondary_information_struct"`
		FollowerStatus                         int    `json:"follower_status"`
		FollowingCount                         int    `json:"following_count"`
		ImRoleIds                              any    `json:"im_role_ids"`
		IsAdFake                               bool   `json:"is_ad_fake"`
		IsBan                                  bool   `json:"is_ban"`
		IsBlockedV2                            bool   `json:"is_blocked_v2"`
		IsBlockingV2                           bool   `json:"is_blocking_v2"`
		IsCf                                   int    `json:"is_cf"`
		LiveHighValue                          int    `json:"live_high_value"`
		MaxFollowerCount                       int    `json:"max_follower_count"`
		Nickname                               string `json:"nickname"`
		NotSeenItemIdList                      any    `json:"not_seen_item_id_list"`
		NotSeenItemIdListV2                    any    `json:"not_seen_item_id_list_v2"`
		OfflineInfoList                        any    `json:"offline_info_list"`
		PersonalTagList                        any    `json:"personal_tag_list"`
		PreventDownload                        bool   `json:"prevent_download"`
		RiskNoticeText                         string `json:"risk_notice_text"`
		SecUid                                 string `json:"sec_uid"`
		Secret                                 int    `json:"secret"`
		ShareInfo                              struct {
			ShareDesc      string `json:"share_desc"`
			ShareDescInfo  string `json:"share_desc_info"`
			ShareQrcodeUrl struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"share_qrcode_url"`
			ShareTitle       string `json:"share_title"`
			ShareTitleMyself string `json:"share_title_myself"`
			ShareTitleOther  string `json:"share_title_other"`
			ShareUrl         string `json:"share_url"`
			ShareWeiboDesc   string `json:"share_weibo_desc"`
		} `json:"share_info"`
		ShortId             string `json:"short_id"`
		Signature           string `json:"signature"`
		SignatureExtra      any    `json:"signature_extra"`
		SpecialFollowStatus int    `json:"special_follow_status"`
		SpecialPeopleLabels any    `json:"special_people_labels"`
		Status              int    `json:"status"`
		TextExtra           any    `json:"text_extra"`
		TotalFavorited      int    `json:"total_favorited"`
		Uid                 string `json:"uid"`
		UniqueId            string `json:"unique_id"`
		UserAge             int    `json:"user_age"`
		UserCanceled        bool   `json:"user_canceled"`
		UserPermissions     any    `json:"user_permissions"`
		VerificationType    int    `json:"verification_type"`
	} `json:"author"`
	AuthorMaskTag int   `json:"author_mask_tag"`
	AuthorUserId  int64 `json:"author_user_id"`
	AwemeAcl      struct {
		DownloadMaskPanel struct {
			Code     int `json:"code"`
			ShowType int `json:"show_type"`
		} `json:"download_mask_panel"`
	} `json:"aweme_acl"`
	AwemeControl struct {
		CanComment     bool `json:"can_comment"`
		CanForward     bool `json:"can_forward"`
		CanShare       bool `json:"can_share"`
		CanShowComment bool `json:"can_show_comment"`
	} `json:"aweme_control"`
	AwemeId   string `json:"aweme_id"`
	AwemeType int    `json:"aweme_type"`
	BookBar   struct {
	} `json:"book_bar"`
	BoostStatus           int   `json:"boost_status"`
	ChallengePosition     any   `json:"challenge_position"`
	ChapterList           any   `json:"chapter_list"`
	CollectStat           int   `json:"collect_stat"`
	CollectionCornerMark  int   `json:"collection_corner_mark"`
	CommentGid            int64 `json:"comment_gid"`
	CommentList           any   `json:"comment_list"`
	CommentPermissionInfo struct {
		CanComment              bool `json:"can_comment"`
		CommentPermissionStatus int  `json:"comment_permission_status"`
		ItemDetailEntry         bool `json:"item_detail_entry"`
		PressEntry              bool `json:"press_entry"`
		ToastGuide              bool `json:"toast_guide"`
	} `json:"comment_permission_info"`
	CommentWordsRecommend struct {
		ZeroComment any `json:"zero_comment"`
	} `json:"comment_words_recommend"`
	CommerceConfigData any    `json:"commerce_config_data"`
	CommonBarInfo      string `json:"common_bar_info"`
	ComponentInfoV2    string `json:"component_info_v2"`
	CoverLabels        any    `json:"cover_labels"`
	CreateScaleType    any    `json:"create_scale_type"`
	CreateTime         int    `json:"create_time"`
	Desc               string `json:"desc"`
	Descendants        struct {
		NotifyMsg string   `json:"notify_msg"`
		Platforms []string `json:"platforms"`
	} `json:"descendants"`
	DiggLottie struct {
		CanBomb  int    `json:"can_bomb"`
		LottieId string `json:"lottie_id"`
	} `json:"digg_lottie"`
	DisableRelationBar     int `json:"disable_relation_bar"`
	DislikeDimensionList   any `json:"dislike_dimension_list"`
	DislikeDimensionListV2 any `json:"dislike_dimension_list_v2"`
	DistributeCircle       struct {
		CampusBlockInteraction bool `json:"campus_block_interaction"`
		DistributeType         int  `json:"distribute_type"`
		IsCampus               bool `json:"is_campus"`
	} `json:"distribute_circle"`
	DuetAggregateInMusicTab bool   `json:"duet_aggregate_in_music_tab"`
	Duration                int    `json:"duration"`
	Geofencing              []any  `json:"geofencing"`
	GeofencingRegions       any    `json:"geofencing_regions"`
	GroupId                 string `json:"group_id"`
	GuideSceneInfo          struct {
		DiamondExposeInfoStr string `json:"diamond_expose_info_str"`
		FeedOriginGidInfoStr string `json:"feed_origin_gid_info_str"`
		GuideSceneType       int    `json:"guide_scene_type"`
	} `json:"guide_scene_info"`
	HybridLabel    any     `json:"hybrid_label"`
	Images         []Image `json:"images"`
	ImgBitrate     any     `json:"img_bitrate"`
	ImpressionData struct {
		GroupIdListA   []any `json:"group_id_list_a"`
		GroupIdListB   []any `json:"group_id_list_b"`
		GroupIdListC   []any `json:"group_id_list_c"`
		SimilarIdListA any   `json:"similar_id_list_a"`
		SimilarIdListB any   `json:"similar_id_list_b"`
	} `json:"impression_data"`
	InteractionStickers  any  `json:"interaction_stickers"`
	IsAds                bool `json:"is_ads"`
	IsCollectsSelected   int  `json:"is_collects_selected"`
	IsDuetSing           bool `json:"is_duet_sing"`
	IsImageBeat          bool `json:"is_image_beat"`
	IsLifeItem           bool `json:"is_life_item"`
	IsSharePost          bool `json:"is_share_post"`
	IsStory              int  `json:"is_story"`
	IsTop                int  `json:"is_top"`
	ItemWarnNotification struct {
		Content string `json:"content"`
		Show    bool   `json:"show"`
		Type    int    `json:"type"`
	} `json:"item_warn_notification"`
	JumpTabInfoList     any `json:"jump_tab_info_list"`
	LabelTopText        any `json:"label_top_text"`
	LiveAppointmentInfo struct {
	} `json:"live_appointment_info"`
	LongVideo any `json:"long_video"`
	MediaType int `json:"media_type"`
	Music     struct {
		Album            string `json:"album"`
		ArtistUserInfos  any    `json:"artist_user_infos"`
		Artists          []any  `json:"artists"`
		AuditionDuration int    `json:"audition_duration"`
		Author           string `json:"author"`
		AuthorDeleted    bool   `json:"author_deleted"`
		AuthorPosition   any    `json:"author_position"`
		AuthorStatus     int    `json:"author_status"`
		AvatarLarge      struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"avatar_large"`
		AvatarMedium struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"avatar_medium"`
		AvatarThumb struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"avatar_thumb"`
		BindedChallengeId int  `json:"binded_challenge_id"`
		CanBackgroundPlay bool `json:"can_background_play"`
		CollectStat       int  `json:"collect_stat"`
		CoverHd           struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"cover_hd"`
		CoverLarge struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"cover_large"`
		CoverMedium struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"cover_medium"`
		CoverThumb struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"cover_thumb"`
		DmvAutoShow          bool   `json:"dmv_auto_show"`
		DspStatus            int    `json:"dsp_status"`
		Duration             int    `json:"duration"`
		EndTime              int    `json:"end_time"`
		ExternalSongInfo     []any  `json:"external_song_info"`
		Extra                string `json:"extra"`
		Id                   int64  `json:"id"`
		IdStr                string `json:"id_str"`
		IsAudioUrlWithCookie bool   `json:"is_audio_url_with_cookie"`
		IsCommerceMusic      bool   `json:"is_commerce_music"`
		IsDelVideo           bool   `json:"is_del_video"`
		IsMatchedMetadata    bool   `json:"is_matched_metadata"`
		IsOriginal           bool   `json:"is_original"`
		IsOriginalSound      bool   `json:"is_original_sound"`
		IsPgc                bool   `json:"is_pgc"`
		IsRestricted         bool   `json:"is_restricted"`
		IsVideoSelfSee       bool   `json:"is_video_self_see"`
		LunaInfo             struct {
			IsLunaUser bool `json:"is_luna_user"`
		} `json:"luna_info"`
		LyricShortPosition             any    `json:"lyric_short_position"`
		Mid                            string `json:"mid"`
		MusicChartRanks                any    `json:"music_chart_ranks"`
		MusicCollectCount              int    `json:"music_collect_count"`
		MusicCoverAtmosphereColorValue string `json:"music_cover_atmosphere_color_value"`
		MusicStatus                    int    `json:"music_status"`
		MusicianUserInfos              any    `json:"musician_user_infos"`
		MuteShare                      bool   `json:"mute_share"`
		OfflineDesc                    string `json:"offline_desc"`
		OwnerHandle                    string `json:"owner_handle"`
		OwnerId                        string `json:"owner_id"`
		OwnerNickname                  string `json:"owner_nickname"`
		PgcMusicType                   int    `json:"pgc_music_type"`
		PlayUrl                        struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlKey  string   `json:"url_key"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"play_url"`
		Position                  any     `json:"position"`
		PreventDownload           bool    `json:"prevent_download"`
		PreventItemDownloadStatus int     `json:"prevent_item_download_status"`
		PreviewEndTime            float64 `json:"preview_end_time"`
		PreviewStartTime          float64 `json:"preview_start_time"`
		ReasonType                int     `json:"reason_type"`
		Redirect                  bool    `json:"redirect"`
		SchemaUrl                 string  `json:"schema_url"`
		SearchImpr                struct {
			EntityId string `json:"entity_id"`
		} `json:"search_impr"`
		SecUid            string `json:"sec_uid"`
		ShootDuration     int    `json:"shoot_duration"`
		SourcePlatform    int    `json:"source_platform"`
		StartTime         int    `json:"start_time"`
		Status            int    `json:"status"`
		TagList           any    `json:"tag_list"`
		Title             string `json:"title"`
		UnshelveCountries any    `json:"unshelve_countries"`
		UserCount         int    `json:"user_count"`
		VideoDuration     int    `json:"video_duration"`
	} `json:"music"`
	NicknamePosition    any   `json:"nickname_position"`
	OriginCommentIds    any   `json:"origin_comment_ids"`
	OriginTextExtra     []any `json:"origin_text_extra"`
	OriginalImages      any   `json:"original_images"`
	PackedClips         any   `json:"packed_clips"`
	PhotoSearchEntrance struct {
		EcomType int `json:"ecom_type"`
	} `json:"photo_search_entrance"`
	Position             any    `json:"position"`
	PreviewTitle         string `json:"preview_title"`
	PreviewVideoStatus   int    `json:"preview_video_status"`
	Promotions           []any  `json:"promotions"`
	Rate                 int    `json:"rate"`
	RefTtsIdList         any    `json:"ref_tts_id_list"`
	RefVoiceModifyIdList any    `json:"ref_voice_modify_id_list"`
	Region               string `json:"region"`
	RelationLabels       any    `json:"relation_labels"`
	ReplySmartEmojis     any    `json:"reply_smart_emojis"`
	SearchImpr           struct {
		EntityId   string `json:"entity_id"`
		EntityType string `json:"entity_type"`
	} `json:"search_impr"`
	SeoInfo struct {
		OcrContent string `json:"ocr_content"`
	} `json:"seo_info"`
	SeriesPaidInfo struct {
		ItemPrice        int `json:"item_price"`
		SeriesPaidStatus int `json:"series_paid_status"`
	} `json:"series_paid_info"`
	ShareInfo struct {
		ShareDesc     string `json:"share_desc"`
		ShareDescInfo string `json:"share_desc_info"`
		ShareLinkDesc string `json:"share_link_desc"`
		ShareUrl      string `json:"share_url"`
	} `json:"share_info"`
	ShareUrl           string `json:"share_url"`
	ShouldOpenAdReport bool   `json:"should_open_ad_report"`
	ShowFollowButton   struct {
	} `json:"show_follow_button"`
	SlidesMusicBeats    any `json:"slides_music_beats"`
	SocialTagList       any `json:"social_tag_list"`
	StandardBarInfoList any `json:"standard_bar_info_list"`
	Statistics          struct {
		AdmireCount  int    `json:"admire_count"`
		AwemeId      string `json:"aweme_id"`
		CollectCount int    `json:"collect_count"`
		CommentCount int    `json:"comment_count"`
		DiggCount    int    `json:"digg_count"`
		PlayCount    int    `json:"play_count"`
		ShareCount   int    `json:"share_count"`
	} `json:"statistics"`
	Status struct {
		AllowShare        bool   `json:"allow_share"`
		AwemeId           string `json:"aweme_id"`
		InReviewing       bool   `json:"in_reviewing"`
		IsDelete          bool   `json:"is_delete"`
		IsProhibited      bool   `json:"is_prohibited"`
		ListenVideoStatus int    `json:"listen_video_status"`
		PartSee           int    `json:"part_see"`
		PrivateStatus     int    `json:"private_status"`
		ReviewResult      struct {
			ReviewStatus int `json:"review_status"`
		} `json:"review_result"`
	} `json:"status"`
	SuggestWords struct {
		SuggestWords []struct {
			ExtraInfo string `json:"extra_info"`
			HintText  string `json:"hint_text"`
			IconUrl   string `json:"icon_url"`
			Scene     string `json:"scene"`
			Words     []struct {
				Info   string `json:"info"`
				Word   string `json:"word"`
				WordId string `json:"word_id"`
			} `json:"words"`
		} `json:"suggest_words"`
	} `json:"suggest_words"`
	TextExtra           []any `json:"text_extra"`
	TtsIdList           any   `json:"tts_id_list"`
	UniqidPosition      any   `json:"uniqid_position"`
	UserDigged          int   `json:"user_digged"`
	UserRecommendStatus int   `json:"user_recommend_status"`
	Video               struct {
		BigThumbs any `json:"big_thumbs"`
		BitRate   []struct {
			FPS       int    `json:"FPS"`
			HDRBit    string `json:"HDR_bit"`
			HDRType   string `json:"HDR_type"`
			BitRate   int    `json:"bit_rate"`
			GearName  string `json:"gear_name"`
			IsBytevc1 int    `json:"is_bytevc1"`
			IsH265    int    `json:"is_h265"`
			PlayAddr  struct {
				DataSize int      `json:"data_size"`
				FileCs   string   `json:"file_cs"`
				FileHash string   `json:"file_hash"`
				Height   int      `json:"height"`
				Uri      string   `json:"uri"`
				UrlKey   string   `json:"url_key"`
				UrlList  []string `json:"url_list"`
				Width    int      `json:"width"`
			} `json:"play_addr"`
			QualityType int    `json:"quality_type"`
			VideoExtra  string `json:"video_extra"`
		} `json:"bit_rate"`
		BitRateAudio  any `json:"bit_rate_audio"`
		CdnUrlExpired int `json:"cdn_url_expired"`
		Cover         struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"cover"`
		CoverOriginalScale struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"cover_original_scale"`
		DownloadAddr struct {
			DataSize int      `json:"data_size"`
			FileCs   string   `json:"file_cs"`
			Height   int      `json:"height"`
			Uri      string   `json:"uri"`
			UrlList  []string `json:"url_list"`
			Width    int      `json:"width"`
		} `json:"download_addr"`
		DownloadSuffixLogoAddr struct {
			DataSize int      `json:"data_size"`
			FileCs   string   `json:"file_cs"`
			Height   int      `json:"height"`
			Uri      string   `json:"uri"`
			UrlList  []string `json:"url_list"`
			Width    int      `json:"width"`
		} `json:"download_suffix_logo_addr"`
		Duration     int `json:"duration"`
		DynamicCover struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"dynamic_cover"`
		GaussianCover struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"gaussian_cover"`
		HasDownloadSuffixLogoAddr bool   `json:"has_download_suffix_logo_addr"`
		HasWatermark              bool   `json:"has_watermark"`
		Height                    int    `json:"height"`
		IsH265                    int    `json:"is_h265"`
		IsSourceHDR               int    `json:"is_source_HDR"`
		Meta                      string `json:"meta"`
		MiscDownloadAddrs         string `json:"misc_download_addrs"`
		OriginCover               struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"origin_cover"`
		PlayAddr struct {
			DataSize int      `json:"data_size"`
			FileCs   string   `json:"file_cs"`
			FileHash string   `json:"file_hash"`
			Height   int      `json:"height"`
			Uri      string   `json:"uri"`
			UrlKey   string   `json:"url_key"`
			UrlList  []string `json:"url_list"`
			Width    int      `json:"width"`
		} `json:"play_addr"`
		PlayAddr265 struct {
			DataSize int      `json:"data_size"`
			FileCs   string   `json:"file_cs"`
			FileHash string   `json:"file_hash"`
			Height   int      `json:"height"`
			Uri      string   `json:"uri"`
			UrlKey   string   `json:"url_key"`
			UrlList  []string `json:"url_list"`
			Width    int      `json:"width"`
		} `json:"play_addr_265"`
		PlayAddrH264 struct {
			DataSize int      `json:"data_size"`
			FileCs   string   `json:"file_cs"`
			FileHash string   `json:"file_hash"`
			Height   int      `json:"height"`
			Uri      string   `json:"uri"`
			UrlKey   string   `json:"url_key"`
			UrlList  []string `json:"url_list"`
			Width    int      `json:"width"`
		} `json:"play_addr_h264"`
		Ratio      string `json:"ratio"`
		VideoModel string `json:"video_model"`
		Width      int    `json:"width"`
	} `json:"video"`
	VideoLabels any `json:"video_labels"`
	VideoTag    []struct {
		Level   int    `json:"level"`
		TagId   int    `json:"tag_id"`
		TagName string `json:"tag_name"`
	} `json:"video_tag"`
	VideoText         []any `json:"video_text"`
	VoiceModifyIdList any   `json:"voice_modify_id_list"`
	YummeRecreason    any   `json:"yumme_recreason"`
}

type DouYinResult struct {
	Url        string `json:"url"`
	AwemeId    string `json:"aweme_id"`
	Desc       string `json:"desc"`
	CreateTime int    `json:"create_time"`
	Type       string `json:"type"`
	Author     struct {
		AvatarThumb struct {
			Height  float64  `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   float64  `json:"width"`
		} `json:"avatar_thumb"`
		CfList          interface{} `json:"cf_list"`
		CloseFriendType int         `json:"close_friend_type"`
		ContactsStatus  int         `json:"contacts_status"`
		ContrailList    interface{} `json:"contrail_list"`
		CoverUrl        []struct {
			Height  float64  `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   float64  `json:"width"`
		} `json:"cover_url"`
		CustomVerify           string      `json:"custom_verify"`
		DataLabelList          interface{} `json:"data_label_list"`
		EndorsementInfoList    interface{} `json:"endorsement_info_list"`
		EnterpriseVerifyReason string      `json:"enterprise_verify_reason"`
		FamiliarVisitorUser    interface{} `json:"familiar_visitor_user"`
		ImRoleIds              interface{} `json:"im_role_ids"`
		IsAdFake               bool        `json:"is_ad_fake"`
		IsBan                  bool        `json:"is_ban"`
		IsBlockedV2            bool        `json:"is_blocked_v2"`
		IsBlockingV2           bool        `json:"is_blocking_v2"`
		Nickname               string      `json:"nickname"`
		NotSeenItemIdList      interface{} `json:"not_seen_item_id_list"`
		NotSeenItemIdListV2    interface{} `json:"not_seen_item_id_list_v2"`
		OfflineInfoList        interface{} `json:"offline_info_list"`
		PersonalTagList        interface{} `json:"personal_tag_list"`
		PreventDownload        bool        `json:"prevent_download"`
		RiskNoticeText         string      `json:"risk_notice_text"`
		SecUid                 string      `json:"sec_uid"`
		ShareInfo              struct {
			ShareDesc      string `json:"share_desc"`
			ShareDescInfo  string `json:"share_desc_info"`
			ShareQrcodeUrl struct {
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
			} `json:"share_qrcode_url"`
			ShareTitle       string `json:"share_title"`
			ShareTitleMyself string `json:"share_title_myself"`
			ShareTitleOther  string `json:"share_title_other"`
			ShareUrl         string `json:"share_url"`
			ShareWeiboDesc   string `json:"share_weibo_desc"`
		} `json:"share_info"`
		ShortId             string      `json:"short_id"`
		Signature           string      `json:"signature"`
		SignatureExtra      interface{} `json:"signature_extra"`
		SpecialFollowStatus int         `json:"special_follow_status"`
		SpecialPeopleLabels interface{} `json:"special_people_labels"`
		Status              int         `json:"status"`
		TextExtra           interface{} `json:"text_extra"`
		TotalFavorited      int         `json:"total_favorited"`
		Uid                 string      `json:"uid"`
		UniqueId            string      `json:"unique_id"`
		UserAge             int         `json:"user_age"`
		UserCanceled        bool        `json:"user_canceled"`
		UserPermissions     interface{} `json:"user_permissions"`
		VerificationType    int         `json:"verification_type"`
	} `json:"author"`
	CoverData struct {
		Cover struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"cover"`
		OriginCover struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"origin_cover"`
		DynamicCover struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"dynamic_cover"`
	} `json:"cover_data"`
	Music struct {
		PlayUrl struct {
			Height  int      `json:"height"`
			Uri     string   `json:"uri"`
			UrlKey  string   `json:"url_key"`
			UrlList []string `json:"url_list"`
			Width   int      `json:"width"`
		} `json:"play_url"`
	} `json:"music"`
	VideoData struct {
		WmVideoUrl    string `json:"wm_video_url"`
		WmVideoUrlHQ  string `json:"wm_video_url_HQ"`
		NwmVideoUrl   string `json:"nwm_video_url"`
		NwmVideoUrlHQ string `json:"nwm_video_url_HQ"`
	} `json:"video_data"`
	Images []Image `json:"images"`
}

type Image struct {
	Height  int      `json:"height"`
	Width   int      `json:"width"`
	URLList []string `json:"url_list"`
	URI     string   `json:"uri"`
}

func FromDouYinDetail(d *DouYinDetail) *DouYinResult {
	ret := &DouYinResult{}

	ret.Type = MediaType(d.AwemeDetail.MediaType).String()

	ret.Url = d.AwemeDetail.Video.PlayAddr.Uri
	ret.AwemeId = d.AwemeDetail.AwemeId
	ret.Desc = d.AwemeDetail.Desc
	ret.CreateTime = d.AwemeDetail.CreateTime

	ret.Author.Uid = d.AwemeDetail.Author.Uid
	ret.Author.ShortId = d.AwemeDetail.Author.ShortId
	ret.Author.Nickname = d.AwemeDetail.Author.Nickname
	ret.Author.Signature = d.AwemeDetail.Author.Signature

	for _, cover := range d.AwemeDetail.Author.CoverUrl {
		ret.Author.CoverUrl = append(ret.Author.CoverUrl, struct {
			Height  float64  `json:"height"`
			Uri     string   `json:"uri"`
			UrlList []string `json:"url_list"`
			Width   float64  `json:"width"`
		}{Height: float64(cover.Height), Uri: cover.Uri, UrlList: cover.UrlList, Width: float64(cover.Width)})
	}
	ret.Author.AvatarThumb.UrlList = d.AwemeDetail.Author.AvatarThumb.UrlList
	ret.Author.AvatarThumb.Uri = d.AwemeDetail.Author.AvatarThumb.Uri
	ret.Author.AvatarThumb.Width = float64(d.AwemeDetail.Author.AvatarThumb.Width)
	ret.Author.AvatarThumb.Height = float64(d.AwemeDetail.Author.AvatarThumb.Height)

	ret.Author.ShareInfo.ShareDescInfo = d.AwemeDetail.Author.ShareInfo.ShareDescInfo
	ret.Author.ShareInfo.ShareDesc = d.AwemeDetail.Author.ShareInfo.ShareDesc
	ret.Author.ShareInfo.ShareQrcodeUrl.Uri = d.AwemeDetail.Author.ShareInfo.ShareQrcodeUrl.Uri
	ret.Author.ShareInfo.ShareQrcodeUrl.UrlList = d.AwemeDetail.Author.ShareInfo.ShareQrcodeUrl.UrlList
	ret.Author.ShareInfo.ShareTitle = d.AwemeDetail.Author.ShareInfo.ShareTitle
	ret.Author.ShareInfo.ShareTitleMyself = d.AwemeDetail.Author.ShareInfo.ShareTitleMyself
	ret.Author.ShareInfo.ShareTitleOther = d.AwemeDetail.Author.ShareInfo.ShareTitleOther
	ret.Author.ShareInfo.ShareUrl = d.AwemeDetail.Author.ShareInfo.ShareUrl
	ret.Author.ShareInfo.ShareWeiboDesc = d.AwemeDetail.Author.ShareInfo.ShareWeiboDesc

	ret.Author.Status = d.AwemeDetail.Author.Status
	ret.Author.UserAge = d.AwemeDetail.Author.UserAge
	ret.Author.UserCanceled = d.AwemeDetail.Author.UserCanceled
	ret.Author.UserPermissions = d.AwemeDetail.Author.UserPermissions
	ret.Author.VerificationType = d.AwemeDetail.Author.VerificationType
	ret.Author.SpecialFollowStatus = d.AwemeDetail.Author.SpecialFollowStatus
	ret.Author.SpecialPeopleLabels = d.AwemeDetail.Author.SpecialPeopleLabels
	ret.Author.UniqueId = d.AwemeDetail.Author.UniqueId
	ret.Author.TotalFavorited = d.AwemeDetail.Author.TotalFavorited
	ret.Author.CfList = d.AwemeDetail.Author.CfList
	ret.Author.CloseFriendType = d.AwemeDetail.Author.CloseFriendType
	ret.Author.ContrailList = d.AwemeDetail.Author.ContrailList
	ret.Author.CustomVerify = d.AwemeDetail.Author.CustomVerify
	ret.Author.FamiliarVisitorUser = d.AwemeDetail.Author.FamiliarVisitorUser
	ret.Author.IsBan = d.AwemeDetail.Author.IsBan
	ret.Author.ImRoleIds = d.AwemeDetail.Author.ImRoleIds
	ret.Author.IsBlockingV2 = d.AwemeDetail.Author.IsBlockedV2
	ret.Author.IsBlockingV2 = d.AwemeDetail.Author.IsBlockingV2
	ret.Author.NotSeenItemIdListV2 = d.AwemeDetail.Author.NotSeenItemIdListV2
	ret.Author.NotSeenItemIdList = d.AwemeDetail.Author.NotSeenItemIdList
	ret.Author.OfflineInfoList = d.AwemeDetail.Author.OfflineInfoList
	ret.Author.SecUid = d.AwemeDetail.Author.SecUid
	ret.Author.SignatureExtra = d.AwemeDetail.Author.SignatureExtra

	ret.VideoData.WmVideoUrl = utils.First(d.AwemeDetail.Video.PlayAddr.UrlList)
	ret.VideoData.WmVideoUrlHQ = utils.First(d.AwemeDetail.Video.PlayAddr.UrlList)

	ret.VideoData.NwmVideoUrl = utils.First(d.AwemeDetail.Video.PlayAddr.UrlList)
	ret.VideoData.NwmVideoUrlHQ = utils.First(d.AwemeDetail.Video.PlayAddrH264.UrlList)

	ret.CoverData.Cover.UrlList = d.AwemeDetail.Video.Cover.UrlList
	ret.CoverData.Cover.Uri = d.AwemeDetail.Video.Cover.Uri
	ret.CoverData.Cover.Width = d.AwemeDetail.Video.Cover.Width
	ret.CoverData.Cover.Height = d.AwemeDetail.Video.Cover.Height

	ret.CoverData.OriginCover.UrlList = d.AwemeDetail.Video.OriginCover.UrlList
	ret.CoverData.OriginCover.Uri = d.AwemeDetail.Video.OriginCover.Uri
	ret.CoverData.OriginCover.Width = d.AwemeDetail.Video.OriginCover.Width
	ret.CoverData.OriginCover.Height = d.AwemeDetail.Video.OriginCover.Height

	ret.CoverData.DynamicCover.Uri = d.AwemeDetail.Video.DynamicCover.Uri
	ret.CoverData.DynamicCover.Width = d.AwemeDetail.Video.DynamicCover.Width
	ret.CoverData.DynamicCover.Height = d.AwemeDetail.Video.DynamicCover.Height
	ret.CoverData.DynamicCover.UrlList = d.AwemeDetail.Video.DynamicCover.UrlList

	ret.Music.PlayUrl.UrlList = d.AwemeDetail.Music.PlayUrl.UrlList
	ret.Music.PlayUrl.UrlKey = d.AwemeDetail.Music.PlayUrl.UrlKey
	ret.Music.PlayUrl.Uri = d.AwemeDetail.Music.PlayUrl.Uri
	ret.Music.PlayUrl.Height = d.AwemeDetail.Music.PlayUrl.Height
	ret.Music.PlayUrl.Width = d.AwemeDetail.Music.PlayUrl.Width

	ret.Images = d.AwemeDetail.Images
	return ret
}
