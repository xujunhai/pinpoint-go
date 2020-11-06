thrift版本v0.13.0

##常见问题
####1.bison >= 2.5
macOS brew install bison
重命名 /Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/bin/bison
cp 安装/bin/bison 以上目录

####2.fatal error: 'openssl/ssl.h' file not found #include <openssl/ssl.h>
brew install openssl 
export LDFLAGS="-L/usr/local/opt/openssl@1.1/lib"
export CPPFLAGS="-I/usr/local/opt/openssl@1.1/include"

####3.fatal error: 'glib.h' file not found

####4.brew install thrift

###使用
thrift -r --gen go trace.thrift
