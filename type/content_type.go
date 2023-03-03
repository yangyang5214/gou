package typeutil

type ContentType = string

const (
	TextHtml                  ContentType = "text/html"
	ImagePng                  ContentType = "image/png"
	ImageJpeg                 ContentType = "image/jpeg"
	ApplicationJson           ContentType = "application/json"
	TextJs                    ContentType = "text/javascript"
	TextCss                   ContentType = "text/css"
	ApplicationJsonUtf8       ContentType = "application/json;charset=utf-8"
	ApplicationFormUrlencoded ContentType = "application/x-www-form-urlencoded"
	MultipartFormData         ContentType = "multipart/form-data"
)
