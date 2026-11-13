package proxy

import "fmt"

type Role string

const (
	Admin    Role = "ADMIN"
	Standard Role = "STANDARD"
)

type User struct {
	Name string
	Role Role
}

type Document interface {
	DisplayReport()
}

type RealDocument struct {
	filePath string
}

func (r *RealDocument) DisplayReport() {
	fmt.Printf("Displaying report contents from: %s\n", r.filePath)
}

func NewRealDocument(filePath string) *RealDocument {
	fmt.Printf("[Heavy I/O] Loading file from disk: %s\n", filePath)
	return &RealDocument{
		filePath: filePath,
	}
}

type DocumentProxy struct {
	filePath string
	user     User
	realDoc  *RealDocument // Lazy initialized reference
}

func (d *DocumentProxy) DisplayReport() {
	if d.user.Role != Admin {
		fmt.Printf("Access Denied: User '%s' (Role: %s) is not authorized to view confidential reports.\n", d.user.Name, d.user.Role)
		return
	}

	if d.realDoc == nil {
		fmt.Println("Proxy: Authorization passed. Initializing RealDocument...")
		d.realDoc = NewRealDocument(d.filePath)
	}

	d.realDoc.DisplayReport()
}

func NewDocumentProxy(filePath string, user User) Document {
	return &DocumentProxy{
		filePath: filePath,
		user:     user,
	}
}
