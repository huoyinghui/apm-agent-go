// Code generated by "go generate". DO NOT EDIT.

package stacktrace

var libraryPackages = map[string]struct{}{
	"archive/tar":                       {},
	"archive/zip":                       {},
	"bufio":                             {},
	"bytes":                             {},
	"compress/bzip2":                    {},
	"compress/flate":                    {},
	"compress/gzip":                     {},
	"compress/lzw":                      {},
	"compress/zlib":                     {},
	"container/heap":                    {},
	"container/list":                    {},
	"container/ring":                    {},
	"context":                           {},
	"crypto":                            {},
	"crypto/aes":                        {},
	"crypto/cipher":                     {},
	"crypto/des":                        {},
	"crypto/dsa":                        {},
	"crypto/ecdsa":                      {},
	"crypto/elliptic":                   {},
	"crypto/hmac":                       {},
	"crypto/internal/cipherhw":          {},
	"crypto/md5":                        {},
	"crypto/rand":                       {},
	"crypto/rc4":                        {},
	"crypto/rsa":                        {},
	"crypto/sha1":                       {},
	"crypto/sha256":                     {},
	"crypto/sha512":                     {},
	"crypto/subtle":                     {},
	"crypto/tls":                        {},
	"crypto/x509":                       {},
	"crypto/x509/pkix":                  {},
	"database/sql":                      {},
	"database/sql/driver":               {},
	"debug/dwarf":                       {},
	"debug/elf":                         {},
	"debug/gosym":                       {},
	"debug/macho":                       {},
	"debug/pe":                          {},
	"debug/plan9obj":                    {},
	"encoding":                          {},
	"encoding/ascii85":                  {},
	"encoding/asn1":                     {},
	"encoding/base32":                   {},
	"encoding/base64":                   {},
	"encoding/binary":                   {},
	"encoding/csv":                      {},
	"encoding/gob":                      {},
	"encoding/hex":                      {},
	"encoding/json":                     {},
	"encoding/pem":                      {},
	"encoding/xml":                      {},
	"errors":                            {},
	"expvar":                            {},
	"flag":                              {},
	"fmt":                               {},
	"go/ast":                            {},
	"go/build":                          {},
	"go/constant":                       {},
	"go/doc":                            {},
	"go/format":                         {},
	"go/importer":                       {},
	"go/internal/gccgoimporter":         {},
	"go/internal/gcimporter":            {},
	"go/internal/srcimporter":           {},
	"go/parser":                         {},
	"go/printer":                        {},
	"go/scanner":                        {},
	"go/token":                          {},
	"go/types":                          {},
	"hash":                              {},
	"hash/adler32":                      {},
	"hash/crc32":                        {},
	"hash/crc64":                        {},
	"hash/fnv":                          {},
	"html":                              {},
	"html/template":                     {},
	"image":                             {},
	"image/color":                       {},
	"image/color/palette":               {},
	"image/draw":                        {},
	"image/gif":                         {},
	"image/internal/imageutil":          {},
	"image/jpeg":                        {},
	"image/png":                         {},
	"index/suffixarray":                 {},
	"internal/cpu":                      {},
	"internal/nettrace":                 {},
	"internal/poll":                     {},
	"internal/race":                     {},
	"internal/singleflight":             {},
	"internal/syscall/unix":             {},
	"internal/syscall/windows":          {},
	"internal/syscall/windows/registry": {},
	"internal/syscall/windows/sysdll":   {},
	"internal/testenv":                  {},
	"internal/testlog":                  {},
	"internal/trace":                    {},
	"io":                                {},
	"io/ioutil":                         {},
	"log":                               {},
	"log/syslog":                        {},
	"math":                              {},
	"math/big":                          {},
	"math/bits":                         {},
	"math/cmplx":                        {},
	"math/rand":                         {},
	"mime":                              {},
	"mime/multipart":                    {},
	"mime/quotedprintable":              {},
	"net":                            {},
	"net/http":                       {},
	"net/http/cgi":                   {},
	"net/http/cookiejar":             {},
	"net/http/fcgi":                  {},
	"net/http/httptest":              {},
	"net/http/httptrace":             {},
	"net/http/httputil":              {},
	"net/http/internal":              {},
	"net/http/pprof":                 {},
	"net/internal/socktest":          {},
	"net/mail":                       {},
	"net/rpc":                        {},
	"net/rpc/jsonrpc":                {},
	"net/smtp":                       {},
	"net/textproto":                  {},
	"net/url":                        {},
	"os":                             {},
	"os/exec":                        {},
	"os/signal":                      {},
	"os/signal/internal/pty":         {},
	"os/user":                        {},
	"path":                           {},
	"path/filepath":                  {},
	"plugin":                         {},
	"reflect":                        {},
	"regexp":                         {},
	"regexp/syntax":                  {},
	"runtime":                        {},
	"runtime/cgo":                    {},
	"runtime/debug":                  {},
	"runtime/internal/atomic":        {},
	"runtime/internal/sys":           {},
	"runtime/pprof":                  {},
	"runtime/pprof/internal/profile": {},
	"runtime/race":                   {},
	"runtime/trace":                  {},
	"sort":                           {},
	"strconv":                        {},
	"strings":                        {},
	"sync":                           {},
	"sync/atomic":                    {},
	"syscall":                        {},
	"testing":                        {},
	"testing/internal/testdeps":      {},
	"testing/iotest":                 {},
	"testing/quick":                  {},
	"text/scanner":                   {},
	"text/tabwriter":                 {},
	"text/template":                  {},
	"text/template/parse":            {},
	"time":                           {},
	"unicode":                        {},
	"unicode/utf16":                  {},
	"unicode/utf8":                   {},
	"unsafe":                         {},
	"vendor/golang_org/x/crypto/chacha20poly1305":                       {},
	"vendor/golang_org/x/crypto/chacha20poly1305/internal/chacha20":     {},
	"vendor/golang_org/x/crypto/cryptobyte":                             {},
	"vendor/golang_org/x/crypto/cryptobyte/asn1":                        {},
	"vendor/golang_org/x/crypto/curve25519":                             {},
	"vendor/golang_org/x/crypto/poly1305":                               {},
	"vendor/golang_org/x/net/http2/hpack":                               {},
	"vendor/golang_org/x/net/idna":                                      {},
	"vendor/golang_org/x/net/internal/nettest":                          {},
	"vendor/golang_org/x/net/lex/httplex":                               {},
	"vendor/golang_org/x/net/nettest":                                   {},
	"vendor/golang_org/x/net/proxy":                                     {},
	"vendor/golang_org/x/text/secure":                                   {},
	"vendor/golang_org/x/text/secure/bidirule":                          {},
	"vendor/golang_org/x/text/transform":                                {},
	"vendor/golang_org/x/text/unicode":                                  {},
	"vendor/golang_org/x/text/unicode/bidi":                             {},
	"vendor/golang_org/x/text/unicode/norm":                             {},
	"github.com/elastic/apm-agent-go":                                   {},
	"github.com/elastic/apm-agent-go/contrib/apmgin":                    {},
	"github.com/elastic/apm-agent-go/contrib/apmhttp":                   {},
	"github.com/elastic/apm-agent-go/contrib/apmsql":                    {},
	"github.com/elastic/apm-agent-go/contrib/apmsql/dsn":                {},
	"github.com/elastic/apm-agent-go/contrib/apmsql/pq":                 {},
	"github.com/elastic/apm-agent-go/contrib/apmsql/pq/pqdsn":           {},
	"github.com/elastic/apm-agent-go/contrib/apmsql/sqlite3":            {},
	"github.com/elastic/apm-agent-go/contrib/apmsql/sqlite3/sqlite3dsn": {},
	"github.com/elastic/apm-agent-go/internal/uuid":                     {},
	"github.com/elastic/apm-agent-go/model":                             {},
	"github.com/elastic/apm-agent-go/stacktrace":                        {},
	"github.com/elastic/apm-agent-go/transport":                         {},
	"github.com/elastic/apm-agent-go/transport/transporttest":           {},
}

// IsLibraryPackage reports whether or not the given package path is
// a well-known library path (stdlib or apm-agent-go).
func IsLibraryPackage(pkg string) bool {
	_, ok := libraryPackages[pkg]
	return ok
}
