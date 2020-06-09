package constant

type serviceKey = string

const (
	randAuth serviceKey = "kzmake.traefikcfa.randauth.v1"
)

// Service はサービスの定義です。
var Service = struct {
	RandAuth serviceKey
}{
	RandAuth: randAuth,
}
