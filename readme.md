# go-svelte-demo

演示使用svelte作为前端搭配go后端。

虽然前后端分离但是是依赖于sveltekit编译出的静态前端，没有单独的前端服务器。

将前端编译结果embed进可执行文件，部署分发时仅单文件。

在安装有go和node.js的环境中，运行 `go generate ./frontend` 编译前端，再 `go build .` 编译后端。

demo可执行文件中便包含前后端，可独立使用。