// +build mipsle

package bolt

// maxMapSize represents the largest mmap size supported by Bolt.
const maxMapSize = 0x7FFFFFF // 128M

// maxAllocSize is the size used when creating array pointers.
const maxAllocSize = 0xFFFFFF

// Are unaligned load/stores broken on this arch?
var brokenUnaligned = false