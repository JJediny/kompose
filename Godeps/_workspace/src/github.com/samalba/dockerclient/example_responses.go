package dockerclient

var haproxyPullOutput = `{"status":"The image you are pulling has been verified","id":"haproxy:1"}
{"status":"Already exists","progressDetail":{},"id":"511136ea3c5a"}{"status":"Already exists","progressDetail":{},"id":"1aeada447715"}{"status":"Already exists","progressDetail":{},"id":"479215127fa7"}{"status":"Already exists","progressDetail":{},"id":"66301eb54a7d"}{"status":"Already exists","progressDetail":{},"id":"e3990b07573f"}{"status":"Already exists","progressDetail":{},"id":"ecb4b23ca7ce"}{"status":"Already exists","progressDetail":{},"id":"f453e940c177"}{"status":"Already exists","progressDetail":{},"id":"fc5ea1bc05ab"}{"status":"Already exists","progressDetail":{},"id":"380557f8f7b3"}{"status":"The image you are pulling has been verified","id":"haproxy:1.4"}
{"status":"Already exists","progressDetail":{},"id":"511136ea3c5a"}{"status":"Already exists","progressDetail":{},"id":"1aeada447715"}{"status":"Already exists","progressDetail":{},"id":"479215127fa7"}{"status":"Already exists","progressDetail":{},"id":"63a1b9929e14"}{"status":"Already exists","progressDetail":{},"id":"af43bf7d176e"}{"status":"Already exists","progressDetail":{},"id":"851aac2d69aa"}{"status":"Already exists","progressDetail":{},"id":"345053a92c95"}{"status":"Already exists","progressDetail":{},"id":"b41231d429c9"}{"status":"The image you are pulling has been verified","id":"haproxy:1.4.25"}
{"status":"Already exists","progressDetail":{},"id":"511136ea3c5a"}{"status":"Already exists","progressDetail":{},"id":"1aeada447715"}{"status":"Already exists","progressDetail":{},"id":"479215127fa7"}{"status":"Already exists","progressDetail":{},"id":"63a1b9929e14"}{"status":"Already exists","progressDetail":{},"id":"af43bf7d176e"}{"status":"Already exists","progressDetail":{},"id":"851aac2d69aa"}{"status":"Already exists","progressDetail":{},"id":"345053a92c95"}{"status":"Already exists","progressDetail":{},"id":"b41231d429c9"}{"status":"The image you are pulling has been verified","id":"haproxy:1.5"}
{"status":"Already exists","progressDetail":{},"id":"511136ea3c5a"}{"status":"Already exists","progressDetail":{},"id":"1aeada447715"}{"status":"Already exists","progressDetail":{},"id":"479215127fa7"}{"status":"Already exists","progressDetail":{},"id":"66301eb54a7d"}{"status":"Already exists","progressDetail":{},"id":"e3990b07573f"}{"status":"Already exists","progressDetail":{},"id":"ecb4b23ca7ce"}{"status":"Already exists","progressDetail":{},"id":"f453e940c177"}{"status":"Already exists","progressDetail":{},"id":"fc5ea1bc05ab"}{"status":"Already exists","progressDetail":{},"id":"380557f8f7b3"}{"status":"The image you are pulling has been verified","id":"haproxy:1.5.10"}
{"status":"Already exists","progressDetail":{},"id":"511136ea3c5a"}{"status":"Already exists","progressDetail":{},"id":"1aeada447715"}{"status":"Already exists","progressDetail":{},"id":"479215127fa7"}{"status":"Already exists","progressDetail":{},"id":"66301eb54a7d"}{"status":"Already exists","progressDetail":{},"id":"e3990b07573f"}{"status":"Already exists","progressDetail":{},"id":"ecb4b23ca7ce"}{"status":"Already exists","progressDetail":{},"id":"f453e940c177"}{"status":"Already exists","progressDetail":{},"id":"fc5ea1bc05ab"}{"status":"Already exists","progressDetail":{},"id":"380557f8f7b3"}{"status":"The image you are pulling has been verified","id":"haproxy:1.5.9"}
{"status":"Already exists","progressDetail":{},"id":"511136ea3c5a"}{"status":"Already exists","progressDetail":{},"id":"1aeada447715"}{"status":"Already exists","progressDetail":{},"id":"479215127fa7"}{"status":"Already exists","progressDetail":{},"id":"66301eb54a7d"}{"status":"Already exists","progressDetail":{},"id":"e3990b07573f"}{"status":"Already exists","progressDetail":{},"id":"3d894e6f7e63"}{"status":"Already exists","progressDetail":{},"id":"4d949c40bc77"}{"status":"Already exists","progressDetail":{},"id":"55e031889365"}{"status":"Already exists","progressDetail":{},"id":"c7aa675e1876"}{"status":"The image you are pulling has been verified","id":"haproxy:latest"}
{"status":"Already exists","progressDetail":{},"id":"511136ea3c5a"}{"status":"Already exists","progressDetail":{},"id":"1aeada447715"}{"status":"Already exists","progressDetail":{},"id":"479215127fa7"}{"status":"Already exists","progressDetail":{},"id":"66301eb54a7d"}{"status":"Already exists","progressDetail":{},"id":"e3990b07573f"}{"status":"Already exists","progressDetail":{},"id":"ecb4b23ca7ce"}{"status":"Already exists","progressDetail":{},"id":"f453e940c177"}{"status":"Already exists","progressDetail":{},"id":"fc5ea1bc05ab"}{"status":"Already exists","progressDetail":{},"id":"380557f8f7b3"}{"status":"Status: Image is up to date for haproxy"}
`

