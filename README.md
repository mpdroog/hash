Hash
=============
Calculate Diablo hashfeed.

Calculation:
```
msgid=<>
hashfeed=1-120/360:8
MD5(msgid)=f1980b66e0a43405f3199d92774695cf
offset(0)=774695cf(hex)
offset(0)=2001114575(10)
modulo=pos % base +1 = 2001114575 % 360 +1
match= modulo >= from && modulo <= to = modulo >= 1 && modulo <= 120
```

To determine if `hashfeed 1-120/360:8` matches `<part29of143.RndMw4FFWQ9TdIWjYDmt@camelsystem-powerpost.local>`
and mismatches `<part6of143.htzsE$6rldAt7WuGTpr9@camelsystem-powerpost.local>`.
