TODAY := $(shell date +'%Y%m%d')
FILEPATH = leetcode/$(TODAY).go

define INIT_GO_CODE
package leetcode

func Run${TODAY}() {
}
endef

export INIT_GO_CODE

.PHONY: new
new:
	@touch $(FILEPATH)
	@echo "$$INIT_GO_CODE" > $(FILEPATH)
	@echo "Initialized today's Go code:\n  $(FILEPATH)"
