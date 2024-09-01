// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package tcpcreatesocket represents the imported interface "wasi:sockets/tcp-create-socket@0.2.0".
package tcpcreatesocket

import (
	"github.com/fermyon/spin-go-sdk/v2/internal/wasi/sockets/v0.2.0/network"
	"github.com/fermyon/spin-go-sdk/v2/internal/wasi/sockets/v0.2.0/tcp"
	"github.com/ydnar/wasm-tools-go/cm"
)

// CreateTCPSocket represents the imported function "create-tcp-socket".
//
// Create a new TCP socket.
//
// Similar to `socket(AF_INET or AF_INET6, SOCK_STREAM, IPPROTO_TCP)` in POSIX.
// On IPv6 sockets, IPV6_V6ONLY is enabled by default and can't be configured otherwise.
//
// This function does not require a network capability handle. This is considered
// to be safe because
// at time of creation, the socket is not bound to any `network` yet. Up to the moment
// `bind`/`connect`
// is called, the socket is effectively an in-memory configuration object, unable
// to communicate with the outside world.
//
// All sockets are non-blocking. Use the wasi-poll interface to block on asynchronous
// operations.
//
// # Typical errors
// - `not-supported`:     The specified `address-family` is not supported. (EAFNOSUPPORT)
// - `new-socket-limit`:  The new socket resource could not be created because of
// a system limit. (EMFILE, ENFILE)
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/socket.html>
// - <https://man7.org/linux/man-pages/man2/socket.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-wsasocketw>
// - <https://man.freebsd.org/cgi/man.cgi?query=socket&sektion=2>
//
//	create-tcp-socket: func(address-family: ip-address-family) -> result<tcp-socket,
//	error-code>
//
//go:nosplit
func CreateTCPSocket(addressFamily network.IPAddressFamily) (result cm.Result[tcp.TCPSocket, tcp.TCPSocket, network.ErrorCode]) {
	addressFamily0 := (uint32)(addressFamily)
	wasmimport_CreateTCPSocket((uint32)(addressFamily0), &result)
	return
}

//go:wasmimport wasi:sockets/tcp-create-socket@0.2.0 create-tcp-socket
//go:noescape
func wasmimport_CreateTCPSocket(addressFamily0 uint32, result *cm.Result[tcp.TCPSocket, tcp.TCPSocket, network.ErrorCode])