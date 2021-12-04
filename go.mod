module github.com/udistrital/plan_adquisiciones_mid

go 1.15

require (
	github.com/astaxie/beego v1.12.3
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/smartystreets/goconvey v1.6.4
	github.com/udistrital/movimientos_crud v0.3.1-0.20211204010755-1c9a5d3f9af4
	github.com/udistrital/utils_oas v0.0.0-20211125230753-1091d2af48e2
	golang.org/x/crypto v0.0.0-20211202192323-5770296d904e // indirect
	golang.org/x/net v0.0.0-20211203184738-4852103109b8 // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/astaxie/beego v1.12.3 => github.com/udistrital/beego v1.12.4-0.20211126032252-ee78ca48b207
