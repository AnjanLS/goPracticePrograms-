strNum=input()
lowest=-2**31
greatest=2**31-1
if strNum[0] == "-":
    remain=strNum[1:]
    num=int(remain)
    rev=int(str(num)[::-1])
    rev=-rev
else:
    num=int(strNum)
    rev=int(str(num)[::-1])
if lowest<rev<greatest:
    print(rev)
else:
    print(0)
