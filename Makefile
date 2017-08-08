PROTO_FILES = $(shell find protobuf/ -type f -name '*.proto')
PROTO_GO_FILES = $(patsubst protobuf/%.proto, protobuf/%.pb.go, $(PROTO_FILES))

IMPORT_PREFIX := github.com/d4l3k/pok/

# To edit in-place without creating a backup file, GNU sed requires a bare -i,
# while BSD sed requires an empty string as the following argument.
SED_INPLACE = sed $(shell sed --version 2>&1 | grep -q GNU && echo -i || echo "-i ''")
$(call make-lazy,SED_INPLACE)

.PHONY:
build: protobuf

.PHONY: protobuf
protobuf: protobuf/tensorflow $(PROTO_GO_FILES)

current_dir = $(shell pwd)
tensorflow_dir = "$(current_dir)/../../tensorflow/tensorflow/"

.PHONY: protobuf/tensorflow
protobuf/tensorflow:
	cd $(tensorflow_dir) && cp -u --parents `find -name \*.proto | sed '/tensorflow\/python/d'` $(current_dir)/protobuf/


%.pb.go: %.proto
	protoc --proto_path=${GOPATH}/src/github.com/d4l3k/pok/protobuf:. -I protobuf/ $< --gogoslick_out=plugins=grpc,import_prefix=$(IMPORT_PREFIX):.
	# Fixup standard packages wrongly imported by gogo because of import_prefix.
	$(SED_INPLACE) -E 's!$(IMPORT_PREFIX)(strings|reflect|math|strconv|bytes|errors|fmt|io|github\.com|golang\.org|google\.golang\.org)!\1!g' $@
	touch $@

