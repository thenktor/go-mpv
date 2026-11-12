//go:build ios

package mpv

/*
#cgo CFLAGS: -I/opt/xcframework/Libmpv.xcframework/ios-arm64/Headers
#cgo LDFLAGS: -F/opt/xcframework/Libmpv.xcframework/ios-arm64 -lmpv-combined
*/
import "C"
