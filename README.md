# 实验结果

## 1.列出信息
命令
```
selpg -s 1 -e 1 in.txt
```
结果 

![image](https://wx2.sinaimg.cn/mw690/b41a0581gy1fw4rk32czfj20jf0aaaa9.jpg)



## 2.列出指定页数间隔信息
命令
```
selpg -s 1 -e 10 in.txt
```
结果
 

![image](https://wx4.sinaimg.cn/mw690/b41a0581gy1fw4rk324f4j20e10o3t8z.jpg)


## 3.导出文件
命令
```
selpg -s1 -e1 in.txt >out.txt
```
结果
 

![image](https://wx2.sinaimg.cn/mw690/b41a0581gy1fw4rk33tsyj20m10b8jsg.jpg)

## 4.导出提示信息
命令
```
selpg -s 1 -e 1 in.txt 2>error.txt
```
结果
 

![image](https://wx1.sinaimg.cn/mw690/b41a0581gy1fw4rk32edmj20in07eaa7.jpg)

## 5.用ps命令
命令
```
ps |  selpg -s 1 -e 1 in.txt 

```
结果 
 

![image](https://wx4.sinaimg.cn/mw690/b41a0581gy1fw4rk3hk4lj20dh06lgli.jpg)

## 6.用cat命令
命令
```
selpg -s 1 -e 1 in.txt | cat -n
```
结果 
 

![image](https://wx2.sinaimg.cn/mw690/b41a0581gy1fw4rk33sjgj20dp085t8q.jpg)

## 7.显示固定行数
命令
```
selpg -s 1 -e 1 -l 11  in.txt
```
结果   

![image](https://wx3.sinaimg.cn/mw690/b41a0581gy1fw4rk3538nj20d3062jr9.jpg)



## 8.出错提示  参数缺省

命令
```
selpg -s 1 -e 
```
结果 

![image](https://wx3.sinaimg.cn/mw690/b41a0581gy1fw4rk35ochj20cc04pq2z.jpg)

## 9.出错提示  参数大小出错
命令

```
selpg -s 5 -e 1 -l 11  in.txt
```
结果 

![image](https://wx2.sinaimg.cn/mw690/b41a0581gy1fw4rk36ftpj20b404vjrg.jpg)
