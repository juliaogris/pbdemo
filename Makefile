protos:
	protoc -I. --descriptor_set_out=hello.pb hello.proto





# use hermit in make even if not activated
define nl


endef
ifndef ACTIVE_HERMIT
$(eval $(subst \n,$(nl),$(shell bin/hermit env -r | sed 's/^\(.*\)$$/export \1\\n/')))
endif
