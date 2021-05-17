package hashtable

type (
	item struct {
		next  *item
		key   string
		value interface{}
	}

	// Bucket The hash Bucket
	Bucket struct {
		node   *item
		length int
	}

	// HashInterface Common hash methods
	HashInterface interface {
		GetTable() []*Bucket
		ResetTable()
		GetValueFromBucket(key string) interface{}
		Insert(key string, value interface{})
	}
)

var table = make([]*Bucket, 100)

// hash function, adds up the rune values of the characters in the key
func hash(key string) int {
	sum := 0
	for i, v := range key {
		sum += int(v) + i
	}
	sum %= 97
	return sum
}

func getItemByKey(key string, bucketItem *item) *item {
	if bucketItem.key == key {
		return bucketItem
	}
	return getItemByKey(key, bucketItem.next)
}

// GetValueFromBucket get the value of the key
func GetValueFromBucket(key string) interface{} {
	index := hash(key)
	if len(table) == 0 {
		return nil
	}
	if table[index] != nil {
		bucketItem := table[index].node
		item := getItemByKey(key, bucketItem)
		return item.value
	}
	return ""
}

// GetHashBucket Get the hash table Bucket
func GetHashBucket() []*Bucket {
	return table
}

//ResetTable Reset the bashtable Bucket
func ResetTable() {
	table = make([]*Bucket, 100)
}

// Insert a key value pair into the table
func Insert(key string, value interface{}) {
	indexToInsert := hash(key)
	newItem := item{nil, key, value}
	if table[indexToInsert] == nil {
		// new entry
		newBucket := Bucket{&newItem, 1}
		table[indexToInsert] = &newBucket
	} else {
		// add next free linked item to the Bucket
		table[indexToInsert].node.next = &newItem
		table[indexToInsert].length++
	}
}
