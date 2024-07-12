package models

type ImageUploadResult struct {
	ImageNo string `json:"image_no"` // 图片编号
}

type FileUploadResult struct {
	FileNo string `json:"file_no"` // 图片编号
}

type ImageRelationParams struct {
	BusinessNo   string   `json:"business_no"`   // 关联业务号,如订单号、会员号等
	ImgList      []string `json:"img_list"`      // 关联图片列表,调用图片上传接口获取image_no,总长不能超过256
	BusinessType int      `json:"business_type"` // 业务类型,1:订单关联,2:会员关联
}

type ImageRelationResult struct {
	BusinessNo string `json:"business_no"` // 关联业务号
}

type DownloadOrderCheckFileParams struct {
	BillDate string `json:"bill_date"` // 对账单日期,yyyy-MM-dd
}
