package golang2

import (
	"regexp"

	gsjs "github.com/bitly/go-simplejson"
	"github.com/flosch/pongo2"
)

//Render 渲染Jinja2模板，pongo2貌似只能渲染一维map，所以value一律按string类型传递
/*
{
	"Name": "name1",
	"Data": "{\"Key1\": \"value1\"}"
}
这种是可以被渲染，但是\"会被替换成HTML转义符&quot;
{
	"Name": "name1",
	"Data" {
		"Key1": "value1"
	}
}
这种Data字段无法正确渲染，这是渲染后的效果
最后返回之前，替换&quot;为"生成正确的字符串

newMap是模板变量key-value的封装，tplBytes是模板的字节切片
返回渲染后的模板字符串
*/
func Render(newMap map[string]interface{}, tplBytes []byte) (string, error) {
	tpl, err := pongo2.FromBytes(tplBytes) //生成模板
	ctx := pongo2.Context(newMap)          //封装模板变量
	tplSting, err := tpl.Execute(ctx)      //执行渲染
	re := regexp.MustCompile(`\&quot;`)
	tplSting = re.ReplaceAllString(tplSting, `"`) //如果有&quot;替换为"
	return tplSting, MyError(err)
}

//JSONRender 适合处理POST请求传递过来的json字符串
/*
使用github.com/bitly/go-simplejson是为了避免定义结构体去转换map
go-simplejson最大的好处是不必用结构体去序列化，而是直接处理json字符串
*/
func JSONRender(jsonStr string, tplBytes []byte) (string, error) {
	gsjs1, err := gsjs.NewJson([]byte(jsonStr))
	map1, err := gsjs1.Map()
	tplString, err := Render(map1, tplBytes)
	return tplString, err
}
