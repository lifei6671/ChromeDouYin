package chrome

type MediaType int

const (
	// MediaTypeImage 图片类型
	MediaTypeImage MediaType = 2
	// MediaTypeVideo 视频类型
	MediaTypeVideo MediaType = 4
)

func (t MediaType) String() string {
	switch t {
	case MediaTypeImage:
		return "image"
	case MediaTypeVideo:
		return "video"
	default:
		return "unknown"
	}
}
