package main

import (
        "flag"
        "fmt"
        "io"
        "log"
        "net"
        "os"
        // "bufio"
        "golang.org/x/crypto/nacl/box"
)


type secureReader struct {
  public *[32]byte
  reader io.Reader
}
func (r *secureReader) Read([]byte) (int, error) {
  return 1, nil
}

type secureWriter struct {
  public *[32]byte
  writer io.Writer
}
func (w *secureWriter) Write([]byte) (int, error){
  return 1, nil
}

// NewSecureReader instantiates a new SecureReader
func NewSecureReader(r io.Reader, priv, pub *[32]byte) io.Reader {
        sr := &secureReader {
          public: pub,
          reader: r,
        }

        return sr
}

// NewSecureWriter instantiates a new SecureWriter
func NewSecureWriter(w io.Writer, priv, pub *[32]byte) io.Writer {
        return newSecureWriter(w, priv, pub)
}

func newSecureWriter(w io.Writer, priv, pub *[32]byte) *secureWriter {
  sw := &secureWriter{
    public: &[32]byte{},
    writer: w,
  }

  box.Precompute(sw.public, pub, priv)

  return sw
}

// Dial generates a private/public key pair,
// connects to the server, perform the handshake
// and return a reader/writer.
func Dial(addr string) (io.ReadWriteCloser, error) {
        return nil, nil
}

// Serve starts a secure echo server on the given listener.
func Serve(l net.Listener) error {
        return nil
}

func main() {
        port := flag.Int("l", 0, "Listen mode. Specify port")
        flag.Parse()

        // Server mode
        if *port != 0 {
                l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
                if err != nil {
                        log.Fatal(err)
                }
                defer l.Close()
                log.Fatal(Serve(l))
        }

        // Client mode
        if len(os.Args) != 3 {
                log.Fatalf("Usage: %s <port> <message>", os.Args[0])
        }
        conn, err := Dial("localhost:" + os.Args[1])
        if err != nil {
                log.Fatal(err)
        }
        if _, err := conn.Write([]byte(os.Args[2])); err != nil {
                log.Fatal(err)
        }
        buf := make([]byte, len(os.Args[2]))
        n, err := conn.Read(buf)
        if err != nil && err != io.EOF {
                log.Fatal(err)
        }
        fmt.Printf("%s\n", buf[:n])
}
