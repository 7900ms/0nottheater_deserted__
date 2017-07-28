
### general

```
找到工作页面，然后

页面 https://version-control.axxxxxxx.edu.xx/svn/axxxxxxx/2017/s2/fcs

svn mkdir -m "create dir" --parents https://version-control.axxxxxxx.edu.xx/svn/axxxxxxx/2017/s2/fcs/svnexercise1
(svn rm https://version-control.axxxxxxx.edu.xx/svn/axxxxxxx/2017/s2/fcs/svnexercise1 -m "delete dir")

新页面 https://version-control.axxxxxxx.edu.xx/svn/axxxxxxx/2017/s2/fcs


cd ~/RHEL-Desktop-svn
svn checkout https://version-control.axxxxxxx.edu.xx/svn/axxxxxxx/2017/s2/fcs/svnexercise1
cd ~/RHEL-Desktop-svn/svnexercise1
la // 有 .svn 文件夹
svn status
touch 1.txt
svn status
svn add 1.txt
svn commit -m 'add 1.txt'



```

### school

~/RHEL-Desktop-svn/svn.md
```
查看区
https://version-control.adelaide.edu.au/svn/a1734530
https://myuni.adelaide.edu.au/courses/25355

低频查看区
https://github.com/7900ms/0nottheater_deserted__/tree/master/Usage_Manual/svn
https://cs.adelaide.edu.au/services/websubmission/
https://version-control.adelaide.edu.au/svn/a1734530/2017/s2/fcs/workshop-03/
https://version-control.adelaide.edu.au/svn/a1734530/2017/s2/fcs/prac-01
https://version-control.adelaide.edu.au/svn/a1734530/2017/s2/fcs/svnexercise
(a1734530, s2)


步骤：
1.
页面

页面
https://version-control.adelaide.edu.au/svn/a1734530/2017/s2/fcs/prac-02

如果404 则在 svn 服务器上新建
svn mkdir -m "create dir" --parents https://version-control.adelaide.edu.au/svn/a1734530/2017/s2/fcs/prac-02

得到
页面
https://version-control.adelaide.edu.au/svn/a1734530/2017/s2/fcs/prac-02

（如果想删除
svn rm https://version-control.adelaide.edu.au/svn/a1734530/2017/s2/fcs/prac-02 -m "delete dir" ）


2.
下载目录

在页面上确认有这个目录
cd ~/RHEL-Desktop-svn
svn checkout https://version-control.adelaide.edu.au/svn/a1734530/2017/s2/fcs/prac-02


3.
修改

cd prac-02
echo "hi" >> 1.txt
svn status

（撤销修改
svn revert 1.txt
（抓取
svn update


4.
上传

svn status
svn add 1.txt
svn status
svn commit -m 'c'

刷新页面

```

然后提交作业

```
https://cs.axxxxxxx.edu.xx/services/websubmission/
进入 fcs -- prac-01 - "Make Submission" - 提交最新版本
会提示一次 确认，然后再点击下一步 提交成功

```

-
