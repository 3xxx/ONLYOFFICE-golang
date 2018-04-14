# ONLYOFFICE-golang
基于golang go语言（beego框架）下的ONLYOFFICE Document Server二次开发。主要功能为文档的管理（上传、列表、删除等）和根据ONLYOFFICE Document Server中的文档进行更新和回调、历史版本数据管理。这个项目来自3xxx/EngineerCMS，完整的代码也可参考EngineerCMS。

Second development of ONLYOFFICE Document Server based on the golang go language (beego framework). The main functions are Document management (upload, list, delete, etc.) and update and callback based on the documents in the ONLYOFFICE Document Server,history changes data manage. This project is extracted from 3xxx/EngineerCMS.

ONLYOFFICE Document Server is a free collaborative online office suite comprising viewers and editors for texts, spreadsheets and presentations, fully compatible with Office Open XML formats: .docx, .xlsx, .pptx and enabling collaborative editing in real time.

ONLYOFFICE文档服务器是一个免费的在线协作办公套件，包括word文本、电子表格和演示文稿的查看器和编辑器，完全兼容office Open XML格式:.docx、.xlsx、.pptx，并实时支持协作编辑。

我很喜欢ONLYOFFICE，它功能强大，只要在局域网内任意一台电脑安装一下，全局的人都能享受多人在线编辑文档。能完全取代桌面程序office的word/excel/ppt,所以我对它进行了二次开发，方便大家在windows环境下部署和使用强大的实时文档协作，改变传统的汇总文档：投标，工作总结，工作量统计，工程进度汇报等，月报等，会议纪要是否有意见，直接发个链接给大家，有意见需要修改的，直接在上面改了，不用再通过邮件发回给汇总负责人，效率大大提高，工作方式也优雅很多。

我的其他资料可供参考：主要是界面汉化，中文字体，二次开发（仅仅是文档管理：增删改权限设置），局域网中的安装和部署等细节。
网盘共享文件地址：[https://pan.baidu.com/s/1gf0ucuR](https://pan.baidu.com/s/1gf0ucuR)

[如何在 Windows 上 使用 ONLYOFFICE 协同编辑文档](http://blog.csdn.net/hotqin888/article/details/79337881)

[onlyoffice的优势和特点](https://www.onlyoffice.com/document-editor-comparison.aspx)

[onlyoffice与ms office online的对比](https://help.onlyoffice.com/products/files/doceditor.aspx?fileid=4476630&doc=ZEFoK2lNMW1ZQjhRNjNyY2JnWk5MaVAvTUU4dmdhV3ZHWGtOL01GUStydz0_IjQ0NzY2MzAi0&_ga=2.247409280.564294813.1518263805-199484730.1518174023)

golang语言开发的文档管理：增删改权限设置
![engineercms onlyoffice](https://user-images.githubusercontent.com/10678867/38768484-0a55e06e-4027-11e8-9871-fc65e1686408.png)

文档的用户权限和角色（用户组）权限
登录用户，对于已经进行了权限设置的文档，将根据权限数据库，比对用户名，当与用户有关时，就显示相对应的权限（编辑、评论、只读），当都与登录用户无关时，则显示拒绝访问；
如果是用户自己上传的，则允许编辑。
对于未登录用户，已经设置了权限的文档，都将显示拒绝访问；
对于登录和未登录用户，未进行权限设置的文档，则显示允许编辑。
![snap2](https://user-images.githubusercontent.com/10678867/38768561-61f67742-4028-11e8-96f5-9bdc24ca4d71.png)
![snap3](https://user-images.githubusercontent.com/10678867/38768560-61b14398-4028-11e8-9c2a-501b5142064f.png)

word文件在线实时协作效果
![onlyoffice word](https://user-images.githubusercontent.com/10678867/36413270-9fc09c00-1658-11e8-817e-3e58021a8253.jpg)

powerpoint文件在线实时协作效果
![onlyoffice powerpoint](https://user-images.githubusercontent.com/10678867/36413278-a61110ee-1658-11e8-9955-14241b8b8bd7.jpg)

excel文件在线实时协作效果
![onlyoffice excel](https://user-images.githubusercontent.com/10678867/36413285-ab0dcd8a-1658-11e8-9a11-4c94709efea8.jpg)

word历史版本控制
![word_history](https://user-images.githubusercontent.com/10678867/37910779-eb9f9c54-3140-11e8-969c-8677abc67dc4.png)

