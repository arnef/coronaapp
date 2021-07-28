GOPATH := $(shell go env GOPATH)
gettext := $(GOPATH)/bin/go-xgettext

all: armhf arm64

clean:
	clickable clean; rm -rf locales; rm -rf build;

pot:
	$(gettext) -o po/coronaapp.pot --keyword=gotext.Get $$(find . -name "*.go")

mo:
	@for f in $$(find po -name "*.po"); do mkdir -p locales/$$(basename $${f} .po); msgfmt $${f} -o locales/$$(basename $${f} .po)/default.mo; done

armhf: mo
	clickable clean --arch armhf; clickable build --arch armhf

arm64: mo
	clickable clean --arch arm64; clickable build --arch arm64
