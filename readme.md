###编译程序说明
####框架内包含两个输入文件(阿姨文件、订单文件)  
####参赛者需编写相应的solution文件，实现其中的process函数。process为处理函数，框架会把输入文件直接传给process函数，process需要读取并处理文件，框架提供名为addSet的回调函数，将结果集合写入addSet接口，作为结果输出  
####执行sh compile.sh 编译（如果该框架包含compile，需要编译（c/go/java））
####执行sh exe.sh 1 2 3 4 运行程序，可以获得输出文件
######1是阿姨文件路径  
######2是订单文件路径  
######3是输出文件路径  
######4是计时文件路径  


###校验程序使用说明
####java -jar gamechecker-1.0-SNAPSHOT.jar 1 2 3
######1为结果文件全路径
######2为订单文件全路径
######3为阿姨文件全路径
#####源码地址https://igit.58corp.com/baojiandong/gamechecker
####执行校验程序需要有本地java环境，jdk版本1.6+
# go_frame_1
