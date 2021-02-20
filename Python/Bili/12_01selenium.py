# selenium基本使用
from selenium import webdriver
from lxml import etree
from time import sleep
#实例化一个浏览器对象（传入浏览器的驱动）
bro = webdriver.Chrome(executable_path='./chromedriver.exe')
#让浏览器发起一个指定url对应请求
bro.get('http://scxk.nmpa.gov.cn:81/xk/')

#获取浏览器当前页面的页面源码数据
page_text=bro.page_source

#解析企业名称
tree = etree.HTML(page_text)
li_list=tree.xpath('//ul[@id="gzlist"]/li')

for li in li_list:
    name = li.xpath('./dl/a/text()')
    print(name)
sleep(3)
bro.quit()

#这个页面有反爬还是咋回事……没整明白
