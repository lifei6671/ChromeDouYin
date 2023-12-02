package chrome

type ImageList struct {
	FileId   string   `json:"fileId"`
	Height   int      `json:"height"`
	InfoList InfoList `json:"infoList"`
	TraceId  string   `json:"traceId"`
	Url      string   `json:"url"`
	Width    int      `json:"width"`
}

type XiaoHongShuDetail struct {
	Desc      string      `json:"desc"`
	ImageList []ImageList `json:"imageList"`
}

type XHSDetail struct {
	Code int      `json:"code"`
	Desc string   `json:"desc"`
	List []string `json:"list"`
}

type InfoList []struct {
	ImageScene string `json:"imageScene"`
	Url        string `json:"url"`
}

func FromXiaoHongShu(d *XiaoHongShuDetail) XHSDetail {
	detail := XHSDetail{
		Desc: d.Desc,
	}
	for _, info := range d.ImageList {
		for _, img := range info.InfoList {
			if img.ImageScene == "CRD_WM_WEBP" {
				detail.List = append(detail.List, img.Url)
			}
		}

	}
	return detail
}
