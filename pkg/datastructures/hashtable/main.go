package hashtable

type (
	item struct {
		next  *item
		key   string
		value string
	}
	bucket struct {
		node   *item
		length int
	}
)

var table = make([]*bucket, 100)

// hash function, adds up the rune values of the characters in the key
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	sum %= 100
	return sum
}

func getItemByKey(key string, bucketItem *item) *item {
	if bucketItem.key == key {
		return bucketItem
	}
	return getItemByKey(key, bucketItem.next)
}

// GetValueFromBucket get the value of the key
func GetValueFromBucket(key string) string {
	index := hash(key)
	bucketItem := table[index].node
	item := getItemByKey(key, bucketItem)
	return item.value
}

// Insert a key value pair into the table
func Insert(key string, value string) {
	indexToInsert := hash(key)
	newItem := item{nil, key, value}
	if table[indexToInsert] == nil {
		// new entry
		newBucket := bucket{&newItem, 1}
		table[indexToInsert] = &newBucket
	} else {
		// add next free linked item to the bucket
		table[indexToInsert].node.next = &newItem
		table[indexToInsert].length++
	}
}
