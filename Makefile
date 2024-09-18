VERSION = 2.3.0-pre0

# ----------------------------------------------------------------------
# Test
# ----------------------------------------------------------------------
.PHONY: test
test: test
	tinygo test -target=wasi -gc=leaking -v ./http
	tinygo test -target=wasi -gc=leaking -v ./redis

.PHONY: test-integration
test-integration:
	go test -v -count=1 .

# ----------------------------------------------------------------------
# Generate C bindings
# ----------------------------------------------------------------------
GENERATED_SPIN_VARIABLES = variables/spin-config.c variables/spin-config.h
GENERATED_OUTBOUND_HTTP  = http/wasi-outbound-http.c http/wasi-outbound-http.h
GENERATED_SPIN_HTTP      = http/spin-http.c http/spin-http.h
GENERATED_OUTBOUND_REDIS = redis/outbound-redis.c redis/outbound-redis.h
GENERATED_SPIN_REDIS     = redis/spin-redis.c redis/spin-redis.h
GENERATED_KEY_VALUE      = kv/key-value.c kv/key-value.h
GENERATED_SQLITE         = sqlite/sqlite.c sqlite/sqlite.h
GENERATED_LLM            = llm/llm.c llm/llm.h
GENERATED_OUTBOUND_MYSQL = mysql/outbound-mysql.c mysql/outbound-mysql.h
GENERATED_OUTBOUND_PG    = pg/outbound-pg.c pg/outbound-pg.h

SDK_VERSION_SOURCE_FILE  = sdk_version/sdk-version-go-template.c

# NOTE: Please update this list if you add a new directory to the SDK:
SDK_VERSION_DEST_FILES   = variables/sdk-version-go.c http/sdk-version-go.c \
			   kv/sdk-version-go.c redis/sdk-version-go.c \
				 sqlite/sdk-version-go.c llm/sdk-version-go.c

# NOTE: To generate the C bindings you need to install a forked version of wit-bindgen.
#
#   cargo install wit-bindgen-cli --git https://github.com/fermyon/wit-bindgen-backport --rev "b89d5079ba5b07b319631a1b191d2139f126c976"
#
.PHONY: generate
generate: $(GENERATED_OUTBOUND_HTTP) $(GENERATED_SPIN_HTTP)
generate: $(GENERATED_OUTBOUND_REDIS) $(GENERATED_SPIN_REDIS)
generate: $(GENERATED_SPIN_VARIABLES) $(GENERATED_KEY_VALUE)
generate: $(GENERATED_SQLITE) $(GENERATED_LLM)
generate: $(GENERATED_OUTBOUND_MYSQL) $(GENERATED_OUTBOUND_PG)
generate: $(SDK_VERSION_DEST_FILES)

$(SDK_VERSION_DEST_FILES): $(SDK_VERSION_SOURCE_FILE)
	export commit="$$(git rev-parse HEAD)"; \
	sed -e "s/{{VERSION}}/${VERSION}/" -e "s/{{COMMIT}}/$${commit}/" < $< > $@

$(GENERATED_SPIN_VARIABLES):
	wit-bindgen c --import wit/spin-config.wit --out-dir ./variables

$(GENERATED_OUTBOUND_HTTP):
	wit-bindgen c --import wit/wasi-outbound-http.wit --out-dir ./http

$(GENERATED_SPIN_HTTP):
	wit-bindgen c --export wit/spin-http.wit --out-dir ./http

$(GENERATED_OUTBOUND_REDIS):
	wit-bindgen c --import wit/outbound-redis.wit --out-dir ./redis

$(GENERATED_SPIN_REDIS):
	wit-bindgen c --export wit/spin-redis.wit --out-dir ./redis

$(GENERATED_KEY_VALUE):
	wit-bindgen c --import wit/key-value.wit --out-dir ./kv

$(GENERATED_SQLITE):
	wit-bindgen c --import wit/sqlite.wit --out-dir ./sqlite

$(GENERATED_LLM):
	wit-bindgen c --import wit/llm.wit --out-dir ./llm

$(GENERATED_OUTBOUND_MYSQL):
	wit-bindgen c --import wit/outbound-mysql.wit --out-dir ./mysql

$(GENERATED_OUTBOUND_PG):
	wit-bindgen c --import wit/outbound-pg.wit --out-dir ./pg

# ----------------------------------------------------------------------
# Cleanup
# ----------------------------------------------------------------------
.PHONY: clean
clean:
	rm -f $(GENERATED_SPIN_CONFIG)
	rm -f $(GENERATED_OUTBOUND_HTTP) $(GENERATED_SPIN_HTTP)
	rm -f $(GENERATED_OUTBOUND_REDIS) $(GENERATED_SPIN_REDIS)
	rm -f $(GENERATED_KEY_VALUE) $(GENERATED_SQLITE)
	rm -f $(GENERATED_LLM)
	rm -f $(GENERATED_OUTBOUND_MYSQL)
	rm -f $(GENERATED_SDK_VERSION)
	rm -f $(SDK_VERSION_DEST_FILES)
