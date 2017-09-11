# GoVM
go实现的虚拟机

代码来自 《自己动手写java虚拟机》。这本书不错，内容比较详实，缺点是用go语言实现，对初学者不太友好，这本书并非Go语言入门书籍。<br/>

全书读下来收获颇多，这本书主要内容集中在Class文件执行，以及一些java的特性，对细节描述不多，对编译器和垃圾收集器没有任何描述。<br/>

我本人也看了一些hotspot的源码，感觉入门JVM，首先要对C++有比较深的了解，尤其是JVM实现多线程，需要用C++处理一些信号，其他书籍我自己也翻了一下，比较推荐这么几本：<br/>
1.周志明的 《深入理解java虚拟机》<br/>

2.java虚拟机规范 （这本书有1.7和1.8的版本，读者根据使用的虚拟机自行购买，可以说，《自己动手写java虚拟机》这本书就是虚拟机规范的go语言翻译）

3.实战java虚拟机<br/>

<br/>

代码已经基本都写完了，但是由于书上是按照知识体系来划分目录的，和实际逻辑上有出入，将来拉个分支，会把目录重新整合一下。<br/>

书中内容加上了许多我自己读写过程中，认识的不太全面或者不太懂的，或者亲身验证，加上了自己理解的地方，如果不对希望大家多多指正，提个issue。<br/>

-Xms750m
-Xmx1000m

“代码缓存”，它是用来存储已编译方法生成的本地代码。
代码缓存确实很少引起性能问题，但是一旦发生其影响可能是毁灭性的。
如果代码缓存被占满，JVM会打印出一条警告消息，并切换到interpreted-only 模式：JIT编译器被停用，字节码将不再会被编译成机器码。
因此，应用程序将继续运行，但运行速度会降低一个数量级，直到有人注意到这个问题。就像其他内存区域一样，我们可以自定义代码缓存的大小。
相关的参数是-XX:InitialCodeCacheSize 和-XX:ReservedCodeCacheSize，它们的参数和上面介绍的参数一样，都是字节值。
-XX:ReservedCodeCacheSize=100m 
-XX:+UseCompressedOops

-server
-XX:+UseParNewGC
-XX:ParallelGCThreads=3
-XX:+UseConcMarkSweepGC

关于以下这两个参数,添加下面两个参数强制在remark阶段和FULL GC阶段之前先在进行一次年轻代的GC，这样需要进行处理的内存量就不会太多。
-XX:+ScavengeBeforeFullGC 

-XX:+CMSScavengeBeforeRemark
-XX:CMSInitiatingOccupancyFraction=73
-XX:SoftRefLRUPolicyMSPerMB=0
-Xverify:none

降低标记停顿，并发remark
-XX:+CMSParallelRemarkEnabled 

显示的执行System.gc()也是并发的
-XX:ExplicitGCInvokesConcurrent

-XX:+PrintGCDetails
-Xloggc:/Users/xiuyuhang/Downloads/gc-idea.log


from区使用到90%，再将对象送入老年代
-XX:TargetSurvivorRatio=90
老年代阈值上线
-XX:MaxTenuringThreshold=10
上面两个参数的使用有多种方式,设置新生代行为，没有通用准则。我们必须清楚以下2中情况：

1 如果从年龄分布中发现，有很多对象的年龄持续增长，在到达老年代阀值之前。这表示 -XX:MaxTenuringThreshold 设置过大

2 如果 -XX:MaxTenuringThreshold 的值大于1，但是很多对象年龄从未大于1.应该看下幸存区的目标使用率。如果幸存区使用率从未到达，这表示对象都被GC回收，这正是我们想要的。如果幸存区使用率经常达到，有些年龄超过1的对象被移动到老年代中。这种情况，可以尝试调整幸存区大小或目标使用率。

本仓库已经不再更新。
后面读周志明的《深入理解java虚拟机》，一些读书笔记会更新新的仓库。
