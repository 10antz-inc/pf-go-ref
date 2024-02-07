package option

// Scan は走査用のオプション
type Scan struct {
	PrivateField bool
	LenientType  bool
}

type ScanOption func(options *Scan)

// WithPrivateField は非公開フィールドもスキャンするかどうかを設定する
func WithPrivateField() ScanOption {
	return func(options *Scan) {
		options.PrivateField = true
	}
}

// WithLenientType は型が一致しなくてもスキャンするかどうかを設定する
func WithLenientType() ScanOption {
	return func(options *Scan) {
		options.LenientType = true
	}
}
