package render

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang-web/internal/models"
)

func TestAddDefaultDate(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, r, "home.page.gohtml", &models.TemplateData{})
	if err != nil {
		t.Error("error writing template to browser")
	}

	err = RenderTemplate(&ww, r, "non-existent.page.gohtml", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template that does not exist")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}

func TestNewTemplates(t *testing.T) {
	NewTemplate(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}

func TestCreateTemplateCache_TemplateParseError(t *testing.T) {
	// Tạo thư mục tạm và file page với nội dung không hợp lệ.
	tmpDir := t.TempDir()
	badPage := filepath.Join(tmpDir, "bad.page.gohtml")
	err := os.WriteFile(badPage, []byte("{{"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	// Gán lại pathToTemplates cho thư mục tạm.
	originalPath := pathToTemplates
	pathToTemplates = tmpDir
	defer func() { pathToTemplates = originalPath }()

	_, err = CreateTemplateCache()
	if err == nil {
		t.Error("expected error from parsing invalid template")
	}
}

// TestCreateTemplateCache_LayoutParseGlobError tạo file layout không hợp lệ để bắt lỗi từ ParseGlob.
func TestCreateTemplateCache_LayoutParseGlobError(t *testing.T) {
	// Tạo thư mục tạm với file page hợp lệ và file layout không hợp lệ.
	tmpDir := t.TempDir()
	goodPage := filepath.Join(tmpDir, "good.page.gohtml")
	err := os.WriteFile(goodPage, []byte("Hello, {{.Name}}!"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	badLayout := filepath.Join(tmpDir, "bad.layout.gohtml")
	err = os.WriteFile(badLayout, []byte("{{"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	originalPath := pathToTemplates
	pathToTemplates = tmpDir
	defer func() { pathToTemplates = originalPath }()

	_, err = CreateTemplateCache()
	if err == nil {
		t.Error("expected error from parsing invalid layout template")
	}
}
