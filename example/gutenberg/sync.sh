mkdir -p data
rsync -av --include='*.txt' --include='*/' --exclude='*' aleph.gutenberg.org::gutenberg ./data