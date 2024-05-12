package chrome

// DouYinAPIResult 通过Douyin_TikTok_Download_API接口返回的结果
type DouYinAPIResult struct {
	Code        int    `json:"code"`
	Router      string `json:"router"`
	AwemeDetail struct {
		ActivityVideoType int `json:"activity_video_type"`
		AnchorInfo        struct {
			Content string `json:"content"`
			Extra   string `json:"extra"`
			Icon    struct {
				Height  int           `json:"height"`
				Uri     string        `json:"uri"`
				UrlKey  string        `json:"url_key"`
				UrlList []interface{} `json:"url_list"`
				Width   int           `json:"width"`
			} `json:"icon"`
			Id        string `json:"id"`
			LogExtra  string `json:"log_extra"`
			MpUrl     string `json:"mp_url"`
			OpenUrl   string `json:"open_url"`
			StyleInfo struct {
				DefaultIcon string `json:"default_icon"`
				Extra       string `json:"extra"`
				SceneIcon   string `json:"scene_icon"`
			} `json:"style_info"`
			Title    string `json:"title"`
			TitleTag string `json:"title_tag"`
			Type     int    `json:"type"`
			WebUrl   string `json:"web_url"`
		} `json:"anchor_info"`
		Anchors             interface{} `json:"anchors"`
		AuthenticationToken string      `json:"authentication_token"`
		Author              struct {
			AvatarThumb struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"avatar_thumb"`
			AwemehtsGreetInfo string      `json:"awemehts_greet_info"`
			CfList            interface{} `json:"cf_list"`
			CloseFriendType   int         `json:"close_friend_type"`
			ContactsStatus    int         `json:"contacts_status"`
			ContrailList      interface{} `json:"contrail_list"`
			CoverUrl          []struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"cover_url"`
			CreateTime                             int         `json:"create_time"`
			CustomVerify                           string      `json:"custom_verify"`
			DataLabelList                          interface{} `json:"data_label_list"`
			EndorsementInfoList                    interface{} `json:"endorsement_info_list"`
			EnterpriseVerifyReason                 string      `json:"enterprise_verify_reason"`
			FavoritingCount                        int         `json:"favoriting_count"`
			FollowStatus                           int         `json:"follow_status"`
			FollowerCount                          int         `json:"follower_count"`
			FollowerListSecondaryInformationStruct interface{} `json:"follower_list_secondary_information_struct"`
			FollowerStatus                         int         `json:"follower_status"`
			FollowingCount                         int         `json:"following_count"`
			ImRoleIds                              interface{} `json:"im_role_ids"`
			IsAdFake                               bool        `json:"is_ad_fake"`
			IsBlockedV2                            bool        `json:"is_blocked_v2"`
			IsBlockingV2                           bool        `json:"is_blocking_v2"`
			IsCf                                   int         `json:"is_cf"`
			LiveHighValue                          int         `json:"live_high_value"`
			MaxFollowerCount                       int         `json:"max_follower_count"`
			Nickname                               string      `json:"nickname"`
			OfflineInfoList                        interface{} `json:"offline_info_list"`
			PersonalTagList                        interface{} `json:"personal_tag_list"`
			PreventDownload                        bool        `json:"prevent_download"`
			RiskNoticeText                         string      `json:"risk_notice_text"`
			RoomData                               string      `json:"room_data"`
			SecUid                                 string      `json:"sec_uid"`
			Secret                                 int         `json:"secret"`
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
		AwemeId               string      `json:"aweme_id"`
		AwemeType             int         `json:"aweme_type"`
		BoostStatus           int         `json:"boost_status"`
		Caption               string      `json:"caption"`
		CfRecheckTs           int         `json:"cf_recheck_ts"`
		ChallengePosition     interface{} `json:"challenge_position"`
		ChapterList           interface{} `json:"chapter_list"`
		CollectStat           int         `json:"collect_stat"`
		CollectionCornerMark  int         `json:"collection_corner_mark"`
		CommentGid            int64       `json:"comment_gid"`
		CommentList           interface{} `json:"comment_list"`
		CommentPermissionInfo struct {
			CanComment              bool `json:"can_comment"`
			CommentPermissionStatus int  `json:"comment_permission_status"`
			ItemDetailEntry         bool `json:"item_detail_entry"`
			PressEntry              bool `json:"press_entry"`
			ToastGuide              bool `json:"toast_guide"`
		} `json:"comment_permission_info"`
		CommerceConfigData     interface{} `json:"commerce_config_data"`
		CommonBarInfo          string      `json:"common_bar_info"`
		ComponentInfoV2        string      `json:"component_info_v2"`
		CoverLabels            interface{} `json:"cover_labels"`
		CreateScaleType        []string    `json:"create_scale_type"`
		CreateTime             int         `json:"create_time"`
		Desc                   string      `json:"desc"`
		DisableRelationBar     int         `json:"disable_relation_bar"`
		DislikeDimensionList   interface{} `json:"dislike_dimension_list"`
		DislikeDimensionListV2 interface{} `json:"dislike_dimension_list_v2"`
		DistributeCircle       struct {
			CampusBlockInteraction bool `json:"campus_block_interaction"`
			DistributeType         int  `json:"distribute_type"`
			IsCampus               bool `json:"is_campus"`
		} `json:"distribute_circle"`
		DuetAggregateInMusicTab  bool `json:"duet_aggregate_in_music_tab"`
		Duration                 int  `json:"duration"`
		EnableCommentStickerRec  bool `json:"enable_comment_sticker_rec"`
		EntertainmentProductInfo struct {
			MarketInfo struct {
				LimitFree struct {
					InFree bool `json:"in_free"`
				} `json:"limit_free"`
			} `json:"market_info"`
		} `json:"entertainment_product_info"`
		FallCardStruct struct {
			RecommendReason   string `json:"recommend_reason"`
			RecommendReasonV2 string `json:"recommend_reason_v2"`
		} `json:"fall_card_struct"`
		FeedCommentConfig struct {
		} `json:"feed_comment_config"`
		FriendRecommendInfo struct {
		} `json:"friend_recommend_info"`
		Geofencing        []interface{} `json:"geofencing"`
		GeofencingRegions interface{}   `json:"geofencing_regions"`
		GroupId           string        `json:"group_id"`
		GuideSceneInfo    struct {
		} `json:"guide_scene_info"`
		HasImageSticker     int         `json:"has_image_sticker"`
		HybridLabel         interface{} `json:"hybrid_label"`
		ImageAlbumMusicInfo struct {
			BeginTime int `json:"begin_time"`
			EndTime   int `json:"end_time"`
			Volume    int `json:"volume"`
		} `json:"image_album_music_info"`
		ImageComment struct {
		} `json:"image_comment"`
		ImageCropCtrl  int         `json:"image_crop_ctrl"`
		ImageInfos     interface{} `json:"image_infos"`
		ImageList      interface{} `json:"image_list"`
		Images         []Image     `json:"images"`
		ImgBitrate     interface{} `json:"img_bitrate"`
		ImpressionData struct {
			GroupIdListA   []interface{} `json:"group_id_list_a"`
			GroupIdListB   []interface{} `json:"group_id_list_b"`
			GroupIdListC   []interface{} `json:"group_id_list_c"`
			SimilarIdListA interface{}   `json:"similar_id_list_a"`
			SimilarIdListB interface{}   `json:"similar_id_list_b"`
		} `json:"impression_data"`
		IncentiveItemType    int         `json:"incentive_item_type"`
		InteractionStickers  interface{} `json:"interaction_stickers"`
		Is24Story            int         `json:"is_24_story"`
		IsAds                bool        `json:"is_ads"`
		IsCollectsSelected   int         `json:"is_collects_selected"`
		IsDuetSing           bool        `json:"is_duet_sing"`
		IsImageBeat          bool        `json:"is_image_beat"`
		IsLifeItem           bool        `json:"is_life_item"`
		IsSharePost          bool        `json:"is_share_post"`
		IsStory              int         `json:"is_story"`
		IsTop                int         `json:"is_top"`
		IsUseMusic           bool        `json:"is_use_music"`
		ItemTitle            string      `json:"item_title"`
		ItemWarnNotification struct {
			Content string `json:"content"`
			Show    bool   `json:"show"`
			Type    int    `json:"type"`
		} `json:"item_warn_notification"`
		LabelTopText        interface{} `json:"label_top_text"`
		LibfinsertTaskId    string      `json:"libfinsert_task_id"`
		LifeAnchorShowExtra struct {
			AnchorType    int  `json:"anchor_type"`
			HasAnchorInfo bool `json:"has_anchor_info"`
			ShouldShow    bool `json:"should_show"`
		} `json:"life_anchor_show_extra"`
		LongVideo            interface{} `json:"long_video"`
		MarkLargelyFollowing bool        `json:"mark_largely_following"`
		MediaType            int         `json:"media_type"`
		Music                struct {
			Album             string        `json:"album"`
			ArtistUserInfos   interface{}   `json:"artist_user_infos"`
			Artists           []interface{} `json:"artists"`
			AuditionDuration  int           `json:"audition_duration"`
			Author            string        `json:"author"`
			AuthorDeleted     bool          `json:"author_deleted"`
			AuthorPosition    interface{}   `json:"author_position"`
			BindedChallengeId int           `json:"binded_challenge_id"`
			CanBackgroundPlay bool          `json:"can_background_play"`
			CollectStat       int           `json:"collect_stat"`
			CoverColorHsv     struct {
				H int `json:"h"`
				S int `json:"s"`
				V int `json:"v"`
			} `json:"cover_color_hsv"`
			CoverHd struct {
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
			DmvAutoShow                    bool          `json:"dmv_auto_show"`
			DspStatus                      int           `json:"dsp_status"`
			Duration                       int           `json:"duration"`
			EndTime                        int           `json:"end_time"`
			ExternalSongInfo               []interface{} `json:"external_song_info"`
			Extra                          string        `json:"extra"`
			Id                             int64         `json:"id"`
			IdStr                          string        `json:"id_str"`
			IsAudioUrlWithCookie           bool          `json:"is_audio_url_with_cookie"`
			IsCommerceMusic                bool          `json:"is_commerce_music"`
			IsDelVideo                     bool          `json:"is_del_video"`
			IsMatchedMetadata              bool          `json:"is_matched_metadata"`
			IsOriginal                     bool          `json:"is_original"`
			IsOriginalSound                bool          `json:"is_original_sound"`
			IsPgc                          bool          `json:"is_pgc"`
			IsRestricted                   bool          `json:"is_restricted"`
			IsVideoSelfSee                 bool          `json:"is_video_self_see"`
			LyricShortPosition             interface{}   `json:"lyric_short_position"`
			Mid                            string        `json:"mid"`
			MusicChartRanks                interface{}   `json:"music_chart_ranks"`
			MusicCollectCount              int           `json:"music_collect_count"`
			MusicCoverAtmosphereColorValue string        `json:"music_cover_atmosphere_color_value"`
			MusicStatus                    int           `json:"music_status"`
			MusicianUserInfos              interface{}   `json:"musician_user_infos"`
			MuteShare                      bool          `json:"mute_share"`
			OfflineDesc                    string        `json:"offline_desc"`
			OwnerHandle                    string        `json:"owner_handle"`
			OwnerNickname                  string        `json:"owner_nickname"`
			PgcMusicType                   int           `json:"pgc_music_type"`
			PlayUrl                        struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlKey  string   `json:"url_key"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"play_url"`
			Position                  interface{} `json:"position"`
			PreventDownload           bool        `json:"prevent_download"`
			PreventItemDownloadStatus int         `json:"prevent_item_download_status"`
			PreviewEndTime            int         `json:"preview_end_time"`
			PreviewStartTime          int         `json:"preview_start_time"`
			ReasonType                int         `json:"reason_type"`
			Redirect                  bool        `json:"redirect"`
			SchemaUrl                 string      `json:"schema_url"`
			SearchImpr                struct {
				EntityId string `json:"entity_id"`
			} `json:"search_impr"`
			ShootDuration int `json:"shoot_duration"`
			Song          struct {
				Artists interface{} `json:"artists"`
				Chorus  struct {
					DurationMs int `json:"duration_ms"`
					StartMs    int `json:"start_ms"`
				} `json:"chorus"`
				Id    int64  `json:"id"`
				IdStr string `json:"id_str"`
				Title string `json:"title"`
			} `json:"song"`
			SourcePlatform int `json:"source_platform"`
			StartTime      int `json:"start_time"`
			Status         int `json:"status"`
			StrongBeatUrl  struct {
				Height  int      `json:"height"`
				Uri     string   `json:"uri"`
				UrlList []string `json:"url_list"`
				Width   int      `json:"width"`
			} `json:"strong_beat_url"`
			TagList           interface{} `json:"tag_list"`
			Title             string      `json:"title"`
			UnshelveCountries interface{} `json:"unshelve_countries"`
			UserCount         int         `json:"user_count"`
			VideoDuration     int         `json:"video_duration"`
		} `json:"music"`
		NicknamePosition    interface{}   `json:"nickname_position"`
		OriginCommentIds    interface{}   `json:"origin_comment_ids"`
		OriginTextExtra     []interface{} `json:"origin_text_extra"`
		Original            int           `json:"original"`
		OriginalImages      interface{}   `json:"original_images"`
		PackedClips         interface{}   `json:"packed_clips"`
		PhotoSearchEntrance struct {
			EcomType int `json:"ecom_type"`
		} `json:"photo_search_entrance"`
		Position           interface{}   `json:"position"`
		PreviewTitle       string        `json:"preview_title"`
		PreviewVideoStatus int           `json:"preview_video_status"`
		Promotions         []interface{} `json:"promotions"`
		Rate               int           `json:"rate"`
		Region             string        `json:"region"`
		RelationLabels     interface{}   `json:"relation_labels"`
		RiskInfos          struct {
			Content  string `json:"content"`
			RiskSink bool   `json:"risk_sink"`
			Type     int    `json:"type"`
			Vote     bool   `json:"vote"`
			Warn     bool   `json:"warn"`
		} `json:"risk_infos"`
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
		ShareRecExtra      string `json:"share_rec_extra"`
		ShareUrl           string `json:"share_url"`
		ShouldOpenAdReport bool   `json:"should_open_ad_report"`
		ShowFollowButton   struct {
		} `json:"show_follow_button"`
		SocialTagList interface{} `json:"social_tag_list"`
		Statistics    struct {
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
		TextExtra []struct {
			CaptionEnd   int    `json:"caption_end"`
			CaptionStart int    `json:"caption_start"`
			End          int    `json:"end"`
			HashtagId    string `json:"hashtag_id"`
			HashtagName  string `json:"hashtag_name"`
			IsCommerce   bool   `json:"is_commerce"`
			Start        int    `json:"start"`
			Type         int    `json:"type"`
		} `json:"text_extra"`
		UniqidPosition      interface{} `json:"uniqid_position"`
		UserDigged          int         `json:"user_digged"`
		UserRecommendStatus int         `json:"user_recommend_status"`
		Video               struct {
			BigThumbs []interface{} `json:"big_thumbs"`
			BitRate   []struct {
				FPS       int    `json:"FPS"`
				HDRBit    string `json:"HDR_bit"`
				HDRType   string `json:"HDR_type"`
				BitRate   int    `json:"bit_rate"`
				Format    string `json:"format"`
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
			BitRateAudio  interface{} `json:"bit_rate_audio"`
			CdnUrlExpired int         `json:"cdn_url_expired"`
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
			Format        string `json:"format"`
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
		VideoControl struct {
			AllowDouplus             bool   `json:"allow_douplus"`
			AllowDownload            bool   `json:"allow_download"`
			AllowDuet                bool   `json:"allow_duet"`
			AllowDynamicWallpaper    bool   `json:"allow_dynamic_wallpaper"`
			AllowMusic               bool   `json:"allow_music"`
			AllowReact               bool   `json:"allow_react"`
			AllowRecord              bool   `json:"allow_record"`
			AllowShare               bool   `json:"allow_share"`
			AllowStitch              bool   `json:"allow_stitch"`
			DisableRecordReason      string `json:"disable_record_reason"`
			DownloadIgnoreVisibility bool   `json:"download_ignore_visibility"`
			DownloadInfo             struct {
				FailInfo struct {
					Code   int    `json:"code"`
					Msg    string `json:"msg"`
					Reason string `json:"reason"`
				} `json:"fail_info"`
				Level int `json:"level"`
			} `json:"download_info"`
			DraftProgressBar     int  `json:"draft_progress_bar"`
			DuetIgnoreVisibility bool `json:"duet_ignore_visibility"`
			DuetInfo             struct {
				Level int `json:"level"`
			} `json:"duet_info"`
			PreventDownloadType   int  `json:"prevent_download_type"`
			ShareGrayed           bool `json:"share_grayed"`
			ShareIgnoreVisibility bool `json:"share_ignore_visibility"`
			ShareType             int  `json:"share_type"`
			ShowProgressBar       int  `json:"show_progress_bar"`
			TimerStatus           int  `json:"timer_status"`
		} `json:"video_control"`
		VideoGameDataChannelConfig struct {
		} `json:"video_game_data_channel_config"`
		VideoLabels          interface{} `json:"video_labels"`
		VideoShareEditStatus int         `json:"video_share_edit_status"`
		VideoTag             []struct {
			Level   int    `json:"level"`
			TagId   int    `json:"tag_id"`
			TagName string `json:"tag_name"`
		} `json:"video_tag"`
		VideoText        []interface{} `json:"video_text"`
		VisualSearchInfo struct {
			IsEcomImg         bool `json:"is_ecom_img"`
			IsShowImgEntrance bool `json:"is_show_img_entrance"`
		} `json:"visual_search_info"`
		XiguaBaseInfo struct {
			ItemId           int `json:"item_id"`
			StarAltarOrderId int `json:"star_altar_order_id"`
			StarAltarType    int `json:"star_altar_type"`
			Status           int `json:"status"`
		} `json:"xigua_base_info"`
	} `json:"data"`
}