var eventsResp = `{"status":"pull","id":"nginx:latest","time":1428620433}{"status":"create","id":"9b818c3b8291708fdcecd7c4086b75c222cb503be10a93d9c11040886032a48b","from":"nginx:latest","time":1428620433}{"status":"start","id":"9b818c3b8291708fdcecd7c4086b75c222cb503be10a93d9c11040886032a48b","from":"nginx:latest","time":1428620433}{"status":"die","id":"9b818c3b8291708fdcecd7c4086b75c222cb503be10a93d9c11040886032a48b","from":"nginx:latest","time":1428620442}{"status":"create","id":"352d0b412aae5a5d2b14ae9d88be59dc276602d9edb9dcc33e138e475b3e4720","from":"52.11.96.81/foobar/ubuntu:latest","time":1428620444}{"status":"start","id":"352d0b412aae5a5d2b14ae9d88be59dc276602d9edb9dcc33e138e475b3e4720","from":"52.11.96.81/foobar/ubuntu:latest","time":1428620444}{"status":"die","id":"352d0b412aae5a5d2b14ae9d88be59dc276602d9edb9dcc33e138e475b3e4720","from":"52.11.96.81/foobar/ubuntu:latest","time":1428620444}{"status":"pull","id":"debian:latest","time":1428620453}{"status":"create","id":"668887b5729946546b3072655dc6da08f0e3210111b68b704eb842adfce53f6c","from":"debian:latest","time":1428620453}{"status":"start","id":"668887b5729946546b3072655dc6da08f0e3210111b68b704eb842adfce53f6c","from":"debian:latest","time":1428620453}{"status":"die","id":"668887b5729946546b3072655dc6da08f0e3210111b68b704eb842adfce53f6c","from":"debian:latest","time":1428620453}{"status":"create","id":"eb4a19ec21ab29bbbffbf3ee2e2df9d99cb749780e1eff06a591cee5ba505180","from":"nginx:latest","time":1428620458}{"status":"start","id":"eb4a19ec21ab29bbbffbf3ee2e2df9d99cb749780e1eff06a591cee5ba505180","from":"nginx:latest","time":1428620458}{"status":"pause","id":"eb4a19ec21ab29bbbffbf3ee2e2df9d99cb749780e1eff06a591cee5ba505180","from":"nginx:latest","time":1428620462}{"status":"unpause","id":"eb4a19ec21ab29bbbffbf3ee2e2df9d99cb749780e1eff06a591cee5ba505180","from":"nginx:latest","time":1428620466}{"status":"die","id":"eb4a19ec21ab29bbbffbf3ee2e2df9d99cb749780e1eff06a591cee5ba505180","from":"nginx:latest","time":1428620469}`
