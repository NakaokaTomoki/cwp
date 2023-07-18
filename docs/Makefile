build_with_hugo: clean
	@mkdir -p $(TARGET_DIR)
	@cp $(SOURCE) $(TARGET_FILE)
	@sed -i -e 's/.\/docs\/static\/images/.\/images/g' $(TARGET_FILE)
	hugo --ignoreCache

start:
	hugo server -D
