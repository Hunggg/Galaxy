package fetcher

func FindMinBlocks(blocks ...uint64) uint64 {
	if len(blocks) == 0 {
		return 0
	}
	minBlock := blocks[0]
	for _, b := range blocks {
		if b < minBlock {
			minBlock = b
		}
	}
	return minBlock
}

func FindMaxBlocks(blocks ...uint64) uint64 {
	if len(blocks) == 0 {
		return 0
	}
	maxBlock := blocks[0]
	for _, b := range blocks {
		if b > maxBlock {
			maxBlock = b
		}
	}
	return maxBlock
}
