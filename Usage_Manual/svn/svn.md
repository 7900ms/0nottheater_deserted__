
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

~/RHEL-Desktop-svn/svn1.md

记得提交作业
```
https://cs.axxxxxxx.edu.xx/services/websubmission/
进入 fcs -- prac-01 - "Make Submission" - 提交最新版本
会提示一次 确认，然后再点击下一步 提交成功
```

-
