module github.com/mapprotocol/atlas

go 1.15

require (
	github.com/MadBase/MadNet v0.5.0
	github.com/VictoriaMetrics/fastcache v1.6.0
	github.com/btcsuite/btcd v0.23.5-0.20231215221805-96c9fd8078fd
	github.com/btcsuite/btcd/btcec/v2 v2.2.0
	github.com/btcsuite/btcd/btcutil v1.1.5
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0
	github.com/buraksezer/consistent v0.9.0
	github.com/cespare/cp v1.1.1
	github.com/cespare/xxhash/v2 v2.1.2
	github.com/davecgh/go-spew v1.1.1
	github.com/deckarep/golang-set v1.7.1
	github.com/dop251/goja v0.0.0-20211011172007-d99e4b8cbf48
	github.com/edsrzf/mmap-go v1.1.0
	github.com/ethereum/go-ethereum v1.10.10
	github.com/fatih/color v1.10.0
	github.com/fjl/memsize v0.0.1
	github.com/gballet/go-libpcsclite v0.0.0-20191108122812-4678299bea08
	github.com/golang/snappy v0.0.4
	github.com/google/uuid v1.1.5
	github.com/gorilla/websocket v1.5.0
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/herumi/bls-eth-go-binary v1.28.1
	github.com/holiman/bloomfilter/v2 v2.0.3
	github.com/holiman/uint256 v1.2.0
	github.com/influxdata/influxdb v1.8.5
	github.com/influxdata/influxdb-client-go/v2 v2.4.0
	github.com/logrusorgru/aurora v2.0.3+incompatible
	github.com/loinfish/azure-storage-go v0.0.1
	github.com/mapprotocol/compass v1.1.0
	github.com/mattn/go-colorable v0.1.12
	github.com/mattn/go-isatty v0.0.14
	github.com/minio/highwayhash v1.0.1
	github.com/minio/sha256-simd v1.0.0
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	github.com/olekukonko/tablewriter v0.0.5
	github.com/onsi/gomega v1.18.1
	github.com/peterh/liner v1.2.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/tsdb v0.10.0
	github.com/prysmaticlabs/fastssz v0.0.0-20220628131814-351fdcbb9964
	github.com/prysmaticlabs/go-bitfield v0.0.0-20210809151128-385d8c5e3fb7
	github.com/prysmaticlabs/gohashtree v0.0.2-alpha
	github.com/rjeczalik/notify v0.9.2
	github.com/rs/cors v1.8.2
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible
	github.com/shopspring/decimal v1.2.0
	github.com/status-im/keycard-go v0.0.0-20211109104530-b0e0482ba91d
	github.com/stretchr/testify v1.8.0
	github.com/supranational/blst v0.3.10
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
	github.com/tyler-smith/go-bip39 v1.0.2
	golang.org/x/crypto v0.1.0
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	golang.org/x/sys v0.1.0
	golang.org/x/term v0.1.0
	golang.org/x/text v0.4.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/olebedev/go-duktape.v3 v3.0.0-20200619000410-60c24ae608a6
	gopkg.in/urfave/cli.v1 v1.20.0
	gotest.tools v2.2.0+incompatible
)

replace github.com/ethereum/go-ethereum v1.10.10 => github.com/mapprotocol/go-ethereum v1.10.10-patch1 // indirect
