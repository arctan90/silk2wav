# silk2wav
Convert wechat audio files (Skype silk format) to wav format

# Build
`go build -o ./bin/wechat2wav  ./cmd` Tested under MacOS X

# Usage
```bash
chmod 755 ./bin/wechat2wav
./bin/wechat2wav inputfile outputfile
```

# License
forked from https://github.com/jdeng/silk2wav
silk source code (/interface & /src) is from https://github.com/collects/silk and under the Skype license.
silk.go is under the Apache license

