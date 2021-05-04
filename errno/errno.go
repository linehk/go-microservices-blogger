package errno

// 自定义的错误码
const (
	Success       = 200
	Error         = 300
	InvalidParams = 400

	TagNameIsExisted  = 10001
	GetExistedTagFail = 10002
	TagIsNotExist     = 10003
	GetAllTagFail     = 10004
	CountTagFail      = 10005
	AddTagFail        = 10006
	EditTagFail       = 10007
	DeleteTagFail     = 10008

	ArticleIsNotExist       = 20001
	CheckArticleIsExistFail = 20002
	AddArticleFail          = 20003
	DeleteArticleFail       = 20004
	EditArticleFail         = 20005
	CountArticleFail        = 20006
	GetArticleListFail      = 20007
	GetArticleFail          = 20008
)

// 错误码对应的错误消息
var Msg = map[int]string{
	Success:       "成功",
	Error:         "错误",
	InvalidParams: "请求参数错误",

	TagNameIsExisted:  "已存在该标签名称",
	GetExistedTagFail: "获取已存在标签失败",
	TagIsNotExist:     "该标签不存在",
	GetAllTagFail:     "获取所有标签失败",
	CountTagFail:      "统计标签失败",
	AddTagFail:        "新增标签失败",
	EditTagFail:       "修改标签失败",
	DeleteTagFail:     "删除标签失败",

	ArticleIsNotExist:       "该文章不存在",
	CheckArticleIsExistFail: "检查文章是否存在失败",
	AddArticleFail:          "新增文章失败",
	DeleteArticleFail:       "删除文章失败",
	EditArticleFail:         "修改文章失败",
	CountArticleFail:        "统计文章失败",
	GetArticleListFail:      "获取多个文章失败",
	GetArticleFail:          "获取单个文章失败",
}
